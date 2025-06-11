// internal/adapter/server/server.go
package server

import (
	"log"
	"net/http"

	"github.com/willians-e-silva/maestro/internal/adapter/graphql"
	"github.com/willians-e-silva/maestro/internal/usecase"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// ServeGraphQL configura e inicia o servidor GraphQL
func ServeGraphQL(port string, userUsecase *usecase.UserUsecase) {
	// Crie o resolvedor com seus casos de uso injetados
	resolver := graphql.NewResolver(userUsecase)

	// Configure o servidor GraphQL
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("Conecte-se ao http://localhost:%s/ para o playground GraphQL", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
