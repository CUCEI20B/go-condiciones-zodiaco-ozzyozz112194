package main

import "fmt"

func main() {

	var dia int
	fmt.Scanln(&dia)

	var mes int
	fmt.Scanln(&mes)

	if mes == 3 && dia >= 21 || mes == 4 && dia <= 19 {
		fmt.Println("aries")
	}

	if mes == 4 && dia >= 20 || mes == 5 && dia <= 20 {
		fmt.Println("tauro")
	}

	if mes == 5 && dia >= 21 || mes == 6 && dia <= 20 {
		fmt.Println("gemenis")
	}

	if mes == 6 && dia >= 21 || mes == 7 && dia <= 22 {
		fmt.Println("cancer")
	}

	if mes == 7 && dia >= 23 || mes == 8 && dia <= 22 {
		fmt.Println("leo")
	}

	if mes == 8 && dia >= 23 || mes == 9 && dia <= 22 {
		fmt.Println("virgo")
	}

	if mes == 9 && dia >= 23 || mes == 10 && dia <= 22 {
		fmt.Println("libra")
	}

	if mes == 10 && dia >= 23 || mes == 11 && dia <= 21 {
		fmt.Println("escorpio")
	}

	if mes == 11 && dia >= 22 || mes == 12 && dia <= 21 {
		fmt.Println("sagitario")
	}

	if mes == 12 && dia >= 22 || mes == 1 && dia <= 19 {
		fmt.Println("capricornio")
	}

	if mes == 1 && dia >= 18 || mes == 2 && dia <= 18 {
		fmt.Println("acuario")
	}

	if mes == 2 && dia >= 19 || mes == 3 && dia <= 20 {
		fmt.Println("picis")
	}

}
