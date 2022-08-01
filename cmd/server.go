package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/y16ra/go-gin-helloworld/middlewares"
	"github.com/y16ra/go-gin-helloworld/routes"
)

var Server = NewServer()

func NewServer() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "server",
		Short: "Run server",
		Long:  `Run server`,
		Run: func(cmd *cobra.Command, args []string) {
			r := gin.Default()
			r.Use(gin.Recovery())
			r.Use(middlewares.Check())
			routes.SetupGreeting(r)
			r.Run()
		},
	}
	return cmd
}

func init() {
	rootCmd.AddCommand(Server)
}
