package app

import (
	"bytes"
	"strings"
	"testing"

	"KPO-1/internal/service"
)

// run запускает приложение с заданным вводом и возвращает весь вывод.
func run(input string) string {
	park := service.NewPark(service.NewStandardServiceCenter(3.0))
	var out bytes.Buffer
	New(park, service.NewReport(park), strings.NewReader(input), &out).Run()
	return out.String()
}

func assertContains(t *testing.T, output string, parts ...string) {
	t.Helper()
	for _, part := range parts {
		if !strings.Contains(output, part) {
			t.Errorf("в выводе нет %q", part)
		}
	}
}

func TestAddVehicleFromMenu(t *testing.T) {
	// Самокат: неверный номер вида (7), потом текст вместо числа, потом верные данные.
	input := "1\n7\n1\nNinebot\nабв\n101\n7\n0,5\n4\n0\n"
	output := run(input)

	assertContains(t, output,
		"Такого варианта нет.",
		"Нужно целое число",
		"ПРИНЯТ: Самокат Ninebot (№101)",
		"№101  Самокат Ninebot (простота 7)",
		"До свидания!",
	)
}

func TestRejectedVehicleFromMenu(t *testing.T) {
	// Велосипед с простотой 15 не проходит техосмотр.
	output := run("1\n2\nStels\n5\n15\n3\n")

	assertContains(t, output, "ОТКЛОНЁН: Велосипед Stels (№5)", "Транспорт: 0 ед.")
}

func TestAddThingFromMenu(t *testing.T) {
	// Зарядный шкаф: сначала неверный расход, потом верный.
	input := "2\n2\nЗШ-12\n202\nмного\n6\n2\n1\nUrban\n202\n5\n3\n"
	output := run(input)

	assertContains(t, output,
		"Нужно число",
		"НА БАЛАНСЕ: Зарядный шкаф ЗШ-12 (№202)",
		"ОТКЛОНЕНО: Шлем Urban (№202)",
		"№202  Зарядный шкаф ЗШ-12",
		"Суммарный суточный расход энергии: 6.00 кВт·ч",
	)
}

func TestEmptyReports(t *testing.T) {
	output := run("4\n5\n")

	assertContains(t, output, "Подходящего транспорта нет.", "На балансе пока ничего нет.")
}

func TestDemoShowsAllScenarios(t *testing.T) {
	output := run("6\n")

	assertContains(t, output,
		"ПРИНЯТ: Самокат Ninebot Max (№101)",
		"ОТКЛОНЁН: Самокат Kugoo Turbo (№104)",
		"ОТКЛОНЁН: Самокат Xiaomi Pro (№105)",
		"инвентарный номер 101 уже занят",
		"НА БАЛАНСЕ: Шлем Urban M (№201)",
		"Транспорт: 3 ед.",
		"Вещи: 3 ед.",
		"Суммарный суточный расход энергии: 8.60 кВт·ч",
		"№102  Велосипед Stels Navigator (простота 9)",
		"№203  Док-станция Dock-1",
	)
}

func TestUnknownMenuItemAndEndOfInput(t *testing.T) {
	// Неизвестный пункт, затем ввод обрывается посреди добавления транспорта.
	output := run("9\n1\n1\nNinebot\n")

	assertContains(t, output, "Нет такого пункта меню.")
	if strings.Contains(output, "ПРИНЯТ") {
		t.Error("при оборванном вводе ничего не должно добавляться")
	}
}
