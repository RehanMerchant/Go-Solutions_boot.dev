package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {

	if _, exists := friendships[username]; !exists {
		return nil
	}

	directFriends := make(map[string]struct{})

	for _, friends := range friendships[username] {

		directFriends[friends] = struct{}{}

	}

	suggested := make(map[string]struct{})

	for _, friend := range friendships[username] {
		for _, friendOfFriend := range friendships[friend] {
			if friendOfFriend != username { // Avoid self-suggestion
				if _, isDirectFriend := directFriends[friendOfFriend]; !isDirectFriend {
					suggested[friendOfFriend] = struct{}{}
				}
			}
		}
	}

	if len(suggested) == 0 {
		return nil // Return nil if there are no suggested friends
	}

	result := make([]string, 0, len(suggested))
	for friend := range suggested {
		result = append(result, friend)
	}

	return result

}
