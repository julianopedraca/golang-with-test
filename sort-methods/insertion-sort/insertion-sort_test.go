package insertionsort_test

import (
	"log"
	"os"
	"reflect"
	"testing"

	insertionsort "github.com/julianopedraca/golang-with-test/sort-methods/insertion-sort"
	misc "github.com/julianopedraca/golang-with-test/utils/misc"
)

const (
	infoColor    = "\033[1;34m%s\033[0m"
	noticeColor  = "\033[1;36m%s\033[0m"
	warningColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
	debugColor   = "\033[0;36m%s\033[0m"
	successColor = "\033[1;32m%s\033[0m"
)

func TestInsertionSort(t *testing.T) {
	t.Run("Test insertion sort", func(t *testing.T) {
		pathFile := "../../utils/arrays/"
		fileNameUnsorted := "random-array.txt"
		fileNameSorted := "sorted-array.txt"

		dataUnsorted, err := os.ReadFile(pathFile + fileNameUnsorted)
		if err != nil {
			log.Fatal(err)
		}

		dataSorted, err := os.ReadFile(pathFile + fileNameSorted)
		if err != nil {
			log.Fatal(err)
		}

		unsortedArray := misc.UnwrapArrayFile(dataUnsorted)
		sortedArray := insertionsort.InsertionSort(unsortedArray)

		want := misc.UnwrapArrayFile(dataSorted)

		if !reflect.DeepEqual(sortedArray, want) {
			t.Errorf(errorColor, "Sort failed")
		} else {
			t.Logf(successColor, "File sorted successfully\n")
		}
	})
}
