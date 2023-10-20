package main

import "fmt"

func main() {
	str := "®qwerty¢qwerty😂qwerty😉qwerty☎"
	// для нормального пебора по индексам, т.к. строка это срез байт, а unicode символ занимает больше 1 байта (4 байта)
	// и нормально перебрать строку с unicode сиволами можно только через range (при обращении по индексам можно
	// получить байт являющий частью символа )
	// rune = int32 позволяет нормально представлять unicode символы
	r := []rune(str)
	l := len(r)

	for i := 0; i < l/2; i++ {
		r[i], r[l-i-1] = r[l-i-1], r[i]
	}

	fmt.Print(string(r))
}
