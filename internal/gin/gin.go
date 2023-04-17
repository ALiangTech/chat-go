package gin

import "github.com/gin-gonic/gin"

func StartServer() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	// ProxyOpenAI(r)
	RequestOpenAi(r)
	r.Run(":9000")
}
