package service

import (
	"errors"
	"fmt"

	"KPO-1/internal/domain"
)

// Park хранит транспорт и вещи, стоящие на балансе. Транспорт принимается только если прошел техосмотр
type Park struct {
	center   ServiceCenter
	vehicles []domain.Vehicle
	things   []domain.Inventoried
}

// NewPark создаёт пустой парк. Техосмотр передаётся снаружи
func NewPark(center ServiceCenter) *Park {
	return &Park{center: center}
}

// AddVehicle отправляет транспорт на техосмотр и если все норм, то ставит на баланс
func (p *Park) AddVehicle(v domain.Vehicle) error {
	if p.hasNumber(v.InventoryNumber()) {
		return fmt.Errorf("инвентарный номер %d уже занят", v.InventoryNumber())
	}
	if err := p.center.Inspect(v); err != nil {
		return fmt.Errorf("техосмотр не пройден: %w", err)
	}
	p.vehicles = append(p.vehicles, v)
	return nil
}

// AddThing ставит вещь (шлем, док-станцию, зарядный шкаф) на баланс
func (p *Park) AddThing(t domain.Inventoried) error {
	if t.InventoryNumber() <= 0 {
		return errors.New("нет инвентарного номера")
	}
	if p.hasNumber(t.InventoryNumber()) {
		return fmt.Errorf("инвентарный номер %d уже занят", t.InventoryNumber())
	}
	p.things = append(p.things, t)
	return nil
}

// Vehicles возвращает принятый транспорт
func (p *Park) Vehicles() []domain.Vehicle {
	return p.vehicles
}

// Things возвращает вещи на балансе
func (p *Park) Things() []domain.Inventoried {
	return p.things
}

// hasNumber проверяет занят ли инвентарный номер, локальная функция
func (p *Park) hasNumber(number int) bool {
	for _, v := range p.vehicles {
		if v.InventoryNumber() == number {
			return true
		}
	}
	for _, t := range p.things {
		if t.InventoryNumber() == number {
			return true
		}
	}
	return false
}
