package domain

// Inventoried – всё что стоит на балансе, траспорт и вещи
type Inventoried interface {
	Name() string
	InventoryNumber() int
}

// EnergyConsumer – вещь, которая потребляет электроэнергию
type EnergyConsumer interface {
	DailyEnergy() float64 // суточный расход энергии
}

// Vehicle – транспорт на прокат
type Vehicle interface {
	Inventoried
	Simplicity() int // простота для новичка, 1 <= x <= 10
}
