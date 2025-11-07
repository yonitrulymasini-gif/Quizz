package main

import "fmt"

func main() {
	var nomdujoueur string

	fmt.Println("Salut, quel est ton pseudo ?")
	fmt.Scan(&nomdujoueur)
	fmt.Println("Bienvenue "+nomdujoueur+", choisis un théme pour ton quizz :")
}