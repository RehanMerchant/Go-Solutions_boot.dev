package main

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

type Bold struct {
	message string
}

type Code struct {
	message string
}

func (e PlainText) Format() string {
	return e.message
}

func (e Bold) Format() string {
	return "**" + e.message + "**"
}

func (e Code) Format() string {
	return "`" + e.message + "`"
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
