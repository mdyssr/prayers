package prayers

import (
	"github.com/mdyssr/prayers/internal/services"
	"github.com/mdyssr/prayers/models"
)

// GetPrayersData returns prayers data or an error
func GetPrayersData() (models.PrayersData, error) {
	prayersData := models.PrayersData{}
	prayersData, err := services.GetPrayersData()
	if err != nil {
		return prayersData, err
	}

	return prayersData, nil
}
