package e2e

import (
	"io/ioutil"
	"testing"

	"epha/pkg/dsl"
	"epha/pkg/generator"
	"epha/pkg/stdlib"
)

func TestCompleteWorkflow(t *testing.T) {
	ctx := stdlib.NewImportContext()
	err := ctx.ImportModule("Deployment", "Deployment")
	if err != nil {
		t.Fatal(err)
	}

	// Assuming the test.epha file is located directly under the test/e2e/resources
	input, err := ioutil.ReadFile("test/e2e/resources/test.epha")
	if err != nil {
		t.Fatal(err)
	}

	ast, err := dsl.Parse(string(input))
	if err != nil {
		t.Fatal(err)
	}

	yaml, err := generator.GenerateYAML(ast)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Generated YAML:", yaml)
}
