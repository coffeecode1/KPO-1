package domain

// Helmet – шлем
type Helmet struct {
	Item
}

// NewHelmet создаёт шлем.
func NewHelmet(model string, number int) *Helmet {
	return &Helmet{Item: Item{name: "Шлем " + model, number: number}}
}

// ChargingCabinet – зарядный шкаф для аккумуляторов, нужна энергия
type ChargingCabinet struct {
	Item
	Energy
}

// NewChargingCabinet создаёт зарядный шкаф
func NewChargingCabinet(model string, number int, dailyEnergy float64) *ChargingCabinet {
	return &ChargingCabinet{
		Item:   Item{name: "Зарядный шкаф " + model, number: number},
		Energy: Energy{kwh: dailyEnergy},
	}
}

// DockStation – док-станция для парковки и зарядки самокатов, нужна энергия
type DockStation struct {
	Item
	Energy
}

// NewDockStation создаёт док-станцию
func NewDockStation(model string, number int, dailyEnergy float64) *DockStation {
	return &DockStation{
		Item:   Item{name: "Док-станция " + model, number: number},
		Energy: Energy{kwh: dailyEnergy},
	}
}
