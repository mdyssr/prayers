package services

import (
	"encoding/json"
	"fmt"
	"github.com/mdyssr/prayers/models"
	"io"
	"net/http"
)

func GetGeoLocation() (models.GeoData, error) {
	geoData := models.GeoData{}
	const base_url = "http://api.ipstack.com/check"
	const api_key = "10c6be7bf5f2704393a672e2c714ca04"
	const fields = "ip,latitude,longitude"
	//const url = "http://api.ipstack.com/check?access_key=10c6be7bf5f2704393a672e2c714ca04&fields=ip,latitude,longitude"
	url := fmt.Sprintf("%s?access_key=%s&fields=%s", base_url, api_key, fields)
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return geoData, err
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return geoData, err
	}

	err = json.Unmarshal(responseData, &geoData)
	if err != nil {
		return geoData, err
	}
	fmt.Println(geoData.IP)
	return geoData, nil
}
