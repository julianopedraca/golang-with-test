package createsortedarray

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	misc "github.com/julianopedraca/golang-with-test/utils/misc"
)

func CreateSortedArray() {

	pathFile := ("./utils/arrays/")
	fileName := ("random-array.txt")
	data, err := os.ReadFile(pathFile + fileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("preparing file to sort")
	// split random sorted file, transform it into int, and add to array
	intArray := misc.UnwrapArrayFile(data)

	fmt.Printf("sorting file using 'sort.Ints'...\n")
	sort.Ints(intArray)
	fmt.Printf("File sorted")

	f, err2 := os.Create(pathFile + "sorted-array.txt")

	if err2 != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(intArray); i++ {

		if i%2 == 0 {
			fmt.Printf("creating sorted file ...\n")
		} else {
			fmt.Printf("creating sorted file.. \n")
		}

		s := strconv.Itoa(intArray[i])

		if i != len(intArray)-1 {
			s = s + ","
		}

		_, err3 := f.Write([]byte(s))

		if err3 != nil {
			log.Fatal(err2)
		}
	}

	fmt.Printf("Sorted file created!")

	defer f.Close()

}
