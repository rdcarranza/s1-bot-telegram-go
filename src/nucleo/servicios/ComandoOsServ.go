package servicios

import (
	"os/exec"

	"github.com/rdcarranza/s1-bot-telegram-go/src/nucleo/puertos"
)

type ComandoOsServ struct {
}

func servicios_ComandoOs() puertos.ComandoOS {
	return &ComandoOsServ{}
}

// EjecutarComando ejecuta un comando del sistema y devuelve la salida.
func (c *ComandoOsServ) EjecutarComando(comando string) (string, error) {
	out, err := exec.Command("sh", "-c", comando).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
