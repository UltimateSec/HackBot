package core

import (
	"encoding/json"
	"github.com/sashabaranov/go-openai"
)

var HTTPXFunc = openai.FunctionDefinition{
	Name:        "httpx",
	Description: "Used for website survival detection and basic website information acquisition, such as website title, HTTP response content, etc.(用于网站存活探测和网站基本信息获取，如网站标题、http响应内容等)",
	Parameters: &openai.JSONSchemaDefinition{
		Type: openai.JSONSchemaTypeObject,
		Properties: map[string]openai.JSONSchemaDefinition{
			"url": {
				Type:        openai.JSONSchemaTypeString,
				Description: "destination address(url)",
			},
		},
		Required: []string{"url"},
	},
}

var SubFinderFunc = openai.FunctionDefinition{
	Name:        "subfinder",
	Description: "Obtain the Subdomain of the root domain name through the passive source.(通过被动信源的方式获取根域名的子域名)",
	Parameters: &openai.JSONSchemaDefinition{
		Type: openai.JSONSchemaTypeObject,
		Properties: map[string]openai.JSONSchemaDefinition{
			"domain": {
				Type:        openai.JSONSchemaTypeString,
				Description: "destination domain",
			},
		},
		Required: []string{"domain"},
	},
}

var KatanaFunc = openai.FunctionDefinition{
	Name:        "katana",
	Description: "Path crawling and website crawling for website HTML content.(用于网站html内容的路径爬取、网站爬虫)",
	Parameters: &openai.JSONSchemaDefinition{
		Type: openai.JSONSchemaTypeObject,
		Properties: map[string]openai.JSONSchemaDefinition{
			"url": {
				Type:        openai.JSONSchemaTypeString,
				Description: "destination address(url)",
			},
		},
		Required: []string{"url"},
	},
}

var NaabuFunc = openai.FunctionDefinition{
	Name:        "naabu",
	Description: "Used for target host port detection.(用于目标主机端口探测)",
	Parameters: &openai.JSONSchemaDefinition{
		Type: openai.JSONSchemaTypeObject,
		Properties: map[string]openai.JSONSchemaDefinition{
			"host": {
				Type:        openai.JSONSchemaTypeString,
				Description: "destination host",
			},
			"ports": {
				Type:        openai.JSONSchemaTypeString,
				Description: "destination ports",
			},
		},
		Required: []string{"host"},
	},
}

var SummarizeFunc = openai.FunctionDefinition{
	Name:        "summarize",
	Description: "Summarize the content of website articles and extract the key points of the article.(总结网站文章内容，并提炼文章的要点)",
	Parameters: &openai.JSONSchemaDefinition{
		Type: openai.JSONSchemaTypeObject,
		Properties: map[string]openai.JSONSchemaDefinition{
			"url": {
				Type:        openai.JSONSchemaTypeString,
				Description: "article address(url)",
			},
		},
		Required: []string{"url"},
	},
}

var MyFunctions = []openai.FunctionDefinition{
	HTTPXFunc,
	SubFinderFunc,
	NaabuFunc,
	SummarizeFunc,
	KatanaFunc,
}

func FuncCall(call *openai.FunctionCall) (content string, err error) {
	switch call.Name {
	case "subfinder":
		var f Subfinder
		if err = json.Unmarshal([]byte(call.Arguments), &f); err != nil {
			return
		}
		content, err = f.Run()
	case "httpx":
		var f Httpx
		if err = json.Unmarshal([]byte(call.Arguments), &f); err != nil {
			return
		}
		content, err = f.Run()
	case "naabu":
		var f Naabu
		err = json.Unmarshal([]byte(call.Arguments), &f)
		if err = json.Unmarshal([]byte(call.Arguments), &f); err != nil {
			return
		}
		content, err = f.Run()
	case "katana":
		var f Katana
		err = json.Unmarshal([]byte(call.Arguments), &f)
		if err = json.Unmarshal([]byte(call.Arguments), &f); err != nil {
			return
		}
		content, err = f.Run()
	case "summarize":
		var f Summarize
		err = json.Unmarshal([]byte(call.Arguments), &f)
		if err = json.Unmarshal([]byte(call.Arguments), &f); err != nil {
			return
		}
		content, err = f.Run()
	}

	return
}
