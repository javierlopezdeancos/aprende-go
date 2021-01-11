# Condicionales e iteradores

Go nos provee de **if**, **if-else**, **if if-else if variantes,** **switch** para controlar el flujo condicional de estado con los que ya estamos familiarizados en otros lenguajes.



Estas herramientas se usan para comprobar ciertas condiciones y ejecutar código dependiendo del estado *true o false* de la comprobación de dichas condiciones.



## 1. Condicional if

Go maneja el condicional if como en la mayoría de los lenguajes de programación.



Go no permite escribir la condición dentro de paréntesis.

Veamos el simple ejemplo siguiente que ilustra el funcionamiento del if en go:

```go
package main

import (
	"fmt"
)

func main() {
    fmt.Println()
	fmt.Println("The if condition")

	var condition = true

	if condition {
		fmt.Println("condition met")
	}

	fmt.Println("if terminated")
}

/* 
// Run this code the console outup is
The if condition
condition met
if terminated
*/
```



## 2. Condicional if-else

En el caso del condicional if-else si la condición de if es false, pasará a ejecutarse el bloque else.



```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("The if else condition")

	a := 2

	if a > 10 {
		fmt.Println("condition met")
	} else {
		fmt.Println("condition did not meet")
	}

	fmt.Println("program terminated")
}

/* 
// Run this code the console outup is
The if else condition
condition did not meet
program terminated
*/
```



## 3. Condicional if else if

El condicional es usado cuando necesitas comprobar múltiples condiciones.



Primero el condicional if es evaluado, si la condición coincide (es evaluado a true), el código dentro del bloque if es ejecutado y ninguna condición más es evaluada.



Pero si la condición evaluada en el if no es true, las  condiciones  else if serán evaluadas una a una hasta que una condición se evalue true, en cuyo caso se ejecutará su  bloque de código correspondiente.




Si ninguna condición else if es evaluada a true, se pasará a ejecutar el bloque de instrucciones else. 



```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("The if else if condition")

	fruit := "orange"

	if fruit == "mango" {
		fmt.Println("fruit is mango")
	} else if fruit == "orange" {
		fmt.Println("fruit is orange")
	} else if fruit == "banana" {
		fmt.Println("fruit is banana")
	} else {
		fmt.Println("I don't know which fruit this is")
	}

	// fruit variable is unavailable here
	fmt.Println(fruit)
}

/* 
// Run this code the console outup is
The if else if condition
fruit is orange
*/
```



### Estado inicial

En Go se puede asignar un estado inicial a una variable justo antes de la evaluación de la condición del if.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Initial condition")

	if fruit := "banana"; fruit == "mango" {
		fmt.Println("fruit is mango")
	} else if fruit == "orange" {
		fmt.Println("fruit is orange")
	} else if fruit == "banana" {
		fmt.Println("fruit is banana")
	} else {
		fmt.Println("I don't know which fruit this is")
	}
}

/* 
// Run this code the console outup is
Initial condition
fruit is banana
orange
*/
```



## Condicional switch

```
switch input {
    case value_1:
	    ...
	case value_1:
	    ...
```

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("The switch conditional statement")

	finger := 2

	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}
}

/* 
// Run this code the console outup is
Initial condition
fruit is banana
orange
*/
```


