package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const openAIEndpoint = "https://api.openai.com/v1/engines/text-davinci-003/completions"

type TextRequest struct {
	Prompt string `json:"prompt"`
}
type TextResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func getTextCompletion(prompt string) (string, error) {
	data := TextRequest{
		Prompt: prompt,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	authenticateRequest(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var result TextResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Choices[0].Text, nil
}
func main() {
	prompt := "Once upon a time, in a land far away, there was a brave knight."
	response, err := getTextCompletion(prompt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Generated story:")
	fmt.Println(response)
}
