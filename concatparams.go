package piscine

func ConcatParams(args []string) string {
	phrase := ""
	for _, taille := range args {
		phrase = phrase + taille + "\n"
	}
	return phrase
}
