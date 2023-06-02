package main

import (
	"log"
	"net/http"
	"os"

	"github.com/39TO/gockerql/graph"
	"github.com/39TO/gockerql/graph/resolver"
	"github.com/39TO/gockerql/infrastructure/database"
	"github.com/39TO/gockerql/infrastructure/persistance"
	"github.com/39TO/gockerql/usecase"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repoTodo := persistance.NewTodoRepository(db)
	ucTodo := usecase.NewTodoUsecase(repoTodo)
	rvTodo := resolver.NewTodoResolver(ucTodo)

	resolver := resolver.Resolver{
		Todo: rvTodo,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
