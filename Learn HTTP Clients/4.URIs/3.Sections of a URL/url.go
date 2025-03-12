package main

type ParsedURL struct {
	protocol string
	username string
	password string
	hostname string
	port     string
	pathname string
	search   string
	hash     string
}
