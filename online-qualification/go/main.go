package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

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
	_, err = fmt.Fscanf(f, "%d %d %d\n", &numBooks, &numLibs, &numDays)
	if err != nil {
		log.Fatal(err)
	}

	// Init array
	books := make([]common.Book, numBooks)
	for i := range books {
		books[i].ID = uint32(i)
		if _, err := fmt.Fscanf(f, "%d", &books[i].Score); err != nil {
			log.Fatal(err)
		}
	}

	/*if _, err = fmt.Fscanln(f); err != nil {
		log.Printf("hi")
		log.Fatal(err)
	}*/

	libraries := make([]common.Library, numLibs)
	for i := range libraries {
		lib := &libraries[i]
		lib.ID = uint32(i)
		var libNumBooks uint32
		_, err = fmt.Fscanf(f, "%d %d %d\n", &libNumBooks, &lib.SignUp, &lib.Ship)
		if err != nil {
			log.Printf("mango")
			log.Fatal(err)
		}
		lib.Books = make([]*common.Book, libNumBooks)
		for j := range lib.Books {
			var bookID uint32
			_, err = fmt.Fscanf(f, "%d", &bookID)
			if err != nil {
				log.Fatal(err)
			}
			lib.Books[j] = &books[j]
		}
	}

	// Debug message before doing any computation
	// fmt.Printf("%+v\n%+v\n", books, libraries)

	// solution
	var solLib []common.Library
	for remainingDays := numDays; remainingDays > 0 && len(libraries) > 0; {

		// sort books in each lib
		for idx := range libraries {
			lib := &libraries[idx]
			sort.Slice(lib.Books, func(i, j int) bool {
				return lib.Books[i].Score > lib.Books[j].Score
			})
			lib.Score = lib.CalcScore(remainingDays)
		}

		// DEBUG
		fmt.Printf("step: %+v\n", libraries)

		// TODO: Sorting by Lib score
		sort.Slice(libraries, func(i, j int) bool {
			return libraries[i].Score > libraries[j].Score
		})

		chosenLib := libraries[0]
		solLib = append(solLib, chosenLib)
		// books are chosen, set their scores to 0
		remainingDays -= chosenLib.SignUp
		canTake := chosenLib.Ship * remainingDays
		for i := 0; i < int(canTake) && i < len(chosenLib.Books); i++ {
			chosenLib.Books[i].Score = 0
		}
		libraries = libraries[1:]
	}

	// DEBUG
	fmt.Printf("%+v\n", solLib)

	/*
		NumLibUsed
		<- for each lib
			LibIndex NumBook
			Order of Books
		ex.
		2
		1 3
		5 2 3
		0 5
		0 1 2 3 4
	*/

	fmt.Printf("%d\n", len(solLib))
	for i := range solLib {
		lib := &solLib[i]
		fmt.Printf("%d %d\n", lib.ID, len(lib.Books))
		var ids []string
		for j := range lib.Books {
			ids = append(ids, strconv.FormatUint(uint64(lib.Books[j].ID), 10))
		}
		fmt.Println(strings.Join(ids, " "))
	}
}
