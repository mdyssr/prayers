package utils

import (
	"fmt"
	"github.com/mdyssr/prayers/internal/models"
	prayerModel "github.com/mdyssr/prayers/models"
	"math"
	"strconv"
	"strings"
)

func GetNearestMethod(coords *models.Coords, methods []models.PrayerMethod) int {

	distances := make(map[int]float64)

	for _, m := range methods {
		latDiff := coords.Latitude - m.Coords.Latitude
		lonDiff := coords.Longitude - m.Coords.Longitude
		distance := math.Sqrt(math.Pow(latDiff, 2) + math.Pow(lonDiff, 2))
		distances[m.ID] = distance
	}

	// the default smallest Id is the first method
	closestMethodID := methods[0].ID
	smallestDistance := distances[closestMethodID]

	for id, distance := range distances {
		if smallestDistance > distance {
			smallestDistance = distance
			closestMethodID = id
		}
	}

	return closestMethodID
}

func PrefixWithZero(value int) string {
	if value > 9 {
		return strconv.Itoa(value)
	} else {
		return "0" + strconv.Itoa(value)
	}
}

func GetStandardPrayerTime(timeString string) *prayerModel.StandardPrayerTime {
	timeStringParts := strings.Split(timeString, ":")
	hours, minutes := timeStringParts[0], timeStringParts[1]

	minutesNumber, _ := strconv.Atoi(minutes)
	hoursNumber, _ := strconv.Atoi(hours)
	designation := GetPrayerTimeDesignation(hoursNumber)
	hoursNumber = hoursNumber % 12
	standardTime := fmt.Sprintf("%s:%s", PrefixWithZero(hoursNumber), PrefixWithZero(minutesNumber))

	return &prayerModel.StandardPrayerTime{
		Value:       standardTime,
		Designation: *designation,
	}
}

func GetPrayerTimeDesignation(hour int) *prayerModel.StandardPrayerTimeDesignation {
	//designation := new(models.StandardPrayerTimeDesignation)
	if hour > 12 {
		return &prayerModel.StandardPrayerTimeDesignation{
			Ar: prayerModel.PrayerTimeLanguageDesignation{
				Abbreviated: "م",
				Expanded:    "مساءًا",
			},
			En: prayerModel.PrayerTimeLanguageDesignation{
				Abbreviated: "pm",
				Expanded:    "After Midday",
			},
		}
	} else {
		return &prayerModel.StandardPrayerTimeDesignation{
			Ar: prayerModel.PrayerTimeLanguageDesignation{
				Abbreviated: "ص",
				Expanded:    "صباحًا",
			},
			En: prayerModel.PrayerTimeLanguageDesignation{
				Abbreviated: "am",
				Expanded:    "Before Midday",
			},
		}
	}
}

func FormatPrayerTiming(name, timing string) prayerModel.FormattedPrayerTiming {
	return prayerModel.FormattedPrayerTiming{
		Name: name,
		Time: prayerModel.PrayerTimeDetails{
			Military: timing,
			Standard: *GetStandardPrayerTime(timing),
		},
	}
}
