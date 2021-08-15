package model

type SlackProfile struct {
	Profile `json:"profile"`
}

type Profile struct {
	StatusEmoji      string `json:"status_emoji"`
	StatusExpiration int64  `json:"status_expiration"`
	StatusText       string `json:"status_text"`
}
