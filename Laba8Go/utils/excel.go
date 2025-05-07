package utils

import (
	"Laba8Go/models"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// LoadUsers загружает пользователей из Excel файла
func LoadUsers() ([]models.User, error) {
	f, err := excelize.OpenFile("users.xlsx") // Убедитесь, что файл существует
	if err != nil {
		return nil, err
	}
	rows, err := f.GetRows("Sheet1") // Предполагаем, что данные на первом листе
	if err != nil {
		return nil, err
	}

	var users []models.User
	for _, row := range rows[1:] { // Пропускаем заголовок
		id := len(users) + 1
		name := row[0]
		age, _ := strconv.Atoi(row[1]) // Здесь нужно обработать ошибку
		users = append(users, models.User{ID: id, Name: name, Age: age})
	}

	return users, nil
}

// SaveUsers сохраняет пользователей в Excel файл
func SaveUsers(users []models.User) error {
	f := excelize.NewFile()
	f.NewSheet("Sheet1")

	// Записываем заголовки
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "Age")

	// Записываем пользователей
	for i, user := range users {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), user.ID)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), user.Name)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), user.Age)
	}

	return f.SaveAs("users.xlsx")
}
