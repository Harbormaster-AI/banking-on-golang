package controller

import (
	"net/http"
	"fmt"
)

func Default__(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "demo app is ok")
}


func Health__(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "demo app is healthy")
}



