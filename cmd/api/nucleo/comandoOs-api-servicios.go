package nucleo_api

import (
	controlador_error "github.com/rdcarranza/s1-bot-telegram-go/src/controladores/error"
	"github.com/rdcarranza/s1-bot-telegram-go/src/nucleo/servicios"
)

type ComandoOs_api_servicios struct {
	cOs_api ComandoOs_api
}

// CONSTRUCTOR
func NuevoComandoOs_api_servicio(c string) *ComandoOs_api_servicios {
	comandoOs_api := NuevoComandoOs(c)
	return &ComandoOs_api_servicios{cOs_api: *comandoOs_api}
}

func (cOs *ComandoOs_api_servicios) EjecutarComandoOs_api_servicio() (string, error) {
	servicio := &servicios.ComandoOsServ{}
	return servicio.EjecutarComando(cOs.cOs_api.comando)
}

func (cOs *ComandoOs_api_servicios) ActualizarComandoOs_api_servicio(resultado string) error {
	if !cOs.cOs_api.getEstado() {
		cOs.cOs_api.ActualizarComandoOs(resultado)
		return nil
	} else {
		error1 := controlador_error.NuevoControladorError()
		error1.CargarControladorError("comandoOs-api-servicios.go", "ActualizarComandosOs_api_servicio()", "El comando que intenta actualizar ya contiene un resultado de una ejecución previa, es necesario refrescar el objeto.")
		return error1
	}
}

func (cOs *ComandoOs_api_servicios) Estado() bool {
	if !cOs.cOs_api.getEstado() && cOs.cOs_api.getSalida() == "" {
		return true
	}
	return false
}
