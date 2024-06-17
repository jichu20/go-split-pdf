package info

import (
	"net/http"
)

type ServiceInfo struct {
	Version string `json:"version"`
	Build   string `json:"build"`
}

func New(version string, build string) ServiceInfo {

	return ServiceInfo{
		Version: version,
		Build:   build,
	}
}

func (si *ServiceInfo) GetInfo() []byte {
	return []byte(`{"version": "` + si.Version + `", "build": "` + si.Build + `"}`)
}

func (si *ServiceInfo) InfoHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(si.GetInfo())
	}
}
