package kubernetes

import "epha/pkg/stdlib"

func init() {
	stdlib.RegisterModule("Deployment", func() stdlib.Module {
		return NewDeployment()
	})
	stdlib.RegisterModule("Service", func() stdlib.Module {
		return NewService()
	})
}

type Deployment struct {
	Name       string
	Replicas   int
	Containers []Container
}

type Container struct {
	Name  string
	Image string
	Ports []int
}

func NewDeployment() *Deployment {
	return &Deployment{}
}

func (d *Deployment) Initialize() error {
	// Initialize deployment-related settings if needed.
	return nil
}

func (d *Deployment) Name() string {
	return "Deployment"
}

type Service struct {
	Name  string
	Type  string
	Ports []ServicePort
}

type ServicePort struct {
	Port       int
	TargetPort int
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Initialize() error {
	// Initialize service-related settings if needed.
	return nil
}

func (s *Service) Name() string {
	return "Service"
}
