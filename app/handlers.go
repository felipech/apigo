package app

import (
	"apigo/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

type Customer struct {
	Nombre string `json:"Nombre_Completo"`
	Ciudad string
}

type CustomerHandlers struct {
	service service.CustomerService
}

type Hoaria struct {
	Current string `json:"current_time"`
}

func saludar(w http.ResponseWriter, request *http.Request) {
	fmt.Println("saludando por consola")
}

func listarTodos(w http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{Nombre: "Felipe", Ciudad: "Santiago"},
	}

	w.Header().Add("Content-Type", "application-json")
	//para reponder en xml
	//xml.NewEncoder(w).Encode(customers)
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, request *http.Request) {
	/*	customers := []Customer{
		{Nombre: "Felipe", Ciudad: "Santiago"},
	}*/
	customers, _ := ch.service.GetAllCustomer()
	w.Header().Add("Content-Type", "application-json")
	//para reponder en xml
	//xml.NewEncoder(w).Encode(customers)
	json.NewEncoder(w).Encode(customers)
}

func obtenerCustomer(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprint(w, vars["customer_id"])
}

func zonaHorariaCustom(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application-json")
	s := request.FormValue("tz")
	if s == "" {
		//solicitando hoario local
		location, _ := time.LoadLocation("America/Santiago")
		hora := time.Now().UTC().In(location)
		zonaHoraira := Hoaria{Current: hora.String()}
		json.NewEncoder(writer).Encode(zonaHoraira)
	} else {
		m := make(map[string]string)
		fmt.Println(s)
		arreglo1 := strings.Split(s, ",")
		//var resultado string
		for _, v := range arreglo1 {
			fmt.Println(v)
			location1, _ := time.LoadLocation(v)
			hora := time.Now().UTC().In(location1)
			m[v] = hora.String()
		}
		//tengo el arreglo de zonas horarias
		re, _ := json.Marshal(m)
		fmt.Fprint(writer, string(re))
	}

}
