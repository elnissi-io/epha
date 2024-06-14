package stdlib

import (
	"fmt"
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)


func registerResource(name string, objType reflect.Type) {
	regLock.Lock()
	defer regLock.Unlock()
	if _, exists := resourceRegistry[name]; exists {
		panic(fmt.Sprintf("resource already registered: %s", name))
	}
	resourceRegistry[name] = objType
}

func init() {
	registerResource("Deployment", reflect.TypeOf(appsv1.Deployment{}))
	registerResource("Pod", reflect.TypeOf(corev1.Pod{}))
	registerResource("ConfigMap", reflect.TypeOf(corev1.ConfigMap{}))
	registerResource("Secret", reflect.TypeOf(corev1.Secret{}))
	// Add more resources as needed
}
