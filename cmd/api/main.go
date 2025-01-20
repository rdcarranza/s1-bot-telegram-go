package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"time"

	controladores_api "github.com/rdcarranza/s1-bot-telegram-go/cmd/api/controladores"
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
	defer cancel() //esta instrucción especifica la función que se ejecutará al finalizar.

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

	fmt.Println("Bot inicializado!")
	b.Start(ctx)

}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		var ca *controladores_api.Ctrl_comandosOsApi

		if update.Message.Voice != nil {
			fileID := update.Message.Voice.FileID
			file, _ := b.GetFile(ctx, &bot.GetFileParams{FileID: fileID})
			//file, _ := b.GetFile(fileID)

			env_ := "./.env"
			botToken, _ := env.GetEnv("telegram_token", env_)

			url := "https://api.telegram.org/file/bot" + botToken + "/" + file.FilePath
			//fmt.Println("url voice: ", url)
			//url:="https://api.telegram.org/file/bot<token>/"
			//https://api.telegram.org/file/bot11111:xxxxxxxxxxxxxxxxxxxxxx/voice/file_0.oga

			// Abre un archivo local para guardar la voz
			outFile, err := os.Create("voice.ogg")
			if err != nil {
				log.Fatal("Error al crear archivo local: ", err)
			}
			defer outFile.Close()

			// Descarga el archivo
			response, err := http.Get(url)
			if err != nil {
				log.Fatal("Error al descargar el archivo de voz: ", err)
			}
			defer response.Body.Close()

			_, err = io.Copy(outFile, response.Body)
			if err != nil {
				log.Fatal("Error al guardando el archivo de voz: ", err)
			}

			exec.Command("ffplay", "-nodisp", "-autoexit", "-volume", "192", "voice.ogg").Run()
			//exec.Command("cvlc", "--intf", "dummy", "--volume", "384", "--play-and-exit", "voice.ogg").Run()
			fmt.Println("Reproduciendo mensaje de voz!")
			/*
				f,_:=b.GetFile(fileID)
				fmt.Println("URL de prueba: "+f.FilePath)
			*/
		}

		if update.Message.Text != "" {
			ca = controladores_api.Controlador_comandosOsApi(update.Message.Text)
			res, err := ca.EjecutarComando()
			if err != nil {
				err.Error()
			}

			if res != "" {
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: update.Message.Chat.ID,
					Text:   res,
				})
			}

		}

		/*
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   update.Message.Text,
			})
		*/
	}
}
