package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello userinfo")
	})

	r.GET("/request_introspect", func(c *gin.Context) {
		req, _ := http.NewRequest("GET", os.Getenv("AUTH_INTERNAL_URL")+"/introspect", nil)

		client := new(http.Client)
		resp, e := client.Do(req)
		if e != nil {
			panic(e)
		}
		defer resp.Body.Close()
		byteArray, _ := ioutil.ReadAll(resp.Body)
		c.String(http.StatusOK, string(byteArray))
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "userinfo OK")
	})
	r.Run(":7002")
}
