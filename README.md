# Car Estimator Program

A simple program written in go to simulate difference use cases when
purchasing a new car.

The main intention of this project was to get more familar with go
programming langugage syntax.

## Prerequisites

1. Make sure you have golang installed. Refer to their official
[documentation](https://golang.org/doc/install) for more details.

2. Clone this repository to your workstation

```
$ git clone https://github.com/ryankwilliams/car-estimator.git
```

## How to run

You can run this in one of two ways:

1. Use `go run` against the `main.go` file.

```
$ cd car-estimator
$ go run main.go
```

2. Use `go build` to build a binary file.

```
$ cd car-estimator
$ go build
$ ./car-estimator
```

*Output*

```
$ ./car-estimator 
** Car Estimator **
Enter the car cost: 50000.00
Enter the trade in price: 39500.00
Enter the state sales tax: 0.0625
Enter the total fees (e.g. title, doc fees, etc): 500.00
Enter the down payment: 10000.00

--------------------------------------------------
New Car Cost                    : $50000
Current Car Trade In Value      : $39500
New Car Total Cost (w/trade)    : $10500
Total Sales Tax (new car total) : $656.25
Total purchasing fees           : $500
Total down payment              : $10000
--------------------------------------------------
Total Purchase Cost             : $1656.25
```