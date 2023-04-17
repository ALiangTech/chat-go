package gin

import (
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	// ProxyOpenAI(r)
	RequestOpenAi(r)
	r.Run(":9000")
}
