package nucleo_api

import "github.com/rdcarranza/s1-bot-telegram-go/src/nucleo/servicios"

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
