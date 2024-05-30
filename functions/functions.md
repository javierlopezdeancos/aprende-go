- [Funciones](#1-funciones)
  - [Qué es una función](#11-qu%C3%A9-es-una-funci%C3%B3n)
  - [Convención de nombres para funciones](#12-convenci%C3%B3n-de-nombres-para-funciones)
  - [Parámetros en funciones.](#13-par%C3%A1metros-en-funciones)
  - [Valor de retorno](#14-valor-de-retorno)
  - [Multiples valores de retorno](#15-multiples-valores-de-retorno)
  - [Valores de retorno nombrados](#16-valores-de-retorno-nombrados)
  - [Función recursiva](#17-funci%C3%B3n-recursiva)
  - [defer keyword](#18-defer-keyword)
  - [Función como tipo](#19-funci%C3%B3n-como-tipo)
  - [Función como valor función anónima](#110-funci%C3%B3n-como-valor-funci%C3%B3n-an%C3%B3nima)
  - [Expresión de función invocada inmediatamente IIFE](#111-expresi%C3%B3n-de-funci%C3%B3n-invocada-inmediatamente-iife)
- [Referencias](#2-referencias)

# 1. Funciones

Al igual que JavaScript, las funciones en Go son *first-class citizens*. Se pueden asignar a variables, pasar como argumento, invocar inmediatamente o diferir para la última ejecución.

## 1.1 Qué es una función

Una función, en general, es una pequeña pieza de código que se dedica a realizar una tarea particular en función de algunos valores de entrada. Creamos una función para que podamos realizar dicha operación cuando queramos desde donde queramos proporcionando algunos valores de entrada.

Estos valores de entrada brindan información adicional a una función y son totalmente opcionales. La ejecución de una función puede o no devolver ningún resultado.

En Go, una función se define usando la palabra clave `func`.

```go
func doSomething() {
    fmt.Println("Hello World!")
}
```

Esta función se puede llamar desde cualquier lugar dentro del cuerpo de una función en el programa. Por ejemplo, tenemos la función `doSomething` que imprime algunos caracteres en la salida estándar.

```
Hello World!
```

[The live example](https://go.dev/play/p/THBF9b1nOr-)

## 1.2 Convención de nombres para funciones

Go recomienda escribir los nombres de las funciones en palabras simples o `camelCase`. Incluso los nombres de las funciones under_score son válidos, pero no son idiomáticos en Go.

## 1.3 Parámetros en funciones

Como vimos anteriormente, una función puede tomar valores de entrada para su ejecución. Estos valores de entrada se proporcionan en una llamada de función, llamada argumentos. También se pueden pasar uno o varios argumentos a una función.

[Ejemplo 1. Imprimir mensaje de saludo](https://go.dev/play/p/E_7DdS_Hw4f)

**Code**

```go
package main

import "fmt"

func greet(user string) {
 fmt.Println("Hello " + user)
}

func main() {
 greet("John Doe")
}
```

**Output**

```
Hello John Doe
```

[Ejemplo 2. Sumar dos enteros](https://play.golang.org/p/QUNdCtGO4sA)

**Code**

```go
package main

import "fmt"

func add(a int, b int) {
 c := a + b
 fmt.Println(c)
}

func main() {
 add(1, 5)
}
```

Puede utilizar la notación abreviada de parámetros en caso de que varios parámetros en sucesión sean del mismo tipo de datos.

```go
package main

import "fmt"

func add(a, b int) {
 c := a + b
 fmt.Println(c)
}

func main() {
 add(1, 5)
}
```

**Output**

```
6
```

*`func add(a, b int, c float32)` también es una sintaxis válida porque `a` y `b` son del tipo de datos `int` mientras que c es del tipo de datos `float32`.*

## 1.4 Valor de retorno

Una función también puede devolver un valor que se puede imprimir o asignar a otra variable.

**Code**

```go
package main

import "fmt"

func add(a, b int) int64 {
 return int64(a + b)
}

func main() {
 result := add(1, 5)
 fmt.Println(result)
}
```

**Output**

```
6
```

En caso de que una función devuelva un valor, debe especificar el tipo de datos de un valor de retorno justo después de los paréntesis del parámetro de función.

En el programa anterior, nos aseguramos de que el valor de retorno coincida con el tipo de retorno de una función al convertir el tipo de resultado (originalmente `int`) en `int64`.

## 1.5 Multiples valores de retorno

A diferencia de otros lenguajes de programación, Go puede devolver múltiples valores de la función. En este caso, debemos especificar los tipos de devolución de los valores (al igual que arriba) entre paréntesis justo después de los paréntesis del parámetro de función.

[Example](https://go.dev/play/p/TSSbha-8g9f)

**Code**

```go
package main

import "fmt"

func addMult(a, b int) (int, int) {
 return a + b, a * b
}

func main() {
 addRes, multRes := addMult(2, 5)
 fmt.Println(addRes, multRes)
}
```

**Output**

```
7 10
```

Para capturar múltiples valores de una función que devuelve múltiples valores, debemos especificar la declaración de variables separadas por comas.

En el caso de múltiples valores devueltos pero solo está interesado en un solo valor devuelto por la función, puede asignar otros valores a _ (*blank identifier*) que almacena el valor en una variable vacía.

*Esto es necesario porque si se define una variable pero no se usa en Go, el compilador se queja.*

[Ejemplo en vivo](https://go.dev/play/p/eqLywlCdAAW)

**Code**

```go
package main

import "fmt"

func addMult(a, b int) (int, int) {
 return a + b, a * b
}

func main() {
 _, multRes := addMult(2, 5)
 fmt.Println(multRes)
}
```

**Output**

```
10
```

## 1.6 Valores de retorno nombrados

Los valores de retorno con nombre son una excelente manera de mencionar explícitamente las variables de retorno en la propia definición de la función.

Estas variables se crearán automáticamente y estarán disponibles dentro del cuerpo de la función. Puede cambiar los valores de estas variables dentro de una función.

**Se necesita una declaración de retorno al final de la función para devolver valores con nombre***. Go devolverá automáticamente estas variables cuando la función llegue a la declaración de devolución.

[Example](https://go.dev/play/p/7mjltVetYS8)

**Code**

```go
package main

import "fmt"

func addMult(a, b int) (add int, mul int) {
 add = a + b
 mul = a * b

 return // necessary
}

func main() {
 addRes, multRes := addMult(2, 5)
 fmt.Println(addRes, multRes)
}
```

**Output**

```
7 10
```

También puede combinar valores devueltos con nombre cuando contienen el mismo tipo de datos que se muestra a continuación. Sin embargo, cuando usamos el valor de retorno con nombre, todos los valores de retorno con nombre deben definirse con sus tipos de datos.

[Ejemplo en vivo](https://go.dev/play/p/UBtP7nnBD8_c)

**Code**

```go
package main

import "fmt"

func addMult(a, b int) (add, mul int) {
 add = a + b
 mul = a * b

 return // necessary
}

func main() {
 addRes, multRes := addMult(2, 5)
 fmt.Println(addRes, multRes)
}
```

**Output**

```
7 10
```

**`func math() (add, mult int, div float32)` is also a valid syntax because `add` and `mult` are of `int` data type while div is of `float32` data type.**

## 1.7 Función recursiva

Una función se llama recursiva cuando se llama a sí misma desde el interior del cuerpo. Una sintaxis simple para la función recursiva es

```go
func r() {
    r()
}
```

Si ejecutamos la función anterior `r`, se repetirá infinitamente. Por lo tanto, en una función recursiva, generalmente usamos una declaración condicional como `if-else` para salir del bucle infinito.

Un ejemplo simple de una función recursiva es el `factorial de n`. ¡Una fórmula recursiva simple para el factorial de `n` es `n * (n-1)!` siempre que `n > 0`.

```go
// n! = n*(n-1)! where n > 0
func getFactorial(num int) int {
 if num > 1 {
  return num * getFactorial(num-1)
 } else {
  return 1 // 1! == 1
 }
}
```

La función `getFactorial` anterior es recursiva, ya que estamos llamando a `getFactorial` desde dentro de la función `getFactorial`.

Entendamos cómo funciona.

Cuando se llama a `getFactorial` con un parámetro `int` num, si `num` es igual a `1`, la función devuelve `1`, de lo contrario va dentro del bloque if y ejecuta `num * getFactorial(num-1)`.

Dado que hay una llamada de función, la función `getFactorial` se vuelve a llamar y el valor devuelto se mantendrá en espera hasta que `getFactorial` devuelva algo. Esta pila se mantendrá en construcción hasta que `getFactorial` devuelva algo, que finalmente es `1`.

Tan pronto como eso suceda, toda la `pila de llamadas` se resolverá una por una y eventualmente se resolverá la primera llamada `getFactorial`.

[Ejemplo en vivo](https://go.dev/play/p/xg1_zTFUsd5)

**Code**

```go
package main

import "fmt"

// n! = n×(n-1)! where n >0
func getFactorial(num int) int {
 if num > 1 {
  return num * getFactorial(num-1)
 }

 return 1 // 1! == 1
}

func main() {
 f := getFactorial(4)
 fmt.Println(f)
}
```

**Output**

```
24
```

## 1.8 `defer` keyword

`defer` es una palabra clave en Go que hace que una función se ejecute al final de la ejecución de la función principal o cuando la función principal llega a su declaración de retorno.

Profundicemos en el ejemplo para entenderlo mejor.

**Code**

```go
package main

import "fmt"

func sayDone() {
 fmt.Println("I am done")
}

func main() {
 fmt.Println("main started")

 defer sayDone()

 fmt.Println("main finished")
}
```

**Output**

```
main started
main finished
I am done
```

[Ejemplo en vivo](https://play.golang.org/p/x4JfXO3DEng)

Cuando `main` se ejecuta en la función principal, imprimirá `main started` intentando después ejecutar `sayDone` pero se mantiene en la lista de espera debido a la keyword `defer`. Luego se imprime `main finished` y cuando la función principal deja de ejecutarse, entonces se ejecuta `sayDone()`.

Podemos pasar parámetros para aplazar la función si es compatible, pero hay un problema oculto. Vamos a crear una función simple con argumentos.

**Code**

```go
package main

import "fmt"

func endTime(timestamp string) {
 fmt.Println("Program ended at", timestamp)
}

func main() {
 time := "1 PM"

 defer endTime(time)

 time = "2 PM"

 fmt.Println("doing something")
 fmt.Println("main finished")
 fmt.Println("time is", time)
}
```

**Output**

```
doing something
main finished
time is 2 PM
Program ended at 1 PM
```

Puedes ver el ejemplo completo [aquí](https://play.golang.org/p/bxUskp0MCvo)

En el programa anterior, aplazamos la ejecución de la función `endTime`, lo que significa que se ejecutará al final de la función principal, pero dado que al final de la función principal, `time === "2 PM"`, esperábamos es mensaje `Program ended at 2 PM`. Aunque debido a la ejecución diferida, la función `endTime` se ejecuta al final de la función principal, se colocó antes en la pila con todos los valores de los argumentos disponibles, cuando la variable de tiempo aún era la 1 PM.

> Podemos preguntarnos, ¿qué es este `stack` al que nos referimos? Un `stack` o pila es como un cuaderno donde el compilador de Go escribe funciones diferidas para ejecutar al final de la ejecución de la función actual. Este stack sigue el orden de ejecución `Last In First Out (LIFO)`. Lo que significa que cualquier tarea aplazada primero, se ejecutará al final.

Escribamos varias tareas diferidas y veamos a qué nos referimos:

**Code**

```go
package main

import "fmt"

func greet(message string) {
 fmt.Println("greeting: ", message)
}

func main() {
 fmt.Println("Call one")

 defer greet("Greet one")

 fmt.Println("Call two")

 defer greet("Greet two")

 fmt.Println("Call three")

 defer greet("Greet three")
}
```

**Output**

```
Call one
Call two
Call three
greeting:  Greet three
greeting:  Greet two
greeting:  Greet one
```

Puedes ver el ejemplo completo [aquí](https://play.golang.org/p/NmeNjRPEmTK)

El uso práctico de diferir se puede ver cuando una función tiene demasiadas condiciones, ya sea declaraciones `if-else` o case, y al final de cada condición, debe hacer algo como `cerrar un archivo` o `enviar una respuesta http`. En lugar de escribir varias llamadas, podemos diferir.

A continuación se muestra un ejemplo de un mal programa.

```go
if cond1 {
    ...
    fs.Close(file)
} else if cond2 {
    ...
    fs.Close(file)
} else if cond3 {
    ...
    fs.Close(file)
} else {
    ...
    fs.Close(file)
}
```

A continuación se muestra un ejemplo de un buen programa.

```go
defer fs.Close(file)

if cond1 {
    ...
} else if cond2 {
    ...
} else if cond3 {
    ...
} else {
    ...
}
```

## 1.9 Función como tipo

Una función en Go también es un tipo. Si dos funciones aceptan los mismos parámetros y devuelven los mismos valores, entonces estas dos funciones son del mismo tipo.

Por ejemplo, `add` y `substract`, que individualmente toma dos enteros de tipo `int` y devuelve un entero de tipo `int`, son del mismo tipo.

Hemos visto algunas definiciones de funciones antes, por ejemplo, la función `append` tenia una definición como:

```go
func append(slice []Type, elems ...Type) []Type
```

Por lo tanto, cualquier función, por ejemplo, `prepend` que agrega elementos al comienzo del slice (matriz expandible) si tiene una definición como:

```go
func prepend(slice []Type, elems ...Type) []Type
```

Entonces `append` y `prepend` tendrán el mismo tipo de

```go
func (slice []Type, elems ...Type) []Type
```

Entonces, en Go, el cuerpo de la función no tiene nada que ver con el tipo de función. Pero, ¿cuál es el uso de este tipo?

Esto puede ser útil si está pasando una función como argumento a otra función o cuando una función devuelve otra función y necesita dar un tipo de retorno en una definición de función.

Vamos a crear una función de suma y resta, y veamos cómo se comparan.

**Code**

```go
package main

import "fmt"

func add(a int, b int) int {
 return a + b
}

func subtract(a int, b int) int {
 return a - b
}

func main() {
 fmt.Printf("Type of function add is   %T\n", add)
 fmt.Printf("Type of function subtract is  %T\n", subtract)
}
```

**Output**

```
Type of function add is   func(int, int) int
Type of function subtract is  func(int, int) int
```

Puedes ver el ejemplo completo [aquí](https://play.golang.org/p/LxOPPRvq4Ta)

Entonces puede ver que tanto la función de `add` como la de `subtract` tienen el mismo tipo `func(int, int) int`.

Vamos a crear una función que tome dos números enteros y el tercer argumento, una función que haga una operación matemática con esos dos números. Usaremos la función de `add` y `subtract` como el tercer parámetro de esta función.

**Code**

```go
package main

import "fmt"

func add(a int, b int) int {
 return a + b
}

func subtract(a int, b int) int {
 return a - b
}

func calc(a int, b int, f func(int, int) int) int {
 r := f(a, b)
 return r
}

func main() {
 addResult := calc(5, 3, add)
 subResult := calc(5, 3, subtract)
 fmt.Println("5+3 =", addResult)
 fmt.Println("5-3 =", subResult)
}
```

**Output**

```
5+3 = 8
5-3 = 2
```

En el programa anterior, hemos definido una función `calc` que toma los argumentos `int` `a` & `b` y el tercer argumento de función `f` de tipo `func(int, int) int`. Entonces estamos llamando a la función `f` con `a` y `b` como argumentos.

Puedes ver el ejemplo completo [aquí](https://play.golang.org/p/z3lwQPAhNLJ)

Podemos crear un `tipo derivado` que simplificará las cosas. Podemos reescribir el programa anterior como

**Code**

```go
package main

import "fmt"

func add(a int, b int) int {
 return a + b
}

func subtract(a int, b int) int {
 return a - b
}

type CalcFunc func(int, int) int

func calc(a int, b int, f CalcFunc) int {
 r := f(a, b) // calling add(a,b) or substract(a,b)
 return r
}

func main() {
 addResult := calc(5, 3, add)
 subResult := calc(5, 3, subtract)
 fmt.Println("5+3 =", addResult)
 fmt.Println("5-3 =", subResult)
}
```

**Output**

```
5+3 = 8
5-3 = 2
```

Puedes ver el ejemplo completo [aquí](https://play.golang.org/p/RvayW75C8yy)

> También podríamos escribir el programa anterior de una manera diferente donde la función `calc` en lugar de tomar el tercer argumento como función toma un comando de `string` como `add` o `substract` y devuelve una función anónima basada en el comando que luego podemos ejecutar.*

Como entendimos que una función tiene su propio tipo, podemos declarar una variable de tipo función y asignarle una capa, como se muestra a continuación.

```go
var add func(int, int) int
```

La sintaxis anterior declarará una variable de tipo función que toma dos argumentos `int` y devuelve un valor `int`. Cuando registre la variable add, devolverá `nil`. Esto se debe a que la función `add` no tiene ningún valor.

```go
fmt.Println(add) // <nil>
```

`nil` es un valor cero para muchos tipos como `function`, `pointer`, `slice`, `interface`, `channel`, `map`, etc. Puede leer sobre nil en [este artículo](https://go101.org/article/nil.html).

## 1.10 Función como valor (función anónima)

Una función en Go también puede ser un valor. Esto significa que puede asignar una función a una variable.

**Code**

```go
package main

import "fmt"

var add = func(a int, b int) int {
 return a + b
}

func main() {
 fmt.Println("5+3 =", add(5, 3))
}
```

**Output**

```
5+3 = 8
```

En el programa anterior, hemos creado una variable global `add` y le hemos asignado una función recién creada. Hemos utilizado la inferencia de tipos de Go para obtener el tipo de función anónima (*ya que no hemos mencionado el tipo de `add`*). En este caso, `add` es una función anónima ya que se creó a partir de una función que no tiene nombre.

Puedes ver el ejemplo completo [aquí](https://play.golang.org/p/UoorB5zCGb2)

## 1.11 Expresión de función invocada inmediatamente (IIFE)

Si viene del mundo de `JavaScript`, sabe qué es la expresión de función invocada inmediatamente, pero no se preocupe si no es así. En Go, podemos crear una función anónima que **se puede definir y ejecutar al mismo tiempo**.

Como hemos visto en un ejemplo anterior, una función anónima definida como

**Code**

```go
package main

import "fmt"

var add = func(a int, b int) int {
 return a + b
}

func main() {
 fmt.Println("5+3 =", add(5, 3))
}
```

Donde, `add` es una función anónima. Algunos pueden argumentar que no es realmente anónimo porque aún podemos referirnos a la función de agregar desde cualquier lugar de la función `main` (*en otros casos, desde cualquier lugar del programa*). Pero no en el caso de que una función se invoque o ejecute inmediatamente. Modifiquemos el ejemplo anterior.

**Code**

```go
package main

import "fmt"

func main() {
 sum := func(a int, b int) int {
  return a + b
 }(3, 5)

 fmt.Println("5+3 =", sum)
}
```

En el programa anterior, observe la definición de la función. La primera parte de `func` a `}` define la función mientras que después `(3, 5)` la ejecuta. Por lo tanto, `sum` es el valor devuelto por la ejecución de una función. Por lo tanto, el programa anterior produce el siguiente resultado:

**Output**

```
5+3 = 8
```

Puedes ver el ejemplo [aquí](https://go.dev/play/p/yxRJr51OxzY)

> La función invocada inmediatamente también se puede usar fuera de la función principal en un contexto global. Esto puede ser útil cuando necesita crear una variable global utilizando el valor de retorno de la ejecución **de una función** y no desea revelar la función a las otras partes de su programa.

# 2. Referencias

[Function types](https://go.dev/ref/spec#Function_types)
[Anatomía de funciones en Go](https://medium.com/rungo/the-anatomy-of-functions-in-go-de56c050fe11)
[Defer statements](https://go.dev/ref/spec#Defer_statements)
