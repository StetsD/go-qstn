package code

// Что выведет код
// Почему? Как нужно изменить функцию changePointer, чтобы вывело 5 и 3
func Code_04() {
	v := 5
	p := &v
	println(*p)
	changePointer(p)
	println(*p)
}

func changePointer(p *int) {
	v := 3
	p = &v
	//*p = v
}

//  5 5
