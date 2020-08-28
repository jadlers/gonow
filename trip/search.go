// Copyright © 2018 Jacob Adlers <jacob.adlers@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package trip

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jadlers/gonow/location"
)

const (
	baseURL = "https://api.sl.se/api2/TravelplannerV3_1/trip.json?"
	langVal = "se"
)

var apiKey string

func init() {
	if apiKey == "" {
		log.Fatal("No api key for trip")
	}
}

// SearchNow finds the next five trips between the specified places
func SearchNow(from, to location.Place) [][]SubTrip {
	fmt.Printf("%s => %s\n", from.Name, to.Name)
	searchURL := searchTripBetweenURL(from, to)
	foundTrips := makeSearchRequest(searchURL)
	return createSubTrips(foundTrips)
}

// SearchTime finds the next trips between the places from the given time
func SearchTime(from, to location.Place, time string) [][]SubTrip {
	fmt.Printf("%s => %s\n", from.Name, to.Name)
	searchURL := searchTripBetweenURL(from, to)
	searchURL += "&time=" + time

	foundTrips := makeSearchRequest(searchURL)
	return createSubTrips(foundTrips)
}

func searchTripBetweenURL(from, to location.Place) string {
	url := baseURL + "key=" + apiKey
	url += "&originExtId=" + strconv.Itoa(from.SiteID)
	url += "&destExtId=" + strconv.Itoa(to.SiteID)
	url += "&lang=" + langVal

	return url
}

func makeSearchRequest(url string) SearchTripResponse {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Could not complete http request\n%s\n", err)
	}
	defer res.Body.Close()

	var searchRes SearchTripResponse
	if err := json.NewDecoder(res.Body).Decode(&searchRes); err != nil {
		log.Println(err)
	}

	return searchRes
}

func createSubTrips(res SearchTripResponse) [][]SubTrip {
	var trips [][]SubTrip

	for i, trip := range res.Trip {
		trips = append(trips, []SubTrip{})
		for _, sub := range trip.LegList.Leg {
			st := SubTrip{
				Date:      sub.Origin.Date,
				DepTime:   formatTime(sub.Origin.Time),
				ArrTime:   formatTime(sub.Destination.Time),
				From:      sub.Origin.Name,
				To:        sub.Destination.Name,
				Direction: sub.Direction,
			}
			if sub.Type == "JNY" {
				st.Type = sub.Product.CatIn + " " + sub.Product.Line
			} else {
				st.Type = sub.Type
			}
			trips[i] = append(trips[i], st)
		}
	}

	return trips
}

func formatTime(time string) string {
	exploded := strings.Split(time, ":")
	return strings.Join(exploded[:2], ":")
}
