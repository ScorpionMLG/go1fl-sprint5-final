package actioninfo

import (
	"fmt"
	"log"
)

// DataParser определяет контракт для парсинга данных и получения статистики.
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info обрабатывает датасет через DataParser: парсит все строки, выводит статистику.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println(err)
		}
		info, err := dp.ActionInfo()
		fmt.Println(info)
		if err != nil {
			log.Println(err)
		}
	}
}
