package models

type GeoData struct {
	IP string `json:"ip"`
	Coords
}

type PrayerMethod struct {
	ID int `json:"id"`
	Coords
}

type Coords struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PrayerTimesParams struct {
	Coords   Coords
	MethodID int
}

type PrayerName struct {
	Ar string
	En string
}

type PrayerNames struct {
	Fajr    PrayerName
	Sunrise PrayerName
	Dhuhr   PrayerName
	Asr     PrayerName
	Maghrib PrayerName
	Isha    PrayerName
}
