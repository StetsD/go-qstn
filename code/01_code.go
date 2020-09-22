package code

// Чем обусловлена неоптимальность ф-ии
func Code_01(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}

// под каждую итерацию аллоцируется доп память
