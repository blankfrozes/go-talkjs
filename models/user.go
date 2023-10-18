package models

type User struct {
	Name             string            `json:"name"`
	Email            []string          `json:"email"`
	WelcomeMessage   string            `json:"welcomeMessage"`
	PhotoURL         string            `json:"photoUrl"`
	Role             string            `json:"role"`
	Phone            []string          `json:"phone"`
	Custom           map[string]string `json:"custom"`
	AvailabilityText string            `json:"availabilityText,omitempty"`
	CreatedAt        int64             `json:"createdAt,omitempty"`
	PushTokens       *map[string]bool  `json:"pushTokens,omitempty"`
}
