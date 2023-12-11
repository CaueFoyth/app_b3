package entity

//Classe da ação
type Asset struct {
	ID           string
	Name         string
	MarketVolume int
}

//Adicionando a classe
func NewAsset(id string, name string, marketVolume int) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		MarketVolume: marketVolume,
	}
}
