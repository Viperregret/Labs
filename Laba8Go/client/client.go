package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// User представляет модель данных пользователя
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Базовый URL API
const baseURL = "http://localhost:8081"

// Получение списка пользователей (GET /users)
func getUsers() {
	resp, err := http.Get(baseURL + "/users")
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) == 0 {
		fmt.Println("Пустой ответ от сервера при запросе списка пользователей")
	} else {
		fmt.Println("Ответ сервера:", string(body))
	}
}

// Получение информации о пользователе по ID (GET /users/{id})
func getUserByID(id int) {
	resp, err := http.Get(baseURL + "/users/" + strconv.Itoa(id))
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) == 0 {
		fmt.Println("Пустой ответ от сервера при запросе пользователя по ID")
	} else {
		fmt.Println("Ответ сервера:", string(body))
	}
}

// Создание нового пользователя (POST /users)
func createUser(name string, age int) {
	user := User{Name: name, Age: age}
	jsonData, _ := json.Marshal(user)

	resp, err := http.Post(baseURL+"/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) == 0 {
		fmt.Println("Пустой ответ от сервера при создании пользователя")
	} else {
		fmt.Println("Ответ сервера:", string(body))
	}
}

// Обновление информации о пользователе по ID (PUT /users/{id})
func updateUser(id int, name string, age int) {
	user := User{ID: id, Name: name, Age: age}
	jsonData, _ := json.Marshal(user)

	req, err := http.NewRequest(http.MethodPut, baseURL+"/users/"+strconv.Itoa(id), bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) == 0 {
		fmt.Println("Пустой ответ от сервера при обновлении пользователя")
	} else {
		fmt.Println("Ответ сервера:", string(body))
	}
}

// Удаление пользователя по ID (DELETE /users/{id})
func deleteUser(id int) {
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/users/"+strconv.Itoa(id), nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) == 0 {
		fmt.Println("Пустой ответ от сервера при удалении пользователя")
	} else {
		fmt.Println("Ответ сервера:", string(body))
	}
}

// Основная функция для выполнения запросов
func main() {
	fmt.Println("=== Запрос списка пользователей ===")
	getUsers()

	fmt.Println("\n=== Создание нового пользователя ===")
	createUser("3", 0)

	fmt.Println("\n=== Запрос списка пользователей после добавления ===")
	getUsers()

	fmt.Println("\n=== Запрос информации о пользователе с ID 1 ===")
	getUserByID(1)

	fmt.Println("\n=== Обновление информации пользователя с ID 1 ===")
	updateUser(1, "Alice Updated", 31)

	fmt.Println("\n=== Запрос списка пользователей после обновления ===")
	getUsers()

	fmt.Println("\n=== Удаление пользователя с ID 1 ===")
	deleteUser(1)

	fmt.Println("\n=== Запрос списка пользователей после удаления ===")
	getUsers()
}
