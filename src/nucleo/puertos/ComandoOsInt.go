package puertos

// ComandoOS define la interfaz para ejecutar comandos del sistema.

type ComandoOS interface {
	EjecutarComando(comando string) (string, error)
}
