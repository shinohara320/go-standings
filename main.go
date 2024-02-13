package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/shinohara320/aptiv-klasemen/input"
	"github.com/shinohara320/aptiv-klasemen/klasemen"
	"github.com/shinohara320/aptiv-klasemen/model"
)

var Clubs []model.Club

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("======================")
	fmt.Println("MAIN MENU")
	fmt.Println("======================")
	fmt.Println("1. Klasemen")
	fmt.Println("2. Input Match")
	fmt.Println("3. Input klub")
	fmt.Println("4. Exit")
	fmt.Println("Pilih Menu [1/2/3/4] : ")
	fmt.Println("======================")
	scanner.Scan()
	inputTeks := scanner.Text()

	switch inputTeks {
	case "1":
		klasemen.ShowKlasemen(Clubs)
		main()
	case "2":
		fmt.Println("1. Single Input")
		fmt.Println("2. Multiple Input")
		fmt.Println("Pilih Menu [1/2]: ")
		scanner.Scan()
		matchInputType := scanner.Text()

		switch matchInputType {
		case "1":
			input.InputMatch(Clubs)
		case "2":
			input.InputMultipleMatch(Clubs)
		default:
			fmt.Println("Pilihan tidak valid")
		}
		main()
	case "3":
		data, err := input.InputClub()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Data klub yang dimasukkan (%d klub):\n", len(data))
		for i, club := range data {
			Clubs = data
			fmt.Printf("Klub %d: %s dari kota %s\n", i+1, club.Name, club.Kota)
		}
		main()
	case "4":
		os.Exit(0)
	default:
		fmt.Println("Nilai tidak dikenali")
	}
}
