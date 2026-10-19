package service

import (
	"math"
	"testing"

	"KPO-1/internal/domain"
)

// newFilledReport создаёт парк с техосмотром-заглушкой, который принимает всё.
func newFilledReport(t *testing.T) *Report {
	t.Helper()
	park := NewPark(&stubServiceCenter{})
	vehicles := []domain.Vehicle{
		domain.NewScooter("A", 1, 7, 0.5),
		domain.NewBicycle("B", 2, 6),
		domain.NewElectricBike("C", 3, 5, 0.9),
	}
	for _, v := range vehicles {
		if err := park.AddVehicle(v); err != nil {
			t.Fatal(err)
		}
	}
	things := []domain.Inventoried{
		domain.NewHelmet("D", 4),
		domain.NewChargingCabinet("E", 5, 6.0),
	}
	for _, thing := range things {
		if err := park.AddThing(thing); err != nil {
			t.Fatal(err)
		}
	}
	return NewReport(park)
}

func TestReportCounts(t *testing.T) {
	report := newFilledReport(t)

	if report.VehicleCount() != 3 {
		t.Errorf("транспорт: ожидали 3, получили %d", report.VehicleCount())
	}
	if report.ThingCount() != 2 {
		t.Errorf("вещи: ожидали 2, получили %d", report.ThingCount())
	}
}

func TestReportTotalDailyEnergy(t *testing.T) {
	report := newFilledReport(t)

	// 0.5 (самокат) + 0.9 (электровелосипед) + 6.0 (зарядный шкаф); велосипед и шлем не считаются.
	want := 7.4
	if got := report.TotalDailyEnergy(); math.Abs(got-want) > 1e-9 {
		t.Errorf("ожидали %.2f кВт·ч, получили %.2f", want, got)
	}
}

func TestReportBeginnerVehiclesUsesThresholdSix(t *testing.T) {
	report := newFilledReport(t)

	beginners := report.BeginnerVehicles()
	if len(beginners) != 2 {
		t.Fatalf("ожидали 2 единицы для новичков, получили %d", len(beginners))
	}
	// Простота 7 и 6 подходят, простота 5 — нет.
	if beginners[0].InventoryNumber() != 1 || beginners[1].InventoryNumber() != 2 {
		t.Errorf("неверный список для новичков: №%d, №%d",
			beginners[0].InventoryNumber(), beginners[1].InventoryNumber())
	}
}

func TestReportInventoryContainsVehiclesAndThings(t *testing.T) {
	report := newFilledReport(t)

	inventory := report.Inventory()
	if len(inventory) != 5 {
		t.Fatalf("ожидали 5 позиций, получили %d", len(inventory))
	}
	for i, item := range inventory {
		if item.InventoryNumber() != i+1 {
			t.Errorf("позиция %d: ожидали №%d, получили №%d", i, i+1, item.InventoryNumber())
		}
	}
}

func TestReportOnEmptyPark(t *testing.T) {
	report := NewReport(NewPark(&stubServiceCenter{}))

	if report.TotalDailyEnergy() != 0 || len(report.BeginnerVehicles()) != 0 || len(report.Inventory()) != 0 {
		t.Error("пустой парк должен давать пустые отчёты")
	}
}
