package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type config struct {
	numberOfDice int
	diceSize     int
	removedDice  int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	macros := getMacros()

	reader := bufio.NewReader(os.Stdin)

	expr := regexp.MustCompile(`(\d*)d(\d+)(?:r(\d+))?`)



	for {
		print("~> ")
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n\r")

		if text == "exit" {
			break
		}

		coolNewText := macros[text]

		if coolNewText != "" {
			text = coolNewText
		}

		results := expr.FindStringSubmatch(text)

		if len(results) == 0 {
			println("Not a diceroll")
			continue
		}



		c := parseData(results)

		rolls := make([]int, 0)

		for i := 0; i < c.numberOfDice; i++ {
			rolls = append(rolls, rand.Intn(c.diceSize)+1)
		}

		sort.Ints(rolls)

		rolls = rolls[c.removedDice:]

		total := 0

		for _, item := range rolls {
			total += item
		}

		log.Println(rolls)
		log.Println("Total:", total)
	}

}

func getMacros() map[string]string {
	file, err := os.Open("macros.txt")
	if err != nil {
		println(err.Error())
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		println(err.Error())
	}

	fileData := string(data)
	parts := strings.Split(fileData, "\n")

	m := make(map[string]string)

	for _, part := range parts {
		part = strings.TrimRight(part, "\n\r")
		if part != "" {
			macroParts := strings.Split(part,":")
			m[macroParts[0]] = macroParts[1]
		}
	}

	return m
}

func parseData(results []string) config {
	c := config{}

	numberOfDice, err := strconv.Atoi(results[1])

	if err != nil {
		numberOfDice = 1
	}

	c.numberOfDice = numberOfDice

	diceSize, _ := strconv.Atoi(results[2])

	c.diceSize = diceSize

	removedDice, err := strconv.Atoi(results[3])

	if err != nil {
		removedDice = 0
	}

	c.removedDice = removedDice

	return c
}
