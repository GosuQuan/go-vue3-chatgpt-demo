package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
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

	server.Run(":9000")
	// client := arkruntime.NewClientWithApiKey(
	// 	os.Getenv("ARK_API_KEY"),
	// 	arkruntime.WithBaseUrl("https://ark.cn-beijing.volces.com/api/v3"),
	// 	arkruntime.WithRegion("cn-beijing"),
	// )

	// ctx := context.Background()

	// fmt.Println("----- standard request -----")
	// req := model.ChatCompletionRequest{
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

	// resp, err := client.CreateChatCompletion(ctx, req)
	// if err != nil {
	// 	fmt.Printf("standard chat error: %v\n", err)
	// 	return
	// }
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
}
