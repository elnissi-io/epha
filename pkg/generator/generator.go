package generator

import (
	"bytes"
	"epha/pkg/dsl/ast"
	"fmt"
	"strings"
)

// GenerateYAML transforms an entire AST (represented by ast.Program) into a YAML string.
func GenerateYAML(program *ast.Program) (string, error) {
	var buf bytes.Buffer

	for _, stmt := range program.Statements {
		switch s := stmt.(type) {
		case *ast.ImportStatement:
			// Handle imports if needed
		case *ast.AssignmentStatement:
			yaml, err := generateYAMLForAssignment(s)
			if err != nil {
				return "", err
			}
			buf.WriteString(yaml + "\n\n")
		case *ast.VariableDeclaration:
			yaml, err := generateYAMLForVariableDeclaration(s)
			if err != nil {
				return "", err
			}
			buf.WriteString(yaml + "\n\n")
		}
	}

	return buf.String(), nil
}

func generateYAMLForAssignment(as *ast.AssignmentStatement) (string, error) {
	var buf bytes.Buffer

	buf.WriteString(as.Name.Value + ":\n")
	yamlValue, err := generateYAMLForExpression(as.Value)
	if err != nil {
		return "", err
	}
	buf.WriteString(indent(yamlValue, 2))

	return buf.String(), nil
}

func generateYAMLForVariableDeclaration(vd *ast.VariableDeclaration) (string, error) {
	var buf bytes.Buffer

	buf.WriteString(vd.Name.Value + ":\n")
	yamlValue, err := generateYAMLForExpression(vd.Value)
	if err != nil {
		return "", err
	}
	buf.WriteString(indent(yamlValue, 2))

	return buf.String(), nil
}

func generateYAMLForExpression(expr ast.Expression) (string, error) {
	switch e := expr.(type) {
	case *ast.StringLiteral:
		return fmt.Sprintf("\"%s\"", e.Value), nil
	case *ast.IntegerLiteral:
		return fmt.Sprintf("%d", e.Value), nil
	case *ast.BooleanLiteral:
		return fmt.Sprintf("%t", e.Value), nil
	case *ast.HashLiteral:
		var pairs []string
		for key, value := range e.Pairs {
			keyYAML, err := generateYAMLForExpression(key)
			if err != nil {
				return "", err
			}
			valueYAML, err := generateYAMLForExpression(value)
			if err != nil {
				return "", err
			}
			pairs = append(pairs, fmt.Sprintf("%s: %s", keyYAML, valueYAML))
		}
		return "{\n" + indent(strings.Join(pairs, "\n"), 2) + "\n}", nil
	case *ast.ArrayLiteral:
		var elements []string
		for _, elem := range e.Elements {
			elemYAML, err := generateYAMLForExpression(elem)
			if err != nil {
				return "", err
			}
			elements = append(elements, "- "+elemYAML)
		}
		return strings.Join(elements, "\n"), nil
	default:
		return "", fmt.Errorf("unsupported expression type in YAML generation")
	}
}

func indent(s string, spaces int) string {
	lines := strings.Split(s, "\n")
	indentedLines := make([]string, len(lines))
	indentation := strings.Repeat(" ", spaces)
	for i, line := range lines {
		indentedLines[i] = indentation + line
	}
	return strings.Join(indentedLines, "\n")
}
