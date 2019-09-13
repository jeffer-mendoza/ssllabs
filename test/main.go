package main

import "log"

func main() {
	test("A+", "A")  // A
	test("A+", "B")  // B
	test("A+", "C")  // C
	test("C", "A")   // C
	test("C+", "A")  // C+
	test("C+", "B+") // C+
	test("C+", " ")  // C+
}

func test(a string, b string) {
	if a > b {
		log.Println(a)
	} else {
		log.Println(b)
	}
}
