package entity

type Users struct {
	ID         int16  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Picture    int16  `json:"picture"`
	Newsletter bool   `json:"newsletter"`
}
