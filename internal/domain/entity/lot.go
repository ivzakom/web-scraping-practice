package entity

import "time"

type Lot struct {
	ID              string    `json:"id" bson:"id"`
	Num             int       `json:"num" bson:"num"`
	NoticeNumber    string    `json:"notice_number" bson:"notice_number"`
	NoticeDate      time.Time `json:"notice_date" bson:"notice_date"`
	Description     string    `json:"description,omitempty" bson:"description"`
	Address         string    `json:"address,omitempty" bson:"address"`
	CadastreNumber  string    `json:"cadastre_number,omitempty" bson:"cadastreNumber"`
	Square          int       `json:"square,omitempty" bson:"square"`
	DocURL          string    `json:"doc_url,omitempty" bson:"docURL"`
	PublicationDate time.Time `json:"publication_date" bson:"publicationDate"`
	Price           float64   `json:"price,omitempty" bson:"price"`
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
	TorgiGovData struct {
		LotNumber        int       `json:"lotNumber"`
		LotStatus        string    `json:"lotStatus"`
		LotName          string    `json:"lotName"`
		PriceMin         float64   `json:"priceMin"`
		BiddStartTime    time.Time `json:"biddStartTime"`
		BiddEndTime      time.Time `json:"biddEndTime"`
		AuctionStartDate time.Time `json:"auctionStartDate"`
		NoticeNumber     string    `json:"noticeNumber"`
		Url              string    `json:"url"`
		Attributes       []struct {
			Code          string `json:"code"`
			FullName      string `json:"fullName"`
			Value         string `json:"value"`
			AttributeType string `json:"attributeType"`
			Group         struct {
				Code             string `json:"code"`
				Name             string `json:"name"`
				DisplayGroupType string `json:"displayGroupType"`
			} `json:"group"`
			SortOrder int `json:"sortOrder"`
		} `json:"attributes"`
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
