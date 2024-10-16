package entity

import "time"

type Lot struct {
	ID              string    `json:"id" bson:"id"`
	Num             int       `json:"num" bson:"num"`
	Description     string    `json:"description,omitempty" bson:"description"`
	Address         string    `json:"address,omitempty" bson:"address"`
	CadastreNumber  string    `json:"cadastre_number,omitempty" bson:"cadastreNumber"`
	Square          int       `json:"square,omitempty" bson:"square"`
	DocURL          string    `json:"doc_url,omitempty" bson:"docURL"`
	PublicationDate time.Time `json:"publication_date" bson:"publicationDate"`
}

type LotView struct {
	ID             string `json:"id"`
	Description    string `json:"description,omitempty"`
	Address        string `json:"address,omitempty"`
	CadastreNumber string `json:"cadastre_number,omitempty"`
	Square         int    `json:"square,omitempty"`
	DocURL         string `json:"doc_url,omitempty"`
}
