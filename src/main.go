package main

import (
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	"os"
	_ "post-service/gateway"
	router2 "post-service/http/router"
	"post-service/infrastructure/cassandra_config"
	"post-service/infrastructure/grpc/client"
	"post-service/interactor"
	_ "post-service/usecase"
)


func main() {
	logger := logger.InitializeLogger("post-service", context.Background())

	notificationClient, err := client.NewNotificationClient("127.0.0.1:8078")

	if err != nil {
		panic(err)
	}

	cassandraSession, _ := cassandra_config.NewCassandraSession(logger)

	i := interactor.NewInteractor(cassandraSession, logger, notificationClient)

	handler := i.NewAppHandler()






	router := router2.NewRouter(handler, logger)
	if os.Getenv("DOCKER_ENV") == "" {
		err := router.RunTLS(":8083", "src/certificate/cert.pem", "src/certificate/key.pem")
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
