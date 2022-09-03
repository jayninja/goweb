package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}")
	//ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("./templates/*.tmpl.html")

	r.GET("/", func(c *gin.Context) {
		var posts []string
		// Read the directory, if it fails error
		files, err := ioutil.ReadDir("./markdown/")
		if err != nil {
			log.Fatal(err)
		}
		// for each file
		for _, file := range files {
			fmt.Println(file.Name())
			posts = append(posts, file.Name())
		}

		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"posts": posts,
		})
	})

	r.Run()
}
