package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"flux-web/models"

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
	FluxUrl:            models.GetEnv("FLUX_URL", "http://flux:3030"),
	SyncApi:            "/api/flux/v6/sync?ref=",
	JobApi:             "/api/flux/v6/jobs?id=",
	UpdateManifestsApi: "/api/flux/v9/update-manifests",
	ListImagesApi:      "/api/flux/v10/images?namespace=",
}

func (this *WorkloadController) ListWorkloads() {
	ns := this.Ctx.Input.Param(":ns")
	l.Printf("in ListWorkloads, executing: " + flux.FluxUrl + flux.ListImagesApi + ns)

	res, err := httplib.Get(flux.FluxUrl + flux.ListImagesApi + ns).Debug(true).Bytes()
	if err != nil {
		l.Panic(err.Error)
	}

	workloads, err := json.Marshal(models.NewWorkloads(res))
	if err != nil {
		l.Panic(err.Error)
	}
	this.Ctx.Output.Body(workloads)
}

func (this *WorkloadController) ReleaseWorkloads() {
	newreleaseRequest, _ := models.NewReleseRequest(this.Ctx.Input.RequestBody)
	releaseRequest := newreleaseRequest.GetReleaseRequestJSON()

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
	releaseResult.Status = models.InRelease

	models.MemPut(releaseResult.GetReleaseResultKey(), releaseResult.Status)

	syncID, err := getSyncID(jobID)
	if err != nil {
		releaseResult.Status = models.ReleaseFaild
		models.MemPut(releaseResult.GetReleaseResultKey(), releaseResult.Status)
		l.Printf(err.Error())
		return
	}

	l.Printf("found new syncID: " + syncID)

	for {
		l.Printf("waiting for syncID: " + syncID + " to finish...")
		resp, err := httplib.Get(flux.FluxUrl + flux.SyncApi + syncID).String()
		if err != nil {
			releaseResult.Status = models.ReleaseFaild
			models.MemPut(releaseResult.GetReleaseResultKey(), releaseResult.Status)
			l.Printf(err.Error())
			break
		}
		if resp == "[]" {
			releaseResult.Status = models.UpToDate
			models.MemPut(releaseResult.GetReleaseResultKey(), releaseResult.Status)
			l.Printf("release for " + newreleaseRequest.Workload + " is done!")
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
			l.Println("Error_getSyncID_02")
			return "", errors.New(job.Err)
		} else {
			l.Printf("job status: " + job.StatusString)
		}
		time.Sleep(time.Second)
	}
	return "", nil
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
			l.Panic(err.Error)
			return "", errors.New(err.Error())
		}
		l.Println(string(bodyBytes))
		jobID := strings.Replace(string(bodyBytes), "\"", "", -1)
		l.Println("job " + jobID + " triggered")
		return string(jobID), nil
	} else {
		return "", errors.New("Job request statuscode is: " + string(resp.StatusCode))
	}
}

func Auth(c *context.Context) {
	if readOnly, err := strconv.ParseBool(models.GetEnv("READ_ONLY", "false")); err != nil {
		c.Abort(401, "Not boolean value for READ_ONLY")
		return
	} else if readOnly {
		c.Abort(401, "Not authorized")
		return
	}
}
