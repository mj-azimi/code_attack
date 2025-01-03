package aiService

import (
	"code_attack/lib"
	"encoding/json"
	"fmt"
	"os"
)

type ApiResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
}

func Generate(code string) (string, error) {
	model, _ := lib.GetConfig("ai_model")

	var promt string = " This is a code. Check how clean it is and how coding principles are followed in it. Tell me what points can be considered for improvement. "
	var prompt2 = " Speak briefly. "

	data := map[string]interface{}{
		"model":  model,
		"prompt": promt + prompt2 + code,
		"stream": false,
	}

	response, err := lib.Post("/generate", data)
	if err != nil {
		fmt.Printf("Error traversing files: %v", err)
	}

	var apiResponse ApiResponse

	json.Unmarshal([]byte(response), &apiResponse)

	if apiResponse.Response == "" {
		fmt.Println("Error Api")
		os.Exit(1)
	}

	return apiResponse.Response, nil
}
