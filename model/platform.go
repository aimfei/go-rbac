package model

import "time"

type Platform struct {
	Id           int16
	PlatformCode string    `json:"platform_code"`
	PlatformName string    `json:"platform_name"`
	SignKey      string    `json:"sign_key"`
	SignType     int       `json:"sign_type"`
	GmtCreate    time.Time `json:"gmt_create"`
	GmtModified  time.Time `json:"gmt_modified"`
}

func Save(model *Platform) {

}
