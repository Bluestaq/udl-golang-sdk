package main

import (
	"context"
	"fmt"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/param"

	"github.com/Bluestaq/udl-golang-sdk"
)

func main() {
	ctx := context.Background()
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL("https://dev.unifieddatalibrary.com"),
		option.WithDebugLog(nil),
	)
	iter := client.SecureMessaging.GetMessagesAutoPaging(ctx, 0, unifieddatalibrary.SecureMessagingGetMessagesParams{
		Topic:      "eop",
		MaxResults: param.NewOpt(int64(100)),
	})

	for iter.Next() {
		elsetAbridged := iter.Current()
		fmt.Printf("%+v\n", elsetAbridged)
	}
	if err := iter.Err(); err != nil {
		panic(err.Error())
	}
}
