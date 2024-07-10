package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func main() {
	// getInstance
	server := gin.Default()

	//set cors
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	server.GET(("/chat"), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})
	server.POST("/chat", func(c *gin.Context) {

		var params struct {
			Text string `json:"text"`
		}
		c.ShouldBindJSON(&params)
		fmt.Printf(params.Text)
		client := arkruntime.NewClientWithApiKey(
			os.Getenv("ARK_API_KEY"),
			arkruntime.WithBaseUrl("https://ark.cn-beijing.volces.com/api/v3"),
			arkruntime.WithRegion("cn-beijing"),
		)
		ctx := context.Background()
		// use if err... to replace the try catch
		req := model.ChatCompletionRequest{
			Model: "ep-20240707141643-75bdr",
			Messages: []*model.ChatCompletionMessage{
				{
					Role: model.ChatMessageRoleSystem,
					Content: &model.ChatCompletionMessageContent{
						StringValue: volcengine.String("你是豆包，是由字节跳动开发的 AI 人工智能助手"),
					},
				},
				{
					Role: model.ChatMessageRoleUser,
					Content: &model.ChatCompletionMessageContent{
						StringValue: volcengine.String(params.Text),
					},
				},
			},
		}
		resp, err := client.CreateChatCompletion(ctx, req)
		fmt.Println(*resp.Choices[0].Message.Content.StringValue)
		if err != nil {

			fmt.Printf("standard chat error: %v\n", err)
			return
		}
		// fmt.Println(*resp.Choices[0].Message.Content.StringValue)
		// stream, err := client.CreateChatCompletionStream(ctx, req)
		// if err != nil {
		// 	fmt.Printf("stream chat error: %v\n", err)
		// 	return
		// }
		// defer stream.Close()

		// for {
		// 	recv, err := stream.Recv()
		// 	if err == io.EOF {
		// 		return
		// 	}
		// 	if err != nil {
		// 		fmt.Printf("Stream chat error: %v\n", err)
		// 		return
		// 	}

		// 	if len(recv.Choices) > 0 {
		// 		fmt.Print(recv.Choices[0].Delta.Content)
		// 	}
		// }
		if err := c.ShouldBindJSON(&params); err != nil {
			fmt.Println("绑定错误信息:", err)
			c.JSON(400, gin.H{
				"error": "invalid params",
			})
			return
		}
		c.JSON(200, gin.H{"success": "OK", "content": params.Text, "answer": resp.Choices[0].Message.Content})
	})
	server.Run(":9000")

	// fmt.Println(*resp.Choices[0].Message.Content.StringValue)

	// fmt.Println("----- streaming request -----")
	// req = model.ChatCompletionRequest{
	// 	Model: "ep-20240707141643-75bdr",
	// 	Messages: []*model.ChatCompletionMessage{
	// 		{
	// 			Role: model.ChatMessageRoleSystem,
	// 			Content: &model.ChatCompletionMessageContent{
	// 				StringValue: volcengine.String("你是豆包，是由字节跳动开发的 AI 人工智能助手"),
	// 			},
	// 		},
	// 		{
	// 			Role: model.ChatMessageRoleUser,
	// 			Content: &model.ChatCompletionMessageContent{
	// 				StringValue: volcengine.String("请问什么是厘米的猜想"),
	// 			},
	// 		},
	// 	},
	// }

}
