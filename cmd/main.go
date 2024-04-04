package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./templates/dir.html", "./templates/file.html")

	router.GET("*filepath", func(c *gin.Context) {
		fpath := c.Param("filepath")
		fileInfo, err := os.Stat(fpath)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if fileInfo.IsDir() {
			dirs, err := os.ReadDir(fpath)
			if err != nil {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			c.HTML(http.StatusOK, "dir.html", gin.H{
				"dirs": dirs,
			})
			return
		} else {
			content, err := os.ReadFile(fpath)
			if err != nil {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			c.HTML(http.StatusOK, "file.html", gin.H{
				"data": string(content),
			})
		}
	})

	router.Run(":8080")
}
