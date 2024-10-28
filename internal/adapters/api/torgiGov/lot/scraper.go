package torgiGov

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ivzakom/web-scraping-practice/internal/apperror"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
	"time"
)

type noticeData struct {
	Content []struct {
		Id           string    `json:"id"`
		PublishDate  time.Time `json:"publishDate"`
		NoticeStatus string    `json:"noticeStatus"`
		BiddForm     struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"biddForm"`
		BiddType struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"biddType"`
		SignedData struct {
			FileId      string    `json:"fileId"`
			SignatureId string    `json:"signatureId"`
			FileName    string    `json:"fileName"`
			FileSize    int       `json:"fileSize"`
			UploadDate  time.Time `json:"uploadDate"`
			Hash        string    `json:"hash"`
		} `json:"signedData"`
		NoticeNumber  string `json:"noticeNumber"`
		Version       int    `json:"version"`
		ProcedureName string `json:"procedureName"`
		BidderOrg     struct {
			Name string `json:"name"`
		} `json:"bidderOrg"`
		BiddStartTime    time.Time `json:"biddStartTime"`
		BiddEndTime      time.Time `json:"biddEndTime"`
		BiddReviewDate   time.Time `json:"biddReviewDate"`
		AuctionStartDate time.Time `json:"auctionStartDate"`
		EtpCode          string    `json:"etpCode,omitempty"`
		Lots             []struct {
			LotNumber  int     `json:"lotNumber"`
			LotStatus  string  `json:"lotStatus"`
			LotName    string  `json:"lotName"`
			PriceMin   float64 `json:"priceMin,omitempty"`
			HasAppeals bool    `json:"hasAppeals"`
			IsStopped  bool    `json:"isStopped"`
			Attributes []struct {
				Code          string      `json:"code"`
				FullName      string      `json:"fullName"`
				Value         interface{} `json:"value,omitempty"`
				AttributeType string      `json:"attributeType"`
				Group         struct {
					Code             string `json:"code"`
					Name             string `json:"name"`
					DisplayGroupType string `json:"displayGroupType"`
				} `json:"group"`
				SortOrder int `json:"sortOrder"`
			} `json:"attributes"`
			IsAnnulled bool `json:"isAnnulled"`
			LotVat     struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"lotVat,omitempty"`
		} `json:"lots"`
		TimezoneOffset              string    `json:"timezoneOffset"`
		TimezoneOffsetAbberviation  string    `json:"timezoneOffsetAbberviation"`
		HasAppeals                  bool      `json:"hasAppeals"`
		FirstVersionPublicationDate time.Time `json:"firstVersionPublicationDate"`
		Attributes                  []struct {
			Code          string `json:"code"`
			FullName      string `json:"fullName"`
			AttributeType string `json:"attributeType"`
			Group         struct {
				Code             string `json:"code"`
				Name             string `json:"name"`
				DisplayGroupType string `json:"displayGroupType"`
			} `json:"group"`
			SortOrder int         `json:"sortOrder"`
			Value     interface{} `json:"value,omitempty"`
		} `json:"attributes"`
		IsAnnulled  bool   `json:"isAnnulled"`
		NpaHintCode string `json:"npaHintCode"`
	} `json:"content"`
	Pageable struct {
		Sort struct {
			Unsorted bool `json:"unsorted"`
			Sorted   bool `json:"sorted"`
			Empty    bool `json:"empty"`
		} `json:"sort"`
		PageNumber int  `json:"pageNumber"`
		PageSize   int  `json:"pageSize"`
		Offset     int  `json:"offset"`
		Paged      bool `json:"paged"`
		Unpaged    bool `json:"unpaged"`
	} `json:"pageable"`
	CategoryFacet []struct {
		Id    string `json:"_id"`
		Count int    `json:"count"`
	} `json:"categoryFacet"`
	TotalPages       int  `json:"totalPages"`
	TotalElements    int  `json:"totalElements"`
	Last             bool `json:"last"`
	NumberOfElements int  `json:"numberOfElements"`
	First            bool `json:"first"`
	Size             int  `json:"size"`
	Number           int  `json:"number"`
	Sort             struct {
		Unsorted bool `json:"unsorted"`
		Sorted   bool `json:"sorted"`
		Empty    bool `json:"empty"`
	} `json:"sort"`
	Empty bool `json:"empty"`
}

type torgiGovScraper struct {
	BaseNoticesURL     string
	BaseNoticesParam   map[string]string
	BaseNoticesViewURL string
	BaseLotcardsURL    string
	BaseLotcardsParam  map[string]string
}

func NewTorgiGovScraper() *torgiGovScraper {

	noticesParams := map[string]string{
		"sd":             "sds",
		"subjRF":         "39",
		"byFirstVersion": "true",
		"withFacets":     "false",
		"size":           "10",
		"sort":           "firstVersionPublicationDate,desc",
		"catCode":        "2",
	}

	lotcardsParams := map[string]string{
		"dynSubjRF":      "43",
		"lotStatus":      "PUBLISHED, APPLICATIONS_SUBMISSION",
		"catCode":        "2",
		"matchPhrase":    "false",
		"byFirstVersion": "true",
		"withFacets":     "true",
		"size":           "10",
		"sort":           "firstVersionPublicationDate, desc",
	}

	return &torgiGovScraper{
		BaseNoticesURL:     "torgi.gov.ru/new/api/public/notices/search",
		BaseNoticesParam:   noticesParams,
		BaseLotcardsURL:    "torgi.gov.ru/new/api/public/lotcards/search",
		BaseLotcardsParam:  lotcardsParams,
		BaseNoticesViewURL: "https://torgi.gov.ru/new/public/notices/view/",
	}
}

func (s *torgiGovScraper) ScrapNotices(ctx context.Context, params map[string]string) ([]TorgiGovLotDto, error) {

	u, err := url.Parse(fmt.Sprint("https://", s.BaseNoticesURL))
	if err != nil {
		fmt.Println("Ошибка разбора URL:", err)
		return nil, err
	}

	// Добавляем параметры запроса
	queryParams := url.Values{}
	for param, value := range s.BaseNoticesParam {
		queryParams.Add(param, value)
	}
	for param, value := range params {
		queryParams.Add(param, value)
	}

	// Добавляем параметры к URL
	u.RawQuery = queryParams.Encode()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(u.String())
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var data noticeData
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.TotalPages == data.Pageable.PageNumber {
		return nil, apperror.ErrorEOL
	}

	var lotsDto []TorgiGovLotDto
	for _, notice := range data.Content {

		noticeNumber := notice.NoticeNumber
		biddStartTime := notice.BiddStartTime
		biddEndTime := notice.BiddEndTime
		auctionStartDate := notice.AuctionStartDate

		for _, lot := range notice.Lots {

			lotDto := TorgiGovLotDto{}
			copier.Copy(&lotDto, &lot)

			lotDto.NoticeNumber = noticeNumber
			lotDto.BiddStartTime = biddStartTime
			lotDto.BiddEndTime = biddEndTime
			lotDto.AuctionStartDate = auctionStartDate

			lotDto.Url = s.createUrl(noticeNumber, lotDto.LotNumber)

			lotsDto = append(lotsDto, lotDto)
		}

	}

	return lotsDto, nil

}

func (s *torgiGovScraper) createUrl(noticeNumber string, lotNumber int) string {
	return fmt.Sprint(s.BaseNoticesViewURL, noticeNumber, "#", lotNumber)
}

// https://torgi.gov.ru/new/api/public/notices/search? // поиск извещений
//https://torgi.gov.ru/new/api/public/notices/search?byFirstVersion=true&withFacets=true&size=10&sort=firstVersionPublicationDate,desc
//https://torgi.gov.ru/new/api/public/notices/search?catCode=2&byFirstVersion=true&withFacets=true&size=10&sort=firstVersionPublicationDate,desc

// https://torgi.gov.ru/new/public/notices/view/22000054080000000226#2 // карточка извещения url

// https://torgi.gov.ru/new/api/public/lotcards/search?dynSubjRF=43&lotStatus=PUBLISHED,APPLICATIONS_SUBMISSION&matchPhrase=false&byFirstVersion=true&withFacets=true&size=10&sort=firstVersionPublicationDate,desc
// https://torgi.gov.ru/new/api/public/lotcards/search?dynSubjRF=43&lotStatus=PUBLISHED,APPLICATIONS_SUBMISSION&matchPhrase=false&byFirstVersion=true&withFacets=false&page=1&size=10&sort=firstVersionPublicationDate,desc
// https://torgi.gov.ru/new/public/lots/lot/21000033490000000304_1/(lotInfo:info)?fromRec=false // лот
