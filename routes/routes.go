package routes

import (
	"go-server-seccion-8/utils"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func LoadHtmlMust(path string, response http.ResponseWriter, data any) *template.Template {
	template := template.Must(template.ParseFiles(path, utils.Frontend))
	template.Execute(response, data)

	return template
}

func Home(response http.ResponseWriter, request *http.Request) {

	LoadHtmlMust("templates/ejemplo/home.html", response, nil)
	// template, err := template.ParseFiles("templates/ejemplo/home.html", "templates/layout/frontend.html")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	template.Execute(response, nil)
	// }

}

// paso data generando un struct o interface en next
// // y luego dandole valor a los datos
type Datos struct {
	Nombre string
	Edad   int
	Perfil int
	// es un slice no un arreglo
	Habilidades []Habilidad
}

type Habilidad struct {
	Nombre string
}

func Estructuras(response http.ResponseWriter, request *http.Request) {

	habilidad1 := Habilidad{"Inteligencia"}
	habilidad2 := Habilidad{"VideoJuegos"}
	habilidad3 := Habilidad{"Programacion"}
	habilidad4 := Habilidad{"Canto"}

	habilidades := []Habilidad{habilidad1, habilidad2, habilidad3, habilidad4}

	LoadHtmlMust("templates/estructuras/estructuras.html", response, Datos{"Sergio", 11, 1, habilidades})

	// template, err := template.ParseFiles("templates/estructuras/estructuras.html")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	template.Execute(response, Datos{"Sergio", 11, 1})
	// }

}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	LoadHtmlMust("templates/nosotros/nosotros.html", response, nil)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	template.Execute(response, nil)
	// }

}
func Pagina404(response http.ResponseWriter, request *http.Request) {
	LoadHtmlMust("templates/pagina404/pagina404.html", response, nil)

}

func Parametros(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	// fmt.Fprintln(response, "id= "+vars["id"]+" | slug= "+vars["slug"]+"")

	texto := "mi muñeca me hablo"
	data := map[string]string{
		"id":    vars["id"],
		"slug":  vars["slug"],
		"texto": texto,
	}
	LoadHtmlMust("templates/parametros/parametros.html", response, data)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	template.Execute(response, data)
	// }

}

func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {

	slug := request.URL.Query().Get("slug")
	id := request.URL.Query().Get("id")

	data := map[string]string{
		"slug": slug,
		"id":   id,
	}

	LoadHtmlMust("templates/parametros-qs/parametros-qs.html", response, data)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	template.Execute(response, data)
	// }

}

// func Nosotros(response http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintln(response, "sobre nosotros aererrrrrrrr")

// }

// func Parametros(response http.ResponseWriter, request *http.Request) {
// 	vars := mux.Vars(request)

// 	fmt.Fprintln(response, "id= "+vars["id"]+" | slug= "+vars["slug"]+"")
// }
// func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {

// 	// esta es la query http://127.0.0.1:8080/parametros-querystring?aer=1&ooo=1
// 	fmt.Fprintln(response, request.URL) // esta rte muetra la url completa
// 	fmt.Fprintln(response, request.URL.RawQuery) // esta solo la query ?aer=1&ooo=1
// 	fmt.Fprintln(response, request.URL.Query()) // este te devuelve la query como un mapa map[aer:[1] ooo:[1]]
// 	fmt.Fprintln(response, request.URL.Query().Get("aer")) //este devuelve solo 1 el valor de aer realmente

// }
