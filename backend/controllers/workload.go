package controllers

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"flux-web/models"

	"github.com/astaxie/beego"
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
}

func (this *WorkloadController) ListWorkloads() {
	ns := this.Ctx.Input.Param(":ns")
	//this.Data["read_only"] = os.Getenv("READ_ONLY")
	//this.Data["namespaces"] = strings.Split(os.Getenv("NAMESPACES"), ";")
	//this.Data["fluxUrl"] = flux.FluxUrl
	//this.Data["workloads"] = GetImages(ns, this.Input().Get("filter"))
	//this.TplName = "main.tpl"
	res, err := httplib.Get(flux.FluxUrl + flux.ListImagesApi + ns).Debug(true).Bytes()
	if err != nil {
		l.Panic(err.Error)
	}
	this.Ctx.Output.Body(res)
}

func (this *WorkloadController) ReleaseWorkloads() {
	jobID, err := triggerJob(this.Ctx.Input.RequestBody)
	if err != nil {
		l.Printf(err.Error())
		this.Ctx.Output.SetStatus(500)
		return
	}
	syncID, err := getSyncID(jobID)
	if err != nil {
		l.Printf(err.Error())
		this.Ctx.Output.SetStatus(500)
		return
	}
	this.Ctx.Output.SetStatus(waitForSync(syncID))
}

func waitForSync(syncID string) int {
	l.Printf("waiting for sync: " + syncID)

	for true {
		resp, err := httplib.Get(flux.FluxUrl + flux.SyncApi + syncID).String()
		if err != nil {
			l.Printf(err.Error())
			break
		}
		if resp == "[]" {
			return 200
			break
		}
		time.Sleep(time.Second)
	}
	return 500
}

func getSyncID(jobID string) (string, error) {
	l.Printf("getting syncID...")

	for true {
		resp, err := httplib.Get(flux.FluxUrl + flux.JobApi + jobID).Bytes()
		if err != nil {
			l.Printf(err.Error())
			return "", errors.New(err.Error())
		}
		job, err := models.NewJob(resp)
		if err != nil {
			l.Panic("Error_getSyncID_01: " + err.Error())
			return "", errors.New(err.Error())
		}
		if job.Result.Revision != "" {
			l.Printf("got syncID: " + job.Result.Revision)
			return job.Result.Revision, nil
		} else if job.Err != "" {
			l.Printf("Error_getSyncID_02: " + job.Err)
			return job.Err, errors.New(job.Err)
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
		l.Printf(string(bodyBytes))
		jobID := strings.Replace(string(bodyBytes), "\"", "", -1)
		l.Printf("job " + jobID + " triggered")
		return string(jobID), nil
	} else {
		return "", errors.New("Job request statuscode is: " + string(resp.StatusCode))
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
		l.Panic(err.Error)
	}

	images, err := models.NewImages(res)
	if err != nil {
		l.Panic(err.Error)
	}
	if len(params) > 1 {
		filter := params[1]
		l.Printf(filter)
		for i := 0; i < len(images); i++ {
			if !strings.Contains(images[i].ID, filter) {
				images = append(images[:i], images[i+1:]...)
				i--
			}
		}
	}
	return images
}

// func (this *WebSocketController) Join() {
// 	// Upgrade from http request to WebSocket.
// 	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
// 	if _, ok := err.(websocket.HandshakeError); ok {
// 		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
// 		return
// 	} else if err != nil {
// 		beego.Error("Cannot setup WebSocket connection:", err)
// 		return
// 	}
// 	// Message receive loop.
// 	for {
// 		mt, message, err := ws.ReadMessage()
// 		if err != nil {
// 			log.Println("read:", err)
// 			break
// 		}
// 		log.Printf("recv: %s", message)

// 		jobID, err := triggerJob(message)
// 		if err != nil {
// 			err = ws.WriteMessage(mt, message)
// 			if err != nil {
// 				log.Println("write:", err)
// 				break
// 			}
// 		}
// 		syncID, err := getSyncID(jobID)
// 		if err != nil {
// 			err = ws.WriteMessage(mt, message)
// 			if err != nil {
// 				log.Println("write:", err)
// 				break
// 			}
// 		}

// 		err = ws.WriteMessage(mt, message)
// 		if err != nil {
// 			log.Println("write:", err)
// 			break
// 		}
// 	}
// }
