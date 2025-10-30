package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Sahil2k07/graphql/internal/configs"
	"github.com/Sahil2k07/graphql/internal/database"
	"github.com/Sahil2k07/graphql/internal/graphql/directives"
	"github.com/Sahil2k07/graphql/internal/graphql/generated"
	"github.com/Sahil2k07/graphql/internal/graphql/resolvers"
	"github.com/Sahil2k07/graphql/internal/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs := configs.LoadConfig()
	database.Connect()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     configs.Server.Origins,
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	e.Use(middlewares.JWTContext())

	gqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolvers.Resolver{},
				Directives: generated.DirectiveRoot{
					Public: directives.AuthDirective(),
				},
			},
		),
	)

	e.POST("/graphql", echo.WrapHandler(gqlHandler))
	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL playground", "/graphql")))

	e.Logger.Fatal(e.Start(configs.Server.ServerPort))
}
