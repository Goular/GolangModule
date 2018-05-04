package main

import "fmt"

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score :%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}
	return g
}

func grade2(score int) string {
	g := ""
	switch score {
	case 60:
		panic(fmt.Sprintf("Wrong Score :%d", score))
	case 70:
		g = "F"
	case 80:
		g = "C"
	case 90:
		g = "B"
	case 100:
		g = "A"
	}
	return g
}

func main() {
	fmt.Println(
		grade(60),
		grade(70),
		grade(80),
		grade(90),
		grade(100),
		grade(100),
	)
	fmt.Println(
		grade2(70),
		grade2(80),
		grade2(90),
		grade2(100),
		grade2(100),
	)
}
