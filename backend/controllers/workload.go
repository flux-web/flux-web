package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/flux-web/flux-web/conf"

	"github.com/flux-web/flux-web/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

const (
	statusUpToDate      = "up to date"
	actionUpdateRelease = "updateRelease"
)

type WorkloadController struct {
	Config conf.Config
	beego.Controller
}

type jobResponse struct {
	Result       interface{} `json:"-"`
	Err          string      `json:"Err"`
	StatusString string      `json:"StatusString"`
}

var l = logs.GetLogger()

var flux = models.Flux{
	FluxUrl:            os.Getenv("FLUX_URL"),
	SyncApi:            "/api/flux/v6/sync?ref=",
	JobApi:             "/api/flux/v6/jobs?id=",
	UpdateManifestsApi: "/api/flux/v9/update-manifests",
	ListImagesApi:      "/api/flux/v10/images?namespace=",
	ListServices:       "/api/flux/v11/services?namespace=",
}

func (this *WorkloadController) ListWorkloads() {
	ns := this.Ctx.Input.Param(":ns")
	res, err := httplib.Get(flux.FluxUrl + flux.ListImagesApi + ns).Debug(true).Bytes()
	if err != nil {
		l.Panic(err.Error())
	}
	var images []models.Image
	images, err = models.NewImages(res)

	services, err := models.NewServices(flux.FluxUrl + flux.ListServices + ns)
	if err != nil {
		l.Printf("Found error: " + err.Error())
		this.Ctx.Output.SetStatus(500)
	}

	var workloads []models.Workload
	workloads = models.NewWorkloads(images, services)

	workloadsResponse, err := json.Marshal(workloads)

	this.Ctx.Output.Body(workloadsResponse)
}

func (this *WorkloadController) ReleaseWorkloads() {
	releaseRequest, err := models.NewReleseRequest(this.Ctx.Input.RequestBody)
	if err != nil {
		l.Printf("Found error: " + err.Error())
		this.Ctx.Output.SetStatus(500)
		return
	}

	services, err := models.NewServices(flux.FluxUrl + flux.ListServices + releaseRequest.Namespace)
	if err != nil {
		l.Printf("Found error: " + err.Error())
		this.Ctx.Output.SetStatus(500)
	}

	currentStatusAutomated := services.GetStatusAutomateByWorkload(releaseRequest.Workload)
	if currentStatusAutomated != releaseRequest.Automated {
		l.Printf("Switch automated workload stature from %t to %t\n", currentStatusAutomated, releaseRequest.Automated)
		err = this.automateWorkload(releaseRequest)
		if err != nil {
			l.Printf("Found error: " + err.Error())
			this.Ctx.Output.SetStatus(500)
			return
		}
	}

	if !services.WantedImageAlreadyDeployed(releaseRequest.Workload, releaseRequest.Target) {
		this.updateWorkload(releaseRequest)
	} else {
		r := models.ReleaseResult{
			Status:    statusUpToDate,
			Workload:  releaseRequest.Workload,
			Container: releaseRequest.Container,
			Tag:       releaseRequest.Target,
			Action:    actionUpdateRelease,
		}
		broadcastReleaseResult(r)
		l.Printf("Image %s is already deployed!\n", releaseRequest.Target)
	}
	this.Ctx.WriteString("Done")
}

func (this *WorkloadController) updateWorkload(releaseRequest models.ReleaseRequest) {
	jsonRequest, err := releaseRequest.GetReleaseRequestJSON(this.Config.FluxUser)
	if err != nil {
		l.Printf("Found error: " + err.Error())
		this.Ctx.Output.SetStatus(500)
		return
	}

	jobID, err := triggerJob(jsonRequest)
	if err != nil {
		l.Printf("Found error: " + err.Error())
		this.Ctx.Output.SetStatus(500)
		return
	}
	this.Ctx.WriteString("Done")

	go func(jobID string, releaseRequest models.ReleaseRequest) {
		waitForSync(jobID, releaseRequest)
	}(jobID, releaseRequest)
}

func (this *WorkloadController) automateWorkload(releaseRequest models.ReleaseRequest) error {
	jsonRequest, err := releaseRequest.GetAutomatedRequestJSON(this.Config.FluxUser)
	if err != nil {
		return err
	}
	jobID, err := triggerJob(jsonRequest)
	if err != nil {
		return err
	}

	timeout := time.After(time.Duration(this.Config.PollTimeout) * time.Second)
	ticker := time.Tick(time.Duration(this.Config.PollInterval) * time.Millisecond)
mainLoop:
	for {
		select {
		case <-timeout:
			l.Printf("jobID: " + jobID + " timed out")
			return errors.New("timeout while automateWorkload")
		case <-ticker:
			l.Printf("waiting for jobID: " + jobID + " to finish...")
			jobStatus, err := fetchJobstatus(flux.FluxUrl + flux.JobApi + jobID)
			if err != nil {
				return err
			}
			if jobStatus.StatusString == statusSucceeded {
				l.Printf("automate for workload" + releaseRequest.Workload + " is done!")
				break mainLoop
			}
		}
	}
	return nil
}

func waitForSync(jobID string, newreleaseRequest models.ReleaseRequest) {
	l.Printf("getting syncID...")

	var releaseResult models.ReleaseResult
	releaseResult.Workload = newreleaseRequest.Workload
	releaseResult.Container = newreleaseRequest.Container
	releaseResult.Tag = newreleaseRequest.Target
	releaseResult.Status = "release failed"
	releaseResult.Action = actionUpdateRelease

	syncID, err := getSyncID(jobID)
	if err != nil {
		l.Printf(err.Error())
		if err.Error() == "no changes found" {
			releaseResult.Status = err.Error()
			broadcastReleaseResult(releaseResult)
		}
		return
	}

	l.Printf("found new syncID: " + syncID)

	for {
		l.Printf("waiting for syncID: " + syncID + " to finish...")
		resp, err := httplib.Get(flux.FluxUrl + flux.SyncApi + syncID).String()
		if err != nil {
			l.Printf(err.Error())
			break
		}
		if resp == "[]" {
			releaseResult.Status = statusUpToDate
			l.Printf("release for " + newreleaseRequest.Workload + " is done!")
			break
		}
		time.Sleep(time.Millisecond * 300)
	}

	broadcastReleaseResult(releaseResult)
}

func getSyncID(jobID string) (string, error) {
	l.Printf("getting syncID...")

	for {
		resp, err := httplib.Get(flux.FluxUrl + flux.JobApi + jobID).Bytes()
		if err != nil {
			l.Println(err.Error())
			return "", errors.New(err.Error())
		}
		job, err := models.NewJob(resp)
		if err != nil {
			return "", errors.New(err.Error())
		}
		if job.Result.Revision != "" {
			l.Println("got syncID: " + job.Result.Revision)
			return job.Result.Revision, nil
		} else if job.Err != "" {
			l.Printf("job error: " + job.Err)
			return "", errors.New(job.Err)
		} else {
			l.Printf("job status: " + job.StatusString)
		}
		time.Sleep(time.Second)
	}
}

func triggerJob(requestBody []byte) (string, error) {
	resp, err := http.Post(flux.FluxUrl+flux.UpdateManifestsApi, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		l.Printf("Error_triggerJob_01: " + err.Error())
		return "", errors.New(err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			l.Panic(err.Error())
			return "", errors.New(err.Error())
		}
		l.Println(string(bodyBytes))
		jobID := strings.Replace(string(bodyBytes), "\"", "", -1)
		l.Println("job " + jobID + " triggered")
		return string(jobID), nil
	} else {
		return "", errors.New("Job request statuscode is: " + strconv.Itoa(resp.StatusCode))
	}
}

func GetImages(params ...string) []models.Image {
	namespace := os.Getenv("DEFAULT_NAMESPACE")
	if len(params) > 0 {
		namespace = params[0]
		l.Printf(namespace)
	}
	res, err := httplib.Get(flux.FluxUrl + flux.ListImagesApi + namespace).Debug(true).Bytes()
	if err != nil {
		l.Panic(err.Error())
	}

	images, err := models.NewImages(res)
	if err != nil {
		l.Panic(err.Error())
	}
	if len(params) > 1 {
		filter := params[1]
		for i := 0; i < len(images); i++ {
			if !strings.Contains(images[i].ID, filter) {
				images = append(images[:i], images[i+1:]...)
				i--
			}
		}
	}
	return images
}

func Auth(c *context.Context) {
	if readOnly, err := strconv.ParseBool(os.Getenv("READ_ONLY")); err != nil {
		c.Abort(401, "Not boolean value for READ_ONLY")
		return
	} else if readOnly {
		c.Abort(401, "Not authorized")
		return
	}
}

func fetchJobstatus(url string) (jobResponse, error) {
	jobres := jobResponse{}
	client := http.Client{
		Timeout: time.Second * 20,
	}
	resp, err := client.Get(url)
	if err != nil {
		return jobres, err
	}
	if resp == nil {
		return jobres, errors.New("fetchJobStatus, http.Get resp is nil")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return jobres, err
	}
	err = json.Unmarshal(data, &jobres)
	if err != nil {
		return jobres, err
	}
	return jobres, nil
}

func broadcastReleaseResult(r models.ReleaseResult) {
	jsonString, err := json.Marshal(r)
	if err != nil {
		l.Println(err)
	}
	h.broadcast <- jsonString
}
