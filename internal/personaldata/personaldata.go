package personaldata

import "fmt"

type Personal struct {
	Name string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.1f кг\n", p.Weight)
	fmt.Printf("Рост: %.1f см\n", p.Height)
}
