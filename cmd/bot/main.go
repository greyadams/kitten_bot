package main

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"github.com/greyadams/kitten_bot/internal/logger"

	"github.com/greyadams/kitten_bot/internal/client"
	"github.com/greyadams/kitten_bot/internal/storage"
)

func main() {
	//Инициализация логгера
	log := logger.InitLogger()

	//Загрузка файла окружения .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("Токен Telegram бота отсутствует в переменных окружения")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %s", err)
	}

	log.Infof("Бот запущен")

	stats := &storage.Stats{}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		//Логирование всех команд от пользователей
		log.WithFields(map[string]interface{}{
			"user_id":  update.Message.From.ID,
			"username": update.Message.From.UserName,
			"command":  update.Message.Text,
		}).Info("Получено сообщение")

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Приветик ! Я бот, который любит скидывать милых котиков 🐱\n Напиши /meow, чтобы получить котика!")
				if _, err := bot.Send(msg); err != nil {
					log.WithError(err).Error("Ошибка при команде /start")
				} else {
					log.Info("Команда /start выполнена успешно")
				}
			case "meow":
				catURL, err := client.GetRandomCatImageURL()
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось получить котика 😿")
					bot.Send(msg)
					log.WithError(err).Error("Ошибка при команде /meow")
					continue
				}

				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(catURL))
				if _, err := bot.Send(photo); err != nil {
					log.WithError(err).Error("Ошибка при отправке фото котика 😿")
				} else {
					log.Info("Фото отправлено")
				}
				photo.Caption = "Вот твой котик! 🐱"

				stats.IncCat()

			/*case "meme":
			url, err := client.GetRandomMemeURL()
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при получении мема 😿")
				bot.Send(msg)
				continue
			}

			photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(url))
			photo.Caption = "😸😸😸"
			bot.Send(photo)
			stats.IncMeme()
			*/
			case "stats":
				cats, memes := stats.GetStats()
				text := fmt.Sprintf("Статистика:\nКотики: %d\nМемы: %d", cats, memes)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)

			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Мяу, я не понимаю 😿")
				if _, err := bot.Send(msg); err != nil {
					log.WithError(err).Error("Ошибка отправки сообщения о неизвестной команде")
				} else {
					log.Warn("Получена неизвестная команда")
				}
			}
		}
	}

}
