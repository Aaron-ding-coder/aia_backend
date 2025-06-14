package handler

import (
	"fmt"
	"github.com/alioygur/gores"
	"net/http"
)

func PingHandler(rw http.ResponseWriter, _ *http.Request) {
	_ = gores.JSON(rw, http.StatusOK, fmt.Sprintf("Welcome to aia program"))
}

func Add() {

}
