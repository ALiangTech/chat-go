package gin

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
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
		Scheme: "https",
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

// 走请求模式 不走代理
type OpenAIBody struct {
	Messages []openai.ChatCompletionMessage `json:"messages"`
}

func RequestOpenAi(r *gin.Engine) {

	r.POST("/chat", func(ctx *gin.Context) {
		// 获取body 参数
		// ChatMessageRoleSystem    = "system"
		// ChatMessageRoleUser      = "user"
		// ChatMessageRoleAssistant = "assistant"
		var body OpenAIBody
		if err := ctx.BindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		config := openai.DefaultConfig("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")
		config.BaseURL = "https://openai.hanjunty.top/v1"
		c := openai.NewClientWithConfig(config)
		octx := context.Background()

		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 20,
			Messages:  body.Messages,
			// Messages: []openai.ChatCompletionMessage{
			// 	{
			// 		Role:    openai.ChatMessageRoleUser,
			// 		Content: "你好",
			// 	},
			// },
			Stream: true,
		}

		stream, err := c.CreateChatCompletionStream(octx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			return
		}
		defer stream.Close()
		fmt.Printf("Stream response: ")
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nstream finished")
				return
			}
			if err != nil {
				fmt.Printf("\nstream error: %v\n", err)
				return
			}
			sse.Encode(ctx.Writer, sse.Event{
				Event: "message",
				Data:  response.Choices[0].Delta.Content,
			})
			// 刷新数据，以通知请求端
			ctx.Writer.Flush()
		}
	})

}
