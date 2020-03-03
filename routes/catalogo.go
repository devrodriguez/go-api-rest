package routes

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

// GetParms obtiene parametros del catalogo
func GetParms(res http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("http://localhost:5012/catalogo/v1/parms/services.credito.consultarTCV?canal=48&modulo=API&lenguaje=ES&pais=CO&kind=api")
	if err != nil {
		log.Fatalln(err)
		json.NewEncoder(res).Encode(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	in := []byte(string(body))

	var raw map[string]interface{}

	if err := json.Unmarshal(in, &raw); err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatalln(err)
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(raw)
}