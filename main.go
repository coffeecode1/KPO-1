package main

import (
	"log"
	"os"

	"KPO-1/internal/app"
	"KPO-1/internal/di"
	"KPO-1/internal/service"
)

func main() {
	container, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(func(park *service.Park, report *service.Report) {
		app.New(park, report, os.Stdin, os.Stdout).Run()
	})
	if err != nil {
		log.Fatal(err)
	}
}
