package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() *logrus.Logger {
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Создаём (или открываем) файл для логов
	file, err := os.OpenFile("bot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		// Создаём (или открываем) файл для логов
		Log.Info("Не удалось открыть файл для логов, логирование происходит в консоль")
	}

	return Log
}
