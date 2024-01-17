package main

import (
	"fmt"

	"github.com/eiannone/keyboard"

	cr "github.com/julianopedraca/golang-with-test/utils/create-random-array"
	cs "github.com/julianopedraca/golang-with-test/utils/create-sorted-array"
)

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	methods := []string{
		"Insertion Sort",
		"Merge Sort",
	}

	cr.CreateRandomNumbers()
	cs.CreateSortedArray()

	index := 0
	exit := 0
	for exit == 0 {
		// cmd := exec.Command("clear")
		// cmd.Stdout = os.Stdout
		// cmd.Run()

		fmt.Println("Press ESC to quit")
		for i, v := range methods {
			if i == index {
				fmt.Printf("> %s \n", v)
			} else {
				fmt.Printf(" %s \n", v)
			}
		}

		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyEsc:
			exit = 1
		case keyboard.KeyArrowDown:
			index += 1
			if index > len(methods)-1 {
				index = 0
			}
		case keyboard.KeyArrowUp:
			index -= 1
			if index < 0 {
				index = len(methods) - 1
			}
		case keyboard.KeyEnter:
			fmt.Printf("%d \n", index)
		}

	}

}
