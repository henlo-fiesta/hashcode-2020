package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/natebwangsut/hashcode-2020-henlo-fiesta/online-qualification/go/common"
)

const (
	// DEBUG FLAG
	DEBUG = false
)

func main() {

	// Static Files
	files := []string{
		"../in/a_example.txt",
		"../in/b_read_on.txt",
		"../in/c_incunabula.txt",
		"../in/d_tough_choices.txt",
		"../in/e_so_many_books.txt",
		"../in/f_libraries_of_the_world.txt",
	}
	out := []string{
		"../out/a_example.txt",
		"../out/b_read_on.txt",
		"../out/c_incunabula.txt",
		"../out/d_tough_choices.txt",
		"../out/e_so_many_books.txt",
		"../out/f_libraries_of_the_world.txt",
	}

	score := uint32(0)
	for i := range files {
		f, err := os.Create(out[i])
		if err != nil {
			log.Fatal(err)
		}

		// Print to STDOUT if DEBUG
		if DEBUG {
			score += doIt(files[i], os.Stdout)
			os.Exit(0)
		} else {
			score += doIt(files[i], f)
		}

		defer f.Close()
	}
	fmt.Printf("total score: %d\n", score)
}

// doIt calls the main prog with IO redirections
func doIt(filename string, out io.Writer) uint32 {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var numBooks, numLibs, numDays uint32
	_, err = fmt.Fscanf(f, "%d %d %d\n", &numBooks, &numLibs, &numDays)
	if err != nil {
		log.Fatal(err)
	}

	// Map books to struct
	books := make([]common.Book, numBooks)
	meanBookScore := float32(0)
	for i := range books {
		books[i].ID = uint32(i)
		if _, err := fmt.Fscanf(f, "%d", &books[i].Score); err != nil {
			log.Fatal(err)
		}
		meanBookScore += float32(books[i].Score)
	}
	meanBookScore /= float32(numBooks)
	booksBackup := make([]common.Book, numBooks)
	copy(booksBackup, books)

	if DEBUG {
		fmt.Println(books)
	}

	/*
		// Was trying to detect last '\n' from books score line
		if _, err = fmt.Fscanln(f); err != nil {
			log.Printf("hi")
			log.Fatal(err)
		}
	*/

	// Map libraries to struct
	libraries := make([]common.Library, numLibs)
	meanSignUp := float32(64)
	for i := range libraries {
		lib := &libraries[i]
		lib.ID = uint32(i)
		var libNumBooks uint32
		_, err = fmt.Fscanf(f, "%d %d %d\n", &libNumBooks, &lib.SignUp, &lib.Ship)
		if err != nil {
			log.Fatal(err)
		}
		meanSignUp += float32(lib.SignUp)
		lib.Books = make([]*common.Book, libNumBooks)
		for j := range lib.Books {
			var bookID uint32
			_, err = fmt.Fscanf(f, "%d", &bookID)
			if err != nil {
				log.Fatal(err)
			}
			lib.Books[j] = &books[bookID]
		}
	}

	// Compute meanSignUp to be used as a scoring mechanism
	meanSignUp /= float32(numLibs)

	// Debug message before doing any computation
	if DEBUG {
		fmt.Printf("%+v\n%+v\n", books, libraries)
	}

	// Solution
	var solLib []common.Library
	for remainingDays := numDays; remainingDays > 0 && len(libraries) > 0; {
		signUpCost := meanSignUp * meanBookScore
		// Sort books in each lib
		for idx := range libraries {
			lib := &libraries[idx]
			sort.Slice(lib.Books, func(i, j int) bool {
				return lib.Books[i].Score > lib.Books[j].Score
			})
			lib.Score = lib.CalcScore(remainingDays, signUpCost)
		}

		// DEBUG
		if DEBUG {
			fmt.Printf("step: %+v\n", libraries)
		}

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
	if DEBUG {
		fmt.Printf("%+v\n", solLib)
	}

	/*
		Score Formatting Logics
		NumLibUsed
		<- for each lib
			LibIndex NumBook
			Order of Books
		ex. in/a_example.txt
		2
		1 3
		5 2 3
		0 5
		0 1 2 3 4
	*/
	fmt.Fprintf(out, "%d\n", len(solLib))
	for i := range solLib {
		lib := &solLib[i]
		fmt.Fprintf(out, "%d %d\n", lib.ID, len(lib.Books))
		var ids []string
		for j := range lib.Books {
			ids = append(ids, strconv.FormatUint(uint64(lib.Books[j].ID), 10))
		}
		fmt.Fprintln(out, strings.Join(ids, " "))
	}

	// Debug print score: does not get included in to the main IO
	// Only prints to stdout
	rd := int(numDays)
	score := uint32(0)
	for i := range solLib {
		lib := &solLib[i]
		canTake := int(lib.Ship * uint32(rd-int(lib.SignUp)))
		if canTake > len(lib.Books) {
			canTake = len(lib.Books)
		}
		if canTake < 0 {
			canTake = 0
		}
		tBooks := lib.Books[:canTake]
		for _, tb := range tBooks {
			score += booksBackup[tb.ID].Score
			booksBackup[tb.ID].Score = 0
		}
		rd -= int(lib.SignUp)
		if rd <= 0 {
			break
		}
	}
	fmt.Printf("score: %d\n", score)
	return score
}
