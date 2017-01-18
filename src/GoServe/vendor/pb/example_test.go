package pb_test

import (
	"time"
	
	//"gopkg.in/cheggaaa/pb.v1" // https://github.com/cheggaaa/pb/tree/v1.0.7 -> pb
	"pb"
)

func Example() {
	count := 5000
	bar := pb.New(count)

	// show percents (by default already true)
	bar.ShowPercent = true

	// show bar (by default already true)
	bar.ShowBar = true

	bar.ShowCounters = true

	bar.ShowTimeLeft = true

	// and start
	bar.Start()
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint("The End!")
}
