package include

import (
	"time"

	"fyne.io/fyne/v2/widget"
)

func UpdateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func InitRecurringFunctionUpdateClock(clock *widget.Label) {
	// w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
}
