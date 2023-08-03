package main

func main() {
	var mp map[string]int = make(map[string]int)
	mp = map[string]int{"Ali": 2, "Mohammad": 10, "Hossein": 3}
	mp["Hamid"] = 100
	delete(mp, "Ali")

	for k, v := range mp {
		println(k, v)
	}
}
