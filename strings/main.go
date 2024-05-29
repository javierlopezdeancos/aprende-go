package main

import (
	"fmt"
)

func main() {
 	// Como declarar un string
 	var s string

 	s = "Hello World"

 	// Como declarar un string de forma abreviada
 	sa := "Hello World"

 	fmt.Println(s)
 	fmt.Println(sa)

	// Imprimir la longitud de un string, funcion len()
	fmt.Println(len(s))

	// Recorrer un string caracter a caracter en un bucle for
	for i := 0; i < len(s); i++ {
  	fmt.Print(s[i], " ")
 	}

	// Imprimir en diferentes formatos el carcater de un string
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%T ", s[i])
	}

	fmt.Println("")

	// Reemplacemos `o` por `õ` que tiene una representación unicode `U+00F5`
	sb := "Hellõ World"
	fmt.Println("len(u)", len(sb))

	for i := 0; i < len(sb); i++ {
		fmt.Printf("%c ", sb[i])
	}

	fmt.Println("")

	for i := 0; i < len(sb); i++ {
		fmt.Printf("%v ", sb[i])
	}

	fmt.Println("")

	for i := 0; i < len(sb); i++ {
		fmt.Printf("%x ", sb[i])
	}

	fmt.Println("")

	// Convertir un string a un slice de runas para calcular la longitud de un string correctamente
	sc := "Hellõ World"
	r := []rune(sc)

	fmt.Println("len(r)", len(r))

	for i := 0; i < len(r); i++ {
		fmt.Printf("%c ", r[i])
	}

	fmt.Println("")

	for i := 0; i < len(r); i++ {
		fmt.Printf("%v ", r[i])
	}

	fmt.Println("")

	for i := 0; i < len(r); i++ {
		fmt.Printf("%x ", r[i])
	}

	fmt.Println("")

	for i := 0; i < len(r); i++ {
		fmt.Printf("%T ", r[i])
	}

	fmt.Println("")

	// Usando un loop for/range en un string
	w := "Hellõ World"

	for index, char := range w {
		fmt.Printf("character at index %d is %c\n", index, char)
	}

	// Imprimir una runa
	ra := 'õ'
	fmt.Printf("%x ", ra)
	fmt.Printf("%v ", ra)
	fmt.Printf("%T", ra)

	fmt.Println()

	// Strings usando backticks
	sd := `Hello,\n
  My Big Blue
  "World"!`

  fmt.Println(sd)

	// Los strings son inmutables
	/**
	// Esto no funcionará y lanzará un error
	s := "Hello World"

  s[0] = 'F'

  fmt.Println(s)
	**/

	// Comparacion de caracteres
	fmt.Printf("value of character a is %v of type %T\n", 'a', 'a')
  fmt.Printf("value of character b is %v of type %T\n", 'b', 'b')
  fmt.Println("hence 'b' > 'a' is", 'b' > 'a')

	fmt.Printf("value of character a is %v of type %T\n", 'a', 'a')
  fmt.Printf("value of character A is %v of type %T\n", 'A', 'A')
  fmt.Println("hence 'A' > 'a' is", 'A' > 'a')

  fmt.Printf("\nvalue of character ℻ is %v of type %T\n", '℻', '℻')
  fmt.Printf("value of character ™ is %v of type %T\n", '™', '™')
  fmt.Println("hence '℻' > '™' is", '℻' > '™')

	// Bucle for recorriendo dos rangos de caracteres
	for i := 'a'; i < 'g'; i++ {
    fmt.Printf("character = '%c' with decimal value %v\n", i, i)
  }
}
