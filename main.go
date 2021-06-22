package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hugobcar/http_requests_fake/prometheus"
	"github.com/hugobcar/http_requests_fake/router"
)

var v int

func main() {
	v = 15
	// router.Value = &i

	go prometheus.Run()
	prometheus.CreateGauges()

	go runEndpoint()

	for {
		run(v)

		time.Sleep(time.Duration(1) * time.Second)

		v = v + 15
	}
}

func run(va int) {
	prometheus.RegistredGauges["http_server_requests_seconds_count"].Total.Set(float64(va))
}

func runEndpoint() {
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":6885", router))
}
