package main

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Format struct {
	Type string `json:"type"`
}

type Payload struct {
	Model       string    `json:"model"`
	Seed        int       `json:"seed"`
	Temperature float32   `json:"temperature"`
	Messages    []Message `json:"messages"`
	Format      Format    `json:"response_format"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choice struct {
	Message      Message `json:"message"`
	LogProbs     int     `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
}

type Response struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}
