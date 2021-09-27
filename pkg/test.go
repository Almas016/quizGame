package pkg

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type myCSV [][]string

func LevelOne(fileName string) {
	csvLines := ReadCSV(fileName)

	fmt.Printf("%d correct answers out of %d\n", correctAnswer(&csvLines), len(csvLines))
}

func LevelTwo(fileName string, timer int) {
	csvLines := ReadCSV(fileName)

	correct := correctAnswer2(&csvLines, timer)
	fmt.Printf("%d correct answers out of %d\n", correct, len(csvLines))
}

func correctAnswer(s *myCSV) int {
	correct := 0
	var answer string
	for _, v := range *s {
		fmt.Printf("%v, your answer?\n", v[0])
		fmt.Scanln(&answer)
		if answer == v[1] {
			correct++
		}
	}
	return correct
}

func correctAnswer2(s *myCSV, n int) int {
	correct := 0
	timer := time.NewTimer(time.Duration(n) * time.Second)
	for _, v := range *s {
		fmt.Printf("%v, your answer?\n", v[0])
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Timed out")
			return correct
		case answer := <-answerCh:
			if answer == v[1] {
				correct++
			}
		}
	}

	return correct
}

func ShuffleFile(s myCSV) myCSV {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}

func ReadCSV(fileName string) myCSV {
	csvFile, err := os.Open(fileName + ".csv")
	if err != nil {
		fmt.Println(error.Error(err))
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(error.Error(err))
	}

	for _, v := range csvLines {
		if v[0] == "" {
			fmt.Println("Your file is missing a question(s)")
			os.Exit(1)
		}
	}

	return ShuffleFile(csvLines)
}
