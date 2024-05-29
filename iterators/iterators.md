- [Iteradores](#1-iteradores)
  - [Bucles for](#11-bucles-for)
    - [Sintaxis del bucle for](#111-sintaxis-del-bucle-for)
    - [Variantes del bucle for](#112-variantes-del-bucle-for)
      - [Opcional init statment](#1121-opcional-init-statment)
      - [Opcional post statment](#1122-opcional-post-statment)
      - [Opcional init y post statement](#1123-opcional-init-y-post-statement)
      - [Sin ningún statment](#1124-sin-ning%C3%BAn-statment)
      - [El break statment](#1125-el-break-statment)
      - [El continue statement](#1126-el-continue-statement)
      - [El return statement](#1127-el-return-statement)
      - [Range](#1128-range)
        - [Range sobre un array](#11281-range-sobre-un-array)
        - [Range sobre un map](#11282-range-sobre-un-map)
          - [Range sobre un map usando keys](#112821-range-sobre-un-map-usando-keys)
          - [Range sobre un map usando key/value](#112822-range-sobre-un-map-usando-keyvalue)
- [References](#2-references)

# 1. Iteradores

## 1.1 Bucles for

### 1.1.1 Sintaxis del bucle for

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

### 1.1.2 Variantes del bucle for

#### 1.1.2.1 Opcional init statment

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

#### 1.1.2.2 Opcional post statment

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

#### 1.1.2.3 Opcional init y post statement

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

#### 1.1.2.4 Sin ningún statment

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

#### 1.1.2.5 El break statment

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

#### 1.1.2.6 El continue statement

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

#### 1.1.2.7 El return statement

Si el bucle for encuentra return statement, la función de ejecución se parará y el valor será devuelto por la función, por ende podemos decir que el return no es algo específico del bucle for.

#### 1.1.2.8 Range

##### 1.1.2.8.1 Range sobre un array

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

##### 1.1.2.8.2 Range sobre un map

###### 1.1.2.8.2.1 Range sobre un map usando keys

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

###### 1.1.2.8.2.2 Range sobre un map usando key/value

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

# 2. References

- [For statements](https://go.dev/ref/spec#For_statements)
- [Go Range](https://www.tutorialspoint.com/go/go_range.htm)
- [Break statements](https://go.dev/ref/spec#Break_statements)
- [Continue statements](https://go.dev/ref/spec#Continue_statements)
