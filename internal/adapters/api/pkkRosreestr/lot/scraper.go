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
		BaseURL: "pkk.rosreestr.ru/api/features/1?text=",
	}
}

func (s *pkkRosreestrScraper) GetLocationPoint(Decription string) (PkkRosreestrLotDto, error) {

	var (
		lotDto PkkRosreestrLotDto
		err    error
	)

	re := regexp.MustCompile(`КН\s*\d{2}:\d{2}:\d{6,7}:\d+`)
	if matches := re.FindStringSubmatch(Decription); len(matches) > 0 {
		CadastreNumber := matches[0][5:]
		lotDto.CadastreNumber = CadastreNumber

		lotDto, err = s.getDataByCadastreNumber(context.Background(), CadastreNumber)
		if err != nil {
			return PkkRosreestrLotDto{}, err
		}

	}

	return lotDto, err

}

func (s *pkkRosreestrScraper) getDataByCadastreNumber(ctx context.Context, CadastreNumber string) (PkkRosreestrLotDto, error) {

	pkkCadastreNumber := normalizeCadastreNumber(CadastreNumber)
	url := fmt.Sprint("https://", s.BaseURL, pkkCadastreNumber)

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
