package main

func reformat(message string, formatter func(string) string) string {
	s := formatter(message)
	s = formatter(s)
	s = formatter(s)

	a := "TEXTIO: " + s
	return a
}
