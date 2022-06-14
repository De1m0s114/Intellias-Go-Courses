package main

import "fmt"

func main() {
	const (
		apple_price = 5.99
		pear_price  = 7
		money       = 23
	)
	var (
		pear  int
		apple int
		res   float64
	)
	res = (money / apple_price)
	pear = money / pear_price
	apple = int(res)
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Println("Ціна 9 яблук:", 9*apple_price)
	fmt.Println("Ціна 8 груш:", 8*pear_price)
	fmt.Println("Скільки груш ми можемо купити?")
	fmt.Println("Можна купити", pear, "груші")
	fmt.Println("Скільки яблук ми можемо купити?")
	fmt.Println("Можна купити", apple, "яблука")
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?")
	if 2*apple_price+2*pear_price <= 23 {
		fmt.Println("Можна ")
	} else {
		fmt.Println("Не можна ")

	}

}
