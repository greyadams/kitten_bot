package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"github.com/greyadams/kitten_bot/internal/client"
	"github.com/greyadams/kitten_bot/internal/storage"
)

func main() {
	stats := &storage.Stats{}

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
		log.Fatal("Ошибка при создании бота: %s", err)
	}

	log.Printf("Бот %s запущен", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Приветик ! Я бот, который любит скидывать милых котиков 🐱\n Напиши /meow, чтобы получить котика!")
				bot.Send(msg)
			case "cat":
				catURL, err := client.GetRandomCatImageURL()
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось получить котика 😿")
					bot.Send(msg)
					continue
				}

				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(catURL))
				photo.Caption = "Вот твой котик! 🐱"
				bot.Send(photo)
				stats.IncCat()

			case "meme":
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

			case "stats":
				cats, memes := stats.GetStats()
				text := fmt.Sprintf("Статистика:\nКотики: %d\nМемы: %d", cats, memes)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)

			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Мяу, я не понимаю 😿")
				bot.Send(msg)
			}
		}
	}

}
