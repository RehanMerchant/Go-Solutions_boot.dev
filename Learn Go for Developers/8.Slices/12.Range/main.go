package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, bad := range badWords {
			if word == bad {
				return i
			}
		}
	}
	return -1
}
