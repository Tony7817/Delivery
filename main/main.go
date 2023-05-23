package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"toaster/types"
	"toaster/util"
)

const file string = "order.db"

func main() {
	var uid int
	flag.IntVar(&uid, "user-id", 0, "user id from 0 to 1000")
	flag.Parse()

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalln(err)
	}

	row, err := db.Query("select * from `order` where uid = ?", uid)
	if err != nil {
		log.Fatalln(err)
	}

	var orders []types.Order
	for row.Next() {
		var order types.Order
		err := row.Scan(&order.Id, &order.Uid, &order.Weight, &order.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		orders = append(orders, order)
	}

	for _, v := range orders {
		price, err := util.GetPrice(v.Weight)
		if err != nil {
			fmt.Println(fmt.Sprintf("id: %d, error: %s\n", v.Id, err.Error()))
			continue
		}
		fmt.Println(fmt.Sprintf("order id: %d, price: %.2f\n", v.Id, price))
	}

	if len(orders) == 0 {
		fmt.Println("no user with id ", uid)
	}
	row.Close()
}
