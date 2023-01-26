package app

import (
	"apigo/domain"
	"apigo/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	//router := http.NewServeMux()
	router := mux.NewRouter()

	//enlace

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	//aplicacion
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	/*ejemplos*/
	router.HandleFunc("/saludar", saludar).Methods(http.MethodGet)
	router.HandleFunc("/listar", listarTodos)
	router.HandleFunc("/customer/{customer_id}", obtenerCustomer)
	//router.HandleFunc("/api/time", zonaHorariaCustom)
	router.HandleFunc("/api/time", zonaHorariaCustom)
	router.Path("/api/time").Queries("tz", "{[0-9]*?}").HandlerFunc(zonaHorariaCustom).Name("zonaHorariaCustom")
	//router.Path("/api/time/tz").HandlerFunc(zonasHorarias)

	log.Fatal(http.ListenAndServe("localhost:8081", router))
}
