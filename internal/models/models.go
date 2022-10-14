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

type StandardPrayerTimeDesignation struct {
	Ar PrayerTimeLanguageDesignation
	En PrayerTimeLanguageDesignation
}
type PrayerTimeLanguageDesignation struct {
	Abbreviated string
	Expanded    string
}

type StandardPrayerTime struct {
	Value       string
	Designation StandardPrayerTimeDesignation
}

type PrayerTimeDetails struct {
	Military string
	Standard StandardPrayerTime
}

type FormattedPrayerTiming struct {
	Name string
	Time PrayerTimeDetails
}

type FormattedPrayerTimings []FormattedPrayerTiming

type HijriDate struct {
	Day     string `json:"day"`
	Weekday struct {
		En string `json:"en"`
		Ar string `json:"ar"`
	} `json:"weekday"`
	Month struct {
		Number int    `json:"number"`
		En     string `json:"en"`
		Ar     string `json:"ar"`
	} `json:"month"`
	Year string `json:"year"`
}
