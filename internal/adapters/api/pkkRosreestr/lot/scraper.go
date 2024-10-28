package pkkRosreestr

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type pkkRosreestrScraper struct {
	BaseURL   string
	BaseParam string
}

func NewPkkRosreestrGovScraper() *pkkRosreestrScraper {
	return &pkkRosreestrScraper{
		BaseURL: "pkk.rosreestr.ru/api/features/",
	}
}

func (s *pkkRosreestrScraper) GetLocationPoint(Decription string) (PkkRosreestrLotDto, error) {

	var (
		lotDto PkkRosreestrLotDto
		err    error
	)

	var CadastreNumber string

	rePlot := regexp.MustCompile(`\d{2}:\d{2}:\d{6,7}:\d{1,5}`)
	reQuarter := regexp.MustCompile(`\d{2}:\d{2}:\d{6,7}`)
	if matches := rePlot.FindStringSubmatch(Decription); len(matches) > 0 {
		CadastreNumber = matches[0]
	} else if matches := reQuarter.FindStringSubmatch(Decription); len(matches) > 0 {
		CadastreNumber = matches[0]
	}

	if CadastreNumber != "" {
		lotDto, err = s.getDataByCadastreNumber(context.Background(), CadastreNumber)
		lotDto.CadastreNumber = CadastreNumber
		if err != nil {
			return PkkRosreestrLotDto{}, err
		}
	}

	return lotDto, err

}

func (s *pkkRosreestrScraper) getDataByCadastreNumber(ctx context.Context, CadastreNumber string) (PkkRosreestrLotDto, error) {

	pkkCadastreNumber := normalizeCadastreNumber(CadastreNumber)
	cadasterCode := cadasterCode(pkkCadastreNumber)
	url := fmt.Sprint("https://", s.BaseURL, cadasterCode, "?text=", pkkCadastreNumber)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return PkkRosreestrLotDto{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PkkRosreestrLotDto{}, errors.New(resp.Status)
	}

	var lotDto PkkRosreestrLotDto
	if err = json.NewDecoder(resp.Body).Decode(&lotDto); err != nil {
		return PkkRosreestrLotDto{}, err
	}

	return lotDto, nil
}

func normalizeCadastreNumber(CadastreNumber string) string {

	var result []string

	NumParts := strings.Split(CadastreNumber, ":")
	for _, part := range NumParts {

		result = append(result, strings.TrimLeft(part, "0"))

	}

	return strings.Join(result, ":")
}

func cadasterCode(CadastreNumber string) (code string) {

	NumParts := len(strings.Split(CadastreNumber, ":"))
	switch NumParts {
	case 4:
		code = "1"
	case 3:
		code = "2"
	case 2:
		code = "3"
	case 1:
		code = "4"
	}

	return

}
