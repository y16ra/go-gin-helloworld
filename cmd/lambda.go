package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var Lambda = NewLambda()

func NewLambda() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "lambda",
		Short: "Run lambda handler",
		Long:  `Run lambda handler`,
		Run: func(cmd *cobra.Command, args []string) {
			// stdout and stderr are sent to AWS CloudWatch Logs
			log.Printf("Gin cold start")
			r := gin.Default()
			r.Use(gin.Recovery())
			r.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "Hello",
				})
			})
			r.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})

			ginLambda = ginadapter.New(r)
			lambda.Start(Handler)
		},
	}
	return cmd
}

func init() {
	rootCmd.AddCommand(Lambda)
}

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}
