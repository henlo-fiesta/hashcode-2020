package main

import (
	"fmt"
	"log"
	"os"

	"github.com/natebwangsut/hashcode-2020-henlo-fiesta/online-qualification/go/common"
)

func main() {
	f, err := os.Open("../samples/a_example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Fuck im too stoopid to open a fucking file
	var numBooks, numLibs, numDays uint32
	fmt.Fscanf(f, "%d %d %d\n", &numBooks, &numLibs, &numDays)

	books := make([]common.Book, numBooks)
	for i := range books {
		books[i].ID = uint32(i)
		fmt.Fscanf(f, "%d", &books[i].Score)
	}
	fmt.Fscanln(f)

	libraries := make([]common.Library, numLibs)
	for i := range libraries {
		lib := &libraries[i]
		var libNumBooks uint32
		fmt.Fscanf(f, "%d %d %d\n", &libNumBooks, &lib.SignUp, &lib.Ship)
		lib.Books = make([]*common.Book, libNumBooks)
		for j := range lib.Books {
			var bookId uint32
			fmt.Fscanf(f, "%d", &bookId)
			lib.Books[j] = &books[j]
		}
		fmt.Fscanln(f)
	}

	fmt.Printf("%+v\n", libraries)
}
