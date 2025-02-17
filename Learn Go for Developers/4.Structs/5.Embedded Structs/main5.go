package main

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}
