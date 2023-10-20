package main

func main() {
	str := "  snow dog sun  "
	res := ""

	for i := 0; i < len(str); i++ {
		// пропуск пробелов
		if str[i] == ' ' {
			continue
		}

		// поиск конца слова
		j := i
		for j < len(str) && str[j] != ' ' {
			j++
		}
		// побавление слова к результату с пробелом
		res = str[i:j] + " " + res
		// переход через слово
		i = j
	}

	println(res)
}
