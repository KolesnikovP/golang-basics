package main

import "fmt"

func main() {
	/*
		Map — тип данных, предназначенный для хранения пар ключ-значение. В других языках эту структуру так же называют:
		хэш-таблица, словарь, ассоциативный массив. Запись и чтение элементов происходят в основном за O(1):
	*/
	// empty map
	var m map[int]string
	// or
	m1 := map[int]string{}

	// recommendation to create map with size
	m2 := make(map[int]string, 10)

	// map with values
	m3 := map[int]string{1: "value1", 2: "value2"}

	// lets add new element
	m3[3] = "value3" // // map[1:value1, 2:value2, 3:value3]

	// read a element
	word := m[1] // value1

	fmt.Println(m3)

	/*
		При чтении элемента по несуществующему ключу возвращается нулевое значение данного типа.
		Это приводит к ошибкам логики, когда используется bool как значение.
		Для решения данной проблемы при чтении используется вторая переменная, в которую записывается наличие элемента в мапе:
	*/

	existedIDs := map[int64]bool{1: true, 2: true}

	idExists, elementExists := existedIDs[2] // true, true

	//idExists, elementExists := existedIDs[225] // false, false
	// Элементы удаляются с помощью встроенной функции delete(m map[Type]Type1, key Type):
	engToRus := map[string]string{"hello": "привет", "world": "мир"}

	delete(engToRus, "world")

	fmt.Println(engToRus) // map[hello:привет]

	// Мапы в Go всегда передаются по ссылке:
	m4 := map[int]string{1: "hello", 2: "world"}

	modifyMap(m4)

	fmt.Println(m4) // вывод: map[1:changed 2:world 200:added]

	// Как и слайс, мапу можно обойти с помощью конструкции for range:
	idToName := map[int64]string{1: "Alex", 2: "Dan", 3: "George"}

	// Стоит учитывать, что порядок ключей в мапе рандомизирован:
	// первый аргумент — ключ, второй — значение
	for id, name := range idToName {
		fmt.Println("id: ", id, "name: ", name)
	}
}

func modifyMap(m map[int]string) {
	m[200] = "added"

	m[1] = "changed"
}
