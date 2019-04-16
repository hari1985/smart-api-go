package main

import (
	"smartapigo/detail-testa/app"
	"smartapigo/detail-testa/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}

//french-latin - detail-testa -> detail test
