package utils

import (
	"fmt"
	"github.com/mdyssr/prayers/models"
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

func GetStandardPrayerTime(timeString string) *models.StandardPrayerTime {
	timeStringParts := strings.Split(timeString, ":")
	hours, minutes := timeStringParts[0], timeStringParts[1]

	minutesNumber, _ := strconv.Atoi(minutes)
	hoursNumber, _ := strconv.Atoi(hours)
	designation := GetPrayerTimeDesignation(hoursNumber)
	hoursNumber = hoursNumber % 12
	standardTime := fmt.Sprintf("%s:%s", PrefixWithZero(hoursNumber), PrefixWithZero(minutesNumber))

	return &models.StandardPrayerTime{
		Value:       standardTime,
		Designation: *designation,
	}
}

func GetPrayerTimeDesignation(hour int) *models.StandardPrayerTimeDesignation {
	//designation := new(models.StandardPrayerTimeDesignation)
	if hour > 12 {
		return &models.StandardPrayerTimeDesignation{
			Ar: models.PrayerTimeLanguageDesignation{
				Abbreviated: "م",
				Expanded:    "مساءًا",
			},
			En: models.PrayerTimeLanguageDesignation{
				Abbreviated: "pm",
				Expanded:    "After Midday",
			},
		}
	} else {
		return &models.StandardPrayerTimeDesignation{
			Ar: models.PrayerTimeLanguageDesignation{
				Abbreviated: "ص",
				Expanded:    "صباحًا",
			},
			En: models.PrayerTimeLanguageDesignation{
				Abbreviated: "am",
				Expanded:    "Before Midday",
			},
		}
	}
}

func FormatPrayerTiming(name, timing string) models.FormattedPrayerTiming {
	return models.FormattedPrayerTiming{
		Name: name,
		Time: models.PrayerTimeDetails{
			Military: timing,
			Standard: *GetStandardPrayerTime(timing),
		},
	}
}
