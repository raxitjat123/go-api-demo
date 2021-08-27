package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	for i := range a {
		fmt.Println("this is first funtion Loop", a[i])
	}

	countryState := map[string]string{
		"India":       "Gujarat",
		"USA":         "London",
		"Afghanistan": "kabul",
	}
	for country := range countryState {
		fmt.Println("country Name is", country, "and state Name ", countryState[country])
	}

}

/* output
this is first funtion Loop 1
this is first funtion Loop 2
this is first funtion Loop 3
this is first funtion Loop 4
this is first funtion Loop 5


country Name is USA and state Name  London
country Name is Afghanistan and state Name  kabul
country Name is India and state Name  Gujarat
*/
