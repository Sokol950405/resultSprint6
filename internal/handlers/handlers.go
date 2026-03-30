package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler - handler from /
// return HTML from file index.html
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// open file index.html
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "Failed to open index.html", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// return file HTML
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UploadHandler - handler from /upload
// return parse morse to text and text to morse
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// parse form (max size 10 MB)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	// receive file from the form
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// reading data from a file
	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file data", http.StatusInternalServerError)
		return
	}

	// convert data
	convertedString, err := service.AutoDetect(fileData)
	if err != nil {
		http.Error(w, "Failed to convert data", http.StatusInternalServerError)
		return
	}

	// create new fule
	fileName := time.Now().UTC().Format("2006-01-02_15-04-05")
	extension := filepath.Ext(header.Filename)
	outputFileName := fileName + extension
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		http.Error(w, "Failed to create output file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// write the file
	_, err = outputFile.WriteString(convertedString)
	if err != nil {
		http.Error(w, "Failed to write to output file", http.StatusInternalServerError)
		return
	}

	// return result
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, convertedString)
}
