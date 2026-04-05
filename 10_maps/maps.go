package main

func main() {
	// map -> key value pairs, unordered, dynamic

	m := map[string]int{"apple": 5, "banana": 3}
	m["orange"] = 2
	// println(m["apple"])
	// println(m["banana"])
	// println(m["orange"])

	value := m["apple"]
	// println("Apple:", value)

	delete(m, "apple")
	value, ok := m["apple"]
	if ok {
		println("Apple:", value)
	} else {
		println("Apple not found")
	}
}
