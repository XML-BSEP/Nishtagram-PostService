package main

import (
	"context"
	"fmt"
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

	repo := i.NewPostUseCase()
	res, err := repo.GenerateUserFeed("424935b1-766c-4f99-b306-9263731518bc", context.Background())
	fmt.Println(res)

	router := router2.NewRouter(handler)


	router.RunTLS("localhost:8083", "certificate/cert.pem", "certificate/key.pem")

}
