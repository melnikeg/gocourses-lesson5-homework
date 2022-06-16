package main

import "fmt"

const (
	priceOfOneAplle float32 = 5.99
	priceOfOnePear  float32 = 7
	myCash          float32 = 23
)

func main() {

	var priceOf9Apples = priceOfOneAplle * 9
	var priceOf8Pears = priceOfOnePear * 8
	var price9ApplesAnd8Pears = priceOf9Apples + priceOf8Pears
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?", price9ApplesAnd8Pears)

	var howManyApplesCanBuy = myCash / priceOfOneAplle
	fmt.Println("2. Скільки груш ми можемо купити?", int(howManyApplesCanBuy))

	var howManyPearsCanBuy = myCash / priceOfOnePear
	fmt.Println("3. Скільки яблук ми можемо купити?", int(howManyPearsCanBuy))

	var priceOf2Apples = priceOfOneAplle * 2
	var priceOf2Pears = priceOf8Pears * 2
	var priceOf2ApplesAnd2Pears = priceOf2Apples + priceOf2Pears
	var enoughCashForFruit = myCash > priceOf2ApplesAnd2Pears
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?", enoughCashForFruit)
}
