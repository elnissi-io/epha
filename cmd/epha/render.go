package main

import (
	"epha/pkg/listener"
	"epha/pkg/parser"
	"epha/pkg/stdlib"
	"fmt"
	"io/ioutil"

	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/cobra"
)

func renderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "render [file]",
		Short: "Render a Kubernetes resource file",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]

			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			input := antlr.NewInputStream(string(content))
			lexer := parser.NewEphaLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewEphaParser(stream)

			crds, err := stdlib.LoadCRDs(".") // Ensure this path is correct
			if err != nil {
				fmt.Println("Error loading CRDs:", err)
				return
			}

			ctx := stdlib.NewImportContext()
			listener := listener.NewEphaListener(crds, ctx)
			antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

			program := listener.GetProgram()
			fmt.Println("Parsed program:", program)
		},
	}
	return cmd
}