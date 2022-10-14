package names

import "github.com/mdyssr/prayers/internal/models"

var PrayerNames = models.PrayerNames{
	Fajr: models.PrayerName{
		Ar: "الفجر",
		En: "Fajr",
	},
	Sunrise: models.PrayerName{
		Ar: "الشروق",
		En: "Sunrise",
	},
	Dhuhr: models.PrayerName{
		Ar: "الظهر",
		En: "Dhuhr",
	},
	Asr: models.PrayerName{
		Ar: "العصر",
		En: "Asr",
	},
	Maghrib: models.PrayerName{
		Ar: "المغرب",
		En: "Maghrib",
	},
	Isha: models.PrayerName{
		Ar: "العشاء",
		En: "Isha",
	},
}
