package main

import (
	"encoding/json"
	"net/http"
	"os"
	"split-pdf/cmd/api/v1/pdf"
	file "split-pdf/internal/files"
	"split-pdf/internal/health"
	"split-pdf/internal/info"
	"split-pdf/pkg/server"
	"time"
)

var Version = "development"
var Build = time.Now().String()

func statusHandler(healthService health.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.MarshalIndent(healthService.Health(), "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}

func Routes() {
	healthService := health.New()

	http.HandleFunc("/_service/health", statusHandler(healthService))

	// Load info Version
	si := info.New(Version, Build)

	// creatt handlers

	infolHandler := http.HandlerFunc(si.InfoHandler())
	http.Handle("/_service/info", server.CreateRequestContext(infolHandler))

	// Map standar routes
	http.HandleFunc("/upload", pdf.UploadFile)
	// http.HandleFunc("/_service/info", si.InfoHandler())
	http.HandleFunc("/", file.HelloWorld)

	// Map static files
	fs := http.FileServer(http.Dir(os.Getenv("STATIC_PATH")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

}

func main() {

	// u.Log.Debug(nil, "This is a debug message")
	// u.Log.Info(nil, "This is a info message")
	// u.Log.Warn(nil, "This is a warn message")
	// u.Log.Error(nil, "This is a error message")

	Routes()
	http.ListenAndServe(":8080", nil)
}
