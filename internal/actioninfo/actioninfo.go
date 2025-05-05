package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse() error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i, data := range dataset {
		// Пропускаем пустые строки
		if data == "" {
			continue
		}

		// Парсим данные
		err := dp.Parse()
		if err != nil {
			fmt.Printf("Ошибка обработки данных (строка %d): %v\n", i+1, err)
			continue
		}

		// Получаем информацию о действии
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Ошибка получения информации (строка %d): %v\n", i+1, err)
			continue
		}

		// Выводим информацию
		fmt.Printf("--- Запись %d ---\n%s\n\n", i+1, info)
	}
}
