package core

type ProduktRepository struct {
	naechsteId int
	produkte   []Produkt
}

func NewProduktRepository() *ProduktRepository {
	return &ProduktRepository{naechsteId: 0, produkte: make([]Produkt, 0, 16)}
}

func MockProductRepository() *ProduktRepository {
	var produkte []Produkt = []Produkt{
		Produkt{0, "Milch 1.5L", 2, false},
		Produkt{1, "Brot", 1, false},
		Produkt{2, "Haferflocken 500g", 1, false},
		Produkt{3, "Olivenoel 1L", 1, false},
		Produkt{4, "Tomaten", 5, false},
		Produkt{5, "Kesselchips", 4, false},
		Produkt{6, "Nudeln 500g", 2, false},
		Produkt{7, "Zwiebeln 250g", 1, false},
	}
	return &ProduktRepository{naechsteId: len(produkte), produkte: produkte}
}

func (repository *ProduktRepository) CreateProdukt(produkt Produkt) {
	produkt.Id = repository.naechsteId
	repository.produkte = append(repository.produkte, produkt)
	repository.naechsteId++
}

func (repository *ProduktRepository) GetAllProdukte() []Produkt {
	return repository.produkte
}

func (repository *ProduktRepository) UpdateProdukt(produkt Produkt) bool {
	for i, eintrag := range repository.produkte {
		if eintrag.Id == produkt.Id {
			repository.produkte[i].Name = produkt.Name
			repository.produkte[i].Menge = produkt.Menge
			repository.produkte[i].Erledigt = produkt.Erledigt
			return true
		}
	}
	return false
}

func (repository *ProduktRepository) DeleteProdukt(id int) bool {
	for i, eintrag := range repository.produkte {
		if eintrag.Id == id {
			repository.produkte = append(repository.produkte[:i], repository.produkte[i+1:]...)
			return true
		}
	}
	return false
}
