package prayers

import (
	"github.com/mdyssr/prayers/internal/services"
	"github.com/mdyssr/prayers/models"
)

// GetPrayersData returns prayers data and an error
func GetPrayersData() (models.PrayersData, error) {
	return services.GetPrayersData()
}
