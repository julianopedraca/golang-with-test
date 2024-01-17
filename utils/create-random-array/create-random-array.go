package createrandomarray

import (
	"log"
	"math/rand"
	"os"

	"strconv"
)

func CreateRandomNumbers() {
	f, err := os.Create("./utils/arrays/random-array.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	v := 50
	for i := 0; i < v; i++ {
		r := rand.Intn(10000)

		s := strconv.Itoa(r)

		if i != v-1 {
			s = s + ","
		}

		_, err2 := f.Write([]byte(s))

		if err != nil {
			log.Fatal(err2)
		}
	}

}
