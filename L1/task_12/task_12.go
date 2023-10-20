package main

import "fmt"

// Само множество и пустое множество называют несобственными подмножествами, остальные подмножества называют собственными
// также в множестве все элементы ункальны => нужно выделить набор уникальных строк
func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	res := make([]string, 0)

	// хранит отметку, что слово уже встречалось
	m := make(map[string]struct{})
	// перебираем все слова
	for _, word := range words {
		// если слово не встречалось
		if _, ok := m[word]; !ok {
			// помечаем
			m[word] = struct{}{}
			// сохраняем
			res = append(res, word)
		}
	}

	fmt.Println(res)
}
