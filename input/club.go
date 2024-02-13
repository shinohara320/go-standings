package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/shinohara320/aptiv-klasemen/model"
)

func InputClub() ([]model.Club, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan jumlah klub yang akan dimasukkan: ")
	scanner.Scan()
	numClubsStr := scanner.Text()

	numClubs, err := strconv.Atoi(numClubsStr)
	if err != nil {
		return nil, err
	}

	clubs := make([]model.Club, numClubs)

	for i := 0; i < numClubs; i++ {
		fmt.Printf("\nInput klub #%d:\n", i+1)

		fmt.Print("Nama klub: ")
		scanner.Scan()
		clubName := scanner.Text()

		fmt.Print("Kota klub: ")
		scanner.Scan()
		clubCity := scanner.Text()

		clubs[i] = model.Club{Name: clubName, Kota: clubCity}
	}

	return clubs, nil
}
