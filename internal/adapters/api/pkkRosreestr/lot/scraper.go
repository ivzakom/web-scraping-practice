package pkkRosreestrScraper

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func (s *pkkRosreestrScraper) Scrap(ctx context.Context, CadastreNumber string) (PkkRosreestrLotDto, error) {

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
