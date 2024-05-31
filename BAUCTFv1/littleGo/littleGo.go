package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var userInput int
	fmt.Print("Enter the password: ")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		os.Exit(1)
	}

	if userInput == 713371339 {
		resp, err := http.Get("http://pastebin.com/raw/MWFU40xz")
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Errorr")
			os.Exit(1)
		}
		fmt.Print(string(body)[:10])
		fmt.Println("...")
	} else {
		fmt.Println(`⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠀⠀⠀⠀⠀⢀⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢠⡄⠀⠀⠀⠀⣇⠀⠀⠀⠀⠀⡸⢐⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢿⡙⠶⢄⣀⠀⢫⠒⢤⣀⠀⠀⢸⠰⠀⠀⣀⡤⣤⠤⠀⠀⠀
⠀⠀⠀⠀⠉⢒⠦⠄⠈⠉⠁⠠⠀⠉⠉⠙⠃⠚⠉⣠⠎⠀⠀⠀⠀⠀
⢀⣄⣠⡐⣈⡀⡄⠀⠀⢠⠀⢀⣴⣠⠀⠀⠀⠀⠀⠻⣀⠀⠀⠀⠀⠀
⠀⠀⢨⠟⢁⢔⡁⢀⠔⠀⠐⣡⣯⠃⢠⠀⡆⢤⠀⡀⢰⡯⡒⠠⠤⠀
⢀⣴⡵⣾⢗⣥⣶⣿⣷⣮⡼⢣⠃⢠⣧⣤⣯⣘⠀⢣⠀⣣⡏⠉⠚⠉
⠟⠁⣸⣣⡃⢿⣿⣿⣿⣿⠷⠾⢶⣿⣿⣿⣿⣿⡆⣿⡀⢿⣸⡀⠀⠀
⠀⢰⠋⠀⠀⠀⠉⠙⠉⠁⢀⣀⡀⠙⠛⠛⠛⠛⠑⡿⣯⣽⠋⣳⡆⠀
⠀⠈⠳⢦⣄⡀⠀⠀⠘⣄⣀⣀⠼⠃⠀⠀⢀⠀⠠⠴⠿⠛⠋⠁⠀⠀
⠀⠀⠀⠀⠀⠉⠉⠓⠒⠒⠤⠤⠤⠤⠔⠚⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀`)
	}
}
