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

	"github.com/flux-web/flux-web/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

type WorkloadController struct {
	beego.Controller
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
	l.Printf("in ListWorkloads, executing: " + flux.FluxUrl + flux.ListImagesApi + ns)
	res, err := httplib.Get(flux.FluxUrl + flux.ListImagesApi + ns).Debug(true).Bytes()
	if err != nil {
		l.Panic(err.Error())
	}
	var images []models.Image
	images, err = models.NewImages(res)

	res2, err2 := httplib.Get(flux.FluxUrl + flux.ListServices + ns).Debug(true).Bytes()
	if err2 != nil {
		l.Panic(err.Error())
	}
	var services []models.Service
	services, err = models.NewServices(res2)

	var workloads []models.Workload
	workloads = models.NewWorkloads(images, services)

	workloadsResponse, err := json.Marshal(workloads)

	this.Ctx.Output.Body(workloadsResponse)
}

func (this *WorkloadController) ReleaseWorkloads() {
	newreleaseRequest, _ := models.NewReleseRequest(this.Ctx.Input.RequestBody)

	releaseRequest, err := newreleaseRequest.GetReleaseRequestJSON()
	if err != nil {
		l.Printf("Found error: " + err.Error())
		this.Ctx.Output.SetStatus(500)
		return
	}

	jobID, err := triggerJob(releaseRequest)
	if err != nil {
		l.Printf("Found error: " + err.Error())
		this.Ctx.Output.SetStatus(500)
		return
	}
	this.Ctx.WriteString("Done")

	go func(jobID string, newreleaseRequest models.ReleaseRequest) {
		waitForSync(jobID, newreleaseRequest)
	}(jobID, newreleaseRequest)
}

func waitForSync(jobID string, newreleaseRequest models.ReleaseRequest) {
	l.Printf("getting syncID...")

	var releaseResult models.ReleaseResult
	releaseResult.Workload = newreleaseRequest.Workload
	releaseResult.Container = newreleaseRequest.Container
	releaseResult.Tag = newreleaseRequest.Target
	releaseResult.Status = "release failed"
	releaseResult.Action = "updateRelease"

	syncID, err := getSyncID(jobID)
	if err != nil {
		l.Printf(err.Error())
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
			releaseResult.Status = "up to date"
			l.Printf("release for" + newreleaseRequest.Workload + " is done!")
			break
		}
		time.Sleep(time.Millisecond * 300)
	}

	jsonString, err := json.Marshal(releaseResult)
	if err != nil {
		l.Println(err)
	}
	h.broadcast <- jsonString
}

func getSyncID(jobID string) (string, error) {
	l.Printf("getting syncID...")

	for {
		url := flux.FluxUrl + flux.JobApi + jobID
		resp, err := httplib.Get(url).Bytes()
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
	url := flux.FluxUrl + flux.UpdateManifestsApi
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
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
