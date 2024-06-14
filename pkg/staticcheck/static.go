package staticcheck

import (
	"epha/pkg/dsl/ast"
	"fmt"
)

// ValidateCRD checks if a CRD is correctly defined.
func ValidateCRD(crd *ast.CallExpression) error {
	// Perform validation logic for CRDs
	// For example, ensuring required fields are present
	if crd.Function.Value != "CustomResourceDefinition" {
		return fmt.Errorf("invalid CRD: %v", crd.Function.Value)
	}

	// Check required fields
	// Assuming the AST node structure has fields like 'metadata' and 'spec'
	for _, arg := range crd.Arguments {
		switch a := arg.(type) {
		case *ast.HashLiteral:
			if _, ok := a.Pairs["metadata"]; !ok {
				return fmt.Errorf("CRD is missing 'metadata' field")
			}
			if _, ok := a.Pairs["spec"]; !ok {
				return fmt.Errorf("CRD is missing 'spec' field")
			}
		default:
			return fmt.Errorf("unexpected argument type for CRD: %T", arg)
		}
	}

	return nil
}
