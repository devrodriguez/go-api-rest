package routes

import (
	"log"
	"time"
	"net/http"
	"github.com/gbrlsnchs/jwt"

	"encoding/json"
	"github.com/devrodriguez/go-api-rest/models"
)

// SignIn retorna un token de autenticacion
func SignIn(res http.ResponseWriter, req *http.Request) {
	log.Println("Sign in...")
	now := time.Now()
 
	var hs = jwt.NewHS256([]byte("dev1986"))

	payload := models.JwtPayload{
		Payload: jwt.Payload{
			Issuer: "devrodriguez",
			Subject: "dev",
			Audience: jwt.Audience{"http://localhost:3000"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "foobar",
		},
		Foo: "foo",
		Bar: 1337,
	} 

	token, err := jwt.Sign(payload, hs)

	log.Println(string(token))

	if err != nil {
		log.Fatal(err)
		json.NewEncoder(res).Encode("Autenticacion fallida")
	}

	json.NewEncoder(res).Encode(string(token))
}

// Login valida el token de Authorization
func Login(res http.ResponseWriter, req *http.Request) {
	var secret = jwt.NewHS256([]byte("dev1986"))
	var payload models.JwtPayload
	var response models.Response

	token := []byte(req.Header.Get("Authorization"))

	hd, err := jwt.Verify(token, secret, &payload)

	log.Println(hd)
	log.Println(err)
	

	if err != nil {
		response.Message = "¡Usuario no autorizado!"
		dataRes, _ := json.Marshal(response)
		
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(string(dataRes))
		return
	}

	response.Message = "Bienvenido a la App"

	dataRes, err := json.Marshal(response)

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(string(dataRes))

}