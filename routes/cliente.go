package routes

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/devrodriguez/go-api-rest/models"
	"github.com/devrodriguez/go-api-rest/interfaces"
)

var clientes []models.Cliente

func GetClientes(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(clientes)
}

func CreateCliente(res http.ResponseWriter, req *http.Request) {

	var cliente models.Cliente

	// params := mux.Vars(req)

	// log.Println(params)

	// cliente.Name = params["name"]
	// cliente.Score = params["score"]

	json.NewDecoder(req.Body).Decode(&cliente)

	clientes = append(clientes, cliente)

	json.NewEncoder(res).Encode(clientes)
}

func SaludoCliente(res http.ResponseWriter, req *http.Request) {
	cliente := models.Cliente{Name: "Camilo", Score: "100"}
	var empleado models.Empleado = models.Empleado{Nombre: "Jose"}

	saludo1 := llamarSaludo(cliente)
	saludo2 := llamarSaludo(empleado)

	json.NewEncoder(res).Encode(fmt.Sprintf("%s - %s", saludo1, saludo2))
}

func llamarSaludo(p interfaces.IPerson, ) string {
	return p.Saludo("John")
}