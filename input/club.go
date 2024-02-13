// input/inputclub.go

package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/shinohara320/aptiv-klasemen/model"
)

var existingClubs []model.Club

func InputClub() ([]model.Club, error) {
	scanner := bufio.NewScanner(os.Stdin)

	clubExists := func(newClub model.Club) bool {
		for _, club := range existingClubs {
			if club.Name == newClub.Name && club.Kota == newClub.Kota {
				return true
			}
		}
		return false
	}

	fmt.Print("Masukkan jumlah klub yang akan dimasukkan: ")
	scanner.Scan()
	numClubsStr := scanner.Text()

	numClubs, err := strconv.Atoi(numClubsStr)
	if err != nil {
		return nil, err
	}

	clubs := make([]model.Club, 0)

	for i := 0; i < numClubs; {
		fmt.Printf("\nInput klub #%d:\n", i+1)

		fmt.Print("Nama klub: ")
		scanner.Scan()
		clubName := scanner.Text()

		fmt.Print("Kota klub: ")
		scanner.Scan()
		clubCity := scanner.Text()

		newClub := model.Club{Name: clubName, Kota: clubCity}

		if !clubExists(newClub) {
			clubs = append(clubs, newClub)
			existingClubs = append(existingClubs, newClub)
			i++
		} else {
			fmt.Printf("Klub %d: %s dari kota %s sudah ada\n", i+1, newClub.Name, newClub.Kota)
		}
	}

	fmt.Print("Apakah inputan sudah benar? [y/n]: ")
	scanner.Scan()
	response := strings.ToLower(scanner.Text())
	if response != "y" {
		fmt.Println("Inputan tidak benar. Ulangi input.")
		existingClubs = nil
		return InputClub()
	}

	return clubs, nil
}
