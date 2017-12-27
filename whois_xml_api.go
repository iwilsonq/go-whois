package whois

import (
	"encoding/json"
	"errors"
	"net/http"
)

type WhoisXMLAPI struct {
	APIKey string
}

type whoisResponse struct {
	WhoisRecord map[string]interface{} `json:"WhoisRecord"`
}

func (w *WhoisXMLAPI) Lookup(domain string) (bool, error) {
	response, err := http.Get("https://www.whoisxmlapi.com/whoisserver/WhoisService?apiKey=" + w.APIKey + "&domainName=" + domain + "&outputFormat=json&da=1")
	if err != nil {
		return false, errors.New("whoisxmlapi: Failed when looking for availablitiy of " + domain + err.Error())
	}

	var data whoisResponse
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return false, err
	}

	return data.WhoisRecord["domainAvailability"] == "UNAVAILABLE", nil
}
