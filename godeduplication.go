package godeduplication

func Dedup_String(str *[]string) {
	found := make(map[string]bool)
	timber := 0
	for id, val := range *str {
		if !found[val] {
			found[val] = true
			(*str)[timber] = (*str)[id]
			timber++
		}
	}
	*str = (*str)[:timber]
}
