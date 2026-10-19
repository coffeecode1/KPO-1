package domain

// transport – структура всех транспортов
type transport struct {
	Item
	simplicity int
}

func newTransport(kind, model string, number, simplicity int) transport {
	return transport{Item: Item{name: kind + " " + model, number: number}, simplicity: simplicity}
}

// Simplicity – простота для новичка от 1 до 10 включительно
func (t transport) Simplicity() int {
	return t.simplicity
}

// Scooter – электросамокат, нужна энергия
type Scooter struct {
	transport
	Energy
}

// NewScooter создаёт электросамокат
func NewScooter(model string, number, simplicity int, dailyEnergy float64) *Scooter {
	return &Scooter{
		transport: newTransport("Самокат", model, number, simplicity),
		Energy:    Energy{kwh: dailyEnergy},
	}
}

// Bicycle – велосипед
type Bicycle struct {
	transport
}

// NewBicycle создаёт велосипед
func NewBicycle(model string, number, simplicity int) *Bicycle {
	return &Bicycle{
		transport: newTransport("Велосипед", model, number, simplicity),
	}
}

// ElectricBike – электровелосипед, нужна энергия
type ElectricBike struct {
	transport
	Energy
}

// NewElectricBike создаёт электровелосипед
func NewElectricBike(model string, number, simplicity int, dailyEnergy float64) *ElectricBike {
	return &ElectricBike{
		transport: newTransport("Электровелосипед", model, number, simplicity),
		Energy:    Energy{kwh: dailyEnergy},
	}
}
