package domain

import "testing"

func TestVehiclesKeepTheirData(t *testing.T) {
	scooter := NewScooter("Ninebot", 1, 7, 0.5)
	if scooter.Name() != "Самокат Ninebot" || scooter.InventoryNumber() != 1 ||
		scooter.Simplicity() != 7 || scooter.DailyEnergy() != 0.5 {
		t.Errorf("самокат сохранил неверные данные: %+v", scooter)
	}

	ebike := NewElectricBike("Eltreco", 2, 4, 0.9)
	if ebike.Name() != "Электровелосипед Eltreco" || ebike.DailyEnergy() != 0.9 {
		t.Errorf("электровелосипед сохранил неверные данные: %+v", ebike)
	}

	bike := NewBicycle("Stels", 3, 9)
	if bike.Name() != "Велосипед Stels" || bike.Simplicity() != 9 {
		t.Errorf("велосипед сохранил неверные данные: %+v", bike)
	}
}

func TestOnlyElectricDevicesConsumeEnergy(t *testing.T) {
	items := map[string]Inventoried{
		"самокат":          NewScooter("A", 1, 5, 0.5),
		"электровелосипед": NewElectricBike("B", 2, 5, 0.8),
		"зарядный шкаф":    NewChargingCabinet("C", 3, 6),
		"док-станция":      NewDockStation("D", 4, 1),
	}
	for name, item := range items {
		if _, ok := item.(EnergyConsumer); !ok {
			t.Errorf("%s должен быть EnergyConsumer", name)
		}
	}

	var bike Inventoried = NewBicycle("E", 5, 5)
	var helmet Inventoried = NewHelmet("F", 6)
	if _, ok := bike.(EnergyConsumer); ok {
		t.Error("велосипед не должен потреблять энергию")
	}
	if _, ok := helmet.(EnergyConsumer); ok {
		t.Error("шлем не должен потреблять энергию")
	}
}

func TestThingsHaveInventoryNumbers(t *testing.T) {
	helmet := NewHelmet("Urban", 10)
	cabinet := NewChargingCabinet("ЗШ-12", 11, 6)
	dock := NewDockStation("Dock-1", 12, 1.2)

	if helmet.Name() != "Шлем Urban" || helmet.InventoryNumber() != 10 {
		t.Errorf("неверный шлем: %+v", helmet)
	}
	if cabinet.Name() != "Зарядный шкаф ЗШ-12" || cabinet.InventoryNumber() != 11 || cabinet.DailyEnergy() != 6 {
		t.Errorf("неверный зарядный шкаф: %+v", cabinet)
	}
	if dock.Name() != "Док-станция Dock-1" || dock.InventoryNumber() != 12 || dock.DailyEnergy() != 1.2 {
		t.Errorf("неверная док-станция: %+v", dock)
	}
}
