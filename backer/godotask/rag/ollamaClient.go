package rag

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func QueryOllama(prompt string) (string, error) {
	payload := map[string]interface{}{
		"model":  "llama3.1:8b",
		"prompt": prompt,
		"stream": false,
	}
	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post("http://host.docker.internal:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if response, ok := result["response"].(string); ok {
		return response, nil
	}
	return "", nil
}