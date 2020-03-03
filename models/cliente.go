package models

import "fmt"

// Cliente es una estructura que tiene Name y Score
type Cliente struct {
	Name string
	Score string
}

// Saludo retorna un mensaje de saludo
func (cli Cliente) Saludo(saluda string) string {
	return fmt.Sprintf("Hola %s, soy %s estoy para servirte! Mi score de atencion es de %s", saluda, cli.Name, cli.Score)
}