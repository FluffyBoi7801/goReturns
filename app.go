package main

import "fmt"

func main() {
	var num int
	text1, text2 := test()

	fmt.Println(test())

	fmt.Println(text1, text2)

	fmt.Scanf("%d", &num)
	fmt.Printf("Итог: %d\n\n\n", loop(num))

	fmt.Printf("Итог: %d\n", whileClone(num))

	defer fmt.Println("Заканчиваю программу...") // Отложенная операция, выполнится после окончания функции. Если несколько - выполнябтся в реверсе
	fmt.Println("Все действия выполнены!")
}

func test() (string, string) {
	a := "Hi"
	b := "Bobik"
	return a, b
}

func loop(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		fmt.Println(result)
		result += i
	}
	return result
}

func whileClone(n int) int {
	result := 0
	for result < n {
		result++
		fmt.Println(result)
		if result == n {
			fmt.Println("Достигнуто!")
		} else {
			fmt.Println("Идем дальше...")
		}
	}
	return result
}

func switches(i int) {
	switch i {
	case 1:

	case 2:

	default:

	}
}
