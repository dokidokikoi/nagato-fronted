package model

var config *Config

type Config struct {
	ID         uint     `json:"id" gorm:"primarykey"`
	DownloadID uint     `json:"download_id"`
	Default    bool     `json:"default" gorm:"default:0"`
	InUse      bool     `json:"in_use" gorm:"default:0"`
	Download   Download `gorm:"foreignKey:DownloadID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" mapstructure:"download"`
	Sqlite     Sqlite   `json:"sqlite" gorm:"-" mapstructure:"sqlite"`
}

func GetConfig() *Config {
	return config
}

func SetConfig(c *Config) {
	config = c
}
