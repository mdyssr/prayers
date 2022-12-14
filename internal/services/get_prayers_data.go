package services

import (
	"encoding/json"
	"fmt"
	"github.com/mdyssr/prayers/internal/data"
	"github.com/mdyssr/prayers/internal/models"
	"github.com/mdyssr/prayers/internal/utils"
	prayerModel "github.com/mdyssr/prayers/models"
	"io"
	"net/http"
	"time"
)

const GetUserIPDataError = Error("Error getting user IP data")
const GetPrayerMethodsError = Error("Error getting prayer methods")
const GetPrayerDataError = Error("Error getting prayer data")

type Error string

func (e Error) Error() string {
	return string(e)
}

type PrayerMethod struct {
	ID int `json:"id"`
	models.Coords
}

type Date struct {
	HijriDate prayerModel.HijriDate `json:"hijri"`
}

type PrayerTimings struct {
	Fajr    string `json:"Fajr"`
	Sunrise string `json:"Sunrise"`
	Dhuhr   string `json:"Dhuhr"`
	Asr     string `json:"Asr"`
	Sunset  string `json:"Sunset"`
	Maghrib string `json:"Maghrib"`
	Isha    string `json:"Isha"`
}

type Data struct {
	Date    Date          `json:"date"`
	Timings PrayerTimings `json:"timings"`
}

type PrayersResponse struct {
	Data Data `json:"data"`
}

func getPrayersDataFromAPI(params models.PrayerTimesParams) (prayerModel.PrayersData, error) {
	var prayerData prayerModel.PrayersData
	client := http.Client{}
	now := time.Now().Unix()
	url := fmt.Sprintf("%s%d?latitude=%g&longitude=%g&method=%d", data.PRAYER_TIMINGS_BASE_URL, now, params.Coords.Latitude, params.Coords.Longitude, params.MethodID)
	response, err := client.Get(url)
	if err != nil {
		return prayerData, err
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return prayerData, err
	}

	prayersResponse := new(PrayersResponse)
	err = json.Unmarshal(bodyBytes, prayersResponse)
	if err != nil {
		return prayerData, err
	}

	prayerData = prayerModel.PrayersData{
		HijriDate: prayersResponse.Data.Date.HijriDate,
		//PrayerTimings: prayersResponse.Data.Timings,
		PrayerTimings: prayerModel.FormattedPrayerTimings{
			utils.FormatPrayerTiming("Fajr", prayersResponse.Data.Timings.Fajr),
			utils.FormatPrayerTiming("Sunrise", prayersResponse.Data.Timings.Sunrise),
			utils.FormatPrayerTiming("Dhuhr", prayersResponse.Data.Timings.Dhuhr),
			utils.FormatPrayerTiming("Asr", prayersResponse.Data.Timings.Asr),
			utils.FormatPrayerTiming("Maghrib", prayersResponse.Data.Timings.Maghrib),
			utils.FormatPrayerTiming("Isha", prayersResponse.Data.Timings.Isha),
		},
	}
	return prayerData, nil
}

func GetMethods() ([]models.PrayerMethod, error) {
	return data.PrayerMethods, nil
}

func GetPrayersData() (prayerModel.PrayersData, error) {
	geoData, err := GetGeoLocation()
	if err != nil {
		return prayerModel.PrayersData{}, GetUserIPDataError
	}

	methods, err := GetMethods()
	if err != nil {
		return prayerModel.PrayersData{}, GetPrayerMethodsError
	}

	nearestMethodID := utils.GetNearestMethod(&geoData.Coords, methods)
	prayerTimesParams := models.PrayerTimesParams{
		Coords:   geoData.Coords,
		MethodID: nearestMethodID,
	}

	prayerTimes, err := getPrayersDataFromAPI(prayerTimesParams)
	if err != nil {
		return prayerModel.PrayersData{}, GetPrayerDataError
	}

	return prayerTimes, nil
}
