package main

import (
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	router2 "post-service/http/router"
	"post-service/infrastructure/cassandra_config"
	"post-service/interactor"
)


func main() {
	logger := logger.InitializeLogger("post-service", context.Background())


	cassandraSession, _ := cassandra_config.NewCassandraSession(logger)

	i := interactor.NewInteractor(cassandraSession, logger)

	handler := i.NewAppHandler()

	router := router2.NewRouter(handler, logger)

	router.RunTLS(":8083", "certificate/cert.pem", "certificate/key.pem")

}
