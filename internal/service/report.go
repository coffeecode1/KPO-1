package service

import "KPO-1/internal/domain"

// MinBeginnerSimplicity – с какой простоты можно выдавать транспорт новичкам
const MinBeginnerSimplicity = 6

// Report отчёты по парку
type Report struct {
	park *Park
}

// NewReport создаёт новый отчёт по парку
func NewReport(park *Park) *Report {
	return &Report{park: park}
}

// VehicleCount возвращает количество транспорта
func (r *Report) VehicleCount() int {
	return len(r.park.Vehicles())
}

// ThingCount возвращает количество вещей на балансе
func (r *Report) ThingCount() int {
	return len(r.park.Things())
}

// TotalDailyEnergy возвращает суммарный суточный расход энергии, кВт*ч
func (r *Report) TotalDailyEnergy() float64 {
	total := 0.0
	for _, item := range r.Inventory() {
		if consumer, ok := item.(domain.EnergyConsumer); ok {
			total += consumer.DailyEnergy()
		}
	}
	return total
}

// BeginnerVehicles возвращает массив транспортов, которые можно выдать новичкам
func (r *Report) BeginnerVehicles() []domain.Vehicle {
	var result []domain.Vehicle
	for _, v := range r.park.Vehicles() {
		if v.Simplicity() >= MinBeginnerSimplicity {
			result = append(result, v)
		}
	}
	return result
}

// Inventory возвращает всю технику и вещи на балансе массивом
func (r *Report) Inventory() []domain.Inventoried {
	var result []domain.Inventoried
	for _, v := range r.park.Vehicles() {
		result = append(result, v)
	}
	result = append(result, r.park.Things()...)
	return result
}
