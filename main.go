package main

import (
	"net/http"

	"github.com/we11done/simple-api-go/app"
)

func main() {
	http.ListenAndServe(":3000", app.NewHandler())
}
