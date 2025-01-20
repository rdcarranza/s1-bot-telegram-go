package controladores_api

import (
	"fmt"
	"strings"

	nucleo_api "github.com/rdcarranza/s1-bot-telegram-go/cmd/api/nucleo"
	controlador_error "github.com/rdcarranza/s1-bot-telegram-go/src/controladores/error"
)

type Ctrl_comandosOsApi struct {
	cOsApiServ nucleo_api.ComandoOs_api_servicios
}

func Controlador_comandosOsApi(ingreso string) *Ctrl_comandosOsApi {
	//if lista vacia()
	cargarListaComando()
	com := isComando(ingreso)
	if com != "" {
		cas := nucleo_api.NuevoComandoOs_api_servicio(com)
		return &Ctrl_comandosOsApi{cOsApiServ: *cas}
	}
	return nil
}

func (contCAS *Ctrl_comandosOsApi) EjecutarComando() (string, error) {

	if !contCAS.cOsApiServ.Estado() {
		error1 := controlador_error.NuevoControladorError()
		error1.CargarControladorError("controlador-comandoOsApi.go", "ejecutarComando()", "El comando que intenta ejecutar ya contiene un resultado de una ejecución previa, es necesario refrescar el objeto.")
		return "", error1
	}

	res, err := contCAS.cOsApiServ.EjecutarComandoOs_api_servicio()
	if err != nil {
		return "", err
	}

	err = contCAS.cOsApiServ.ActualizarComandoOs_api_servicio(res)
	if err != nil {
		return res, err
	}

	return res, nil
}

type nodoComando struct {
	Nombre  string `json:"nombre"`
	Comando string `json:"comando"`
}

var ListaComandos []nodoComando

func cargarListaComando() {
	c1 := nodoComando{"ip", "ip a"}
	c2 := nodoComando{"host", "hostnamectl status"}
	ListaComandos = append(ListaComandos, c1)
	ListaComandos = append(ListaComandos, c2)
	fmt.Println("Lista de Comandos cargada: ", ListaComandos)
}

func buscarComando(nom string) string {
	for _, i := range ListaComandos {
		if i.Nombre == nom {
			return i.Comando
		}
	}
	return ""
}

func isComando(ingreso string) string {
	ing := strings.Trim(ingreso, "")
	fmt.Println("Ingreso limpio: " + ing)
	com := buscarComando(ing)
	if com != "" {
		return com
	} else {
		return ""
	}
}
