// Package app — консольный интерфейс учёта парка «ВышКат».
package app

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"KPO-1/internal/domain"
	"KPO-1/internal/service"
)

const menu = `
1. Принять транспорт (с техосмотром)
2. Поставить вещь на баланс
3. Отчёт: число единиц и суточная энергия
4. Транспорт для новичков
5. Инвентарный список (техника и вещи)
6. Демо: все сценарии на примерах
0. Выход
`

// App читает команды пользователя и выводит результаты
type App struct {
	park   *service.Park
	report *service.Report
	in     *bufio.Scanner
	out    io.Writer
}

// New создаёт консольное приложение, ввод и вывод передаются снаружи
func New(park *service.Park, report *service.Report, in io.Reader, out io.Writer) *App {
	return &App{park: park, report: report, in: bufio.NewScanner(in), out: out}
}

// Run показывает меню, пока пользователь не выберет «Выход» или ввод не закончится
func (a *App) Run() {
	a.printf("=== ВышКат: учёт парка. «Доедем до дедлайна» ===\n")
	for {
		a.printf(menu)

		choice, ok := a.ask("Выберите пункт: ")
		if !ok {
			return
		}
		switch choice {
		case "1":
			a.addVehicle()
		case "2":
			a.addThing()
		case "3":
			a.printSummary()
		case "4":
			a.printBeginners()
		case "5":
			a.printInventory()
		case "6":
			a.runDemo()
		case "0":
			a.printf("До свидания!\n")
			return
		default:
			a.printf("Нет такого пункта меню.\n")
		}
	}
}

// addVehicle добавляет транспорт (велосипед, электровелосипед, электросамокат)
func (a *App) addVehicle() {
	titles := make([]string, len(vehicleKinds))
	for i, k := range vehicleKinds {
		titles[i] = k.title
	}
	index, ok := a.askChoice("Вид транспорта:", titles)
	if !ok {
		return
	}
	kind := vehicleKinds[index]

	model, ok := a.ask("Модель: ")
	if !ok {
		return
	}
	number, ok := a.askInt("Инвентарный номер: ")
	if !ok {
		return
	}
	simplicity, ok := a.askInt("Простота для новичка (1-10): ")
	if !ok {
		return
	}
	energy := 0.0
	if kind.electric {
		if energy, ok = a.askFloat("Суточный расход энергии, кВт·ч: "); !ok {
			return
		}
	}
	a.acceptVehicle(kind.create(model, number, simplicity, energy))
}

// addThing добавляет вещь (шлем, док-станцию, зарядный шкаф)
func (a *App) addThing() {
	titles := make([]string, len(thingKinds))
	for i, k := range thingKinds {
		titles[i] = k.title
	}
	index, ok := a.askChoice("Вид вещи:", titles)
	if !ok {
		return
	}
	kind := thingKinds[index]

	model, ok := a.ask("Модель: ")
	if !ok {
		return
	}
	number, ok := a.askInt("Инвентарный номер: ")
	if !ok {
		return
	}
	energy := 0.0
	if kind.electric {
		if energy, ok = a.askFloat("Суточный расход энергии, кВт·ч: "); !ok {
			return
		}
	}
	a.acceptThing(kind.create(model, number, energy))
}

// acceptVehicle принимает транспорт (велосипед, электровелосипед, электросамокат)
func (a *App) acceptVehicle(v domain.Vehicle) {
	if err := a.park.AddVehicle(v); err != nil {
		a.printf("ОТКЛОНЁН: %s (№%d) — %v\n", v.Name(), v.InventoryNumber(), err)
		return
	}
	a.printf("ПРИНЯТ: %s (№%d)\n", v.Name(), v.InventoryNumber())
}

// acceptThing принимает вещь (шлем, док-станцию, зарядный шкаф)
func (a *App) acceptThing(t domain.Inventoried) {
	if err := a.park.AddThing(t); err != nil {
		a.printf("ОТКЛОНЕНО: %s (№%d) — %v\n", t.Name(), t.InventoryNumber(), err)
		return
	}
	a.printf("НА БАЛАНСЕ: %s (№%d)\n", t.Name(), t.InventoryNumber())
}

// Суммарный отчет
func (a *App) printSummary() {
	a.printf("--- Отчёт по парку ---\n")
	a.printf("Транспорт: %d ед.\n", a.report.VehicleCount())
	a.printf("Вещи: %d ед.\n", a.report.ThingCount())
	a.printf("Суммарный суточный расход энергии: %.2f кВт·ч\n", a.report.TotalDailyEnergy())
}

// Транспорт для новичков
func (a *App) printBeginners() {
	a.printf("--- Транспорт для новичков (простота >= %d) ---\n", service.MinBeginnerSimplicity)
	vehicles := a.report.BeginnerVehicles()
	if len(vehicles) == 0 {
		a.printf("Подходящего транспорта нет.\n")
		return
	}
	for _, v := range vehicles {
		a.printf("№%d  %s (простота %d)\n", v.InventoryNumber(), v.Name(), v.Simplicity())
	}
}

// Список транспорта или вещей на балансе
func (a *App) printInventory() {
	a.printf("--- Инвентарный список ---\n")
	items := a.report.Inventory()
	if len(items) == 0 {
		a.printf("На балансе пока ничего нет.\n")
		return
	}
	for _, item := range items {
		a.printf("№%d  %s\n", item.InventoryNumber(), item.Name())
	}
}

// runDemo показывает все сценарии: приём, отказы техосмотра, вещи и отчёты и повторы номеров
func (a *App) runDemo() {
	a.printf("--- Демо: приём транспорта ---\n")
	a.acceptVehicle(domain.NewScooter("Ninebot Max", 101, 7, 0.5))
	a.acceptVehicle(domain.NewBicycle("Stels Navigator", 102, 9))
	a.acceptVehicle(domain.NewElectricBike("Eltreco XT", 103, 4, 0.9))
	a.acceptVehicle(domain.NewScooter("Kugoo Turbo", 104, 8, 4.5)) // слишком большой расход
	a.acceptVehicle(domain.NewScooter("Xiaomi Pro", 105, 11, 0.4)) // простота вне 1..10
	a.acceptVehicle(domain.NewBicycle("Forward", 101, 8))          // номер уже занят

	a.printf("--- Демо: вещи на баланс ---\n")
	a.acceptThing(domain.NewHelmet("Urban M", 201))
	a.acceptThing(domain.NewChargingCabinet("ЗШ-12", 202, 6.0))
	a.acceptThing(domain.NewDockStation("Dock-1", 203, 1.2))
	a.acceptThing(domain.NewHelmet("Urban L", 201)) // номер уже занят

	a.printSummary()
	a.printBeginners()
	a.printInventory()
}

// printf печатает в вывод приложения
func (a *App) printf(format string, args ...any) {
	_, _ = fmt.Fprintf(a.out, format, args...)
}

// ask выводит подсказку и читает строку. ok = false, если ввод закончился
func (a *App) ask(prompt string) (string, bool) {
	a.printf("%s", prompt)
	if !a.in.Scan() {
		a.printf("\n")
		return "", false
	}
	return strings.TrimSpace(a.in.Text()), true
}

// askInt читает целое число, переспрашивая при ошибке
func (a *App) askInt(prompt string) (int, bool) {
	for {
		text, ok := a.ask(prompt)
		if !ok {
			return 0, false
		}
		if n, err := strconv.Atoi(text); err == nil {
			return n, true
		}
		a.printf("Нужно целое число, попробуйте ещё раз.\n")
	}
}

// askFloat читает дробное число, переспрашивая при ошибке
func (a *App) askFloat(prompt string) (float64, bool) {
	for {
		text, ok := a.ask(prompt)
		if !ok {
			return 0, false
		}
		if f, err := strconv.ParseFloat(strings.Replace(text, ",", ".", 1), 64); err == nil {
			return f, true
		}
		a.printf("Нужно число, например 0.5, попробуйте ещё раз.\n")
	}
}

// askChoice показывает пронумерованный список и возвращает индекс выбранного варианта
func (a *App) askChoice(title string, options []string) (int, bool) {
	a.printf("%s\n", title)
	for i, option := range options {
		a.printf("  %d. %s\n", i+1, option)
	}
	for {
		n, ok := a.askInt("Номер: ")
		if !ok {
			return 0, false
		}
		if n >= 1 && n <= len(options) {
			return n - 1, true
		}
		a.printf("Такого варианта нет.\n")
	}
}
