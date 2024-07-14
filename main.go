package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go-server-seccion-8/routes"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := mux.NewRouter()
	// rutas
	mux.HandleFunc("/", routes.Home)
	mux.HandleFunc("/nosotros", routes.Nosotros)
	mux.HandleFunc("/parametros/{id:.*}/{slug:.*}", routes.Parametros)
	mux.HandleFunc("/parametros-qs", routes.ParametrosQueryString)
	mux.HandleFunc("/estructuras", routes.Estructuras)
	// formularios rutas
	mux.HandleFunc("/formularios", routes.Formularios_get)
	mux.HandleFunc("/formularios-post", routes.Formularios_post).Methods("POST")

	
	// archivos estaticos hacia mux con esta config reconocera recursos estaticos
	// css, javascript lo que sea
	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	mux.PathPrefix("/public/").Handler(s)
	//error
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(routes.Pagina404).GetHandler()


	errorVariables := godotenv.Load()

	fmt.Println(os.Getenv("PORT"))
	if errorVariables != nil {

		fmt.Println("entro al distntino de nil")
		panic(errorVariables)

	}

	server := &http.Server{
		Addr:         "localhost:" + os.Getenv("PORT"),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {

// 	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
// 		fmt.Fprintln(response, "hola mundo")
// 	})

// 	fmt.Println("corriendo servidor desde http://localhost:8081")
// 	log.Fatal(http.ListenAndServe("localhost:8081", nil))
// }
