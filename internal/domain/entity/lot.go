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
	RosreestrData   struct {
		Total         int    `json:"total"`
		TotalRelation string `json:"total_relation"`
		Features      []struct {
			Center struct {
				Y float64 `json:"y"`
				X float64 `json:"x"`
			} `json:"center"`
			Extent struct {
				Xmin float64 `json:"xmin"`
				Xmax float64 `json:"xmax"`
				Ymax float64 `json:"ymax"`
				Ymin float64 `json:"ymin"`
			} `json:"extent"`
			Sort  int64 `json:"sort"`
			Type  int   `json:"type"`
			Attrs struct {
				Address      string `json:"address"`
				CategoryType string `json:"category_type"`
				Cn           string `json:"cn"`
				Id           string `json:"id"`
			} `json:"attrs"`
		} `json:"features"`
	}
}

type LotView struct {
	Description     string    `json:"description,omitempty"`
	Address         string    `json:"address,omitempty"`
	CadastreNumber  string    `json:"cadastre_number,omitempty"`
	Square          int       `json:"square,omitempty"`
	DocURL          string    `json:"doc_url,omitempty"`
	PublicationDate time.Time `json:"publication_date" bson:"publicationDate"`
}
