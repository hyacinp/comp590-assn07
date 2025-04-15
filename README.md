# Sleeping Barber Go Program

**TEAM: Kibby Hyacinth Pangilinan**

## Design Rationale

Because the assignment asks us to implement a barber, receptionist and cutomer processes, I implemented each as a function. The `barber`, `customerProcess`, `receptionist` functions are invoked as goroutines, allowing them to behave as concurrent processes.

The assignment also asks for a waiting room. To do so, I made the `waitingRoom` as a buffered channel, naturally enforcing a FIFO queue behavior where each customer is assisted in the order in which they arrived. 

### Customers Arriving at Random Times

To ensure that each customer arrives at random time intervals, I used Go's `math/rand` and `time.Sleep`. I used the two to randomize the time between intervals to be ranging from 300ms to 1000ms; this way, the customers arrived at staggered intervals, and not fixed intervals, reflecting more of a real world scenario. This can be seen within the `main` function at line 53.

### Random Duration for Haircut

To model a realistic situation where haircut duration varies from customer to customer, I took the same steps as mentioned above. Rather, here I made duration of haircuts vary between 1 and 4 seconds through Go's `math/rand`. This ensures that the haircuts do not all take up the same amount of time. This randomization is implemented within the`barber` function at line 36.

### The simulation runs forever

To satisfy the requirement of having the simulation run forever, creating new customers, I created an infinite for loop. Each iteration of this loop generates a new customer with an id, and initiates a goroutine to process the new customer.


## Running the Program

In the terminal, run `go run main.go` and watch the program run indefinitely.

To exit the program, enter `Ctrl+C` in the terminal.