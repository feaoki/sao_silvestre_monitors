package genai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Response struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func Check() {
	credFile := "../.credenciais/credenciais.json"
	credData, err := os.ReadFile(credFile)
	if err != nil {
		panic(fmt.Sprintf("Erro ao ler o arquivo de credenciais: %v", err))
	}
	var creds map[string]map[string]string
	if err := json.Unmarshal(credData, &creds); err != nil {
		panic(fmt.Sprintf("Erro ao fazer unmarshal das credenciais: %v", err))
	}
	apiKey := creds["desafio05"]["api"]

	body := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": "As inscrições para a São Silvestre de 2025 já estão abertas? Responda com 'SIM' ou 'NAO'."},
				},
			},
		},
	}

	jsonData, _ := json.Marshal(body)
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-goog-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respData, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(respData))
	var response Response
	if err := json.Unmarshal(respData, &response); err != nil {
		panic(fmt.Sprintf("Erro ao fazer unmarshal da resposta: %v", err))
	}
	inscricoesAbertas := false
	if response.Candidates[0].Content.Parts[0].Text == "SIM" {
		inscricoesAbertas = true
	}
	fmt.Println("Inscrições abertas:", inscricoesAbertas)

}
