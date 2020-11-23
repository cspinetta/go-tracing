package main

import (
	"github.com/cspinetta/go-tracing/src/app"
	"github.com/cspinetta/go-tracing/src/base"
)

func main() {
	db, err := base.GetDB()
	if err != nil {
		panic("Error trying to connect to DB")
	}
	appRunner := app.NewApp(db)
	appRunner.Start("8080")
}
