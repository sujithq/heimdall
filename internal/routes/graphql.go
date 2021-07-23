package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/moducate/heimdall/internal/env"
	"github.com/moducate/heimdall/internal/graph"
	"github.com/moducate/heimdall/internal/graph/generated"
)

func Graphql(e *env.Env, r *gin.RouterGroup) {
	r.POST("/", graphqlHandler(e))
	r.GET("/", playgroundHandler())
}

func graphqlHandler(e *env.Env) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Env: e,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
