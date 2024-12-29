package main

import (
	"fmt"
	"log"

	"github.com/rdcarranza/s1-bot-telegram-go/cmd/api"
	"github.com/rdcarranza/s1-bot-telegram-go/src/controladores/env"
)

func main() {
	//cargar las variables de entorno.
	env_ := "./.env"
	env_copia := "./src/controladores/env/env.copia"
	if env.VerificarEnv(env_, env_copia) {
		estado_env, err := env.GetEnv("estado_env", env_)
		if err == nil && estado_env == "1" {
			fmt.Println("Archivo env: " + env_ + " - cargado correctamente.")
		} else {
			if err == nil {
				log.Fatal("Verificar configuración de archivo env")
			} else {
				log.Fatal("Verificar configuración de archivo env: " + err.Error())
			}

		}
	}

	//iniciar API
	api.Iniciar()

}
