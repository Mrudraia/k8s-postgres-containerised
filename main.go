package main

import (
	"fmt"

	"github.com/mrudraia/k8s-postgres-containerised/main_app"
)

func main() {
	fmt.Println("welcome to go postgres application")

	main_app.Application()
}
