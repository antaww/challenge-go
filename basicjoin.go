package piscine

func BasicJoin(elems []string) string {
	phrase := ""
	for _, taille := range elems {
		phrase = phrase + taille
	}
	return phrase
}
