package actioninfo

import (
	"log"
)

// ActionInfo интерфейс для работы с информацией о действиях
type ActionInfo interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// ProcessAction обрабатывает действие, используя интерфейс ActionInfo
func ProcessAction(action ActionInfo, data string) (string, error) {
	if err := action.Parse(data); err != nil {
		return "", err
	}
	return action.ActionInfo()
}

// DataParser определяет интерфейс для парсинга данных о тренировках и прогулках
type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

// Info обрабатывает набор данных о тренировках или прогулках
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Парсим данные
		if err := dp.Parse(data); err != nil {
			log.Printf("Ошибка при парсинге данных: %v", err)
			continue
		}

		// Получаем и выводим информацию
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка при получении информации: %v", err)
			continue
		}

		log.Print(info)
	}
}
