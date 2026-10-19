// Package service содержит бизнес-логику: техосмотр, учёт парка и отчёты
package service

import (
	"errors"
	"fmt"

	"KPO-1/internal/domain"
)

// ServiceCenter определяет, можно ли принять транспорт в парк
type ServiceCenter interface {
	Inspect(v domain.Vehicle) error
}

// StandardServiceCenter – обычный техосмотр
type StandardServiceCenter struct {
	maxDailyEnergy float64 // максимально допустимый суточный расход, кВт*ч
}

// NewStandardServiceCenter создаёт новый техосмотр с лимитом
func NewStandardServiceCenter(maxDailyEnergy float64) *StandardServiceCenter {
	return &StandardServiceCenter{maxDailyEnergy: maxDailyEnergy}
}

// Inspect проверяет инвентарный номер, простоту и расход энергии
func (c *StandardServiceCenter) Inspect(v domain.Vehicle) error {
	if v.InventoryNumber() <= 0 {
		return errors.New("нет инвентарного номера")
	}
	if v.Simplicity() < 1 || v.Simplicity() > 10 {
		return fmt.Errorf("простота %d вне диапазона от 1 до 10", v.Simplicity())
	}
	if consumer, ok := v.(domain.EnergyConsumer); ok {
		energy := consumer.DailyEnergy()
		if energy < 0 || energy > c.maxDailyEnergy {
			return fmt.Errorf("расход %.2f кВт*ч выходит за пределы [0; %.2f]", energy, c.maxDailyEnergy)
		}
	}
	return nil
}
