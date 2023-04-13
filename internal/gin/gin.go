package gin

import "github.com/gin-gonic/gin"

func StartServer() {
	r := gin.Default()
	ProxyOpenAI(r)
	r.Run(":9000")
}
