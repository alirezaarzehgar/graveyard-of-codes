package main

func main() {
	var chs [5]chan int
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int)
	}

	go func() {
		for i := 0; i < len(chs); i++ {
			chs[i] <- i
		}
	}()

	for i := 0; i < len(chs); i++ {

		select {
		case msg := <-chs[0]:
			println(msg)
		case msg := <-chs[1]:
			println(msg)
		case msg := <-chs[2]:
			println(msg)
		case msg := <-chs[3]:
			println(msg)
		case msg := <-chs[4]:
			println(msg)
		}
	}
}
