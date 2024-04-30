package main

import (
	"log"
	"manav402/graphql/database"
	"manav402/graphql/graph"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("err initializing env variable ",err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db,err := database.InitDatabase()
	if err != nil {
		log.Println("err initializing database",err)
		return
	}

	database.DB = db
	log.Println("database initialized succesfully")

	
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
