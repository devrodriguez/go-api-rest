package models

type Check struct {
	Address string `firestore:"address"`
	Company string `firestore:"company"`
	Date    string `firestore:"date"`
	Email   string `firestore:"email"`
	Hour    string `firestore:"hour"`
	Type    string `firestore:"type"`
	//Latitude  string  `firestore: "latitude"`
	//Longitude float32 `firestore: "longitude"`
	//CompanyId float32 `firestore: "company_id"`
}
