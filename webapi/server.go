package webapi

import (
	"fizz-buzz/webHandler"

	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	// init structure to stock requests & their number of hits

	// init server
	router := gin.Default()

	// using middleware to count number of hits by endpoint (fullpath)
	router.Use(webHandler.HitEndPoint)

	// init paths
	// fizzbuzz : service
	router.GET("/fizzbuzz", webHandler.Handler_FizzBuzz)

	// return the parameters corresponding to the most used request
	router.GET("/kpi_mostusedrequest", webHandler.Handler_KPI_MostUsedRequest)

	// return
	return router
}
