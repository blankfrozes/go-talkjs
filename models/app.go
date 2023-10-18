package models

type App struct {
	ID            string            `json:"id"`
	DefaultLocale string            `json:"defaultLocale"`
	Custom        map[string]string `json:"custom"`
}
