package klasemen

import (
	"fmt"
	"sort"

	"github.com/shinohara320/aptiv-klasemen/model"
)

func ShowKlasemen(clubs []model.Club) {
	sortKlasemen(clubs)
	fmt.Println("Klasemen:")
	fmt.Println("================================================================================")
	fmt.Println("No  | Club      |  Ma   |   Me   |   S    |    K    |   GM   |   GK    | Score  ")
	fmt.Println("================================================================================")
	for i, club := range clubs {
		fmt.Printf("%d.    %s      %d         %d        %d        %d        %d         %d         %d/\n", i+1, club.Name, club.Main, club.Menang, club.Seri, club.Kalah, club.GM, club.GK, club.Score)
	}
}

func sortKlasemen(clubs []model.Club) {
	// Menggunakan fungsi sort.Slice untuk mengurutkan clubs berdasarkan poin (descending)
	sort.Slice(clubs, func(i, j int) bool {
		return clubs[i].Score > clubs[j].Score
	})
}
