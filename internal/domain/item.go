package domain

// Item – данные любого объекта на балансе: название и инвентарный номер
type Item struct {
	name   string
	number int
}

// Name возвращает название объекта
func (i Item) Name() string {
	return i.name
}

// InventoryNumber возвращает инвентарный номер объекта
func (i Item) InventoryNumber() int {
	return i.number
}

// Energy – суточный расход энергии, делает все устройства EnergyConsumer
type Energy struct {
	kwh float64
}

// DailyEnergy возвращает расход энергии за сутки в кВт*ч
func (e Energy) DailyEnergy() float64 {
	return e.kwh
}
