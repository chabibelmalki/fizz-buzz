package webHandler

import (
	"fizz-buzz/library"
	"fizz-buzz/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Handler_FizzBuzz(c *gin.Context) {
	var textError string

	// retrieve parameters
	string1 := c.Query("string1")
	string2 := c.Query("string2")
	p_int1 := c.Query("int1")
	p_int2 := c.Query("int2")
	p_limit := c.Query("limit")

	// check parameters validity
	// if int1 is integer
	bResult, int1 := library.IsInt(p_int1)
	if !bResult {
		textError += "The parameter 'int1' is not a valid integer.\n"
	}
	// if int2 is integer
	bResult, int2 := library.IsInt(p_int2)
	if !bResult {
		textError += "The parameter 'int2' is not a valid integer.\n"
	}
	// if int3 is integer
	bResult, limit := library.IsInt(p_limit)
	if !bResult || limit < 1 {
		textError += "The parameter 'limit' should be an integer and greater than 0."
	}

	// everything's okay
	if textError == "" {
		result := service.Fizzbuzz(string1, string2, int1, int2, limit)

		fmt.Println(result)
		//Return results
		HTTPDeclare_Success(c, result)
	} else {
		HTTPDeclare_BadRequest(c, textError)
	}

}
