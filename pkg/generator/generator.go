package generator

import (
	"bytes"
	"epha/pkg/dsl/ast"
	"fmt"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/yaml"
)

type Generator struct {
	program *ast.Program
	crds    []*v1.CustomResourceDefinition
}

func NewGenerator(program *ast.Program, crds []*v1.CustomResourceDefinition) *Generator {
	return &Generator{program: program, crds: crds}
}

func (g *Generator) Generate() (string, error) {
	var output bytes.Buffer

	for _, stmt := range g.program.Statements {
		resource := stmt.(*ast.Resource)
		yamlString, err := g.generateResource(resource)
		if err != nil {
			return "", err
		}
		output.WriteString(yamlString + "\n---\n")
	}

	return output.String(), nil
}

func (g *Generator) generateResource(resource *ast.Resource) (string, error) {
	crd := g.getCRDForResource(resource.Type)
	if crd == nil {
		return "", fmt.Errorf("CRD not found for resource: %s", resource.Type)
	}

	k8sResource := map[string]interface{}{
		"apiVersion": crd.Spec.Group + "/" + crd.Spec.Versions[0].Name,
		"kind":       crd.Spec.Names.Kind,
		"metadata": map[string]interface{}{
			"name": resource.Name,
		},
		"spec": g.buildSpec(resource.Properties),
	}

	yamlData, err := yaml.Marshal(k8sResource)
	if err != nil {
		return "", err
	}

	return string(yamlData), nil
}

func (g *Generator) buildSpec(properties []*ast.Property) map[string]interface{} {
	spec := make(map[string]interface{})
	for _, prop := range properties {
		if len(prop.Values) > 0 {
			var values []interface{}
			for _, val := range prop.Values {
				values = append(values, val.Value)
			}
			spec[prop.Name] = values
		} else if len(prop.SubProps) > 0 {
			spec[prop.Name] = g.buildSpec(prop.SubProps)
		} else {
			spec[prop.Name] = prop.Value.Value
		}
	}
	return spec
}

func (g *Generator) getCRDForResource(name string) *v1.CustomResourceDefinition {
	for _, crd := range g.crds {
		if crd.Spec.Names.Kind == name {
			return crd
		}
	}
	return nil
}
