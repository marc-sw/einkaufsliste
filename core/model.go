package core

import "fmt"

type Produkt struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Menge    int    `json:"quantity"`
	Erledigt bool   `json:"done"`
}

func (p Produkt) String() string {
	var erledigtString string
	if p.Erledigt {
		erledigtString = "Done"
	} else {
		erledigtString = "Unfinished"
	}
	return fmt.Sprintf("%-25s [%-2dx] [%-10s]", p.Name, p.Menge, erledigtString)
}
