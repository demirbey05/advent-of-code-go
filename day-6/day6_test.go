package day_6

import (
	"log"
	"math"
	"testing"
)

func DistanceEquation(pressTime, totalTime int) int {
	return (pressTime * totalTime) - (pressTime * pressTime)
}

func FindDiscriminant(a, b, c int) int {
	return (b * b) - (4 * a * c)
}
func FindRootFromDiscriminant(a, b, c int) []float64 {
	discriminant := FindDiscriminant(a, b, c)
	var root []float64
	root = append(root, (float64(-b)+math.Sqrt(float64(discriminant)))/float64(2*a))
	root = append(root, (float64(-b)-math.Sqrt(float64(discriminant)))/float64(2*a))
	log.Println("Root: for a:", root, a)
	return root
}

func FindNumberOfWayToWin(totalTime, record int) int {
	roots := FindRootFromDiscriminant(1, -totalTime, record)
	count := 0
	upper := int(math.Ceil(roots[0]))
	lower := int(math.Floor(roots[1]))
	for i := lower; i <= upper; i++ {
		if DistanceEquation(i, totalTime) > record {
			count++
		}
	}
	return count

}

func TestNumberOfWay(t *testing.T) {
	result1 := FindNumberOfWayToWin(7, 9)
	result2 := FindNumberOfWayToWin(15, 40)
	result3 := FindNumberOfWayToWin(30, 200)

	if result1 != 4 {
		t.Errorf("Expected 4 but got %d", result1)
	}
	if result2 != 8 {
		t.Errorf("Expected 8 but got %d", result2)
	}
	if result3 != 9 {
		t.Errorf("Expected 9 but got %d", result3)
	}

}

func TestSolution(t *testing.T) {
	result1 := FindNumberOfWayToWin(60, 601)
	result2 := FindNumberOfWayToWin(80, 1163)
	result3 := FindNumberOfWayToWin(86, 1559)
	result4 := FindNumberOfWayToWin(76, 1300)

	log.Println("The result is ", result1*result2*result3*result4)
}

func TestKerning(t *testing.T) {

	result1 := FindNumberOfWayToWin(60808676, 601116315591300)
	log.Println("The result is ", result1)

}
