package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFile("/", "index.html")
	r.StaticFS("/static", http.Dir("static/"))
	r.Run(":8000")
}
