package main

import (
	"flag"

	"github.com/Almas016/quizGame/pkg"
)

func main() {
	level := flag.Int("lvl", 1, "quiz level")
	fileName := flag.String("fn", "test1", "file name")
	timer := flag.Int("time", 30, "time allocated for the test")
	flag.Parse()
	if *level == 2 {
		pkg.LevelTwo(*fileName, *timer)
	} else {
		pkg.LevelOne(*fileName)
	}
}
