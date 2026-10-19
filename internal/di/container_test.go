package di

import (
	"errors"
	"testing"

	"KPO-1/internal/domain"
	"KPO-1/internal/service"
)

// rejectAllCenter – подмена техосмотра, которая отклоняет любой транспорт.
type rejectAllCenter struct{}

func (rejectAllCenter) Inspect(domain.Vehicle) error {
	return errors.New("техосмотр закрыт")
}

func TestContainerBuildsParkWithStandardServiceCenter(t *testing.T) {
	c, err := NewContainer()
	if err != nil {
		t.Fatal(err)
	}

	err = c.Invoke(func(park *service.Park, report *service.Report) {
		if err := park.AddVehicle(domain.NewScooter("A", 1, 7, 0.5)); err != nil {
			t.Errorf("исправный самокат должен приниматься: %v", err)
		}
		if err := park.AddVehicle(domain.NewScooter("B", 2, 7, MaxVehicleDailyEnergy+1)); err == nil {
			t.Error("самокат с расходом выше лимита должен отклоняться")
		}
		// Отчёт получает тот же вариант парка
		if report.VehicleCount() != 1 {
			t.Errorf("ожидали 1 единицу в отчёте, получили %d", report.VehicleCount())
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestContainerServiceCenterCanBeReplaced(t *testing.T) {
	c, err := NewContainer()
	if err != nil {
		t.Fatal(err)
	}

	// Подменяем зарегистрированный техосмотр на заглушку прямо в контейнере
	err = c.Decorate(func(service.ServiceCenter) service.ServiceCenter {
		return rejectAllCenter{}
	})
	if err != nil {
		t.Fatal(err)
	}

	err = c.Invoke(func(park *service.Park) {
		if err := park.AddVehicle(domain.NewScooter("A", 1, 7, 0.5)); err == nil {
			t.Error("с подменённым техосмотром транспорт должен отклоняться")
		}
		if len(park.Vehicles()) != 0 {
			t.Error("парк должен остаться пустым")
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}
