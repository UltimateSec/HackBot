package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"hackbot/core"
	"os"
	"strings"
)

func main() {
	var opt core.Options
	flag.StringVar(&opt.Proxy, "p", "", "Proxy")
	flag.StringVar(&opt.AutoToken, "t", "", "Openai Token")
	flag.Parse()

	var gpt core.GptClient
	gpt.InitClient(&opt)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		flag := false
		fmt.Print("Input: ")
		scanner.Scan()
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			return
		} else if strings.ToLower(input) == "clear" {
			gpt.ClearMessage()
		}
		//funcInput := core.FuncPrompt + input + "\nA: "
		//fmt.Println(funcInput)
		//flag, err := gpt.RequestFunc(funcInput)
		//if err != nil {
		//	continue
		//}
		fmt.Print("是否调用Function(y/n): ")
		scanner.Scan()
		funcFlag := scanner.Text()
		if strings.ToLower(funcFlag) == "y" {
			flag = true
		}
		messages := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		}
		gpt.AddMessage(messages)
		gpt.Request(flag)
	}
}
