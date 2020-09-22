package code

func Join(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}
