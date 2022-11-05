package entity

type Users struct {
	ID         int16  `json:"id"`
	Name       string `json:"name" example:"Name"`
	Email      string `json:"email" example:"email@gmail.com"`
	Password   string `json:"password" example:"aB@123456"`
	Picture    int16  `json:"picture" example:"1"`
	Newsletter bool   `json:"newsletter" example:"true"`
}
