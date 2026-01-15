package main

func printboxStar (n int) {
	for i := 0 ; i < n ; i++ {
		for star := 0 ; star < n ; star++ {
			print("* ")
		}
		println();
	}
}

func main() {
	printboxStar(5)
}