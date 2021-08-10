package Helper

type User []struct {
	ID         string   `json:"id"`
	Password   string   `json:"password"`
	IsActive   bool     `json:"isActive"`
	Balance    string   `json:"balance"`
	Age        string   `json:"age"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	Tags       []string `json:"tags"`
	Friends    []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"friends"`
	Data string `json:"data"`
}

type Users []User