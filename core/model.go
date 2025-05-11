package core

import "fmt"

type Produkt struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Menge    int    `json:"menge"`
	Erledigt bool   `json:"erledigt"`
}

func (p Produkt) String() string {
	var erledigtString string
	if p.Erledigt {
		erledigtString = "Erledigt"
	} else {
		erledigtString = "Unerledigt"
	}
	return fmt.Sprintf("%-25s [%-2dx] [%-10s]", p.Name, p.Menge, erledigtString)
}
