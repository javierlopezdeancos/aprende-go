- [Condicionales](#1-condicionales)
  - [Condicional if](#11-condicional-if)
  - [Condicional if-else](#12-condicional-if-else)
  - [Condicional if else if](#13-condicional-if-else-if)
    - [Estado inicial](#131-estado-inicial)
  - [Condicional switch](#14-condicional-switch)
    - [Default case](#141-default-case)
    - [Múltiples valores en el case](#142-m%C3%BAltiples-valores-en-el-case)
    - [Inicial statement](#143-inicial-statement)
    - [Expressionless switch statement](#144-expressionless-switch-statement)
    - [Fallthrough statement](#145-fallthrough-statement)
- [Iteradores](#2-iteradores)
  - [Bucles for](#21-bucles-for)
    - [Sintaxis del bucle for](#211-sintaxis-del-bucle-for)
    - [Variantes del bucle for](#212-variantes-del-bucle-for)
      - [Opcional init statment](#2121-opcional-init-statment)
      - [Opcional post statment](#2122-opcional-post-statment)
      - [Opcional init y post statement](#2123-opcional-init-y-post-statement)
      - [Sin ningún statment](#2124-sin-ning%C3%BAn-statment)
      - [El break statment](#2125-el-break-statment)
      - [El continue statement](#2126-el-continue-statement)
      - [El return statement](#2127-el-return-statement)
      - [Range](#2128-range)
        - [Range sobre un array](#21281-range-sobre-un-array)
        - [Range sobre un map](#21282-range-sobre-un-map)
          - [Range sobre un map usando keys](#212821-range-sobre-un-map-usando-keys)
          - [Range sobre un map usando key/value](#212822-range-sobre-un-map-usando-keyvalue)
- [References](#3-references)

# 1. Condicionales

Go nos provee de **if**, **if-else**, **if if-else if variantes,** **switch** etc para controlar el flujo condicional de ejecución de manera con la que ya estábamos familiarizados en otros lenguajes.

Estas herramientas se usan para comprobar ciertas condiciones y ejecutar código dependiendo del estado *true o false* de  dichas comprobaciones.##

## 1.1 Condicional if

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

## 1.2 Condicional if-else

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

## 1.3 Condicional if else if

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

### 1.3.1 Estado inicial

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

## 1.4. Condicional switch

```go
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
The switch conditional statement
Index
*/
```

### 1.4.1 Default case

A veces necesitamos un equivalente al condicional else para ejecutar una pieza de código en el caso en que la evaluación en el switch no coincida con ningún case, en este caso usamos el **default** block.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println()
    fmt.Println("The default case")

    finger := 6

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
    default:
        fmt.Println("No fingers matched")
    }
}

/*
// Run this code the console outup is
The switch conditional statement
No fingers matched
*/
```

El default case no necesariamente debe estar definido como el último  case, puede estar en cualquier posición del bloque switch.

### 1.4.2 Múltiples valores en el case

Go nos permite presentar múltiples valores de evaluación para un mismo case.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println()
    fmt.Println("Multiple case values")

    letter := "i"

    switch letter {
    case "a", "e", "i", "o", "u":
        fmt.Println("Letter is a vovel.")
    default:
        fmt.Println("Letter is not a vovel.")
    }
}

/*
// Run this code the console outup is
Multiple case values
Letter is a vovel.
*/
```

### 1.4.3 Inicial statement

Hay otra variante para el switch que también teníamos en el if donde se puede añadir un statement inicial.

```go
switch statement; input {
    ...
}
```

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println()
    fmt.Println("The initial statement")

    switch letter := "i"; letter {
    case "a", "e", "i", "o", "u":
        fmt.Println("Letter is a vovel.")
    default:
        fmt.Println("Letter is not a vovel.")
    }
}

/*
// Run this code the console outup is
The initial statement
Letter is a vovel.
*/
```

### 1.4.4 Expressionless switch statement

El valor de entrada a un switch en go es opcional.

Cuando no hay un valor de entrada definido el case será una expresión que devolverá un valor booleano.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println()
    fmt.Println("Expressionless switch statement")

    number := 20

    switch {
    case number <= 5:
        fmt.Println("number is less than or equal to 5")
    case number > 5:
        fmt.Println("number is greater than 5")
    case number > 10:
        fmt.Println("number is greater than 10")
    case number > 15:
        fmt.Println("number is greater than 15")
    }
}

/*
// Run this code the console outup is
Expressionless switch statement
number is greater than 5
*/
```

Mientras la variable `number` es mayor que 20 , la condición `number > 5` en el segundo case es la que ejecutará su bloque de código, al ser el primer case evaluado que retorna true.

### 1.4.5 Fallthrough statement

Normalmente cuando un case block es ejecutado, no se intenta evaluar ningún otro case.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println()
    fmt.Println("The fallthrough statement")

    switch number := 20; {
    case number <= 5:
        fmt.Println("number is less than or equal to 5")
        fallthrough
    case number > 5:
        fmt.Println("number is greater than 5")
        fallthrough
    case number > 10:
        fmt.Println("number is greater than 10")
        fallthrough
    case number > 15:
        fmt.Println("number is greater than 15")
    }
}

/*
// Run this code the console outup is
The fallthrough statement
number is greater than 5
number is greater than 10
number is greater than 15
*/
```

En este ejemplo, `number` es mayor que 5, 10 y 15 pero mientras `number > 5` y su bloque de código es ejecutado, podemos usar la instrucción `fallthrough` que pasará a ejecutar el bloque del siguiente case.

Si el compilador encuentra `fallthrough` , el siguiente bloque es ejecutado sin siquiera evaluar el case de dicho bloque.

La instrucción `fallthrough` debe ser la última en un case block, de lo contrario el compilador de Go lanzará un error.

# 2. Iteradores

## 2.1 Bucles for

### 2.1.1 Sintaxis del bucle for

```go
for init; condition; post {
  ...
}
```

- `init` inicia cualquier variable que pueda necesitarse  al empezar el bucle o usar al finalizarlo.

  La declaración puede ser llamada una vez cuando el bucle empieza a ejecutarse.

- `condition` validaciones del estado para una condición.

  Si la condición devuelve true entonces el código dentro del loop se ejecutará.

- Al final de la ejecución del código el `post` será ejecutado.

  En esta ejecución podremos modificar las variables definidas en el `init` . Después de la ejecución en el `post` la condición de estado se volverá a evaluar de nuevo.

  Si esta evaluación devuelve true, el código dentro del bucle se ejecutará de nuevo, así sucesivamente hasta que el loop termine.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Pritln()
    fmt.Println("The for loops")

    for i := 1; i <= 6; i++ {
        fmt.Printf("Current number is %d \n", i)
    }
}

/*
// Run this code the console outup is
The for loops
Current number is 1
Current number is 2
Current number is 3
Current number is 4
Current number is 5
Current number is 6
*/
```

### 2.1.2 Variantes del bucle for

#### 2.1.2.1 Opcional init statment

Go permite colocar el init statment y ejecutarlo fuera del loop pero en ese caso la variable estará accesible fuera del loop.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println()
    fmt.Println("Optional init statement")

    q := 1

    for ; q <= 6; q++ {
        fmt.Printf("Current number is %d \n", q)
    }
}

/*
// Run this code the console outup is
Optional init statement
Current number is 1
Current number is 2
Current number is 3
Current number is 4
Current number is 5
Current number is 6
*/
```

#### 2.1.2.2 Opcional post statment

Go permite colocar y ejecutar dentro del código del bucle el código post para obtener los mismos resultados.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println()
    fmt.Println("Optional post statement")

    for j := 1; j <= 6; {
        fmt.Printf("Current number is %d \n", j)
        j++
    }
}

/*
// Run this code the console outup is
Optional post statement
Current number is 1
Current number is 2
Current number is 3
Current number is 4
Current number is 5
Current number is 6
*/
```

#### 2.1.2.3 Opcional init y post statement

Un bucle for puede tener solo la condición que comprueba el statment, extrayendo fuera del bucle el init y ubicando dentro del bucle el post.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println()
    fmt.Println("Optional init and post statement")

    x := 1
    for x <= 6 {
        fmt.Printf("Current number is %d \n", x)
        x++
    }
}

/*
Optional init and post statement
Current number is 1
Current number is 2
Current number is 3
Current number is 4
Current number is 5
Current number is 6
*/
```

#### 2.1.2.4 Sin ningún statment

Un bucle for sin init, post o check statetment es un bucle for cuya condition statement siempre es true. Esto quiere decir que estará iterando indefinidamente hasta que sea terminado manualmente dentro del código del bucle. Para esta operación tenemos la palabra reservada break para terminar el loop.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println()
    fmt.Println("Without all statements")

    y := 1
    for {
        fmt.Printf("Current number is %d \n", y)

        if y == 6 {
            break
        }
        y++
    }
}

/*
Optional init and post statement
Without all statements
Current number is 1
Current number is 2
Current number is 3
Current number is 4
Current number is 5
Current number is 6
*/
```

#### 2.1.2.5 El break statment

El break statement es usado dentro del bucle for para terminarlo.

Cualquier código después del break no será ejecutado.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println()
    fmt.Println("The break statement ")

    for l := 1; l <= 10; l++ {
        if l > 6 {
            break
        }

        fmt.Printf("Current number is %d \n", l)
    }
}

/*
Optional init and post statement
The break statement
Current number is 1
Current number is 2
Current number is 3
Current number is 4
Current number is 5
Current number is 6
*/
```

#### 2.1.2.6 El continue statement

El continue statement es usado para saltar a la siguiente iteracción del bucle for, simplemente:

- Ignora la iteracción actual

- Ejecuta el post  statment

- Empieza la siguiente iteracción

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println()
    fmt.Println("The continue statement")

    for f := 1; f <= 10; f++ {
        if f%2 != 0 {
            continue
        }

        fmt.Printf("Current number is %d \n", f)
    }
}

/*
Optional init and post statement
The continue statement
Current number is 2
Current number is 4
Current number is 6
Current number is 8
Current number is 10
*/
```

#### 2.1.2.7 El return statement

Si el bucle for encuentra return statement, la función de ejecución se parará y el valor será devuelto por la función, por ende podemos decir que el return no es algo específico del bucle for.

#### 2.1.2.8 Range

##### 2.1.2.8.1 Range sobre un array

```go
package main

import "fmt"

func main() {
   /* create a slice */
   numbers := []int{0,1,2,3,4,5,6,7,8}

   /* print the numbers */
   for i:= range numbers {
      fmt.Println("Slice item",i,"is",numbers[i])
   }
}
```

##### 2.1.2.8.2 Range sobre un map

###### 2.1.2.8.2.1 Range sobre un map usando keys

```go
package main

import "fmt"

func main() {
   /* create a map*/
   countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo"}

    /* print map using keys*/
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
}
```

###### 2.1.2.8.2.2 Range sobre un map usando key/value

```go
package main

import "fmt"

func main() {
   /* create a map*/
   countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo"}

   /* print map using key-value*/
   for country,capital := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",capital)
   }
}
```

# 3. References

[Anatomy of Conditional Statements and Loops in Go](https://medium.com/rungo/anatomy-of-conditional-statements-and-loops-in-go-aa84352cc34d)
[Go Range](https://www.tutorialspoint.com/go/go_range.htm)
