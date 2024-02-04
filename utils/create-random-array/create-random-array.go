package createrandomarray

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"strconv"
)

func CreateRandomNumbers() {
	fmt.Printf("Creating file...\n")
	fileName := "random-array.txt"
	pathName := "./utils/arrays/"
	f, err := os.Create(pathName + fileName)
	fmt.Printf("%s created at %s \n", fileName, pathName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fmt.Printf("start populating file \n")
	//set size of array
	var sizeOfArray int
	var randInt int
	fmt.Printf("Write the size of the array: \n")
	fmt.Scan(&sizeOfArray)
	fmt.Printf("Write the larger number of the array: \n")
	fmt.Scan(&randInt)
	for i := 0; i < sizeOfArray; i++ {
		if i%2 == 0 {
			fmt.Printf("populating..\n")
		} else {
			fmt.Printf("populating...\n")
		}

		randomNumber := rand.Intn(randInt)

		s := strconv.Itoa(randomNumber)

		if i != sizeOfArray-1 {
			s = s + ","
		}

		_, err2 := f.Write([]byte(s))

		if err != nil {
			log.Fatal(err2)
		}
	}
	fmt.Printf("File poluated with %d \n", sizeOfArray)
}
