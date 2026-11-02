package actioninfo

import (
	"fmt"
	"log"
)

// DataParser определяет методы для обработки данных об активности.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает данные об активности и выводит результаты.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println("Ошибка парсинга:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка получения информации:", err)
			continue
		}

		fmt.Print(info)
	}
}
