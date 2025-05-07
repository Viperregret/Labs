package main

import (
	"fmt"
)

func minStairFlights(n, t int, floors []int, k int) int {
	targetFloor := floors[k-1] // этаж уходящего сотрудника (индексация с 1)
	minFloor := floors[0]
	maxFloor := floors[n-1]

	// Вариант 1: Начинаем с minFloor, идём к maxFloor
	// Успеваем ли дойти до targetFloor за <= t шагов?
	dist1 := maxFloor - minFloor
	timeToTarget1 := targetFloor - minFloor
	if timeToTarget1 <= t {
		// Успеваем, общее расстояние = dist1
	} else {
		// Не успеваем, нужно выбрать другой вариант
		dist1 = -1 // помечаем как невалидный
	}

	// Вариант 2: Начинаем с maxFloor, идём к minFloor
	// Успеваем ли дойти до targetFloor за <= t шагов?
	dist2 := maxFloor - minFloor
	timeToTarget2 := maxFloor - targetFloor
	if timeToTarget2 <= t {
		// Успеваем, общее расстояние = dist2
	} else {
		// Не успеваем
		dist2 = -1
	}

	// Вариант 3: Начинаем на targetFloor
	// Сначала идём к ближайшему крайнему этажу, потом к другому
	dist3 := (targetFloor - minFloor) + (maxFloor - minFloor)
	timeToTarget3 := 0 // уже на targetFloor
	if timeToTarget3 <= t {
		// Успеваем
	} else {
		dist3 = -1
	}

	// Выбираем минимальное расстояние из возможных вариантов
	minDist := maxFloor - minFloor // худший случай
	if dist1 != -1 && dist1 < minDist {
		minDist = dist1
	}
	if dist2 != -1 && dist2 < minDist {
		minDist = dist2
	}
	if dist3 != -1 && dist3 < minDist {
		minDist = dist3
	}

	return minDist
}

func main() {
	var n, t, k int
	fmt.Scan(&n, &t)
	floors := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&floors[i])
	}
	fmt.Scan(&k)

	fmt.Println(minStairFlights(n, t, floors, k))
}
