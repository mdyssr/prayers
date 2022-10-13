package prayertimes

import (
	"github.com/mdyssr/prayers/models"
	"github.com/mdyssr/prayers/services"
	"github.com/mdyssr/prayers/utils"
)

const GetUserIPDataError = Error("Error getting user IP data")
const GetPrayerMethodsError = Error("Error getting prayer methods")
const GetPrayerDataError = Error("Error getting prayer data")

type Error string

func (e Error) Error() string {
	return string(e)
}

func GetPrayersData() (models.PrayersData, error) {
	geoData, err := services.GetGeoLocation()
	if err != nil {
		return models.PrayersData{}, GetUserIPDataError
	}

	methods, err := services.GetMethods()
	if err != nil {
		return models.PrayersData{}, GetPrayerMethodsError
	}

	nearestMethodID := utils.GetNearestMethod(&geoData.Coords, methods)
	prayerTimesParams := models.PrayerTimesParams{
		Coords:   geoData.Coords,
		MethodID: nearestMethodID,
	}

	prayerTimes, err := services.GetPrayersData(prayerTimesParams)
	if err != nil {
		return models.PrayersData{}, GetPrayerDataError
	}

	return prayerTimes, nil
}
