package torgiGov

import "time"

type TorgiGovLotDto struct {
	LotNumber        int       `json:"lotNumber"`
	LotStatus        string    `json:"lotStatus"`
	LotName          string    `json:"lotName"`
	PriceMin         float64   `json:"priceMin"`
	HasAppeals       bool      `json:"hasAppeals"`
	IsStopped        bool      `json:"isStopped"`
	BiddStartTime    time.Time `json:"biddStartTime"`
	BiddEndTime      time.Time `json:"biddEndTime"`
	AuctionStartDate time.Time `json:"auctionStartDate"`
	NoticeNumber     string    `json:"noticeNumber"`
	PublishDate      time.Time `json:"publishDate"`
	IsAnnulled       bool      `json:"isAnnulled"`
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
