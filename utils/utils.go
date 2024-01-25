package utils

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func Check(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CheckPanic(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

func GetRequestBody(r *http.Request) []byte {
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
	}
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes
}

//go:embed templates/*.html
var embedFiles embed.FS

func Templates() (*template.Template, error) {
	templates, err := template.ParseFS(embedFiles, "templates/*.html")
	if err != nil {
		return nil, err
	}
	return templates, nil
}
