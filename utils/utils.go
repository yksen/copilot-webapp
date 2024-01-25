package utils

import (
	"bytes"
	"fmt"
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
