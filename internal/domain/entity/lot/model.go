package lot

import "time"

type Lot struct {
	UUID            string    `json:"uuid"`
	Description     string    `json:"description,omitempty"`
	Address         string    `json:"address,omitempty"`
	CadastreNumber  string    `json:"cadastre_number,omitempty"`
	Square          int       `json:"square,omitempty"`
	DocURL          string    `json:"doc_url,omitempty"`
	PublicationDate time.Time `json:"publication_date"`
}
