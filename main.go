package main

import (
	"os"

	"github.com/y16ra/go-gin-helloworld/cmd"
)

func main() {
	if isLambda() {
		os.Args = []string{os.Args[0], cmd.Lambda.Name()}
	} else {
		os.Args = []string{os.Args[0], cmd.Server.Name()}
	}
	cmd.Execute()
}

func isLambda() bool {
	return os.Getenv("LAMBDA_TASK_ROOT") != ""
}
