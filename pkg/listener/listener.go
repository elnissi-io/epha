package listener

import (
	"epha/pkg/dsl/ast"
	"epha/pkg/parser"
	"epha/pkg/stdlib"
	"strings"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type EphaListener struct {
	*parser.BaseEphaListener
	program *ast.Program
	crds    []*v1.CustomResourceDefinition
	ctx     *stdlib.ImportContext
}

func NewEphaListener(crds []*v1.CustomResourceDefinition, ctx *stdlib.ImportContext) *EphaListener {
	return &EphaListener{
		program: &ast.Program{},
		crds:    crds,
		ctx:     ctx,
	}
}

func (l *EphaListener) GetProgram() *ast.Program {
	return l.program
}

func (l *EphaListener) ExitImportStatement(ctx *parser.ImportStatementContext) {
	moduleName := ctx.IDENTIFIER(0).GetText()
	alias := moduleName
	if len(ctx.AllIDENTIFIER()) > 1 {
		alias = ctx.IDENTIFIER(1).GetText()
	}
	if err := l.ctx.ImportModule(moduleName, alias); err != nil {
		panic(err)
	}
}

func (l *EphaListener) ExitResourceDefinition(ctx *parser.ResourceDefinitionContext) {
	resourceName := ctx.IDENTIFIER(0).GetText()
	resourceType := ctx.IDENTIFIER(1).GetText()

	resourceTypeObj, err := stdlib.GetResourceType(resourceType)
	if err != nil {
		panic(err)
	}

	resource := &ast.Resource{
		Name: resourceName,
		Type: resourceTypeObj.Name(),
	}

	for _, child := range ctx.GetChildren() {
		if propCtx, ok := child.(*parser.ResourcePropertyContext); ok {
			prop := l.buildResourceProperty(propCtx)
			resource.Properties = append(resource.Properties, prop)
		}
	}

	l.program.Statements = append(l.program.Statements, resource)
}

func (l *EphaListener) buildResourceProperty(ctx *parser.ResourcePropertyContext) *ast.Property {
	propName := l.buildPropertyKey(ctx.PropertyKey())
	if ctx.Value() != nil {
		return &ast.Property{
			Name:  propName,
			Value: l.buildValue(ctx.Value().(*parser.ValueContext)),
		}
	} else if ctx.ValueList() != nil {
		var values []ast.Value
		for _, valCtx := range ctx.ValueList().AllValue() {
			values = append(values, l.buildValue(valCtx.(*parser.ValueContext)))
		}
		return &ast.Property{
			Name:   propName,
			Values: values,
		}
	} else if ctx.ResourcePropertyBody() != nil {
		var subProps []*ast.Property
		for _, subPropCtx := range ctx.ResourcePropertyBody().AllResourceProperty() {
			subProps = append(subProps, l.buildResourceProperty(subPropCtx.(*parser.ResourcePropertyContext)))
		}
		return &ast.Property{
			Name:     propName,
			SubProps: subProps,
		}
	}
	return nil
}

func (l *EphaListener) buildPropertyKey(ctx parser.IPropertyKeyContext) string {
	var parts []string
	for _, id := range ctx.AllIDENTIFIER() {
		parts = append(parts, id.GetText())
	}
	return strings.Join(parts, ".")
}

func (l *EphaListener) buildValue(ctx *parser.ValueContext) ast.Value {
	if ctx.STRING() != nil {
		return ast.Value{
			Type:  ast.StringType,
			Value: strings.Trim(ctx.STRING().GetText(), `"`), // Remove quotes from string
		}
	} else if ctx.NUMBER() != nil {
		return ast.Value{
			Type:  ast.NumberType,
			Value: ctx.NUMBER().GetText(),
		}
	}
	return ast.Value{}
}
