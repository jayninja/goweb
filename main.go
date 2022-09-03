package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func formatAsString() string {
	return "pow"
}

func main() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}")
	r.SetFuncMap(template.FuncMap{
		"formatAsString": formatAsString,
	})
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
			posts = append(posts, "markdown/"+file.Name())
		}

		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"posts": posts,
		})
	})
	r.GET("/markdown", func(c *gin.Context) {
		contentTemp, err := os.ReadFile("markdown/hello-world.md")
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		content := string(contentTemp)

		c.HTML(http.StatusOK, "test.tmpl.html", gin.H{
			"content": content,
		},
		)

	})

	r.Run()
}
