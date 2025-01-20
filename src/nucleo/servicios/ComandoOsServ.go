package servicios

import (
	"os/exec"

	"github.com/rdcarranza/s1-bot-telegram-go/src/nucleo/puertos"
)

type ComandoOsServ struct {
}

func Servicios_ComandoOs() puertos.ComandoOS {
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

func (c *ComandoOsServ) ReproducirOggCvlc(arch_ogg string) error {
	e := exec.Command("cvlc", "--intf", "dummy", "--play-and-exit", arch_ogg).Run()
	if e != nil {
		return e
	}
	return nil
}

func (c *ComandoOsServ) ReproducirOggFfplay(arch_ogg string) error {
	e := exec.Command("ffplay", "-nodisp", "-autoexit", arch_ogg).Run()
	if e != nil {
		return e
	}
	return nil
}
