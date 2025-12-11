package main

import (
	"flag"
	"fmt"
	"net/http"
)

type config struct {
	port    int
	env     string
	version string
}

type application struct {
	config config
}

func main() {
	var config config

	flag.IntVar(&config.port, "port", 4000, "API server port")
	flag.StringVar(&config.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&config.version, "version", "1.0.0", "API version")
	flag.Parse()

	app := &application{
		config: config,
	}

	fmt.Printf("Start %s API on :%d\n", config.env, config.port)

	address := fmt.Sprintf(":%d", config.port)

	err := http.ListenAndServe(address, app.routes())
	if err != nil {
		fmt.Println("Error Starting server: ", err)
	}
}
