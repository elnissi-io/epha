package validator

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ValidateManifest(manifest string) error {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return fmt.Errorf("failed to build Kubernetes config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %v", err)
	}

	obj := &unstructured.Unstructured{}
	_, gvk, err := unstructured.UnstructuredJSONScheme.Decode([]byte(manifest), nil, obj)
	if err != nil {
		return fmt.Errorf("failed to decode manifest: %v", err)
	}

	gvr := schema.GroupVersionResource{
		Group:    gvk.Group,
		Version:  gvk.Version,
		Resource: gvk.Kind, // Assuming Kind maps directly to resource; adjust as necessary
	}

	// Perform validation logic here...
	// For example, checking if the resource exists
	_, err = clientset.Dynamic().Resource(gvr).Namespace(obj.GetNamespace()).Get(obj.GetName(), metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to validate resource: %v", err)
	}

	return nil
}
func ValidateCRD(manifest string) error {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return fmt.Errorf("failed to build Kubernetes config: %v", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create dynamic client: %v", err)
	}

	obj := &unstructured.Unstructured{}
	_, gvk, err := unstructured.UnstructuredJSONScheme.Decode([]byte(manifest), nil, obj)
	if err != nil {
		return fmt.Errorf("failed to decode manifest: %v", err)
	}

	gvr := schema.GroupVersionResource{
		Group:    gvk.Group,
		Version:  gvk.Version,
		Resource: gvk.Kind, // You may need to map this appropriately
	}

	_, err = dynamicClient.Resource(gvr).Namespace(obj.GetNamespace()).Create(context.TODO(), obj, v1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to validate CRD: %v", err)
	}

	return nil
}