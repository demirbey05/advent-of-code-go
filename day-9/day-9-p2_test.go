package day_9

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
)

func CraftAuxArray2(initialArray *[]int, n int, r int) int {

	for i := n - 1; i > r; i-- {
		(*initialArray)[i] = (*initialArray)[i] - (*initialArray)[i-1]
	}
	return r + 1
}
func ZeroCheck2(initialArray *[]int, n int, r int) bool {
	ok := true
	for i := n - 1; i >= r; i-- {
		if (*initialArray)[i] != 0 {
			ok = false
		}
	}
	return ok
}

func Predict2(initialArray *[]int, n int) int {
	var size int
	for size = CraftAuxArray2(initialArray, n, 0); size < n && !ZeroCheck2(initialArray, n, size); size = CraftAuxArray2(initialArray, n, size) {
	}

	prediction := 0
	for i := size - 1; i >= 0; i-- {
		prediction = (*initialArray)[i] - prediction
	}
	return prediction
}
func TestPredict2(t *testing.T) {
	var initialArray []int = []int{3, 11, 19, 27, 35, 43, 51, 59, 67, 75, 83, 91, 99, 107, 115, 123, 131, 139, 147, 155, 163}
	prediction := Predict2(&initialArray, len(initialArray))
	if prediction != -5 {
		t.Errorf("Expected 5, got %d", prediction)
	}
}

func TestPart2(t *testing.T) {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup
	results := make(chan int)
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, " ")
		arr, err := ConvertToInt(splitted)
		if err != nil {
			panic(err)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- Predict2(&arr, len(arr))

		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	var sum int
	for res := range results {
		sum += res

	}
	log.Println(sum)
}
