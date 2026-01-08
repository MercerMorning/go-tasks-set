package main

import (
	"fmt"
	"time"
)

func say(id int, phrase string) {
	time.Sleep(20 * time.Millisecond)
	fmt.Printf("Worker %d says: %s\n", id, phrase)
}
func makePool(poolSize int, handler func(int, string)) (func(string), func()) {
}

func main() {
	phrases := []string{}
	for i := range 100 {
		phrases = append(phrases, fmt.Sprintf("phrase %d", i))
	}

	handle, wait := makePool(5, say)
	for _, phrase := range phrases {
		handle(phrase)
	}
	wait()
	fmt.Println("Done!")
}
