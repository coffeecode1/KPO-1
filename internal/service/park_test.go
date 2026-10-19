package service

import (
	"errors"
	"testing"

	"KPO-1/internal/domain"
)

// stubServiceCenter — подмена техосмотра: возвращает заранее заданный ответ
// и считает, сколько раз его вызвали.
type stubServiceCenter struct {
	err   error
	calls int
}

func (s *stubServiceCenter) Inspect(domain.Vehicle) error {
	s.calls++
	return s.err
}

func TestParkAcceptsVehicleWhenInspectionPassed(t *testing.T) {
	center := &stubServiceCenter{}
	park := NewPark(center)

	if err := park.AddVehicle(domain.NewScooter("A", 1, 7, 0.5)); err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
	if center.calls != 1 {
		t.Errorf("техосмотр должен вызываться один раз, вызван %d", center.calls)
	}
	if len(park.Vehicles()) != 1 {
		t.Errorf("в парке должна быть 1 единица, а их %d", len(park.Vehicles()))
	}
}

func TestParkRejectsVehicleWhenInspectionFailed(t *testing.T) {
	center := &stubServiceCenter{err: errors.New("тормоза неисправны")}
	park := NewPark(center)

	// Даже «идеальный» самокат не попадает в парк, если техосмотр отказал.
	err := park.AddVehicle(domain.NewScooter("A", 1, 10, 0.1))
	if err == nil {
		t.Fatal("ожидали отказ техосмотра")
	}
	if len(park.Vehicles()) != 0 {
		t.Error("отклонённый транспорт не должен попадать в парк")
	}
}

func TestParkRejectsDuplicateInventoryNumber(t *testing.T) {
	center := &stubServiceCenter{}
	park := NewPark(center)

	if err := park.AddThing(domain.NewHelmet("Urban", 5)); err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
	if err := park.AddVehicle(domain.NewBicycle("B", 5, 8)); err == nil {
		t.Error("транспорт с занятым номером должен отклоняться")
	}
	if err := park.AddThing(domain.NewHelmet("Urban L", 5)); err == nil {
		t.Error("вещь с занятым номером должна отклоняться")
	}
	if center.calls != 0 {
		t.Error("при занятом номере техосмотр не нужен")
	}
}

func TestParkRejectsThingWithoutInventoryNumber(t *testing.T) {
	park := NewPark(&stubServiceCenter{})

	if err := park.AddThing(domain.NewHelmet("Urban", 0)); err == nil {
		t.Error("вещь без инвентарного номера должна отклоняться")
	}
	if len(park.Things()) != 0 {
		t.Error("вещь без номера не должна попадать на баланс")
	}
}
