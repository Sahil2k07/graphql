package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Sahil2k07/graphql/internal/configs"
	"github.com/Sahil2k07/graphql/internal/database"
	"github.com/Sahil2k07/graphql/internal/graphql/directives"
	"github.com/Sahil2k07/graphql/internal/graphql/generated"
	"github.com/Sahil2k07/graphql/internal/middlewares"
	"github.com/Sahil2k07/graphql/internal/web"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var echoLambda *echoadapter.EchoLambda

func lambdaHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.Proxy(req)
}

func init() {
	configs := configs.LoadConfig()
	database.Connect()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: configs.Server.Origins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Use(middlewares.JWTContext())

	gqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: web.Resolvers(),
				Directives: generated.DirectiveRoot{
					Public: directives.AuthDirective(),
				},
			},
		),
	)

	e.POST("/graphql", echo.WrapHandler(gqlHandler))
	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL playground", "/graphql")))

	echoLambda = echoadapter.New(e)
}

func main() {
	lambda.Start(lambdaHandler)
}
