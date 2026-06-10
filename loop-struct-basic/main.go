package main

import "fmt"

type Item struct {
	Name  string
	Price float64
}

func main() {
	// Create a list of shopping items
	cart := []Item{
		{Name: "Spaghetti", Price: 12.50},
		{Name: "Mercedes Benz", Price: 9000},
		{Name: "HDMI Cable", Price: 12.50},
	}

	var total float64 = 0

	fmt.Println("---- Your shopping cart ----")

	// 'range' loops through arrays, slices or maps easily
	for _, item := range cart {
		fmt.Printf("- %s: $%.2f\n", item.Name, item.Price)
		total += item.Price
	}

	fmt.Printf("\nTotal Due: $%.2f\n", total)

	// A quick conditional check
	if total > 100 {
		fmt.Println("🎉 You qualify for free shopping!")
	} else {
		fmt.Println("Add more items for free shopping")
	}

}
