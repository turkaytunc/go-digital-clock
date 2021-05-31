package clock

import (
	"fmt"
	"time"

	"github.com/turkaytunc/digital-clock/digits"
)

func New() {

	hour, min, sec := time.Now().Clock()
	ds := digits.New()
	sep := digits.Placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	clockArr := []digits.Placeholder{
		ds[hour/10],
		ds[hour%10],
		sep,
		ds[min/10],
		ds[min%10],
		sep,
		ds[sec/10],
		ds[sec%10],
	}

	for line := range clockArr[0] {
		for c := range clockArr {
			fmt.Print(" ", clockArr[c][line], "  ")
		}
		fmt.Println()
	}

}
