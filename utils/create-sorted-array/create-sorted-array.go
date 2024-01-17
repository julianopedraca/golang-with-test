package createsortedarray

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CreateSortedArray() {
	data, err := os.ReadFile("./utils/arrays/random-array.txt")

	if err != nil {
		log.Fatal(err)
	}

	a := strings.Split(string(data), ",")

	intArray := []int{}

	for i := 0; i < len(a); i++ {
		num, _ := strconv.Atoi(a[i])
		intArray = append(intArray, num)
	}

	sort.Ints(intArray)

	f, err2 := os.Create("./utils/arrays/sorted-array.txt")

	if err2 != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(intArray); i++ {

		s := strconv.Itoa(intArray[i])

		if i != len(intArray)-1 {
			s = s + ","
		}

		_, err3 := f.Write([]byte(s))

		if err3 != nil {
			log.Fatal(err2)
		}
	}

	defer f.Close()

}
