package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/greyadams/kitten_bot/internal/logger"
)

type CatAPIResponse struct {
	URL string `json:"url"`
}

func GetRandomCatImageURL() (string, error) {
	resp, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		logger.Log.WithError(err).Error("Ошибка отправки запроса к TheCatAPI")
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Log.WithField("status", resp.StatusCode).Error("Неожиданный статус код от TheCatAPI: %d", resp.StatusCode)
		return "", errors.New("не удалось получить картинку котика")
	}

	var data []CatAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		logger.Log.WithError(err).Error("Ошибка декодирования ответа от TheCatAPI")
		return "", err
	}

	if len(data) == 0 {
		logger.Log.Warn("Пустой ответ от TheCatAPI")
		return "", errors.New("пустой ответ от TheCatAPI")
	}

	logger.Log.WithField("image_url", data[0].URL).Info("Картинка котика успешно получена")
	return data[0].URL, nil
}
