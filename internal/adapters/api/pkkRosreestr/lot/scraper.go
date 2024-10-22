package pkkRosreestrScraper

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type pkkRosreestrScraper struct {
	BaseURL   string
	BaseParam string
}

func NewPkkRosreestrGovScraper() *pkkRosreestrScraper {
	return &pkkRosreestrScraper{
		//BaseURL: "rosreestr.gov.ru/api/online/fir_object/",
		BaseURL: "pkk.rosreestr.ru/api/features/5?text=",
	}
}

func (s *pkkRosreestrScraper) Scrap(ctx context.Context, CadastreNumber string) (PkkRosreestrLotDto, error) {

	//url := "https://pkk.rosreestr.ru/api/features/1?text=39:3:90910:91"
	//
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//client := &http.Client{Transport: tr}
	//
	//resp, err := client.Get(url)
	//if err != nil {
	//	fmt.Println("Error making request:", err)
	//	return PkkRosreestrLotDto{}, err
	//}
	//defer resp.Body.Close()
	//
	//// Читаем ответ
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error reading response body:", err)
	//	return PkkRosreestrLotDto{}, err
	//}
	//
	//// Выводим статус и тело ответа
	//fmt.Println("Status Code:", resp.Status)
	//fmt.Println("Response Body:", string(body))

	//https: //rosreestr.gov.ru/api/online/fir_object/2:56:30302:639
	// 'https://pkk.rosreestr.ru/api/features/5?text={cadastral_id}&limit={limit}&tolerance={tolerance}'
	//pkkCadastreNumber := normalizeCadastreNumber(CadastreNumber)
	//url := fmt.Sprint("https://", s.BaseURL, pkkCadastreNumber)
	//
	//client := &http.Client{
	//	Timeout:   time.Second * 10,
	//	Transport: createTransport(),
	//}
	//
	//request, reqErr := pkkRequest(ctx, url)
	//if reqErr != nil {
	//	return PkkRosreestrLotDto{}, reqErr
	//}
	//
	//response, err := client.Do(request)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return PkkRosreestrLotDto{}, err
	//}
	//defer response.Body.Close()
	//
	//if response.StatusCode != http.StatusOK {
	//	return PkkRosreestrLotDto{}, errors.New(response.Status)
	//}
	//
	//var lotDto PkkRosreestrLotDto
	//if err = json.NewDecoder(response.Body).Decode(&lotDto); err != nil {
	//	return PkkRosreestrLotDto{}, err
	//}
	//
	//return lotDto, nil
	return PkkRosreestrLotDto{}, nil
}

func pkkRequest(ctx context.Context, url string) (*http.Request, error) {

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	//request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	//request.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	//request.Header.Set("Accept-Language", "ru,en;q=0.9")
	//request.Header.Set("Cache-Control", "max-age=0")
	//request.Header.Set("Connection", "keep-alive")
	//request.Header.Set("Cookie", "USER_REGION_ID=428; USER_REGION_ID_NAME=%D0%9C%D0%BE%D1%81%D0%BA%D0%B2%D0%B0; USER_CITY=%D0%9C%D0%BE%D1%81%D0%BA%D0%B2%D0%B0; USER_REGION_ID_INNER_ID=77; uid=CoHkfWcScUo11h/l9aTdAg==; BITRIX_CONVERSION_CONTEXT_s1=%7B%22ID%22%3A14%2C%22EXPIRE%22%3A1729285140%2C%22UNIQUE%22%3A%5B%22conversion_visit_day%22%5D%7D; BX_USER_ID=2438d7740f5f46789dfe523f1f75dbca; session-cookie=1800047aa765f072ed3dcc58b68b8c5b3396b97aca165c14e40853d7660ee28c375f44eddc3278418cc7301b1d5df32a")
	//request.Header.Set("Host", "rosreestr.gov.ru")
	//request.Header.Set("Sec-Fetch-Dest", "document")
	//request.Header.Set("Sec-Fetch-Mode", "navigate")
	//request.Header.Set("Sec-Fetch-Site", "none")
	//request.Header.Set("Sec-Fetch-User", "?1")
	//request.Header.Set("Upgrade-Insecure-Requests", "1")
	//request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 YaBrowser/24.7.0.0 Safari/537.36")
	//request.Header.Set("sec-ch-ua", "\"Not/A)Brand\";v=\"8\", \"Chromium\";v=\"126\", \"YaBrowser\";v=\"24.7\", \"Yowser\";v=\"2.5\"")
	//request.Header.Set("sec-ch-ua-mobile", "?0")
	//request.Header.Set("sec-ch-ua-platform", "Linux")
	request.Header.Set("verify", "/home/igor/GolandProjects/learning/web-scraping-practice/internal/adapters/api/pkkRosreestr/cacert.pem")

	return request, err
}

// Функция для создания настроенного транспортного уровня с минимальной версией TLS 1.2
func createTransport() *http.Transport {

	// Загрузка сертификата CA
	caCert, err := os.ReadFile("internal/adapters/api/pkkRosreestr/cacert.pem")
	if err != nil {
		fmt.Printf("Ошибка загрузки сертификата CA: %v\n", err)
	}

	// Создание пула сертификатов CA
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		fmt.Println("Ошибка добавления сертификата CA в пул")
	}

	// Создаем TLS конфигурацию
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12, // Устанавливаем минимальную версию TLS 1.2
		RootCAs:    caCertPool,
	}

	// Устанавливаем уровень безопасности SSL контекста
	// Это аналог SECLEVEL=1 в OpenSSL
	tlsConfig.CipherSuites = []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	}

	// Настраиваем HTTP транспорт с использованием этой TLS конфигурации
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return transport
}

func normalizeCadastreNumber(CadastreNumber string) string {

	var result []string

	NumParts := strings.Split(CadastreNumber, ":")
	for _, part := range NumParts {

		result = append(result, strings.TrimLeft(part, "0"))

	}

	return strings.Join(result, ":")
}
