- [Slices en go](#1-slices-en-go)
  - [Introducción, que es un slice](#11-introducci%C3%B3n-que-es-un-slice)
  - [El Slice es una referencia a un array](#12-el-slice-es-una-referencia-a-un-array)
  - [Longitud y capacidad de un slice](#13-longitud-y-capacidad-de-un-slice)
  - [Slice es una Struct](#14-slice-es-una-struct)
  - [Append function](#15-append-function)
  - [Anonymous array slice](#16-anonymous-array-slice)
  - [Funcion copy](#17-funcion-copy)
  - [Funcion make](#18-funcion-make)
  - [Unpack operator](#19-unpack-operator)
  - [Extract operator](#110-extract-operator)
  - [Slice iteration](#111-slice-iteration)
  - [Pasado por referencia](#112-pasado-por-referencia)
  - [Delete slice elements](#113-elete-slice-elements)
  - [Comparacion de slices](#114-comparacion-de-slices)
  - [Slices multi-dimensionales](#115-slices-multi-dimensionales)
  - [Optimizacion de memoria](#116-optimizacion-de-memoria)
- [Referencias](#2-referencias)

# 1. Slices en go

## 1.1 Introducción, que es un slice

Un slice es como un array que es un contenedor de elementos del mismo tipo de datos, pero el slice puede variar en tamaño.

> [!NOTE]
> El slice es un tipo de datos compuesto y porque está compuesto de un tipo de datos primitivo.

La sintaxis para definir un slice es bastante similar a la de un array pero sin especificar el recuento de elementos. Por tanto, `s` es una slice.

```go
var s []int
```

El código anterior creará un `slice` de tipo de datos `int`, lo que significa que contendrá elementos de tipo de datos `int`. Pero, ¿cual es el `zero value` de un slice? Como vimos en los arrays, el `zero value` de un array es un array en el que todos sus elementos tienen un valor cero del tipo de datos que contiene.

Al igual que una array de elementos enteros `int` con tamaño `n` tendrá `n` ceros como elementos debido a que el `zero value` de `int` es `0`. Pero en el caso de un slice, el `zero value` del slice definido con la sintaxis anterior es nulo. El siguiente programa devolverá `true`.

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s == nil)
}
```

```text
true
```

[Ejemplo en vivo](https://go.dev/play/p/JI6ikCK2f9x)

Pero ¿por qué nada?, te preguntarás. Porque el slice es solo una referencia a un array, no es el array en si. El `zero value` de una referencia es nulo `nil`.

`nil` o no, el slice tiene el tipo `[]Type`. En el ejemplo anterior, el slice `s` tiene el tipo `[]int`.

## 1.2 El Slice es una referencia a un array

Esto puede parecer extraño, pero el slice no contiene ningún dato. Más bien almacena datos en un array. Pero entonces te preguntarás, ¿cómo es eso posible si la longitud del array es fija?

`slice`, cuando es necesario para almacenar más datos, crea un nuevo array de la longitud necesaria behind de scene para acomodar más datos.

Cuando se crea un slice mediante una sintaxis simple `var s []int`, no hace referencia a un array, por lo que su valor es nulo. Veamos ahora cómo hace referencia a un array.

Creemos un array y copiemos algunos de los elementos de ese array a un slice.

```go
package main

import "fmt"

func main() {
	// define empty slice
	var s []int

	fmt.Println("s == nil", s == nil)

	// create an array of int
	a := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// creates new slice
	s = a[2:4]
	fmt.Println("s == nil", s == nil, "and s = ", s)
}
```

```text
s == nil true
s == nil false and s =  [3 4]
```

[Ejemplo en vivo](https://go.dev/play/p/1naC_0qQz_E)

En el programa anterior, hemos definido un slice `s` de tipo `int` pero este slice no hace referencia a un array. Por lo tanto, es nulo y la primera declaración `Println` se imprimirá como `true`.

Más tarde, creamos un array `a` de tipo `int` y asignamos a `s` un slice devuelto desde `a[2:4]`. una sintaxis `[2:4]` devuelve un segmento del array a partir del elemento de índice `2` al elemento de índice `3`. Explicaremos el operador `[:]` más adelante.

Ahora, dado que `s` hace referencia al array `a`, no debe ser nulo, lo cual es cierto desde el segundo `Println` y `s` es `[3,4]`.

Dado que un slice siempre hace referencia a un array, podemos modificar un array y verificar si eso se refleja en el slice.

En el programa anterior, cambiemos el valor del tercer y cuarto elemento del array `a` (índice `2` y `3` respectivamente) y verifiquemos el valor del slice `s`.

```go
package main

import "fmt"

func main() {
	var s []int
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = a[2:4]

	a[2] = 33
	a[3] = 44

	fmt.Println(s)
}
```

```text
[33 44]
```

[Ejemplo en vivo](https://go.dev/play/p/9xi8b8TTqHY)

A partir del resultado anterior, estamos convencidos de que el slice es solo una referencia a un array y cualquier cambio en este array se reflejará en el slice.

## 1.3 Longitud y capacidad de un slice

Como hemos visto en la lección sobre arrays, para encontrar la longitud de un tipo de datos, usamos la función `len`. También estamos usando la misma función `len` para los slices.

```go
package main

import "fmt"

func main() {
	var s []int

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s = a[2:4]

	fmt.Println("Length of s =", len(s))
}
```

```text
Length of s = 2
```

[Ejercicio en vivo](https://go.dev/play/p/tKJaxdY7dYp)

El programa anterior imprimirá `Length of s = 2` en la consola, lo cual es correcto porque **hace referencia solo a 2 elementos del array `a`**, de la posicion `2` a la `4`, es decir, el `2` y `3`.

La **capacidad** de un slice es la cantidad de elementos que puede contener. Go proporciona una built-in function `cap` para obtener este número de capacidad.

```go
package main

import "fmt"

func main() {
	var s []int

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s = a[2:4]

	fmt.Println("Capacity of s =", cap(s))
}
```

```text
Capacity of s = 7
```

[Ejemplo en vivo](https://go.dev/play/p/eAbelmHUkZK)

El programa anterior devuelve `7`, que es la capacidad del slice. Dado que el segmento hace referencia a un array, podría haber hecho referencia a un array hasta el final. Dado que a partir del índice 2 en el ejemplo anterior, hay 7 elementos en el arrayu, la capacidad del array es 7.

¿Eso significa que podemos hacer crecer los slices más allá de su capacidad natural? Claro que si, es una referencia y podemos cambiar la referencia cuanto queramos sobre el array al que estamos referenciando. Para ello utilizaremos la function de go `append`.

## 1.4 Slice es una Struct

Aprenderemos que es un `struct` en su [propia seccion](../structs/structs.md), pero un `struct` es un tipo compuesto por diferentes campos de diferentes tipos a partir de los cuales se crean variables de ese tipo `struct`.

Un slice `struct-type` se veria asi:

```go
type slice struct {
    zerothElement *type
    len int
    cap int
}
```

Un `struct` de slice se compone de un puntero `zerothElement `que apunta al primer elemento de un array a la que hace referencia el slice. `len` y `cap` son la longitud y la capacidad del slice respectivamente. `type` es el tipo de elementos que se componen debajo del array (referenciado)

Por lo tanto, cuando se define un nuevo slice, el puntero `zerothElement` se establece en su `zero value`, que es nulo. Pero cuando un slice hace referencia a un array, ese puntero no será nulo.

Aprenderemos más sobre los punteros en su propia seccion](../pointers/pointers.md), pero el siguiente ejemplo mostrará que la dirección de `a[2]` y `s[0]` es la misma, lo que significa que son exactamente el mismo elemento en la memoria.

```go
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]

	fmt.Println("address of a[2]", &a[2])
	fmt.Println("address of s[0]", &s[0])
	fmt.Println("&a[2] == &s[0] is", &a[2] == &s[0])
}
```

```text
address of a[2] 0xc00011c010
address of s[0] 0xc00011c010
&a[2] == &s[0] is true
```

[Ejemplo en vivo](https://go.dev/play/p/0jUjmjhTCos)

`0xc00011c010` es un valor hexadecimal de la ubicación de la memoria. Es posible que obtenga un resultado diferente.

¿Qué pasará con el array si cambio el valor de un elemento en el slice? Esa es una muy buena pregunta. Como sabemos, el slice no contiene ningún dato, sino que los datos se encuentran en un array. Si cambiamos algunos valores de elementos en el slice, eso debería reflejarse en el array.

```go
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]
	fmt.Println("before -> a[2] =", a[2])

	s[0] = 33
	fmt.Println("after -> a[2] =", a[2])
}
```

```text
before -> a[2] = 3
after -> a[2] = 33
```

[Ejemplo en vivo](https://go.dev/play/p/eEChIs0-66G)

## 1.5 Append function

Puede agregar nuevos valores al slice usando la función build-in `append`. La firma de la función de `append` es

```go
func append(slice []Type, elems ...Type) []Type
```

Esto significa que la función `append` toma un slice como primer argumento, uno o muchos elementos como argumentos adicionales para agregar al slice y devuelve un nuevo slice del mismo tipo de datos. Por lo tanto, el slice es una **variadic function**.

Dado que `append` no muta el slice original, veamos cómo funciona.

```go
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]
	ns := append(s, 55, 66)

	fmt.Printf("s=%v, ns=%v\n", s, ns)
	fmt.Printf("len=%d, cap=%d\n", len(ns), cap(ns))
	fmt.Printf("a=%v", a)
}
```

```text
s=[3 4], ns=[3 4 55 66]
len=4, cap=7
a=[1 2 3 4 55 66 7 8 9]
```

[Ejemplo en vivo](https://go.dev/play/p/dSA5x7TkFeS)

Como podemos ver en los resultados anteriores, `s` permanece sin cambios y se copiaron dos nuevos elementos en el slice `ns`, pero lo interesante es lo que le sucede al array `a`. Se cambió. la funcion `append` muta el array referenciado por el slice `s`.

Esto es absolutamente horrible. Por tanto, los slices no son algo sencillo. Utilice `append` solo para asignar el nuevo slice a si mismo como `s = append(s, ...)` que es más manejable.

**¿Qué pasará si agrego más elementos que la capacidad de un slice?** De nuevo, gran pregunta. ¿Qué tal si lo intentamos primero?

```go
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]
	fmt.Printf("before -> s=%v\n", s)
	fmt.Printf("before -> a=%v\n", a)
	fmt.Printf("before -> len=%d, cap=%d\n", len(s), cap(s))
	fmt.Println("&a[2] == &s[0] is", &a[2] == &s[0])

	s = append(s, 50, 60, 70, 80, 90, 100, 110)
	fmt.Printf("after -> s=%v\n", s)
	fmt.Printf("after -> a=%v\n", a)
	fmt.Printf("after -> len=%d, cap=%d\n", len(s), cap(s))
	fmt.Println("&a[2] == &s[0] is", &a[2] == &s[0])
}
```

```text
before -> s=[3 4]
before -> a=[1 2 3 4 5 6 7 8 9]
before -> len=2, cap=7
&a[2] == &s[0] is true
after -> s=[3 4 50 60 70 80 90 100 110]
after -> a=[1 2 3 4 5 6 7 8 9]
after -> len=9, cap=14
&a[2] == &s[0] is false
```

[Ejemplo en vivo](https://go.dev/play/p/qKtVAka498Z)

Primero creamos un array `a` de `int` y la inicializamos con un montón de valores. Luego creamos el slice `s` a partir del array `a` comenzando desde el índice `2` al `3`.

Desde el primer conjunto de declaraciones `Print`, verificamos los valores de `s` y `a`. Luego nos aseguramos de que `s` haga referencia al array `a` haciendo comprobando que coinciden la dirección de memoria de sus respectivos elementos. También observamos la longitud y la capacidad de los slices en este punto del programa antes de hacer el `append`.

Luego agregamos al slice `7` valores más. Entonces esperamos que el slice `s` tenga `9` elementos, por lo tanto su longitud es `9`, pero no tenemos idea de su nueva capacidad. A partir de una declaración `Print`, descubrimos que el slice `s` creció más que su capacidad inicial de `7` a `14` y su nueva longitud es `9`. Pero el array `a` permanece sin cambios.

Esto parece extraño al principio pero algo sorprendente. Go descubre por sí solo los cálculos de que estamos tratando de enviar más valores al slice que el array al que referencia puede contener, por lo que crea un nuevo array con mayor longitud y copia en ella los valores de slice antiguos. Luego, se agregan nuevos valores del anexo a ese array y el array de origen permanece sin cambios ya que no se realizó ninguna operación en el.

## 1.6 Anonymous array slice

Hasta ahora, vimos un slice que hace referencia a un array que definimos deliberadamente. Pero casi siempre, lo haremos con un array que está oculto y no es accesible.

De manera similar a un array, el slice se puede definir de manera similar con un valor inicial. En este caso, Go creará un array oculto para contener los valores.

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("s=", s)
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))
}
```

```text
s= [1 2 3 4 5 6]
len=6, cap=6
```

[Ejemplo en vivo](https://go.dev/play/p/l_uhlR5KjNY)

Es bastante obvio que la capacidad de este slice es `6` porque el array es creada por Go y Go prefirió crear un array de longitud `6` ya que estamos creando un slice de `6` elementos. Pero, ¿qué pasará cuando agreguemos dos elementos más?


```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	s = append(s, 7, 8)

	fmt.Println("s=", s)
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))
}
```

```text
s= [1 2 3 4 5 6 7 8]
len=8, cap=12
```

[Ejemplo en vivo](https://go.dev/play/p/dmcnLc6Ys8c)

Go creó un array de `12` de longitud porque cuando insertamos `2` elementos nuevos en el slice, el array original de longitud `6` no era suficiente para contener `8` elementos. No se creará ningun nuevo array si agregamos nuevos elementos al slice a menos que el slice exceda la longitud de `12`.

## 1.7 Funcion copy

Go provee una funcion built-in `copy` para copiar los elementos de un slice a otro. La firma de la funcion `copy` es:

```go
func copy(dst []Type, src []Type) int
```

Donde `dst` es el slice de destino y el slice de origen `src`. La función de `copy` devolverá el número de elementos copiados, que es el mínimo de `len(dst)` y `len(src)`.

```go
package main

import "fmt"

func main() {
	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6, 7}
	s4 := []int{1, 2, 3}

	n1 := copy(s1, s2)
	fmt.Printf("n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
	fmt.Println("s1 == nil", s1 == nil)

	n2 := copy(s2, s3)
	fmt.Printf("n2=%d, s2=%v, s3=%v\n", n2, s2, s3)

	n3 := copy(s3, s4)
	fmt.Printf("n3=%d, s3=%v, s4=%v\n", n3, s3, s4)
}
```

```text
n1=0, s1=[], s2=[1 2 3]
s1 == nil true
n2=3, s2=[4 5 6], s3=[4 5 6 7]
n3=3, s3=[1 2 3 7], s4=[1 2 3]
```

[Playground](https://go.dev/play/p/MkFRMZl-v1B)

En el programa anterior, hemos definido el segmento nulo `nil` `s1` y los slices no vacíos `s2` y `s3`. La primera declaración de `copy` intenta copiar `s2` a `s1`, pero como `s1` es un segmento nulo, no sucederá nada y `s1` seguira siendo nulo.

Ese no será el caso con `append`. Como Go está listo para crear una nueva matriz si es necesario, agregar en un segmento nulo `nill` funcionará como se espera.

En la segunda declaración de `copy`, estamos copiando `s3` en `s2`, dado que `s3` contiene `4` elementos y `s2` contiene `3` elementos, solo se copiarán `3` (mínimo de `3` y `4`). Debido a que `copy` no agrega nuevos elementos, solo los reemplaza.

En la tercera declaración de `copy`, estamos copiando `s4` en `s3`. Dado que `s3` contiene `4` elementos y `s4` contiene `3`, solo se reemplazarán `3` elementos en `s3`.

## 1.8 Function make

En el ejemplo anterior, vimos que `s1` permaneció sin cambios porque era un slice nulo y no podia albergar los nuevos elementos del slice que le pretendiamos copiar. Pero hay una diferencia entre un slice nulo y un slice vacío. El slice nulo es un slice al que le falta la referencia a un array y **el slice vacío es un slice con una referencia a un array vacía o cuando el array está vacío**.

make es una `built-in` function que nos ayuda a crear un slice vacío. La firma de la función make es la siguiente. La función make puede crear muchos tipos compuestos vacíos.

```go
func make(t Type, size ...IntegerType) Type
```

En el caso del slice, la función `make` se ve como se muestra a continuación.

```go
s := make([]type, len, cap)
```

Aquí, `type` es el tipo de datos de los elementos del slice, `len` es la longitud del slice y `cap` es la capacidad del slice.

Probemos el ejemplo anterior con `s1` como un slice vacío.

```go
package main

import "fmt"

func main() {
	s1 := make([]int, 2, 4)
	s2 := []int{1, 2, 3}

	fmt.Printf("before => s1=%v, s2=%v\n", s1, s2)
	fmt.Println("before => s1 == nil", s1 == nil)

	n1 := copy(s1, s2)
	fmt.Printf("after => n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
	fmt.Println("after => s1 == nil", s1 == nil)
}
```

```text
before => s1=[0 0], s2=[1 2 3]
before => s1 == nil false
after => n1=2, s1=[1 2], s2=[1 2 3]
after => s1 == nil false
```

[Ejemplo en vivo](https://go.dev/play/p/z0tlrRYLhMu)

El resultado anterior demuestra que se creó un slice vacío y que **la función de `copy` no agrega valores al slice más allá de su longitud, incluso cuando su capacidad es mayor**.

## 1.9 Unpack operator

Algunas personas llaman al operador para `unpack` (desempaquetar ) o `spread` (expandir), para mí `spread` suena más natural. Si ve la sintaxis de la función `append`, acepta más de un argumento para agregar elementos a un slice.

¿Qué sucede si tiene un slice y necesita agregar valores de él a otro segmento? En ese caso el `...` operador es útil porque `append` no acepta un slice como argumento, solo el tipo del que está formado el elemento de slice.

```go
package main

import "fmt"

func main() {
	s1 := make([]int, 0, 10)
	fmt.Println("before -> s1=", s1)

	s2 := []int{1, 2, 3}
	s1 = append(s1, s2...)
	fmt.Println("after -> s1=", s1)
}
```

```text
before -> s1= []
after -> s1= [1 2 3]
```

[Ejemplo en vivo](https://go.dev/play/p/JfLgynyqVYc)

## 1.10 Extract operator

Go proporciona un sorprendente operador `[start:end]` (me gusta llamarlo operador de extracción) que puedes usar fácilmente para extraer cualquier parte de un slice. Tanto el `start` como `end` son índices opcionales.

El `start` es un índice inicial del slice, mientras que el `end` es el último índice hasta el cual se deben extraer los elementos, por lo que no se incluye el índice final. **Esta sintaxis devuelve un nuevo slice**.

```go
package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("s[:]", s[:])
	fmt.Println("s[2:]", s[2:])
	fmt.Println("s[:4]", s[:4])
	fmt.Println("s[2:4]", s[2:4])
}
```

```text
s[:] [0 1 2 3 4 5 6 7 8 9]
s[2:] [2 3 4 5 6 7 8 9]
s[:4] [0 1 2 3]
s[2:4] [2 3]
```

[Ejemplo en vivo](https://go.dev/play/p/lNhNx5KGVrR)

En el ejemplo anterior, tenemos un slice simple de números enteros que van del `0` al `9`.

- *`s[:]` significa extraer todos los elementos de `s` desde el índice `0` hasta el final. Por tanto, devuelve todos los elementos de `s`.*
- *`s[2:]` significa extraer elementos de `s` desde el segundo índice hasta el final. Por lo tanto devuelve `[2 3 4 5 6 7 8 9]`*
- *`s[:4]` significa extraer elementos de `s` comenzando desde el índice `0` hasta el índice `4`, pero sin incluir el índice `4`. Por lo tanto, devuelve `[0 1 2 3]`*
- *`s[2:4]` significa extraer elementos de `s` comenzando desde el segundo índice hasta el cuarto índice pero sin incluir el índice `4`. Por lo tanto, devuelve `[2 3]`*

**Lo importante a recordar es que cualquier slice creado por el operador de extracción todavía hace referencia al mismo array subyacente. Para evitarlo siempre puedes usar las funciones `copy`, `make` o `append` conjuntamente.**

## 1.11 Slice iteration

No hay diferencia como tal entre array y slice cuando se trata de iteración. Prácticamente, un slice es como un array, con la misma estructura; puede usar todas las funciones del array estas iterando sobre slices.

## 1.12 Pasado por referencia

Los sectores todavía se pasan por valor a una función, pero como hacen referencia a un array, parece que se pasan por referencia. Pero en realidad estamos pasando el valor de la direccion de memoria del array al que hace referencia el slice.

```go
package main

import "fmt"

func makeSquares(slice []int) {
	for index, elem := range slice {
		slice[index] = elem * elem
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	makeSquares(s)

	fmt.Println(s)
}
```

En el ejemplo anterior, hemos definido `makeSquares` que toma un slice y reemplaza los elementos del slice de entrada con sus cuadrados. Esto producirá el siguiente resultado.

```text
[0 1 4 9 16 25 36 49 64 81]
```

[Ejemplo en vivo](https://go.dev/play/p/p6O0Uqeww1g)


Esto demuestra que, aunque el slice se pasa por valor, ya que hace referencia a un array, podemos cambiar el valor de los elementos de ese arraya traves del slice y de su valor, la direccion al array, que pasamos a la funcion `makeSquares`.

> Por qué estamos tan seguros de que el slice se pasa por valor, si cambiamos la función del ejemplo `makeSquares` por `func makeSquares(slice []int) { slice = slice[1:5] }` la cual no cambia `s` en la función principal.

Veamos qué pasará si usamos el programa anterior con un array como argumento de entrada a la función.

```go
package main

import "fmt"

func makeSquares(array [10]int) {
	for index, elem := range array {
		array[index] = elem * elem
	}
}

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	makeSquares(a)

	fmt.Println(a)
}
```

```text
[0 1 2 3 4 5 6 7 8 9]
```

[Ejemplo en vivo](https://go.dev/play/p/qE8grYQ8Q0s)

El programa anterior dará como resultado `[0 1 2 3 4 5 6 7 8 9]` lo que significa que `makeSquares` recibió solo una copia del mismo.

## 1.13 Delete slice element(s)

Go no proporciona ninguna palabra clave o función para eliminar elementos de un slice directamente. Necesitamos usar algunos trucos para llegar a hacer esto. Como eliminar un elemento de un slice es como unir un slice detrás y delante del elemento que debe eliminarse, veamos cómo lo hariamos de manera practica.

```go
package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// delete element at index 2 (== 2)
	s = append(s[:2], s[3:]...)
	fmt.Println(s)
}
```

```text
[0 1 3 4 5 6 7 8 9]
```

[Ejemplo en vivo](https://go.dev/play/p/LfLGN2m-uSm)

En el programa anterior, extrajimos un slice de `s` comenzando desde el índice `0` hasta el índice `2`, pero sin incluirlo, y le agregamos un slice que comienza desde el índice `3` hasta el final.

Esto creará un nuevo slice sin índice `2`. El programa anterior imprimirá `[0 1 3 4 5 6 7 8 9]`. Usando esta misma técnica, podemos eliminar múltiples elementos de cualquier parte del slice.

## 1.14 Comparacion de slices

Si probamos el siguiente ejemplo

```go
package main

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 3}
	s2 := []int{0, 1, 2, 3}

	fmt.Println(s1 == s2)
}
package main

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 3}
	s2 := []int{0, 1, 2, 3}

	fmt.Println(s1 == s2)
}
```

```text
./prog.go:9:14: invalid operation: s1 == s2 (slice can only be compared to nil)
```

[Ejemplo en vivo](https://go.dev/play/p/kZ7-SyCBvpt)

Obtendrá un error de `operación no válida: s1 == s2 (el slice solo se puede comparar con nulo)`, lo que significa que los slices solo se pueden comparar para determinar si la condición es nula o no.

Si realmente necesita comparar dos slices, utilice el bucle `for` `range` para hacer coincidir cada elemento de los dos slices o utilice la función `DeepEqual` del paquete `reflect`.


## 1.15 Slices multi-dimensionales

De **manera similar al array**, los slices también pueden ser multi-dimensionales. La sintaxis para definir slices multi-dimensionales es bastante similar a la de los arrays, pero sin mencionar el tamaño del elemento.

```go
s1 := [][]int{
    []int{1, 2},
    []int{3, 4},
    []int{5, 6},
}

// type inference like arrays
s2 := [][]int{
    {1, 2},
    {3, 4},
    {5, 6},
}
```

## 1.16 Optimizacion de memoria

Como sabemos, el slice hace referencia a un array. Si hay una función que devuelve un slice, ese slice podría hacer referencia a un array de gran tamaño. Mientras ese slice esté en la memoria, el array no se puede recolectar como basura y contendrá una gran parte de la memoria del sistema.

A continuación se muestra un mal programa.

```go
package main

import "fmt"

func getCountries() []string {
	countries := []string{
		"United states",
		"United kingdom",
		"Austrilia",
		"India",
		"China",
		"Russia",
		"France",
		"Germany",
		"Spain"
	} // can be much more

	return countries[:3]
}

func main() {
    countries := getCountries()
    fmt.Println(cap(countries)) // 9
}
```

Como puede ver, la capacidad de `countries` es `9`, lo que significa que debajo del array hay `9` elementos.

Para evitarlo, debemos crear un nuevo slice de un array anónimo cuya longitud sea manejable. El siguiente programa es un buen programa.


```go
package main

import "fmt"

func getCountries() (c []string) {
	countries := []string{
		"United states",
		"United kingdom",
		"Austrilia",
		"India",
		"China",
		"Russia",
		"France",
		"Germany",
		"Spain"
	} // can be much more

	c = make([]string, 3) // made empty of length and capacity 3
	copy(c, countries[:3]) // copied to `c`

	return
}

func main() {
	countries := getCountries()
	fmt.Println(cap(countries)) // 3
}
```

# 2. Referencias

- [Slice types](https://go.dev/ref/spec#Slice_types)
- [Slice expressions](https://go.dev/ref/spec#Slice_expressions)
- [Appending and copying slices](https://go.dev/ref/spec#Appending_and_copying_slices)
- [Making slices, maps and channels](https://go.dev/ref/spec#Making_slices_maps_and_channels)
- [Slices tricks](https://github.com/golang/go/wiki/SliceTrick)
- [Arrays and slices](https://golangbot.com/arrays-and-slices/)
- [Arrays and slices / Learn go with go](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices)
