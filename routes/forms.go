package routes

import (
	"fmt"
	"go-server-seccion-8/utils"
	"go-server-seccion-8/validaciones"
	"net/http"
)

func Formularios_get(response http.ResponseWriter, request *http.Request) {

	css_sesion, css_mensaje := utils.RetornarMensajesFlash(response, request)

	data := map[string]string{
		"css":     css_sesion,
		"mensaje": css_mensaje,
	}
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

	if validaciones.ValidarPassword(request.FormValue("password")) == false {
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
