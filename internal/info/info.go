package info

import (
	"net/http"
	u "split-pdf/internal/utils"
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
		context := r.Context()
		// dummy := middleware.FromContext(context)

		u.Log.Debug(context, "Mensaje de prueba ")
		// fmt.Println("Dummy: ", dummy)
		w.Header().Set("Content-Type", "application/json")
		w.Write(si.GetInfo())
	}
}
