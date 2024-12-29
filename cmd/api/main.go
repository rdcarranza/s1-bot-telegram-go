package api

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/rdcarranza/s1-bot-telegram-go/src/controladores/env"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Iniciar() {
	fmt.Println("iniciando API...")
	env_ := "./.env"
	botToken, err := env.GetEnv("telegram_token", env_)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithCheckInitTimeout(5 * time.Second),
	}

	b, err := bot.New(botToken, opts...)
	if err != nil {
		panic(err)
	}

	/*
		b, err := bot.New(botToken)
		if nil != err {
			// panics for the sake of simplicity.
			// you should handle this error properly in your code.
			panic(err)
		}
	*/

	//verificar el getme
	me, err := b.GetMe(ctx)
	//me, err := b.GetMe(context.Background())
	if err != nil {
		panic(fmt.Sprintf("error call getMe: %v", err))
	}
	fmt.Printf("Bot: %+v\n", me.Username)

	b.Start(ctx)
	fmt.Println("Bot inicializado!")
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   update.Message.Text,
		})
	}
}
