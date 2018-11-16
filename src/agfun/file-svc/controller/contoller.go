package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)

	// Upload the file to specific dst.
	// c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
