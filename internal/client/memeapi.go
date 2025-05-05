package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MemeAPIResponse struct {
	URL string `json:"url"`
}

func GetRandomMemeURL() (string, error) {
	resp, err := http.Get("https://meme-api.com/gimme")
	if err != nil {
		return "", fmt.Errorf("Ошибка запроса: статус %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Не удалось получить мем 😢: статус %d", resp.StatusCode)
	}

	var data struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("Ошибка запроса: статус %d", resp.StatusCode)
	}

	return data.URL, nil
}
