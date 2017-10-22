package main

import (
	"log"
	"net/http"
	"os"

	"stringsvc/service"
	"stringsvc/transport"

	kitLog "github.com/go-kit/kit/log"
)

func main() {
	mainLogger := kitLog.NewLogfmtLogger(os.Stderr)
	svc := service.MyStringService{}
	loggedSvc := service.LoggedService{mainLogger, svc}

	uppercaseHandler := transport.NewUppercaseHandler(loggedSvc)
	countHandler := transport.NewCountHandler(loggedSvc)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
