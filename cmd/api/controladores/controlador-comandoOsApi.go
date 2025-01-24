package controladores_api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	nucleo_api "github.com/rdcarranza/s1-bot-telegram-go/cmd/api/nucleo"
	controlador_error "github.com/rdcarranza/s1-bot-telegram-go/src/controladores/error"
	"github.com/rdcarranza/s1-bot-telegram-go/src/nucleo/servicios"
)

type Ctrl_comandosOsApi struct {
	cOsApiServ nucleo_api.ComandoOs_api_servicios
}

func Controlador_comandosOsApi(ingreso string) *Ctrl_comandosOsApi {
	if listaComandoVacia() {
		cargarListaComando()
	}
	com := isComando(ingreso)
	fmt.Println("comando: ", com)
	if com != "" {
		cas := nucleo_api.NuevoComandoOs_api_servicio(com)
		return &Ctrl_comandosOsApi{cOsApiServ: *cas}
	}
	return nil
}

func Controlador_comandosOsApi_comando(ingreso string, remitente string) error {
	ing := strings.Split(ingreso, " ")
	if strings.HasPrefix(ing[0], "/") {
		com := strings.ReplaceAll(ing[0], "/", "")
		param := strings.ReplaceAll(ingreso, ing[0], "")
		if remitente != "" {
			param = "Mensaje de: " + remitente + " - " + param
		}

		err := selectorComando(com, param)
		if err != nil {
			return err
		}

		return nil

	} else {
		cerror := controlador_error.NuevoControladorError()
		err := cerror.CargarControladorError("controlador-comandosOsApi", "Controlador_comandosOsApi_comando", "El ingreso no contiene un comando.")
		return err
	}
}

func Controlador_comandosOsApi_voz(url string) error {
	// Abre un archivo local para guardar la voz
	outFile, err := os.Create("voice.ogg")
	if err != nil {
		log.Println("Error al crear archivo local: ", err)
		return err
	}
	defer outFile.Close()

	// Descarga el archivo
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error al descargar el archivo de voz: ", err)
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Println("Error al guardando el archivo de voz: ", err)
		return err
	}

	//exec.Command("ffplay", "-nodisp", "-autoexit", "-volume", "192", "voice.ogg").Run()
	//exec.Command("cvlc", "--intf", "dummy", "--volume", "384", "--play-and-exit", "voice.ogg").Run()
	servicio := &servicios.ComandoOsServ{}
	servicio.ReproducirOggFfplay("voice.ogg")
	fmt.Println("Reproduciendo mensaje de voz!")
	/*
		f,_:=b.GetFile(fileID)
		fmt.Println("URL de prueba: "+f.FilePath)
	*/
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

func listaComandoVacia() bool {
	return len(ListaComandos) == 0
}

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
	return com
}

func selectorComando(comando string, parametros string) error {
	servicio := &servicios.ComandoOsServ{}
	switch comando {
	case "leer":
		return servicio.LecturaTexto(parametros)
	default:
		cer := controlador_error.NuevoControladorError()
		return cer.CargarControladorError("controlador-comandosOsApi", "selectorComando", "El comando ingresado no es reconocido.")
	}

}
