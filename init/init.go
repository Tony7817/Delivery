package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"toaster/types"
)

const idNum int = 1000
const maxWeight int = 100

func main() {
	generateSql()
}

func generateSql() {
	rand.Seed(time.Now().UnixNano())

	var weightIntegers []int
	var ids []int

	// generate user id from 1 to 1000.
	for i := 1; i <= idNum; i++ {
		ids = append(ids, i)
	}

	// generate weights.
	var counter int
	for i := 1; i <= maxWeight; i++ {
		for j := 0; j < 19305/i; j++ {
			counter++
			weightIntegers = append(weightIntegers, i)
		}
	}

	var weights []float64
	for {
		if len(weightIntegers) == 0 {
			break
		}
		randIndex := rand.Intn(len(weightIntegers))
		weights = append(weights, rand.Float64()+float64(weightIntegers[randIndex]-1))
		weightIntegers = append(weightIntegers[:randIndex], weightIntegers[randIndex+1:]...)
	}

	// generate orders.
	os.Stdout.Write([]byte("create table `order` ( id primary key not null, uid int not null, weight int not null, created_at datetime not null);\n"))
	for i := 1; i <= 100000; i++ {
		uidIndex := rand.Intn(len(ids))
		uid := ids[uidIndex]
		order := types.Order{
			Id:        i,
			Uid:       uid,
			Weight:    weights[i],
			CreatedAt: time.Now(),
		}
		os.Stdout.Write([]byte(fmt.Sprintf("insert into `order` (id, uid, weight, created_at) values (%d, %d, %.1f, time('now'));\n", order.Id, order.Uid, order.Weight)))
	}
}
