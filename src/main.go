package main

import (
	"log"
	router2 "post-service/http/router"
	"post-service/infrastructure/cassandra_config"
	"post-service/interactor"
)

func main() {
	cassandraSession, err := cassandra_config.NewCassandraSession()
	if err != nil {
		log.Println(err)
	}

	i := interactor.NewInteractor(cassandraSession)

	handler := i.NewAppHandler()

	router := router2.NewRouter(handler)

	router.RunTLS(":8083", "certificate/cert.pem", "certificate/key.pem")

}
