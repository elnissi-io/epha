package ast

import "bytes"

type Resource struct {
	Name       string
	Type       string
	Properties []*Property
}

// Implement the Statement interface
func (r *Resource) statementNode()       {}
func (r *Resource) TokenLiteral() string { return "" }
func (r *Resource) String() string {
	var out bytes.Buffer
	out.WriteString(r.Type + " " + r.Name + " {\n")
	for _, p := range r.Properties {
		out.WriteString(p.String())
	}
	out.WriteString("}\n")
	return out.String()
}

type Property struct {
	Name     string
	Value    Value
	Values   []Value
	SubProps []*Property
}

func (p *Property) String() string {
	var out bytes.Buffer
	if len(p.Values) > 0 {
		out.WriteString(p.Name + " = [")
		for _, v := range p.Values {
			out.WriteString(v.String() + ", ")
		}
		out.WriteString("]\n")
	} else if len(p.SubProps) > 0 {
		out.WriteString(p.Name + " {\n")
		for _, sp := range p.SubProps {
			out.WriteString(sp.String())
		}
		out.WriteString("}\n")
	} else {
		out.WriteString(p.Name + " = " + p.Value.String() + "\n")
	}
	return out.String()
}

type Value struct {
	Type  ValueType
	Value string
}

func (v *Value) String() string {
	return v.Value
}

type ValueType string

const (
	StringType ValueType = "string"
	NumberType ValueType = "number"
)
