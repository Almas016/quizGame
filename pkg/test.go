package pkg

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type myCSV [][]string

func LevelOne(fileName string) {
	csvLines := ReadCSV(fileName)

	fmt.Printf("%d correct answers out of %d\n", correctAnswer(&csvLines), len(csvLines))
}

func LevelTwo(fileName string, timer int) {
	csvLines := ReadCSV(fileName)

	if correct, err := correctAnswer2(&csvLines, timer); err != nil {
		fmt.Println(error.Error(err))
	} else {
		fmt.Printf("%d correct answers out of %d\n", correct, len(csvLines))
	}

}

func correctAnswer(s *myCSV) int {
	correct := 0
	var answer int
	for _, v := range *s {
		fmt.Printf("%v, your answer?\n", v[0])
		fmt.Scanln(&answer)
		trueAnswer, err := strconv.Atoi(v[1])
		if err != nil {
			fmt.Println(error.Error(err))
		}
		if answer == trueAnswer {
			correct++
		}
	}
	return correct
}

func correctAnswer2(s *myCSV, n int) (int, error) {
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
			return correct, errors.New("Timed out")
		case answer := <-answerCh:
			if answer == v[1] {
				correct++
			}
		}
	}

	return correct, nil
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

	return ShuffleFile(csvLines)
}
