package main

import (
	"context"
	"fmt"

	unifieddatalibrary "github.com/Bluestaq/udl-golang-sdk"
	lib "github.com/Bluestaq/udl-golang-sdk/lib"
)

func main() {
	// Set you environment variables for with your personal preference
	// UDL_AUTH_PASSWORD,UDL_AUTH_USERNAME, and UDL_BASE_URL will be
	// obtained automatically
	client := unifieddatalibrary.NewClient()

	// Build query and get request options directly
	queryOpts := lib.NewQueryBuilder(unifieddatalibrary.ElsetAbridged{}).Add("SatNo", lib.EqualTo, "25544").ToRequestOptions()

	page, err := client.Elsets.Current.List(context.TODO(), unifieddatalibrary.ElsetCurrentListParams{}, queryOpts...)

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Response: %+v\n", page)

}
