package test

import (
	"epha/pkg/dsl/parser"
	"epha/pkg/generator"
	"epha/pkg/listener"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestParser(t *testing.T) {
	input := `
import my_module from other_module import func1, func2

myVar = "Hello, World!"
helm my_chart {
  version: "1.0"
  repo: "stable"
}

k8s my_resource {
  kind: "Pod"
  metadata: {
    name: "my-pod"
    namespace: "default"
  }
}
`
	is := antlr.NewInputStream(input)
	lexer := parser.NewEphaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewEphaParser(stream)
	listener := listener.NewEphaListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

	program := listener.GetProgram()
	yaml, err := generator.GenerateYAML(program)
	if err != nil {
		t.Fatalf("Failed to generate YAML: %v", err)
	}

	expectedOutput := `...` // Define the expected YAML output
	if yaml != expectedOutput {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expectedOutput, yaml)
	}
}
