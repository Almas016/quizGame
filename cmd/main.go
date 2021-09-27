package main

import (
	"flag"
	"fmt"

	"github.com/Almas016/quizGame/pkg"
)

func main() {
	level := flag.Int("lvl", 1, "quiz level")
	fileName := flag.String("fn", "test1", "file name")
	timer := flag.Int("time", 30, "time allocated for the test")
	flag.Parse()
	switch *level {
	case 1:
		pkg.LevelOne(*fileName)
	case 2:
		pkg.LevelTwo(*fileName, *timer)
	default:
		fmt.Println("This level does not exist, please try again by changing the value")
	}
}
