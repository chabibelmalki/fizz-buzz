package webHandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// one struct type for all results of this app
type StJsonResult struct {
	ErrorStatus  bool        `json:"error"`
	ErrorMessage string      `json:"message"`
	LogTime      time.Time   `json:"logTime"`
	Data         interface{} `json:"data"`
}

func HTTPDeclare_BadRequest(c *gin.Context, message string) {
	// call http declare with bad request status
	httpDeclare(c, http.StatusBadRequest, message, nil)
}

func HTTPDeclare_Success(c *gin.Context, data interface{}) {
	// call http declare with accepted status
	httpDeclare(c, http.StatusAccepted, "accepted", data)
}

// declare HTTP result and abort context
func httpDeclare(c *gin.Context, errorType int, message string, data interface{}) {
	//abort context
	c.Abort()

	bResult := (errorType < 299)

	// init & return result
	sReturnResult := StJsonResult{
		ErrorStatus:  bResult,
		ErrorMessage: message,
		LogTime:      time.Now(),
		Data:         data,
	}
	c.IndentedJSON(errorType, sReturnResult)
}
