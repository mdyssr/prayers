package models

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

type PrayersData struct {
	PrayerTimings FormattedPrayerTimings
	HijriDate     HijriDate
}
