package utils

import (
	"log"

	//devices "github.com/kubeedge/kubeedge/tree/master/pkg/apis/devices/v1alpha2"

	//devices "github.com/kubeedge/kubeedge/pkg/apis/devices/v1alpha2"
	devices "github.com/kubeedge/kubeedge/pkg/apis/devices/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// constants.go
const (
	// MergePatchType is patch type
	MergePatchType = "application/merge-patch+json"
	// ResourceTypeDevices is plural of device resource in apiserver
	ResourceTypeDevices = "devices"
)

// crdclient.go
// NewCRDClient is used to create a restClient for crd
func NewCRDClient(cfg *rest.Config) (*rest.RESTClient, error) {
	scheme := runtime.NewScheme()
	schemeBuilder := runtime.NewSchemeBuilder(addDeviceCrds)

	err := schemeBuilder.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}

	config := *cfg
	//APIPath=/apis, GroupVersion=devices.kubeedge.io/v1alpha2, ContentType=application/json
	config.APIPath = "/apis"
	config.GroupVersion = &devices.SchemeGroupVersion
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.NewCodecFactory(scheme)

	//test
	log.Printf("[function NewCRDClient]: APIPath=%s, GroupVersion=%s, ContentType=%s, NegotiatedSerializer=%s",
		config.APIPath, config.GroupVersion, config.ContentType, config.NegotiatedSerializer)

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		log.Fatalf("Failed to create REST Client due to error %v", err)
		return nil, err
	}

	return client, nil
}

func addDeviceCrds(scheme *runtime.Scheme) error {
	// Add Device
	scheme.AddKnownTypes(devices.SchemeGroupVersion, &devices.Device{}, &devices.DeviceList{})
	v1.AddToGroupVersion(scheme, devices.SchemeGroupVersion)
	// Add DeviceModel
	scheme.AddKnownTypes(devices.SchemeGroupVersion, &devices.DeviceModel{}, &devices.DeviceModelList{})
	v1.AddToGroupVersion(scheme, devices.SchemeGroupVersion)

	return nil
}

// kubeclient.go
var KubeMaster = ""
var Kubeconfig = ""
var KubeQPS = float32(5.000000)
var KubeBurst = 10
var KubeContentType = "application/vnd.kubernetes.protobuf"

// KubeConfig from flags
func KubeConfig() (conf *rest.Config, err error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags(KubeMaster, Kubeconfig)
	if err != nil {
		return nil, err
	}
	kubeConfig.QPS = KubeQPS
	kubeConfig.Burst = KubeBurst
	kubeConfig.ContentType = KubeContentType
	return kubeConfig, err
}
