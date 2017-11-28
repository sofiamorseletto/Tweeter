package domain

type User struct {
	Name           string
	Tweets         []Tweeter
	Following      []*User
	DirectMessages []Tweeter
}

func NewUser(name string) *User {
	u := User{
		name,
		make([]Tweeter, 0),
		make([]*User, 0),
		make([]Tweeter, 0),
	}

	return &u
}
