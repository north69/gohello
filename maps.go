package main

import "fmt"

type DonutShop struct {
	donuts    map[string]int
	customers map[string]float32
}

func (s *DonutShop) init() {
	s.donuts = map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}
	s.customers = make(map[string]float32)
}

func (s DonutShop) calculatePrice(count int) float32 {
	return float32(count) * 1.50
}

func (s DonutShop) placeOrder(name string, kind string, count int) {
	s.customers[name] = s.calculatePrice(count)
	s.donuts[kind] = s.donuts[kind] - count
}

func (s DonutShop) checkout(name string) {
	fmt.Printf("%s please pay %f\n", name, s.customers[name])
}

func placeOrder() {
	var name, kind string
	var count int

	fmt.Println("Please input a name of a guest")
	fmt.Scan(&name)

	fmt.Println("Please input a product name [frosted, chocolate, jelly]")
	fmt.Scan(&kind)

	fmt.Println("Please input a quantity of a product")
	fmt.Scan(&count)

	var donutShop = new(DonutShop)

	donutShop.init()
	donutShop.placeOrder(name, kind, count)
	donutShop.checkout(name)
	fmt.Println(donutShop.customers)
	fmt.Println(donutShop.donuts)
}

func main() {
	for {
		placeOrder()
	}
}
