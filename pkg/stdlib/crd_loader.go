// pkg/stdlib/crd_loader.go
package stdlib

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

// LoadCRDs loads CRDs from the specified directory.
func LoadCRDs(crdDir string) ([]*v1.CustomResourceDefinition, error) {
	var crds []*v1.CustomResourceDefinition
	files, err := ioutil.ReadDir(crdDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".yaml" {
			crd, err := loadCRD(filepath.Join(crdDir, file.Name()))
			if err != nil {
				return nil, err
			}
			crds = append(crds, crd)
		}
	}
	return crds, nil
}

func loadCRD(path string) (*v1.CustomResourceDefinition, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	obj, _, err := scheme.Codecs.UniversalDeserializer().Decode(data, nil, nil)
	if err != nil {
		return nil, err
	}

	crd, ok := obj.(*v1.CustomResourceDefinition)
	if !ok {
		return nil, fmt.Errorf("file %s is not a valid CRD", path)
	}

	return crd, nil
}
