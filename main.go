package main

// Team: Kibby Hyacinth Pangilinan

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	waitingRoomCap = 6
)

type Customer struct {
	id int
}

func receptionist(customer Customer, waitingRoom chan Customer) {
	fmt.Printf("Receptionist greets Customer #%d\n", customer.id)
	select {
	case waitingRoom <- customer:
		fmt.Printf("Customer #%d sits in the waiting room\n", customer.id)
	default:
		fmt.Printf("Waiting room is at capacity --- Customer #%d leaves the shop\n", customer.id)
	}
}
func customerProcess(customer Customer, waitingRoom chan Customer) {
	receptionist(customer, waitingRoom)
}

func barber(waitingRoom chan Customer) {
	for {
		customer := <-waitingRoom
		fmt.Printf("Barber starts cutting hair of Customer #%d\n", customer.id)

		time.Sleep(time.Duration(rand.Intn(3000)+1000) * time.Millisecond)

		fmt.Printf("Barber finished cutting hair of Customer #%d\n", customer.id)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	waitingRoom := make(chan Customer, waitingRoomCap)
	customerID := 0
	var wg sync.WaitGroup

	go barber(waitingRoom)

	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(700)+300) * time.Millisecond)
			customerID++
			go customerProcess(Customer{id: customerID}, waitingRoom)
		}
	}()

	wg.Add(1)
	wg.Wait()
}
