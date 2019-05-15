package main

import (
	"fmt"
	"math/rand"
)

var arr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d"}

func main() {
	i := 10
	// fmt.Printf("%s", arr[rand.Int31n(14)],arr[rand.Int31n(14)],arr[rand.Int31n(14)],arr[rand.Int31n(14)],arr[rand.Int31n(14)],arr[rand.Int31n(14)],)
	// fmt.Printf("%s", arr[rand.Int31n(14)])
	// fmt.Printf("%s", arr[rand.Int31n(14)])
	// fmt.Printf("%s", arr[rand.Int31n(14)])
	// fmt.Printf("%s", arr[rand.Int31n(14)])
	// fmt.Printf("%s", arr[rand.Int31n(14)])

	fmt.Printf("<html>\n")
	fmt.Printf("<style>\n")
	for j := 0; j < i; j++ {
		fmt.Printf("#div%d {\nposition :relative;\nanimation :anim", j)
		fmt.Printf("%d", j)
		fmt.Printf(" 5s infinite;\n}\n")
		fmt.Printf("@keyframes anim")
		fmt.Printf("%d", j)
		// fmt.Printf("%s", arr[rand.Int31n(14)])
		fmt.Printf(" {\n from {height:%d%%;width:%d%%;top:%d%%;left:%d%%;background:#%s%s%s%s%s%s;transform :rotate(%ddeg);}\n", rand.Int31n(100), rand.Int31n(100), rand.Int31n(100), rand.Int31n(100), arr[rand.Int31n(14)], arr[rand.Int31n(14)], arr[rand.Int31n(14)], arr[rand.Int31n(14)], arr[rand.Int31n(14)], arr[rand.Int31n(14)], rand.Int31n(1000))
		fmt.Printf(" to {height:%d%%;width:%d%%;top:%d%%;left:%d%%;background:#%s%s%s%s%s%s;transform :rotate(%ddeg);}\n}\n", rand.Int31n(100), rand.Int31n(100), rand.Int31n(100), rand.Int31n(100), arr[rand.Int31n(14)], arr[rand.Int31n(14)], arr[rand.Int31n(14)], arr[rand.Int31n(14)], arr[rand.Int31n(14)], arr[rand.Int31n(14)], rand.Int31n(1000))
	}
	fmt.Printf("</style>\n")
	for j := 0; j < i; j++ {
		fmt.Printf("<div id=\"div%d\"></div>\n", j)
	}
	fmt.Printf("</html>")
}
