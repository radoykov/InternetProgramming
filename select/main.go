package main

import "fmt"

func main() {
	var ch1, ch2, ch3 chan int 
	
	sum := 0

	for {
		if ch1 == nil && ch2 == nil && ch3 == nil {
			break
		}

		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil // Disable this case branch forever
				continue
			}
			sum += v

		case v, ok := <-ch2:
			if !ok {
				ch2 = nil // Disable this case branch forever
				continue
			}
			sum += v

		case v, ok := <-ch3:
			if !ok {
				ch3 = nil // Disable this case branch forever
				continue
			}
			sum += v
		}
	}

	fmt.Printf("Total Sum: %d\n", sum)
}
