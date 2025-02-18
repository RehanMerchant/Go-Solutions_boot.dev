package main

func (e email) cost() int {
	if e.isSubscribed == true {
		return len(e.body) * 2
	}
	return len(e.body) * 5
}

func (e email) format() string {
	s := "Not Subscribed"
	if e.isSubscribed == true {
		s = "Subscribed"
	}

	return "'" + e.body + "' | " + s
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
