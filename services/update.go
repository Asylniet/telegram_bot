package services

import (
	"github.com/Asylniet/controllers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func Updates() {
	controllers.LoadEnvVariables()

	unsplashKey := os.Getenv("UNSPLASH_TOKEN")
	imageController := controllers.ImageController{
		UnsplashKey: unsplashKey,
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		go func() {
			if update.Message != nil {
				chatID := update.Message.Chat.ID

				switch update.Message.Text {
				case "/start":
					msg := tgbotapi.NewMessage(chatID, "Hello, I'm Imager bot!")

					if _, err := bot.Send(msg); err != nil {
						log.Panic(err)
					}
				case "/image":
					photoConfig := imageController.GetPhoto(chatID)

					if _, err := bot.Send(photoConfig); err != nil {
						log.Panic(err)
					}
				default:
					msg := tgbotapi.NewMessage(chatID, "Unknown command, please enter /image")

					if _, err := bot.Send(msg); err != nil {
						log.Panic(err)
					}
				}

			}
		}()
	}
}
