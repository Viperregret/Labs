package fun

import (
	"fmt"
	"unicode/utf8"
)

func init() {}

func IsOdd(number int) { //четность
	if number%2 == 0 {
		fmt.Println("Числое четное")
	} else {
		fmt.Println("Число нечетное")
	}
}

func DefineSign(number int) string { //знак
	if number > 0 {
		return "Positive"
	} else if number < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

func PrintFromZeroToTen() { // вывод от 0 до 10
	fmt.Println()
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()
}

func StringLen(str string) int { //длина строки
	return utf8.RuneCountInString(str)
}
func Average(firstNum int, secondNum int) float32 { // среднее
	return (float32(firstNum) + float32(secondNum)) / 2
}
