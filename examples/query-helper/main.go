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
	params := unifieddatalibrary.ElsetCurrentListParams{}
	queryParams := lib.NewQueryBuilder(unifieddatalibrary.ElsetAbridged{}).Add("SatNo", lib.EqualTo, "25544").ToParams()
	extraFields := make(map[string]any)
	for k, v := range queryParams {
		extraFields[k] = v
	}
	params.SetExtraFields(extraFields)
	page, err := client.Elsets.Current.List(context.TODO(), params)

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf(("%+v\n"), page)

}
