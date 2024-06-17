package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"split-pdf/internal/health"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

var Version = "latest"
var Build = time.Now().String()

// Funcion para la subida de contenido via API
func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	// tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	tempFile, err := os.CreateTemp(os.Getenv("TEMP_PATH"), "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	// Dividir PDF
	conf := model.NewDefaultConfiguration()
	conf.ValidationMode = model.ValidationRelaxed
	destPath := os.Getenv("TEMP_PATH")
	if destPath == "" {
		fmt.Println("TEMP_PATH is not set")
		return
	}
	f, err := os.Open(tempFile.Name())

	if err != nil {
		fmt.Println(err)
		return
	}

	err = api.Split(f, destPath, tempFile.Name(), 1, conf)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

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

func setupRoutes() {
	// Map standar routes
	http.HandleFunc("/upload", uploadFile)
	healthService := health.New()
	http.HandleFunc("/_service/health", statusHandler(healthService))
	http.HandleFunc("/", helloWorld)

	// Map static files
	fs := http.FileServer(http.Dir(os.Getenv("STATIC_PATH")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Healcheck
	// http.Handle("/dummy", health)

	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Version: ", Version)
	fmt.Println("Build Time: ", Build)
	setupRoutes()
}
