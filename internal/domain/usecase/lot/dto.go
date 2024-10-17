package lot_usecase

type LotViewDto struct {
	Description     string `json:"description,omitempty"`
	Address         string `json:"address,omitempty"`
	CadastreNumber  string `json:"cadastre_number,omitempty"`
	Square          int    `json:"square,omitempty"`
	DocURL          string `json:"doc_url,omitempty"`
	PublicationDate string `json:"publication_date"`
}
