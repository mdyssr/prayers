package names

import "github.com/mdyssr/prayers/internal/models"

var PrayerNames = map[string]models.PrayerName{
	"Fajr": {
		Ar: "الفجر",
		En: "Fajr",
	},
	"Sunrise": {
		Ar: "الشروق",
		En: "Sunrise",
	},
	"Dhuhr": {
		Ar: "الظهر",
		En: "Dhuhr",
	},
	"Asr": {
		Ar: "العصر",
		En: "Asr",
	},
	"Maghrib": {
		Ar: "المغرب",
		En: "Maghrib",
	},
	"Isha": {
		Ar: "العشاء",
		En: "Isha",
	},
}
