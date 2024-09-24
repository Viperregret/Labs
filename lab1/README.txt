Лабораторная работа #1
Задание 1: Написать программу, которая выводит текущее время и дату:
today := time.Now()
fmt.Println(today)
При запуске кода мы получаем дату и время в формате:
ГГ-ММ-ДД ЧЧ:ММ:СС.МС добавление часового пояса, погрешность милисекунд

Задание 2: Создать переменные различных типов (int,float64,string,bool) и вывести их на экран
A := int(20)
B := float64(1.454231)
S := string("Hello world!")
C := bool(true)
fmt.Println(A)
fmt.Println(B)
fmt.Println(S)
fmt.Println(C)
Созданы переменные различных типов(A,B,S,C) и выведены на экран при помощи команды fmt.Println()

Задание 3: Использовать краткую форму оюъявления и вывода переменных:
A1 := 15
fmt.Println(A1)
B1 := 1.4512
fmt.Println(B1)
Используеться краткая форма объявления переменных через := и вывод fmt.Println()

Задание 4: Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов:
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
В данном примере используеться структура switch case с проверкой строковой переменной на символ (заданный пользователем), используеться метод fmt.Scan() для записи переменной.

Задание 5: Реализовать функцию для вычисления суммы и разности двух чисел с плавающей точкой
В данном задании используется пакет функций fun1 с кодом 
package fun1

func Sum(a float32, b float32) float32 {
	result := a + b
	return result
}

func Sub(a float32, b float32) float32 {
	result := a - b
	return result
}
Код main: 
var f1 float32
var f2 float32
fmt.Println("Введите первое число:")
fmt.Scan(&f1)
fmt.Println("Введите второе число:")
fmt.Scan(&f2)
fmt.Println("Сумма равна:" + fmt.Sprint(funct.Sum(f1, f2)))
fmt.Println("Разница равна:" + fmt.Sprint(funct.Sub(f1, f2)))
При запуске у пользователя запрашиваеться 2 числа и с помощью функций Sum и Sub выводятся на экран сумма и Разница

Задание 6: Написать программу, которая вычисляет среднее значение трех чисел
Код:
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

При запуске пользователь вводит три числа, на экран выводиться их среднее арифметическое