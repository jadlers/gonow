// Copyright © 2018 Jacob Adlers <jacob.adlers@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the \"Software\"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package location

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const (
	baseURL         = "https://api.sl.se/api2/typeahead.json?"
	stationsOnlyVal = true
	maxResultsVal   = 1
)

var apiKey = ""

func init() {
	if apiKey == "" {
		log.Fatal("No api key for location")
	}
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file which is needed for API calls.")
	// }
	// apiKey = os.Getenv("PLATSUPPSLAG_KEY")
}

// SearchPlaces looks for places named after the arguments and returns Place
// structs
func SearchPlaces(names ...string) []Place {
	var places []Place

	for _, name := range names {
		places = append(places, searchPlace(name))
	}

	return places
}

func searchPlace(name string) Place {
	url := baseURL + "searchstring=" + name
	url += "&stationsonly=" + strbool(stationsOnlyVal)
	url += "&maxresults=" + strconv.Itoa(maxResultsVal)
	url += "&key=" + apiKey

	res := makeSearchRequest(url)

	return res.ResponseData[0]
}

func strbool(b bool) string {
	if b == true {
		return "True"
	}
	return "False"
}

func makeSearchRequest(url string) SearchResponse {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Could not complete http request\n%s\n", err)
	}
	defer res.Body.Close()

	var searchRes SearchResponse
	if err := json.NewDecoder(res.Body).Decode(&searchRes); err != nil {
		log.Println(err)
	}

	return searchRes
}
