package nucleo_api

type ComandoOs_api_servicios struct {
	comandoOs_api ComandoOs_api
}

func NuevoComandoOs_api_servicio(comandoOs_api ComandoOs_api) *ComandoOsServ {
	return &ComandoOs_api_servicios{
		comandoOs_api: comandoOs_api,
	}
}
