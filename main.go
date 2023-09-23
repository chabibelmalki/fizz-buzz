package main

import (
	"fizz-buzz/app"
	"fizz-buzz/webapi"
)

func main() {
	//load configuration
	err := app.LoadConfiguration()

	if err == nil {
		//get server and start it
		router := webapi.GetServer()

		router.Run(app.GetApiAndPort())
	} else {
		panic(err.Error())
	}

}
