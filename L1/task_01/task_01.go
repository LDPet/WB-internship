package main

type Human struct {
	Hp int
}

func (h *Human) eat() {
	h.Hp++
}

func (h *Human) fall() {
	h.Hp--
}

type Action struct {
	// встроили структуру Human
	// теперь имеем достук к её полям и мотодам через Action
	Human
}

// что-то типа override
// если совпадут названия у встраиваемых структур, то компилятор выдаст ошибку
// если совпадут названия у "родительских" и "дочерних" структур, то предпочтение отдаётся методу дочерней
func (a *Action) fall() {
	a.Hp -= 10
}

func main() {
	a := Action{}
	println(a.Hp)
	a.eat()
	println(a.Hp)
	a.fall()
	println(a.Hp)
}
