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

package trip

// CompleteTrip from point A to B containing at leas one sub trip
// type CompleteTrip struct {
// 	StartDate string
// 	Steps     []SubTrip
// }

// SubTrip is one part of a trip
type SubTrip struct {
	Date      string
	DepTime   string
	ArrTime   string
	From      string
	To        string
	Direction string
	Type      string
}

// SearchTripResponse is the response to a trip request
type SearchTripResponse struct {
	Trip []struct {
		ServiceDays []struct {
			PlanningPeriodBegin string `json:"planningPeriodBegin"`
			PlanningPeriodEnd   string `json:"planningPeriodEnd"`
			SDaysR              string `json:"sDaysR"`
			SDaysI              string `json:"sDaysI"`
			SDaysB              string `json:"sDaysB"`
		} `json:"ServiceDays"`
		LegList struct {
			Leg []struct {
				Origin struct {
					Name          string  `json:"name"`
					Type          string  `json:"type"`
					ID            string  `json:"id"`
					ExtID         string  `json:"extId"`
					Lon           float64 `json:"lon"`
					Lat           float64 `json:"lat"`
					PrognosisType string  `json:"prognosisType"`
					Time          string  `json:"time"`
					Date          string  `json:"date"`
					Track         string  `json:"track"`
					RtTime        string  `json:"rtTime"`
					RtDate        string  `json:"rtDate"`
					HasMainMast   bool    `json:"hasMainMast"`
					MainMastID    string  `json:"mainMastId"`
					MainMastExtID string  `json:"mainMastExtId"`
				} `json:"Origin"`
				Destination struct {
					Name          string  `json:"name"`
					Type          string  `json:"type"`
					ID            string  `json:"id"`
					ExtID         string  `json:"extId"`
					Lon           float64 `json:"lon"`
					Lat           float64 `json:"lat"`
					PrognosisType string  `json:"prognosisType"`
					Time          string  `json:"time"`
					Date          string  `json:"date"`
					RtTime        string  `json:"rtTime"`
					RtDate        string  `json:"rtDate"`
					HasMainMast   bool    `json:"hasMainMast"`
					MainMastID    string  `json:"mainMastId"`
					MainMastExtID string  `json:"mainMastExtId"`
				} `json:"Destination"`
				JourneyDetailRef struct {
					Ref string `json:"ref"`
				} `json:"JourneyDetailRef"`
				JourneyStatus string `json:"JourneyStatus"`
				Product       struct {
					Name         string `json:"name"`
					Num          string `json:"num"`
					Line         string `json:"line"`
					CatOut       string `json:"catOut"`
					CatIn        string `json:"catIn"`
					CatCode      string `json:"catCode"`
					CatOutS      string `json:"catOutS"`
					CatOutL      string `json:"catOutL"`
					OperatorCode string `json:"operatorCode"`
					Operator     string `json:"operator"`
					Admin        string `json:"admin"`
				} `json:"Product"`
				Idx       string `json:"idx"`
				Name      string `json:"name"`
				Number    string `json:"number"`
				Category  string `json:"category"`
				Type      string `json:"type"`
				Reachable bool   `json:"reachable"`
				Direction string `json:"direction"`
			} `json:"Leg"`
		} `json:"LegList"`
		TariffResult struct {
			FareSetItem []struct {
				FareItem []struct {
					Name  string `json:"name"`
					Desc  string `json:"desc"`
					Price int    `json:"price"`
					Cur   string `json:"cur"`
				} `json:"fareItem"`
				Name string `json:"name"`
				Desc string `json:"desc"`
			} `json:"fareSetItem"`
		} `json:"TariffResult"`
		Idx      int    `json:"idx"`
		TripID   string `json:"tripId"`
		CtxRecon string `json:"ctxRecon"`
		Duration string `json:"duration"`
		Checksum string `json:"checksum"`
	} `json:"Trip"`
	ServerVersion  string `json:"serverVersion"`
	DialectVersion string `json:"dialectVersion"`
	RequestID      string `json:"requestId"`
	ScrB           string `json:"scrB"`
	ScrF           string `json:"scrF"`
}
