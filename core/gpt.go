package core

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type GptClient struct {
	Client         *openai.Client
	MessageHistory []openai.ChatCompletionMessage
}

func (g *GptClient) AddMessage(msg openai.ChatCompletionMessage) {
	//remove old message
	if len(g.MessageHistory) == 3 {
		g.MessageHistory = g.MessageHistory[1:]
	}
	g.MessageHistory = append(g.MessageHistory, msg)
}

func (g *GptClient) DisplayMessage(msg openai.ChatCompletionMessage) {
	fmt.Println("HackBot:\n", msg.Content)
}

func (g *GptClient) ClearMessage() {
	g.MessageHistory = []openai.ChatCompletionMessage{}
	fmt.Println("HackBot:", "新的会话")
}

func (g *GptClient) InitClient(opt *Options) {
	config := openai.DefaultConfig(opt.AutoToken)

	if opt.Proxy != "" {
		proxyURL, err := url.Parse(opt.Proxy)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}

		httpClient := &http.Client{
			Transport: transport,
		}
		config.HTTPClient = httpClient
	}

	g.Client = openai.NewClientWithConfig(config)
}

var FuncPrompt = "帮我判断下面的问题是否需要调用安全工具进行处理,当需要调用工具时回复yes,当需要对数据进行处理时回复no,只用英文回答yes或no\n例子1: \n帮我对www.baidu.com进行网站识别。\nA: yes\n例子2: \n帮我对刚刚收集的数据进行处理。\nA: no\n问题: \n"

func (g *GptClient) RequestFunc(input string) (flag bool, err error) {
	messages := []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: input}}
	resp, err := g.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo0613,
			Messages: messages,
		},
	)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}

	fmt.Println(resp.Choices[0].Message.Content)
	if strings.Contains(resp.Choices[0].Message.Content, "yes") {
		return true, err
	}

	return false, err
}

func (g *GptClient) Request(funcMode bool) {
	if funcMode {
		resp, err := g.Client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:        openai.GPT3Dot5Turbo0613,
				Messages:     g.MessageHistory,
				Functions:    MyFunctions,
				FunctionCall: "auto",
			},
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		respMessage := resp.Choices[0].Message
		if respMessage.FunctionCall == nil {
			fmt.Println("Error: The input information is too little to call the function")
			return
		}

		funcName := resp.Choices[0].Message.FunctionCall
		fmt.Println("HackBot: 正在调用", funcName.Name)

		content, err := FuncCall(funcName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if funcName.Name == "summarize" {
			fmt.Println("FuncResponse: 文章读取完毕")
		} else {
			fmt.Println("FuncResponse:\n", content)
		}

		//4000 token
		if len(content) >= 4000 {
			content = content[0:4000]
		}

		g.AddMessage(respMessage)
		message := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleFunction,
			Name:    respMessage.FunctionCall.Name,
			Content: content,
		}
		g.AddMessage(message)
	} else {
		resp, err := g.Client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo0613,
				Messages: g.MessageHistory,
			},
		)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		g.DisplayMessage(resp.Choices[0].Message)
		g.AddMessage(resp.Choices[0].Message)

	}
}
