package services

import (
	"encoding/json"
	"fmt"
	"github.com/mdyssr/prayers/data"
	"github.com/mdyssr/prayers/models"
	"github.com/mdyssr/prayers/utils"
	"io"
	"net/http"
	"time"
)

func GetPrayersData(params models.PrayerTimesParams) (models.PrayersData, error) {
	var prayerData models.PrayersData
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

	prayersResponse := new(models.PrayersResponse)
	err = json.Unmarshal(bodyBytes, prayersResponse)
	if err != nil {
		return prayerData, err
	}

	prayerData = models.PrayersData{
		HijriDate: prayersResponse.Data.Date.HijriDate,
		//PrayerTimings: prayersResponse.Data.Timings,
		PrayerTimings: models.FormattedPrayerTimings{
			utils.FormatPrayerTiming("Fajr", prayersResponse.Data.Timings.Fajr),
			utils.FormatPrayerTiming("Sunrise", prayersResponse.Data.Timings.Sunrise),
			utils.FormatPrayerTiming("Dhuhr", prayersResponse.Data.Timings.Dhuhr),
			utils.FormatPrayerTiming("Asr", prayersResponse.Data.Timings.Asr),
			utils.FormatPrayerTiming("Sunset", prayersResponse.Data.Timings.Sunset),
			utils.FormatPrayerTiming("Maghrib", prayersResponse.Data.Timings.Maghrib),
			utils.FormatPrayerTiming("Isha", prayersResponse.Data.Timings.Isha),
		},
	}
	return prayerData, nil
}

func GetMethods() ([]models.PrayerMethod, error) {
	return data.PrayerMethods, nil
}
