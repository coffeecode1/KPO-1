package service

import (
	"testing"

	"KPO-1/internal/domain"
)

func TestStandardServiceCenterInspect(t *testing.T) {
	center := NewStandardServiceCenter(3.0)

	tests := []struct {
		name     string
		vehicle  domain.Vehicle
		accepted bool
	}{
		{"исправный самокат", domain.NewScooter("A", 1, 7, 0.5), true},
		{"велосипед без расхода энергии", domain.NewBicycle("B", 2, 9), true},
		{"расход ровно на лимите", domain.NewElectricBike("C", 3, 5, 3.0), true},
		{"расход выше лимита", domain.NewScooter("D", 4, 7, 4.5), false},
		{"нулевой расход допустим", domain.NewScooter("E", 5, 7, 0), true},
		{"отрицательный расход", domain.NewScooter("E2", 8, 7, -1), false},
		{"простота больше 10", domain.NewBicycle("F", 6, 11), false},
		{"простота меньше 1", domain.NewBicycle("G", 7, 0), false},
		{"нет инвентарного номера", domain.NewBicycle("H", 0, 5), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := center.Inspect(tt.vehicle)
			if accepted := err == nil; accepted != tt.accepted {
				t.Errorf("ожидали принят=%v, получили ошибку: %v", tt.accepted, err)
			}
		})
	}
}
