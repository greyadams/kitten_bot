package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CatAPIResponse struct {
	URL string `json:"url"`
}

func GetRandomCatImageURL() (string, error) {
	resp, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		return "", fmt.Errorf("Ошибка запроса: статус %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ошибка запроса (не удалось получить котика): статус %d", resp.StatusCode)
	}

	var data []CatAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("Ошибка ответа: статус %d", resp.StatusCode)
	}

	if len(data) == 0 {
		return "", fmt.Errorf("Пустой ответ от TheCatAPI: статус %d", resp.StatusCode)
	}

	return data[0].URL, nil
}
