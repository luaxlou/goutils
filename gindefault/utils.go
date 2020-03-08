package gindefault

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ReturnSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"status": 1,
		"data":   data,
	})

}

func ReturnFail(c *gin.Context, msg string) {

	fmt.Println(msg)
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    msg,
	})

	c.Abort()

}

func ReturnFailWithCode(c *gin.Context, msg string, code int) {

	c.JSON(200, gin.H{
		"status": 0,
		"code":   code,
		"msg":    msg,
	})

	c.Abort()

}

func ReturnError(c *gin.Context, err error) {

	fmt.Println(err.Error())
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    err.Error(),
	})

	c.Abort()

}
