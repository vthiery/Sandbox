package main

import (
	"log"
	"net/http"

	"stringsvc/service"
	"stringsvc/transport"
)

func main() {
	svc := service.MyStringService{}

	uppercaseHandler := transport.NewUppercaseHandler(svc)
	countHandler := transport.NewCountHandler(svc)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
