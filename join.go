package piscine

func Join(strs []string, sep string) string {
	var phrase string
	for _, taille := range strs {
		phrase = phrase + taille + sep
	}
	return phrase
}
