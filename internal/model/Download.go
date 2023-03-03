package model

type Download struct {
	ID    uint   `json:"id" gorm:"primarykey"`
	Path  string `json:"path" mapstructure:"path"`
	Cache string `json:"cache" mapstructure:"cache"`
}
