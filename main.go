package main

import "fmt"

func main() {
	// Bir string değişkeni tanımlama ve atama
	message := "Merhaba, Dünya!"
	fmt.Println(message)

	// Bir tamsayı değişkeni tanımlama ve atama
	number := 10
	fmt.Println(number)

	// Bir float değişkeni tanımlama ve atama
	pi := 3.14
	fmt.Println(pi)

	// Bir bool değişkeni tanımlama ve atama
	isTrue := true
	fmt.Println(isTrue)

	// İki tamsayı değişkeninin toplamını hesaplama
	a := 5
	b := 3
	sum := a + b
	fmt.Println("Toplam:", sum)

	// Bir karmaşık sayı değişkeni tanımlama ve atama
	complexNum := complex(3, 2)
	fmt.Println(complexNum)
}
