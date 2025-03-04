package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	return [3]string{
			primary,
			secondary,
			tertiary,
		}, [3]int{
			len(primary),
			len(secondary + primary),
			len(secondary + primary + tertiary),
		}
}
