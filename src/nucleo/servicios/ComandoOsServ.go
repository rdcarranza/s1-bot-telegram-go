package servicios

import (
	"fmt"
	"log"
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
		log.Println("error ComandoOsServ - EjecutarComando() - ", err)
		return "", err
	}
	return string(out), nil
}

func (c *ComandoOsServ) ReproducirOggCvlc(arch_ogg string) error {
	e := exec.Command("cvlc", "--intf", "dummy", "--volume", "384", "--play-and-exit", arch_ogg).Run()
	if e != nil {
		log.Println("error ComandoOsServ - ReproducirOggCvlc() - ", e)
		return e
	}
	return nil
}

func (c *ComandoOsServ) ReproducirOggFfplay(arch_ogg string) error {
	e := exec.Command("ffplay", "-nodisp", "-volume", "192", "-autoexit", arch_ogg).Run()
	if e != nil {
		log.Println("error ComandoOsServ - ReproducirOggFfplay() - ", e)
		return e
	}
	return nil
}

func (c *ComandoOsServ) LecturaTexto(texto string) error {
	fmt.Println("LecturaTexto(): texto -> ", texto)
	e := exec.Command("espeak", "-v", "es-la+m3", "-s", "100", "-a", "100", "-g", "10", "-p", "35", texto).Run()
	if e != nil {
		log.Println("error ComandoOsServ - LecturaTexto() - ", e)
		return e
	} else {
		exec.Command("espeak", "-v", "es-la+m3", "-s", "100", "-a", "100", "-g", "10", "-p", "35", "Repito - "+texto).Run()
	}

	return nil
}
