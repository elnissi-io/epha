package stdlib

import (
	"epha/pkg/stdlib/kubernetes"
	"testing"
)

func TestImportModule(t *testing.T) {
	ctx := NewImportContext()
	err := ctx.ImportModule("Deployment", "Deployment")
	if err != nil {
		t.Fatalf("Failed to import module: %s", err)
	}
	if _, exists := ctx.imports["Deployment"]; !exists {
		t.Fatalf("Deployment not imported correctly")
	}
	if _, ok := ctx.imports["Deployment"].(*kubernetes.Deployment); !ok {
		t.Fatalf("Imported module is not of type *Deployment")
	}
}
