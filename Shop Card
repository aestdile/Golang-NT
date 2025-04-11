


package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type ShoppingCart struct {
	Items []Product
}

func (cart *ShoppingCart) AddItem(product Product, quantity int) {
	for i, item := range cart.Items {
		if item.Name == product.Name {
			cart.Items[i].Quantity += quantity
			return
		}
	}
	product.Quantity = quantity
	cart.Items = append(cart.Items, product)
}

func (cart *ShoppingCart) RemoveItem(productName string) error {
	for i, item := range cart.Items {
		if item.Name == productName {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found in the cart")
}

func (cart *ShoppingCart) UpdateQuantity(productName string, newQuantity int) error {
	for i, item := range cart.Items {
		if item.Name == productName {
			cart.Items[i].Quantity = newQuantity
			return nil
		}
	}
	return errors.New("product not found in the cart")
}

func (cart *ShoppingCart) CalculateTotal() float64 {
	total := 0.0
	for _, item := range cart.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func (cart *ShoppingCart) PrintCart() {
	fmt.Println("Savatcha:")
	for _, item := range cart.Items {
		fmt.Printf("Mahsulot: %s, Miqdori: %d, Narxi: $%.2f, Jami: $%.2f\n", item.Name, item.Quantity, item.Price, item.Price*float64(item.Quantity))
	}
}

func main() {
	var mahsulotNomi string
	var narxi float64
	var miqdori int

	fmt.Println("Mahsulot ma'lumotlarini kiriting:")
	fmt.Print("Nomi: ")
	fmt.Scanln(&mahsulotNomi)
	fmt.Print("Narxi: $")
	fmt.Scanln(&narxi)
	fmt.Print("Miqdori: ")
	fmt.Scanln(&miqdori)

	savatcha := ShoppingCart{}
	savatcha.AddItem(Product{Name: mahsulotNomi, Price: narxi}, miqdori)

	savatcha.PrintCart()

	total := savatcha.CalculateTotal()
	fmt.Printf("Jami: $%.2f\n", total)
}


































