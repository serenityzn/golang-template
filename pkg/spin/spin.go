package spin

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ASCII frames for the animation
var emojiFrames = []string{"⡿", "⣟", "⣯", "⣷", "⣾", "⣽", "⣻", "⣿"}

func Spin(task func(), spinTime time.Duration) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	done := make(chan struct{})
	go func(spintTime time.Duration) {
		for {
			select {
			case <-done:
				return
			default:
				for i := 0; ; i++ {
					fmt.Printf("\r\033[32m%s\033[0m...", emojiFrames[i%len(emojiFrames)])
					time.Sleep(spinTime * time.Millisecond) // Control frame speed
				}
			}
		}
	}(spinTime)

	// Execute the provided task
	task()

	// Signal that the task is done
	close(done)

	// Clean up the spinner display
	fmt.Print("\rDone!        \n")
}
