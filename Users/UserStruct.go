package User

type User struct {
	ID         string    `bson:"id"`
	Password   string    `bson:"password"`
	IsActive   bool      `bson:"isActive"`
	Balance    string    `bson:"balance"`
	Age        interface{} `bson:"age"`
	Name       string    `bson:"name"`
	Gender     string    `bson:"gender"`
	Company    string    `bson:"company"`
	Email      string    `bson:"email"`
	Phone      string    `bson:"phone"`
	Address    string    `bson:"address"`
	About      string    `bson:"about"`
	Registered string    `bson:"registered"`
	Latitude   float64   `bson:"latitude"`
	Longitude  float64   `bson:"longitude"`
	Tags       []string  `bson:"tags"`
	Friends    []Friends `bson:"friends"`
	Data       string    `bson:"data"`
}
type Friends struct {
	ID   int    `bson:"id"`
	Name string `bson:"name"`
}

