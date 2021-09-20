package piscine

func Join(strs []string, sep string) string {
	var phrase string
	for i := 0; i < len(strs); i++ {
		if i < len(strs)-1 {
			phrase = phrase + sep
		} else {
			phrase = phrase + strs[i]
		}
	}
	return phrase
}
