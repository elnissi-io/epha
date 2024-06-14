package main

import (
	"epha/pkg/generator"
	"epha/pkg/listener"
	"epha/pkg/parser"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/cobra"
)

func renderCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "render [file]",
		Short: "Render the defined objects as YAML",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			content, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				os.Exit(1)
			}

			input := antlr.NewInputStream(string(content))
			lexer := parser.NewEphaLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
			p := parser.NewEphaParser(stream)

			listener := listener.NewEphaListener()
			antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

			program := listener.GetProgram()

			yaml, err := generator.GenerateYAML(program)
			if err != nil {
				fmt.Printf("Error generating YAML: %v\n", err)
				os.Exit(1)
			}

			fmt.Println(yaml)
		},
	}
}
