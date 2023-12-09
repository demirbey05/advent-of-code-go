package day_9

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func CraftAuxArray(initialArray *[]int, n int) int {

	for i := 0; i < n-1; i++ {
		(*initialArray)[i] = (*initialArray)[i+1] - (*initialArray)[i]
	}
	return n - 1
}

func ZeroCheck(initialArray *[]int, n int) bool {
	ok := true
	for i := 0; i < n; i++ {
		if (*initialArray)[i] != 0 {
			ok = false
		}
	}
	return ok
}

func Predict(initialArray *[]int, n int) int {
	var size int
	for size = CraftAuxArray(initialArray, n); size > 0 && !ZeroCheck(initialArray, size); size = CraftAuxArray(initialArray, size) {

	}
	prediction := 0

	for i := n - 1; i >= size; i-- {
		prediction += (*initialArray)[i]
	}
	return prediction
}

func TestPredict(t *testing.T) {
	var initialArray []int = []int{10, 13, 16, 21, 30, 45}
	prediction := Predict(&initialArray, 6)
	if prediction != 68 {
		t.Errorf("Expected 68, got %d", prediction)
	}
}
func ConvertToInt(splitted []string) ([]int, error) {
	var arrayPointer []int
	for _, value := range splitted {
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		arrayPointer = append(arrayPointer, num)
	}
	return arrayPointer, nil
}

func TestPart1(t *testing.T) {
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
			results <- Predict(&arr, len(arr))

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
