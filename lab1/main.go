package main

import (
	"fmt"
	funct "lab1/fun1"
	"time"
)

func main() {
	fmt.Println("Задание 1")
	today := time.Now()
	fmt.Println(today)
	fmt.Println("Задание 2")
	A := int(20)
	B := float64(1.454231)
	S := string("Hello world!")
	C := bool(true)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(S)
	fmt.Println(C)
	fmt.Println("Задание 3")
	A1 := 15
	fmt.Println(A1)
	B1 := 1.4512
	fmt.Println(B1)
	fmt.Println("Задание 4")
	var a2 int64
	var b2 int64
	var k string
	fmt.Println("Введите первое целое число:")
	fmt.Scan(&a2)
	fmt.Println("Выберете действие(+;-;*;/;%):")
	fmt.Scan(&k)
	fmt.Println("Введите второе целое число:")
	fmt.Scan(&b2)
	switch {
	case k == "+":
		fmt.Println("Сумма чисел равна " + fmt.Sprint(a2+b2))
	case k == "-":
		fmt.Println("Разница чисел равна " + fmt.Sprint(a2-b2))
	case k == "*":
		fmt.Println("Произведение чисел равно " + fmt.Sprint(a2*b2))
	case k == "/":
		fmt.Println("Деление чисел равно " + fmt.Sprint(a2/b2))
	case k == "%":
		fmt.Println("Остаток от деления равен " + fmt.Sprint(a2%b2))
	}
	fmt.Println("Задание 5")
	var f1 float32
	var f2 float32
	fmt.Println("Введите первое число:")
	fmt.Scan(&f1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&f2)
	fmt.Println("Сумма равна:" + fmt.Sprint(funct.Sum(f1, f2)))
	fmt.Println("Разница равна:" + fmt.Sprint(funct.Sub(f1, f2)))

	fmt.Println("Задание 6")
	var l float32
	var m float32
	var h float32
	fmt.Println("Введите первое число:")
	fmt.Scan(&l)
	fmt.Println("Введите второе число:")
	fmt.Scan(&m)
	fmt.Println("Введите третье число:")
	fmt.Scan(&h)
	fmt.Println("Среднее арифметическое равно " + fmt.Sprint((l+m+h)/3))
}
