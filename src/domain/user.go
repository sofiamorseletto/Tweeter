package domain

type User struct {
	Name      string
	Tweets    []*Tweet
	Following []*User
}

func NewUser(name string) *User {
	u := User{
		name,
		make([]*Tweet, 0),
		make([]*User, 0),
	}

	return &u
}
