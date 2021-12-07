package mockagent

import (
	"github.com/c3sr/dlframework"
)

var FrameworkManifest = dlframework.FrameworkManifest{
	Name:    "mock",
	Version: "1.2.3",
	Container: map[string]*dlframework.ContainerHardware{
		"amd64": {
			Cpu: "c3sr/carml-mxnet:amd64-cpu",
			Gpu: "c3sr/carml-mxnet:amd64-gpu",
		},
		"ppc64le": {
			Cpu: "c3sr/carml-mxnet:ppc64le-cpu",
			Gpu: "c3sr/carml-mxnet:ppc64le-gpu",
		},
	},
}