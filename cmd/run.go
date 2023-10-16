package main

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/chaowen112/gin-template/app/api"
)

//go:generate swag init -g ./app/api/router.go -o ./docs --parseDependency

func init() {
}

func main() {
	port := 8080
	var err error
	if os.Getenv("PORT") != "" {
		port, err = strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatalf("unable to use port %s", port)
		}
	}
	log.WithField("port", port).
		Info("start to run gin server")
	api.Run()
}
