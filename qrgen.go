package main

import (
	"fmt"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	for {
		fmt.Println("qrcode")
		var link string
		fmt.Println("Enter your link for qrcode or write exit for exiting the app: ")
		fmt.Scan(&link)
		if link == "exit" {
			fmt.Println("Application exiting...")
			break
		}
		png, err := qrcode.Encode(link, qrcode.Medium, 256)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile("qr.png", png, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println("qrcode generated successfully")
	}
}
