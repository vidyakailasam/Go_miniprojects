package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

const (
	seleniumPath = `/Users/vidyakailasam/Downloads/chromedriver 2`
	port         = 4444
)

type hotel struct {
	Name  string
	price int
}
type flight_list []hotel

func (flights flight_list) Len() int {
	return len(flights)
}
func (flights flight_list) Swap(i, j int) {
	flights[i], flights[j] = flights[j], flights[i]
}
func (flights flight_list) Less(i, j int) bool {
	return flights[i].price < flights[j].price
}
func main() {
	flights := make(flight_list, 15)
	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	wd, err := selenium.NewRemote(caps, "")
	defer wd.Quit()
	if err != nil {
		panic(err)
	}
	if err := wd.Get("https://www.kayak.co.in/hotels/Jaipur,Rajasthan,India-c8030/2021-10-07/2021-10-08/2adults?sort=rank_a"); err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
	wes, err := wd.FindElements(selenium.ByCSSSelector, "div.FLpo-big-name")
	if err != nil {
		panic(err)
	}
	wep, err := wd.FindElements(selenium.ByCSSSelector, ".zV27-price-section")
	if err != nil {
		panic(err)
	}
	//Loop to get information for each element
	for i, we := range wes {
		text, err := we.Text()
		text1, err1 := wep[i].Text()
		text1 = text1[4:]
		text1 = strings.ReplaceAll(text1, ",", "")
		text_, _ := strconv.Atoi(text1)
		if err != nil {
			panic(err)
		}
		if err1 != nil {
			panic(err)
		}
		flights[i] = hotel{
			Name:  text,
			price: text_,
		}
	}
	defer service.Stop()
	sort.Sort(flights)
	//for _, k := range flights {
	//fmt.Printf("%v\t%v\n", k.Name, k.price)
	//}
	fmt.Println("The cheapest Flight available from Madurai to Bangalore :")
	fmt.Printf("Airline Name :%v  Price :%d", flights[0].Name, flights[0].price)
}

//.Flights-Results-FlightPriceSection.sleek .Theme-featured-large .price
//.Flights-Results-FlightPriceSection.right-alignment .multibook-dropdownpackage main
