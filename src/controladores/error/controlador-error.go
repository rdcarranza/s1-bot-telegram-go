package controlador_error

import "fmt"

// Error controlado
type ErrorC struct {
	archivo     string
	modulo      string
	descripcion string
	mensaje     string
}

// implementación de interfaz
func controlador_Error() error {
	return &ErrorC{}
}

// constructor

func NuevoControladorError() *ErrorC {
	return controlador_Error().(*ErrorC)
}

func (c *ErrorC) CargarControladorError(a string, mod string, desc string) *ErrorC {
	c.archivo = a
	c.modulo = mod
	c.descripcion = desc
	return controlador_Error().(*ErrorC)
}

func (c *ErrorC) Error() string {
	c.mensaje = fmt.Sprintf("Error en el archivo: %s,  módulo -> %s. Descripción: %s", c.archivo, c.modulo, c.descripcion)
	return c.mensaje
}
