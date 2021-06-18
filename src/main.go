package main

import (
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	"os"
	_ "post-service/gateway"
	router2 "post-service/http/router"
	"post-service/infrastructure/cassandra_config"
	"post-service/interactor"
	_ "post-service/usecase"
)


func main() {
	logger := logger.InitializeLogger("post-service", context.Background())


	cassandraSession, _ := cassandra_config.NewCassandraSession(logger)

	i := interactor.NewInteractor(cassandraSession, logger)

	handler := i.NewAppHandler()

	router := router2.NewRouter(handler, logger)
	if os.Getenv("DOCKER_ENV") == "" {
		err := router.RunTLS(":8083", "certificate/cert.pem", "certificate/key.pem")
		if err != nil {
			return 
		}
	} else {
		err := router.Run(":8083")
		if err != nil {
			return 
		}
	}

}
