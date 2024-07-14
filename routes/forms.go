package routes

import (
	"fmt"
	"go-server-seccion-8/utils"
	"go-server-seccion-8/validaciones"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func Formularios_get(response http.ResponseWriter, request *http.Request) {

	css_sesion, css_mensaje := utils.RetornarMensajesFlash(response, request)

	data := map[string]string{
		"css":     css_sesion,
		"mensaje": css_mensaje,
	}
	// aqui le envio al template las variables css y mnensaje en le mapa data
	LoadHtmlMust("templates/formularios/formularios.html", response, data)
	// template, err := template.ParseFiles("templates/ejemplo/home.html", "templates/layout/frontend.html")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	template.Execute(response, nil)
	// }

}

// func Formularios_post(response http.ResponseWriter, request *http.Request) {

// 	// lo que diga en el atributo name es lo que traera en el formulario
// 	fmt.Fprintln(response, request.FormValue("nombre"))
// }

func Formularios_post(response http.ResponseWriter, request *http.Request) {

	mensaje := ""

	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + " El campo nombre esta vacío"
	}
	if len(request.FormValue("correo")) == 0 {
		mensaje = mensaje + " El campo email esta vacío"
	}

	if validaciones.Regex_correo.FindStringSubmatch(request.FormValue("correo")) == nil {
		mensaje = mensaje + " . El Email ingresado no es valido"
	}

	if !validaciones.ValidarPassword(request.FormValue("password")) {
		mensaje = mensaje + " . La contraseña debe tener al menos 1 número, una mayúscula, y un largo entre 6 y 8 caracteres"

	}
	if mensaje != "" {
		utils.CrearMensajeFlash(response, request, "danger", mensaje)
		http.Redirect(response, request, "/formularios", http.StatusSeeOther)
		return
	}

	fmt.Fprintln(response,
		" Nombre:"+request.FormValue("nombre")+
			" | Email:  "+request.FormValue("correo")+
			" | Teléfono: "+request.FormValue("telefono")+
			" | Password: "+request.FormValue("password"))
}

func Formularios_upload(response http.ResponseWriter, request *http.Request) {
	css_sesion, css_mensaje := utils.RetornarMensajesFlash(response, request)

	data := map[string]string{
		"css":     css_sesion,
		"mensaje": css_mensaje,
	}
	// aqui le envio al template las variables css y mnensaje en le mapa data
	LoadHtmlMust("templates/formularios/upload.html", response, data)
}

// como sabe el formulario a donde apuntar por que en el action se declara a donde debe apuntar
// a su vez en las rutas esta declarado a que apunta cada ruta
func Formularios_upload_post(response http.ResponseWriter, request *http.Request) {

	file, handler, err := request.FormFile("foto")

	if err != nil {
		utils.CrearMensajeFlash(response, request, "danger", "Ocurrio un error inesperado")

	}

	// tal como  dice la variable esta es la extension
	var extension = strings.Split(handler.Filename, ".")[1]
	// convertir time a string
	time := strings.Split(time.Now().String(), " ")
	// crear el nombre del archivo
	foto := string(time[4][6:14]) + "." + extension
	// rutear archivo
	var archivo string = "public/uploads/fotos/" + foto
	// guardar archivo
	f, errCopy := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0777)

	if errCopy != nil {
		utils.CrearMensajeFlash(response, request, "danger", "Ocurrio un error inesperado")

	}

	// aca lo guardariamos en la bd

	_, errCopiar := io.Copy(f, file)

	if errCopiar != nil {
		utils.CrearMensajeFlash(response, request, "danger", "Ocurrio un error inesperado")

	}

	utils.CrearMensajeFlash(response, request, "success", "Se subió el archivo "+foto+" exitósamente")
	http.Redirect(response, request, "/formularios/upload", http.StatusFound)

}

// flujo mensajes flash

// 1. instalar go get github.com/gorilla/sessions
// 2. crear dos metodos en utils.go crearmensaje y retornarmensaje
// 3. en forms.go en las funciones formularios_post llamar a las validaciones previamente creadas y
// 4. asignar un mensaje si no las pasa
// 5. luego poner la condicion si el mensaje es distinto de vacio quiere decir que hay un error
// 6. ahi es cuando llamammos a lafuncion de crearmensaje de utils.go y pasarle el mensaje creado a la
// mas un redirect al mismo formulario
// 7. en la funcion formularios_get que da la ruta hacia el formulario le pasamos la data de un mapa
// creado a partir de la funcion retornamensaje  de utils.go la cual va a buscar las cookies y obtiene el mensaje creado
// equi se le pasan como data al html y en htmo se pregunta si el css (que deberia ser danger) viene distinto de vacio
// mostrar el popup con el mensaje y ademas añadir a la clase de bootstrap este mismo css asi alert-{{css}}
