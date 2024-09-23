package main

import "fmt"

func main() {
	msgCh := make(chan string)

	// list of article -> comment_id
	// findCommentByID(comment_id)

	// list of user -> user_id
	// findArticleByUserID

	go sendMsg(msgCh)

	for txt := range msgCh {
		fmt.Println(txt)
	}
}

func sendMsg(ch chan<- string) {
	for i := 1; i <= 10; i++ {
		text := fmt.Sprintf("Hello from task %d", i)
		ch <- text
	}
	// Memberi signal kepada proses for-range bahwa tidak ada lagi goroutine yg mengirimkan pesan
	close(ch)
}
