package app

import "KPO-1/internal/domain"

// vehicleKind описывает вид транспорта для меню.
// Чтобы добавить новый вид, достаточно создать тип в domain и добавить строку сюда
type vehicleKind struct {
	title    string
	electric bool
	create   func(model string, number, simplicity int, energy float64) domain.Vehicle
}

var vehicleKinds = []vehicleKind{
	{"Электросамокат", true, func(model string, number, simplicity int, energy float64) domain.Vehicle {
		return domain.NewScooter(model, number, simplicity, energy)
	}},
	{"Велосипед", false, func(model string, number, simplicity int, _ float64) domain.Vehicle {
		return domain.NewBicycle(model, number, simplicity)
	}},
	{"Электровелосипед", true, func(model string, number, simplicity int, energy float64) domain.Vehicle {
		return domain.NewElectricBike(model, number, simplicity, energy)
	}},
}

// thingKind описывает вид вещи для меню
type thingKind struct {
	title    string
	electric bool
	create   func(model string, number int, energy float64) domain.Inventoried
}

var thingKinds = []thingKind{
	{"Шлем", false, func(model string, number int, _ float64) domain.Inventoried {
		return domain.NewHelmet(model, number)
	}},
	{"Зарядный шкаф", true, func(model string, number int, energy float64) domain.Inventoried {
		return domain.NewChargingCabinet(model, number, energy)
	}},
	{"Док-станция", true, func(model string, number int, energy float64) domain.Inventoried {
		return domain.NewDockStation(model, number, energy)
	}},
}
