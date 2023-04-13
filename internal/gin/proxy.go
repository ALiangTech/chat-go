package gin

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type CustomTransport struct {
	http.RoundTripper
}

func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// 修改请求数据
	req.Header.Set("Authorization", "Custom-Value")

	// 发送请求到代理服务器
	return t.RoundTripper.RoundTrip(req)
}

// 发送请求到代理域名上面

// 代理openai

func ProxyOpenAI(r *gin.Engine) {
	// 创建自定义传输对象
	transport := &CustomTransport{
		RoundTripper: http.DefaultTransport,
	}
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "openai.hanjunty.top",
	})
	// 设置传输对象
	proxy.Transport = transport
	// 代理路由
	r.Any("v1/*path", func(c *gin.Context) {
		// 将请求转发到代理服务器
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}
