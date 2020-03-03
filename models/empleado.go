package models

import (
	"fmt"
	"log"
)

// Empleado estructura para definir empleados
type Empleado struct {
	Nombre string
	Telefono string
	Direccion string
}

// Saludo retorna un mensaje de saludo
func (emp Empleado) Saludo(qsaluda string) string {
	var s []string 

	s = append(s, "John Rodriguez")
	s = append(s, "Sandra Triana")
	s = append(s, "Diana Mendez")

	for index, item := range s {
		log.Printf("Index: %d", index)
		log.Printf("Item: %s", item)
	}

	log.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	log.Printf("Metodo desde Empleado")
	return fmt.Sprintf("Hola %s, como estas? Mi nombre es %s", qsaluda, emp.Nombre)
}