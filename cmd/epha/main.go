package main

import (
	"epha/pkg/generator"
	"epha/pkg/listener"
	"epha/pkg/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input := antlr.NewInputStream(`import k8s from "kubernetes" import deployment, service`)
	lexer := parser.NewEphaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewEphaParser(stream)

	listener := listener.NewEphaListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

	program := listener.GetProgram()

	yaml, err := generator.GenerateYAML(program)
	if err != nil {
		fmt.Println("Error generating YAML:", err)
	} else {
		fmt.Println("Generated YAML:\n", yaml)
	}
}
