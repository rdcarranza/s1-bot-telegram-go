package nucleo_api

type ComandoOs_api struct {
	comando string
	salida  string
	estado  bool
}

// constructor
func (cos *ComandoOs_api) NuevoComandoOs(c string) {
	cos.comando = c
	cos.salida = ""
	cos.estado = false
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
