package main

import "fmt"

//Кілограмів їжі на 1 кг ваги тварини
const (
	cowFood = 25
	dogFood = 2
	catFood = 7
)

func main() {
	animals := []animal{
		dog{"Patron", 4},
		cow{"Byrenka", 700},
		cat{"Barsik", 4},
		cat{"Pusha", 3},
		dog{"Tyzik", 8},
	}
	var sum int
	for _, t := range animals {
		res := t.getWeightKG()
		sum += res
		fmt.Printf("\nДля %v, з вагою %v треба %v кг корму ", t.getName(), t.getWeight(), res)
	}
	fmt.Printf("\n")
	fmt.Println("Всього корму треба ", sum)
}

type dog struct {
	name   string
	weight int
}

func (d dog) getName() string {
	return d.name
}
func (d dog) getWeightKG() int {
	return d.weight * dogFood
}
func (d dog) getWeight() int {
	return d.weight
}

type cat struct {
	name   string
	weight int
}

func (c cat) getName() string {
	return c.name
}
func (c cat) getWeightKG() int {
	return c.weight * catFood
}
func (c cat) getWeight() int {
	return c.weight
}

type cow struct {
	name   string
	weight int
}

func (c cow) getName() string {
	return c.name
}
func (c cow) getWeightKG() int {
	return c.weight * cowFood
}
func (c cow) getWeight() int {
	return c.weight
}

//Повертаэ ім'я тварини
type nameGetter interface {
	getName() string
}

//Повертає кількість їжі на місяц для цієї тварини
type weightGetterKG interface {
	getWeightKG() int
}

//Маса тварини
type weightGetter interface {
	getWeight() int
}

type animal interface {
	weightGetterKG
	weightGetter
	nameGetter
}
