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

package cmd

import (
	"fmt"
	"gonow/location"
	"gonow/trip"
	"os"

	"github.com/spf13/cobra"
)

var posA, posB string
var outCompact bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gonow",
	Short: "Search your next trip with SL",

	Run: func(cmd *cobra.Command, args []string) {
		places := location.SearchPlaces(posA, posB)

		tripList := trip.SearchNow(places[0], places[1])
		if outCompact {
			trip.PrintCompact(tripList)
		} else {
			trip.Print(tripList)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&posA, "from", "a", "", "Traveling from")
	rootCmd.Flags().StringVarP(&posB, "to", "b", "", "Traveling to")
	rootCmd.MarkFlagRequired("from")
	rootCmd.MarkFlagRequired("to")

	rootCmd.Flags().BoolVarP(&outCompact, "compact", "c", false, "Compact output")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
