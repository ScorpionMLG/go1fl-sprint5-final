package personaldata

import "fmt"

// Personal содержит персональные данные для расчёта энергозатрат.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит персональные данные в формате "Имя, Вес, Рост".
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)
}
