package control_api

import (
	"backend/models/res"
	"backend/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	devices "github.com/kubeedge/kubeedge/pkg/apis/devices/v1beta1"
	"k8s.io/client-go/rest"
)

// DeviceStatus is used to patch device status
type DeviceStatus struct {
	Status devices.DeviceStatus `json:"status"`
}

// The device name
var deviceName = "modbusplc"

// The default namespace in which the counter device instance resides
var namespace = "default"

// The default status of the counter
var originCmd = "OFF"

// The CRD client used to patch the device instance.
var crdClient *rest.RESTClient

// The twin value map
var statusMap = map[string]string{
	"ON":  "true",
	"OFF": "false",
}

type Controller struct {
	Temperature string `json:"temperature"`
	Humidity    string `json:"humidity"`
	Light       string `json:"light"`
	//Actuator    string `json:"actuator"`
}

// init()是一个默认调用的函数，在程序运行之前就已经执行了该函数
// 打包镜像的时候，记得改为init
func Init() {
	// Create a client to talk to the K8S API server to patch the device CRDs
	kubeConfig, err := utils.KubeConfig()
	if err != nil {
		log.Fatalf("[function init]: Failed to create KubeConfig, error : %v", err)
	}
	log.Println("[function init]: Get kubeConfig successfully")

	crdClient, err = utils.NewCRDClient(kubeConfig)
	if err != nil {
		log.Fatalf("[function init]: Failed to create device crd client , error : %v", err)
	}
	log.Println("[function init]: Get crdClient successfully")
}

// 生成desired value
func UpdateStatus() map[string]string {
	result := DeviceStatus{}
	raw, err := crdClient.Get().Namespace(namespace).Resource(utils.ResourceTypeDevices).Name(deviceName).DoRaw(context.TODO())
	if err != nil {
		log.Println("[function UpdateStatus]: Get err = ", err)
	}
	err = json.Unmarshal(raw, &result)
	if err != nil {
		log.Println("[function UpdateStatus]: Unmarshal err = ", err)
	}
	//get data from GET
	fmt.Printf("[function UpdateStatus]: raw data is %s\n", raw)

	status := make(map[string]string)

	log.Print("[function UpdateStatus]: ")
	for i, twin := range result.Status.Twins {
		status[twin.PropertyName] = twin.Reported.Value
		log.Printf("(%d)reported_value = %s, desired_value = %s\t", i, twin.Reported.Value, twin.ObservedDesired.Value)
		// status["status"] = twin.Desired.Value
		// status["value"] = twin.Reported.Value
	}
	return status
}

// UpdateDeviceTwinWithDesiredTrack patches the desired state of
// the device twin with the command.
func UpdateDeviceTwinWithDesiredTrack(cmd string, temp string, humi string) bool {
	if cmd == originCmd {
		return true
	}

	status := buildStatusWithDesiredTrack(cmd, temp, humi)
	deviceStatus := &DeviceStatus{Status: status}
	body, err := json.Marshal(deviceStatus)
	//body是一个 []byte 类型，不打印
	//log.Printf("[function UpdateDeviceTwinWithDesiredTrack]: body =  %v", body)
	log.Printf("[function UpdateDeviceTwinWithDesiredTrack]: deviceStatus =  %v", deviceStatus)

	if err != nil {
		log.Printf("[function UpdateDeviceTwinWithDesiredTrack]: Failed to marshal device status %v", deviceStatus)
		return false
	}
	result := crdClient.Patch(utils.MergePatchType).Namespace(namespace).Resource(utils.ResourceTypeDevices).Name(deviceName).Body(body).Do(context.TODO())
	if result.Error() != nil {
		log.Printf("function UpdateDeviceTwinWithDesiredTrack]: Failed to patch device status %v of device %v in namespace %v \n error:%+v", deviceStatus, deviceName, namespace, result.Error())
		return false
	} else {
		log.Printf("[funtion UpdateDeviceTwinWithDesiredTrack]: Light = %s, Temp = %s, Humi = %s", cmd, temp, humi)
	}
	originCmd = cmd

	return true
}

func buildStatusWithDesiredTrack(cmd string, temp string, humi string) devices.DeviceStatus {
	metadata := map[string]string{
		"timestamp": strconv.FormatInt(time.Now().Unix()/1e6, 10),
		"type":      "string",
	}
	twins := []devices.Twin{
		{PropertyName: "temperature",
			ObservedDesired: devices.TwinProperty{Value: temp, Metadata: metadata},
			Reported:        devices.TwinProperty{Value: "reported-temp", Metadata: metadata}},
		{PropertyName: "humidity",
			ObservedDesired: devices.TwinProperty{Value: humi, Metadata: metadata},
			Reported:        devices.TwinProperty{Value: "reported-humi", Metadata: metadata}},
		{PropertyName: "light",
			//反了？
			ObservedDesired: devices.TwinProperty{Value: cmd, Metadata: metadata},
			Reported:        devices.TwinProperty{Value: statusMap[cmd], Metadata: metadata}},
	}
	devicestatus := devices.DeviceStatus{Twins: twins}
	return devicestatus
}

// 实现控制下行流
func (ControlApi) ControlALL(c *gin.Context) {
	var controller Controller
	if err := c.ShouldBind(&controller); err != nil {
		res.Error("绑定json失败", c)
		return
	}
	fmt.Println("[function ControlALL]: light = ", controller.Light, "temperature = ", controller.Temperature, "humidity = ", controller.Humidity)
	//fmt.Println("Lighgt = ", controller.Light, "Actuator = ", controller.)
	//res.OKWithData(controller, c)
	status := UpdateStatus()
	fmt.Println("[function ControlALL]: status = ", status)
	if controller.Light == "OFF" {
		UpdateDeviceTwinWithDesiredTrack("false", controller.Temperature, controller.Humidity)
		res.OK(status, "开关成功关闭...", c)
	} else if controller.Light == "ON" {
		UpdateDeviceTwinWithDesiredTrack("true", controller.Temperature, controller.Humidity)
		res.OK(status, "开关成功打开...", c)
	}
}
