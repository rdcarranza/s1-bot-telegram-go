package nucleo_api

import (
	controlador_error "github.com/rdcarranza/s1-bot-telegram-go/src/controladores/error"
	"github.com/rdcarranza/s1-bot-telegram-go/src/nucleo/servicios"
)

type ComandoOs_api_servicios struct {
	cOs_api ComandoOs_api
}

func NuevoComandoOs_api_servicio(comandoOs_api ComandoOs_api, c string) *ComandoOs_api_servicios {
	comandoOs_api.NuevoComandoOs(c)
	return &ComandoOs_api_servicios{cOs_api: comandoOs_api}
}

func (cOs *ComandoOs_api_servicios) EjecutarComandoOs_api_servicio(comandoOs_api ComandoOs_api) (string, error) {
	servicio := &servicios.ComandoOsServ{}
	return servicio.EjecutarComando(comandoOs_api.comando)
}

func (cOs *ComandoOs_api_servicios) ActualizarComandoOs_api_servicio(resultado string) (bool, error) {
	if !cOs.cOs_api.getEstado() {
		cOs.cOs_api.ActualizarComandoOs(resultado)
		return true, nil
	} else {
		error1 := controlador_error.NuevoControladorError()
		error1.CargarControladorError("comandoOs-api-servicios.go", "ActualizarComandosOs_api_servicio()", "El comando que intenta actualizar ya contiene un resultado de una ejecución previa, es necesario refrescar el objeto.")
		return false, error1
	}
}
