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
	"fmt"
	"gonow/location"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	baseURL = "api.sl.se/api2/TravelplannerV3/trip.json?"
	langVal = "se"
)

var apiKey string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file which is needed for API calls.")
	}
	apiKey = os.Getenv("RESEPLANERARE_KEY")
}

// SearchNow finds the next five trips between the specified places
func SearchNow(from, to location.Place) {
	fmt.Printf("Will find trip: %s => %s\n", from.Name, to.Name)
	searchTrip(from, to)
}

func searchTrip(from, to location.Place) {
	url := baseURL + "key=" + apiKey
	url += "&originExtId=" + strconv.Itoa(from.SiteID)
	url += "&destExtId=" + strconv.Itoa(to.SiteID)
	url += "&lang=" + langVal

	fmt.Println(url)
}
