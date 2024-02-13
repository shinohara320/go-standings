package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/shinohara320/aptiv-klasemen/model"
)

type MatchResult struct {
	Club         model.Club
	GolDicetak   int
	GolKebobolan int
}

func InputMatch(clubs []model.Club) {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Masukkan hasil pertandingan:")
		fmt.Print("Klub Home: ")
		scanner.Scan()
		club1Name := scanner.Text()

		club1, err := findClubByName(club1Name, clubs)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Print("Gol Klub Home: ")
		scanner.Scan()
		gol1Str := scanner.Text()
		gol1, err := strconv.Atoi(gol1Str)
		if err != nil || gol1 < 0 {
			fmt.Println("Error: Masukkan angka yang valid untuk skor gol klub 1.")
			continue
		}

		fmt.Print("Klub Away: ")
		scanner.Scan()
		club2Name := scanner.Text()

		club2, err := findClubByName(club2Name, clubs)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Print("Gol Klub Away: ")
		scanner.Scan()
		gol2Str := scanner.Text()
		gol2, err := strconv.Atoi(gol2Str)
		if err != nil || gol2 < 0 {
			fmt.Println("Error: Masukkan angka yang valid untuk skor gol klub 2.")
			continue
		}

		fmt.Printf("Hasil pertandingan: %s %d-%d %s\n", club1.Name, gol1, gol2, club2.Name)

		fmt.Print("Apakah data di atas sudah benar? [y/n]: ")
		scanner.Scan()
		confirmation := scanner.Text()
		if confirmation != "y" && confirmation != "Y" {
			fmt.Println("Inputan dihapus, silakan masukkan kembali.")
			continue
		}

		updateClubScores([]MatchResult{{Club: club1, GolDicetak: gol1, GolKebobolan: gol2},
			{Club: club2, GolDicetak: gol2, GolKebobolan: gol1}}, clubs)
		break
	}
}

func findClubByName(name string, clubs []model.Club) (model.Club, error) {
	var updatedClubs []model.Club

	for _, club := range clubs {
		if club.Name == name {
			updatedClubs = append(updatedClubs, club)
		} else {
			fmt.Println("klub tidak ditemukan")
		}
	}

	return updatedClubs[0], nil
}

func updateClubScores(matchResults []MatchResult, clubs []model.Club) {
	for _, matchResult := range matchResults {
		for i := range clubs {
			if clubs[i].Name == matchResult.Club.Name {
				updateMatchResult(&clubs[i], matchResult.GolDicetak, matchResult.GolKebobolan)
			}
		}
	}
}

func updateMatchResult(club *model.Club, golDicetak, golKebobolan int) {
	club.Main++
	club.GM += golDicetak
	club.GK += golKebobolan

	if golDicetak > golKebobolan {
		club.Menang++
		club.Score += 3
	} else if golDicetak == golKebobolan {
		club.Seri++
		club.Score++
	} else {
		club.Kalah++
	}
}
