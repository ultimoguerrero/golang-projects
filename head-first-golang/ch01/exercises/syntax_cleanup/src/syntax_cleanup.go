package main

import "fmt"

func main() {

	var pebbleWeight, rockWeight, boulderWeight float64 = 0.1, 1.2, 502.4
	totalWeight := pebbleWeight + rockWeight + boulderWeight
	fmt.Println(totalWeight)
}
