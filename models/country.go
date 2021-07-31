package models

type Country struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	TitleRU string `json:"title_ru"`
}
