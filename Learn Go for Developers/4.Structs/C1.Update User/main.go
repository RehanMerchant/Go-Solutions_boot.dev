package main

type Membership struct {
	MessageCharLimit int
	Type             string
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType string) User {

	limit := 100
	if membershipType == "premium" {
		limit = 1000
	}

	return User{
		Name: name,
		Membership: Membership{
			MessageCharLimit: limit,
			Type:             membershipType,
		},
	}

}
