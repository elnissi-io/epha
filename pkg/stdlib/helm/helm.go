// epha/pkg/stdlib/helm/helm.go
package helm

import (
	"epha/pkg/stdlib"
)

func init() {
	stdlib.RegisterModule("Chart", func() stdlib.Module {
		return NewChart()
	})
}

// Chart represents a Helm chart
type Chart struct {
	Name    string
	Version string
	// Define other attributes as needed
}

// NewChart creates a new Chart instance
func NewChart() *Chart {
	return &Chart{}
}

// Initialize initializes the Chart instance
func (c *Chart) Initialize() error {
	// Initialize chart-related settings if needed.
	return nil
}

// Name returns the name of the Chart module
func (c *Chart) Name() string {
	return "Chart"
}
