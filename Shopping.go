





package main

import (
	"fmt"
	"time"
)

type Product struct {
	id          int
	name        string
	price       int
	expiredDate string
	isNew       bool
	barCode     int
	madeIn      string
}

type Order struct {
	id    int
	date  string
	prod  []*Product
}

type User struct {
	id     int
	name   string
	orders []*Order
}

func main() {
	products := []*Product{
		{
			id:          1,
			name:        "Fanta",
			price:       12000,
			expiredDate: "2024-04-03",
			isNew:       true,
			barCode:     534641,
			madeIn:      "Uzbekistan",
		},
		{
			id:          2,
			name:        "Coca-Cola",
			price:       12000,
			expiredDate: "2024-03-27",
			isNew:       false,
			barCode:     534642,
			madeIn:      "Uzbekistan",
		},
		{
			id:          3,
			name:        "Pepsi",
			price:       12000,
			expiredDate: "2024-04-01",
			isNew:       true,
			barCode:     534643,
			madeIn:      "Uzbekistan",
		},
		{
			id:          4,
			name:        "Dinay",
			price:       8000,
			expiredDate: "2024-02-16",
			isNew:       true,
			barCode:     534644,
			madeIn:      "Uzbekistan",
		},
		{
			id:          5,
			name:        "Moxito",
			price:       7000,
			expiredDate: "2024-01-14",
			isNew:       false,
			barCode:     534645,
			madeIn:      "Uzbekistan",
		},
	}

	var user User
	user.id = 1
	user.name = "John"

	var basket []*Product
	var orders []*Order

	for {
		fmt.Println("1. Mahsulot tanlash\n2. Savatni ko'rish\n3. Buyurtmani rasmiylashtirish\n4. Buyurtmalarim")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			getAllProducts(products)
			fmt.Println("Mahsulot idsini kiriting:") // Mahsulot id sini kiritganingizdan keyin, "Buyurtmani rasmiylashtirish" ni tanlang, aks holda, "Buyurtmalarim" da chiqmaydi.
			fmt.Scan(&choice)

			prod := getProduct(products, choice)
			basket = append(basket, &prod)
		case 2:
			fmt.Println("Savat:")
			for _, item := range basket {
				fmt.Printf("%+v\n", *item)
			}
		case 3:
			order := Order{
				id:   len(orders) + 1,
				date: time.Now().Format("2006-01-02"),
				prod: make([]*Product, len(basket)),
			}
			copy(order.prod, basket)
			orders = append(orders, &order)
			basket = nil
			fmt.Println("Buyurtma rasmiylashtirildi.")
		case 4:
			fmt.Println("Buyurtmalarim:")
			for _, ord := range orders {
				fmt.Printf("Buyurtma #%d, Sana: %s\n", ord.id, ord.date)
				for _, prod := range ord.prod {
					fmt.Printf("\t%+v\n", *prod)
	
				}
			}
		}
	}
}

func getAllProducts(prs []*Product) {
	for _, v := range prs {
		fmt.Printf("%+v\n", *v)
	}
}

func getProduct(pr []*Product, id int) Product {
	for _, v := range pr {
		if v.id == id {
			return *v
		}
	}
	return Product{}
}








































