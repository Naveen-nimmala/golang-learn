package main

import "fmt"

func main() {
	india := country{
		Name:       "India",
		Population: 130,
		courency:   "Rupees",
	}
	malaysia := country{
		Name:       "malaysia",
		Population: 2,
		courency:   "Ringgt",
	}
	asia := region{
		Name:      "Asia",
		countries: 10,
	}
	africa := region{
		Name:      "Africa",
		countries: 50,
	}
	india.prints("naveen")
	malaysia.prints("sahithi")
	asia.prints()
	africa.prints()
	fmt.Println(india)

}
