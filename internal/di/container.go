// Package di собирает зависимости приложения в DI-контейнере (go.uber.org/dig).
package di

import (
	"go.uber.org/dig"

	"KPO-1/internal/service"
)

// MaxVehicleDailyEnergy – лимит суточного расхода транспорта на техосмотре, кВт*ч
const MaxVehicleDailyEnergy = 3.0

// NewContainer регистрирует все зависимости: техосмотр, парк и отчёты
func NewContainer() (*dig.Container, error) {
	c := dig.New()

	providers := []any{
		func() service.ServiceCenter {
			return service.NewStandardServiceCenter(MaxVehicleDailyEnergy)
		},
		service.NewPark,
		service.NewReport,
	}
	for _, p := range providers {
		if err := c.Provide(p); err != nil {
			return nil, err
		}
	}
	return c, nil
}
