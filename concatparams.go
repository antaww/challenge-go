package piscine

func ConcatParams(args []string) string {
	phrase := ""
	j := 0
	for _, i := range args {
		phrase = phrase + i
		if j < len(args) {
			phrase += "\n"

		}

	}
	return phrase
}
