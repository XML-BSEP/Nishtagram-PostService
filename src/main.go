package main

import (
	"context"
	gin "github.com/gin-gonic/gin"
	"log"
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

	repo := i.NewPostRepo()
	repo.GetPostsByUserId("43420055-3174-4c2a-9823-a8f060d644c3", context.Background())

	g := gin.Default()
	g.GET("/ping", handler.AddPost)

	g.Run("localhost:8083")
}
