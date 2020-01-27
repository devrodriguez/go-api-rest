package models

type Check struct {
	Address   string  `firestore: "address"`
	Company   string  `firestore: "company"`
	CompanyId float32 `firestore: "company_id"`
	Date      string  `firestore: "date"`
	Email     string  `firestore: "email"`
	Hour      string  `firestore: "hour"`
	Latitude  string  `firestore: "latitude"`
	Longitude float32 `firestore: "longitude"`
	Type      string  `firestore: "type"`
}
