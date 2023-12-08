package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"
)

type Node struct {
	label string
	left  string
	right string
}

var world map[string]*Node = make(map[string]*Node)

func ApplyPattern(input string, start *Node, condition *regexp.Regexp) (int, *Node) {

	cursor := start
	step := 0
	for _, char := range input {
		if condition.MatchString(cursor.label) {
			return step, cursor
		}
		if char == 'L' {
			cursor = world[cursor.left]
		} else if char == 'R' {
			cursor = world[cursor.right]
		}
		step++
	}
	return step, cursor

}

func SolveWithPattern(input string, start *Node, condition *regexp.Regexp) int {
	step := 0
	for progress, cursor := ApplyPattern(input, start, condition); progress != 0; progress, cursor = ApplyPattern(input, cursor, condition) {
		step += progress
	}
	return step

}

func CreateNode(label string) {
	resultString := strings.ReplaceAll(label, "(", "")
	resultString = strings.ReplaceAll(resultString, ")", "")
	resultString = strings.ReplaceAll(resultString, " ", "")
	parts := strings.Split(resultString, "=")
	label = parts[0]
	conns := strings.Split(parts[1], ",")
	world[label] = &Node{label, conns[0], conns[1]}
}

func InitWorld(filename string) string {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	regexPattern := "^[RL]*$"
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		panic(err)
	}
	var pattern string = ""
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if regex.MatchString(line) {
			pattern += line
		} else {
			CreateNode(line)
		}
	}
	return pattern
}

func TestFirstProblem(t *testing.T) {
	pattern := InitWorld("input.txt")
	regexpPattern := `ZZZ$`
	re, err := regexp.Compile(regexpPattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	log.Println(SolveWithPattern(pattern, world["AAA"], re))
}

/* Gcd and Lcm functions are taken from https://github.com/svenwiltink/aoc/blob/main/common/numbers.go*/
func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Lcm(integers ...int) int {
	result := integers[0] * integers[1] / Gcd(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = Lcm(result, integers[i])
	}

	return result
}

func TestSecondProblem(t *testing.T) {
	pattern := InitWorld("input.txt")
	rePattern := `A$`
	regexpPattern := `Z$`
	re, err := regexp.Compile(regexpPattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	re2, err := regexp.Compile(rePattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	var results []int
	for k, _ := range world {
		if re2.MatchString(k) {
			results = append(results, SolveWithPattern(pattern, world[k], re))
		}
	}
	log.Println(Lcm(results...))
}
