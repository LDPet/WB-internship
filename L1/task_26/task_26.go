package main

import "strings"

func main() {
	uniqStr := "qWerTy"
	repeatStr := "qwertEy"

	println(isUniqCharacters(uniqStr))
	println(isUniqCharacters(repeatStr))
}

func isUniqCharacters(str string) bool {
	// приводим к всё к нижнему регистру для регистронезависимости
	str = strings.ToLower(str)
	// в маке омечаются найденные символы
	m := make(map[int32]struct{})

	for _, c := range str {
		// если сивол был, то false
		if _, ok := m[c]; ok {
			return false
		}
		// иначе добавляем в мапу
		m[c] = struct{}{}
	}

	// не было найдено повторов
	return true
}
