- [Arrays en go](#1-arrays-en-go)
  - [Que es un array](#11-que-es-un-array)
  - [Como declarar un array](#12-como-declarar-un-array)
  - [Asignacion de valores a un array](#13-asignacion-de-valores-a-un-array)
  - [Inicializacion de un array](#14-inicializacion-de-un-array)
  - [Inicializacion de un array multilinea](#15-inicializacion-de-un-array-multilinea)
  - [Declaracion automatica de la longitud de un array](#16-declaracion-automatica-de-la-longitud-de-un-array)
  - [Encuentra la longitud de un array](#17-encuentra-la-longitud-de-un-array)
  - [Comparacion de arrays](#18-comparacion-de-arrays)
  - [Iteracion de arrays](#19-iteracion-de-arrays)
  - [Arrays multidimensionales](#110-arrays-multidimensionales)
  - [Pasar un array a una funcion](#111-pasar-un-array-a-una-funcion)
- [References](#2-references)

# 1 Arrays en go

## 1.1 Que es un array

Un `array` es una colección del mismo tipo de datos. Por ejemplo, un array de números enteros o un array de strings. Dado que go es un lenguaje de tipo estático, no se permite mezclar diferentes valores que pertenecen a diferentes tipos de datos en un array.

En go, un `array` tiene una longitud fija. Una vez definido con un tamaño particular, el tamaño del array no se puede aumentar ni disminuir. Pero ese problema se puede resolver usando `slices` que aprenderemos en su propio capitulo `slices`.

Un array es un tipo de datos compuesto o abstracto porque se compone de tipos de datos primitivos o concretos como `int`, `string`, `bool`, etc.

## 1.2 Como declarar un array

Un array es un tipo en sí mismo. Este tipo se define como `[n]T` donde `n` es el número de elementos que puede contener un array y `T` es un tipo de datos como `int` o `string`.

Por lo tanto, para definir una variable que es un `array` de 3 elementos del tipo `int`, tiene la siguiente sintaxis.

```go
package main

import "fmt"

func main() {
  var a [3]int //int array with length 3
  fmt.Println(a)
}
```

```text
[0 0 0]
```

[Ejemplo en vivo](https://go.dev/play/p/zqIixyUmVFN)

En el programa de ejemplo anterior, `a` es un array de 3 elementos enteros pero no hemos asignado ningún valor a los elementos individuales del array. ¿Cuál es su suposición sobre la declaración `Println`? **¿Cuál es el valor de un array vacío?**

Como no le hemos asignado ningún valor a `a`, simplemente definimos el array pero no el valor de los elementos del mismo. Por lo tanto, tendrá un `zero value` en su tipo de datos. Para `int`, su `zero value` es `0`, por lo tanto, la declaración `Println` imprimirá un array de 3 ceros.


## 1.3 Asignacion de valores a un array

Podemos asignar valores individualmente a cada elemento de un array usando su posición en el array, también conocido como `index`. En Go, el `index` del array comienza desde 0, que es el primer elemento, por lo tanto, el `index` del último elemento será `n-1`, donde `n` es la longitud del array.

Para acceder a cualquier elemento del array, necesitamos usar una sintaxis `a[index]` donde `a` es un array variable. Así que tomemos nuestro ejemplo anterior y modifiquémoslo un poco.

```go
package main

import "fmt"

func main() {
	var a [3]int

	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Println("array a => ", a)
	fmt.Println("elements => ", a[0], a[1], a[2])
}
```

imprime una salida

```text
array a =>  [1 2 3]
elements =>  1 2 3
```

[Ejemplo en vivo](https://play.golang.org/p/_4DMJ3iRqRF)

En el ejemplo anterior, asignamos nuevos valores a los 3 elementos del array usando  la sintaxis `a[índex]`.

## 1.4 Inicializacion de un array

Sería bastante difícil asignar un valor a cada elemento de un array si el array es grande, no es una forma que escale muy bien. Por lo tanto, go proporciona una sintaxis abreviada para definir un array con un valor inicial o elementos del array con valores predefinidos. La sintaxis para definir un array con valores iniciales es la siguiente:

```go
var a [n]T = [n]T{V1,V2,...,Vn}
```

En el ejemplo anterior, si los valores de los elementos del array fueran 1, 2, 3, entonces la sintaxis para definir un array sería la siguiente

```go
var a [3]int = [3]int{1, 2, 3}
```

También puede eliminar la declaración del tipo de datos de la declaración de la izquierda y go inferirá el tipo a partir de la definición del array.

```go
var a = [3]int{1, 2, 3}
```

O podrías usar la sintaxis abreviada `:=`, eliminando `var`

```go
a := [3]int{1, 2, 3}
```

No es absolutamente necesario definir todos los elementos de un array. En el ejemplo anterior, podríamos haber definido los dos primeros elementos, dejando el tercer elemento con su `zero value`.

```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2}

	fmt.Println(a)
}
```

```text
[1 2 0]
```

[Ejemplo en vivo](https://go.dev/play/p/J4x7X8Vg73G)

## 1.5 Inicializacion de un array multilinea

Puede definir un array con valores iniciales en varias líneas, pero como estos valores están separados por comas, debe **asegurarse de agregar una coma al final del último elemento**.

```go
package main

import "fmt"

func main() {
	greetings := [4]string{
		"Good morning!",
		"Good afternoon!",
		"Good evening!",
		"Good night!", // must have comma
	}

	fmt.Println(greetings)
}
```

```text
[Good morning! Good afternoon! Good evening! Good night!]
```

[Ejemplo en vivo](https://go.dev/play/p/IcwpC-fQQBj)

Fíjate bien, hemos utilizado la coma `,` al final del último elemento del array. Esta coma es necesaria ya que si no estuviera allí, go habría agregado un punto y coma `;`, lo que habría bloqueado el programa

## 1.6 Declaracion automatica de la longitud de un array

A veces, no sabemos la longitud de un array mientras escribimos sus elementos. Por lo tanto, go proporciona el operador `...` para colocarlo en lugar de `n` en la sintaxis del tipo array `[n]T`. El compilador go encontrará la longitud por sí solo. Solo puede utilizar este operador cuando define un array con un valor inicial.

```go
package main

import "fmt"

func main() {
	greetings := [...]string{
		"Good morning!",
		"Good afternoon!",
		"Good evening!",
		"Good night!",
	}

	fmt.Println(greetings)
}
```

```text
[Good morning! Good afternoon! Good evening! Good night!]
```

[Ejemplo en vivo](https://go.dev/play/p/UKl32kgngnF)

El programa anterior imprimirá el mismo resultado porque el compilador go adivina el valor de 4 a partir del número de elementos del array que son 4.

## 1.7 Encuentra la longitud de un array

Go proporciona una función incorporada `len` que se usa para calcular la longitud de muchos tipos de datos aquí, en este caso, podemos usarla para calcular la longitud de un array.

```go
package main

import "fmt"

func main() {
	greetings := [...]string{
		"Good morning!",
		"Good afternoon!",
		"Good evening!",
		"Good night!",
	}

	fmt.Println(len(greetings))
}
```

```text
4
```

[Ejemplo en vivo](https://play.golang.org/p/wV9edsKhuh4)

## 1.8 Comparacion de arrays

Como comentamos anteriormente en la definición de array, el array es un tipo en sí mismo. `[3]int` es diferente de `[4]int`, que es muy diferente de `[4]string`. Es como comparar `int == string` o `apple == orange`, lo cual no es válido y no tiene sentido. Por lo tanto, estos arrays no se pueden comparar entre sí, a diferencia de otros lenguajes de programación.

Mientras que `[3]int` se puede comparar con `[3]int` incluso si los elementos de su array no coinciden, porque tienen el mismo tipo de datos.

Para que un array sea igual que un segundo array, **ambos arrays deben ser del mismo tipo, deben tener los mismos elementos y todos los elementos deben estar en el mismo orden**. En ese caso, la comparacion `==` será `true`. Si una o más de estas condiciones no coinciden, devolverá `false`.

Go compara primero el tipo de datos y luego cada elemento del array con un elemento de por el `index`.

Echemos un vistazo al siguiente ejemplo.

```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 3, 2}
	c := [3]int{1, 1, 1}
	d := [3]int{1, 2, 3}

	fmt.Println("a == b", a == b)
	fmt.Println("a == c", a == c)
	fmt.Println("a == d", a == d)
}
```

[Ejemplo en vivo](https://go.dev/play/p/U933frU9Gql)

Entonces `a`, `b`, `c` y `d`, todos tienen el mismo tipo de datos de `[3]int`. En la primera comparación, `a == b`, dado que `a` y `b` contienen el mismo elemento pero en diferente orden, esta condición será falsa. En la segunda comparación `a == c`, dado que `c` contiene elementos completamente diferentes que `a`, esta condición será falsa.

Pero en el caso de la tercera comparación, dado que tanto `a` como `d` contienen los mismos elementos con el mismo orden, la condición `a == d` será verdadera. Por lo tanto, el programa anterior imprime el resultado siguiente.

```text
a == b false
a == c false
a == d true
```

## 1.9 Iteracion de arrays

Para iterar sobre un array, podemos usar el bucle `for`.

```go
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}

	for index := 0; index < len(a); index++ {
		fmt.Printf("a[%d] = %d\n", index, a[index])
	}
}
```

```text
a[0] = 1
a[1] = 2
a[2] = 3
a[3] = 4
a[4] = 5
```

[Ejemplo en vivo](https://go.dev/play/p/v5oGN2qQbgN)

En el ejemplo anterior, dado que el índice del elemento del array siempre es menor que `len(a)`, podemos imprimir cada elemento del array usando un bucle `for` simple.

Go a parte proporciona un operador `range` de rango que devuelve el `index` índice y el valor de cada elemento del array en el bucle `for`.

```go
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}

	for index, value := range a {
		fmt.Printf("a[%d] = %d\n", index, value)
	}
}
```

```text
a[0] = 1
a[1] = 2
a[2] = 3
a[3] = 4
a[4] = 5
```

[Ejemplo en vivo](https://go.dev/play/p/XY9pONQaYng)

El operador de rango `range` devuelve el índice `index` y el valor del elemento asociado hasta que todos los elementos del array estén terminados. Si no está interesado en el índice `index`, podemos asignarlo al identificador en blanco `black identifier` `_`.

```go
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}

	for _, value := range a {
		fmt.Println(value)
	}
}
```

```text
1
2
3
4
5
```

[Ejemplo en vivo](https://go.dev/play/p/z0_v0CKquXU)

## 1.10 Arrays multidimensionales

Cuando los elementos de un array son arrays, se denomina array multidimensional. A partir de la definición de array, si un array es una colección de los mismos tipos de datos y un array es un tipo en sí mismo, un array multidimensional debe tener arrays que pertenezcan al mismo tipo de datos.

La sintaxis para escribir un array multidimensional es `[n][m]T` donde `n` es el número de elementos del array y `m` es el número de elementos del array interno. Técnicamente, podemos decir que la matriz contiene `n` elementos de tipo `[m]T`.

```go
package main

import "fmt"

func main() {
	a := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}

	fmt.Println(a)
}
```

```text
[[1 2] [3 4] [0 0]]
```

[Ejemplo en vivo](https://go.dev/play/p/7Hfeb30HD-H)

Dado que podemos usar el operador `...` para adivinar el tamaño de un array, el programa anterior también se puede escribir como

```go
package main

import "fmt"

func main() {
	a := [...][2]int{
		[...]int{1, 2},
		[...]int{3, 4},
		[...]int{5, 6},
	}

	fmt.Println(a)
}
```

```text
[[1 2] [3 4] [5 6]]
```

[Ejemplo en vivo](https://go.dev/play/p/w7oyAb_Acik)

Pero go proporciona una sintaxis corta sorprendente para escribir matrices multidimensionales.

```go
package main

import "fmt"

func main() {
	a := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	fmt.Printf("Array is %v and type of array element is %T", a, a[0])
	fmt.Println()
}
```

```text
Array is [[1 2] [3 4] [5 6]] and type of array element is [2]int
```

[Ejemplo en vivo](https://go.dev/play/p/_ZMipOrWy1R)

En el programa anterior, go ya conoce el tipo de elementos de la matriz que es `[2]int`, por lo que no es necesario mencionarlo nuevamente. También puedes usar `...` operador como

```go
a := [...][2]int{{1, 2}, {3, 4}, {5, 6}}
```

Iterar sobre un array multidimensional no es diferente a la de un array simple. Solo en el caso de un array multidimensional, el valor del elemento del array también es un array que debe recorrerse nuevamente. Por lo tanto, verá un bucle `for` anidado debajo de otro bucle `for` como se muestra a continuación.

```go
for _, child := range parent {
	for _, elem := range child {
		...
	}
}
```

## 1.11 Pasar un array a una funcion

Cuando pasa un array a una función, se pasan por valor como tipo de datos `int` o `cadena`. **La función recibe solo una copia del mismo**. Por lo tanto, cuando realiza cambios en un array dentro de una función, no se reflejará en el array original.

# 2. References

- [Arrays types](https://go.dev/ref/spec#Array_types)
- [Go by examples, arrays](https://gobyexample.com/arrays)
- [Learn go with tests, arrays and slices](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices)
- [Golang arrays and slices](https://golangbot.com/arrays-and-slices/)
- [The anatomy of Arrays in Go](https://webcache.googleusercontent.com/search?q=cache:https://medium.com/rungo/the-anatomy-of-arrays-in-go-24429e4491b7&strip=0&vwsrc=1&referer=medium-parser)
