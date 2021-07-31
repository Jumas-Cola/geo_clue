package models

type City struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	CountryID uint    `json:"country_id"`
	TitleRU   string  `json:"title_ru"`
	AreaRU    *string `json:"area_ru"`
	RegionRU  *string `json:"region_ru"`
}
