package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" swaggignore:"true"`
	// Add more fields if needed
}

type Profile struct {
	Username    string            `json:"username"`
	Email       string            `json:"email"`
	FirstName   string            `json:"first_name,omitempty"`
	LastName    string            `json:"last_name,omitempty"`
	Age         int               `json:"age,omitempty"`
	Gender      string            `json:"gender,omitempty"`
	Address     string            `json:"address,omitempty"`
	City        string            `json:"city,omitempty"`
	Country     string            `json:"country,omitempty"`
	PhoneNumber string            `json:"phone_number,omitempty"`
	Interests   []string          `json:"interests,omitempty"`
	SocialLinks map[string]string `json:"social_links,omitempty"`
}
