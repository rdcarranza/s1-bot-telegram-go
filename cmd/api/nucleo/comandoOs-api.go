package nucleo_api

type ComandoOs_api struct {
	comando string
	salida  string
	estado  bool
}

// constructor
func NuevoComandoOs(c string) *ComandoOs_api {
	return &ComandoOs_api{
		comando: c,
		salida:  "",
		estado:  false,
	}
}

func (cos *ComandoOs_api) ActualizarComandoOs(s string) {
	cos.salida = s
	cos.estado = true
}

func (cos *ComandoOs_api) ResultadoComandoOs() string {
	return cos.salida
}

func (cos *ComandoOs_api) getEstado() bool {
	return cos.estado
}

func (cos *ComandoOs_api) getSalida() string {
	return cos.salida
}
