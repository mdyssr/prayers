package prayers

import (
	"github.com/mdyssr/prayers/internal/models"
	"github.com/mdyssr/prayers/internal/services"
)

type PrayersData struct {
	PrayerTimings models.FormattedPrayerTimings
	HijriDate     models.HijriDate
}

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

// GetPrayersData returns prayers data or an error
func GetPrayersData() (PrayersData, error) {
	prayersData := PrayersData{}
	prayersData, err := services.GetPrayersData()
	if err != nil {
		return prayersData, err
	}

	return prayersData, nil
}
