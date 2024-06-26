- [Variadic functions in go](#1-variadic-functions-in-go)
  - [Que es una variadic function](#11-que-es-una-variadic-function)
    - [Append es una variadic function](#111-append-es-una-variadic-function)
  - [Crear una variadic function](#12-crear-una-variadic-function)
  - [Como pasar un slice a una variadic function](#13-como-pasar-un-slice-a-una-variadic-function)
  - [Slice arguments o Variadic arguments](#14-slice-arguments-o-variadic-arguments)
- [Referencias](#2-referencias)

# 1. Variadic functions in go

Una `variadic function` acepta un número infinito de argumentos y todos estos argumentos se almacenan en un parámetro de tipo slice.

## 1.1 Que es una variadic function

Como hemos visto en una lección de [funciones](../functions/functions.md), una función es un fragmento de código dedicado a realizar un trabajo en particular. Una función toma uno o varios argumentos y puede devolver uno o varios valores.

Las **`variadic functions`** también son funciones, pero pueden tomar un número **infinito** o **variable** de argumentos. Hemos visto esto en la lección de `slices` cuando la función `append` aceptó un número variable de argumentos.

> [!NOTE]
> Solo el ultimo parámetro de una función puede ser `variadic`.

```go
func f(elem ...Type)
```

Una sintaxis típica de una `variadic function` se parece al anterior `...` operador llamado `pack operator` indica a Go como almacenar todos los argumentos del tipo `Type` del parametro `elem`, en un slice. Con esta sintaxis, Go crea una variable `elem` del tipo `[]Type` que es un `slice`. Por lo tanto, todos los argumentos pasados ​​a esta función se almacenan en un `slice` `elem`.

### 1.1.1 Append es una variadic function

¿Alguna vez te has preguntado cómo la función [`append`](https://golang.org/pkg/builtin/#append) utilizada para agregar valores a un `slice` acepta cualquier cantidad de argumentos? Es porque es una `variadic function`.

Veamos la firma de la función `append`

```go
func append(slice []Type, elems ...Type) []Type
```

En esta definición o definición, `elems` es un `variadic params`. Por lo tanto, `append` puede aceptar un número variable de argumentos.

Veamos un ejemplo

```go
append([]Type, args, arg2, argsN)
```

La función `append` espera que el primer argumento sea un slice del tipo `Type`, mientras que puede haber un número variable de argumentos después de ese primer argumento. Si tenemos un `slice` `s2` que queremos agregar a un `slice` `s1`, ¿cómo podriamos hacerlo?

A partir de la sintaxis de la función `append`, no podemos pasar otro `slice` como argumento, tiene que ser uno o varios argumentos de tipo `Type`. Por lo tanto, en su lugar, usaremos el operador de `unpack operator` `...` para **descomprimir el `slice` en la serie de argumentos** (*lo cual es aceptable mediante la función `append`*).

```go
append(s1, s2...)
```

> [!IMPORTANT]
> `...` significa `pack operator` y `unpack operator`, pero si hay tres puntos al final, desempaquetará un `slice` mientras que si hay tres puntos al inicio, empaquetará los argumentos en un `slice`.

Aquí `s1` y `s2` son dos `slices` del mismo tipo. Por lo general, conocemos los parámetros de la función y cuántos argumentos puede aceptar una función. Entonces, ¿cómo sabe la función `append` cuántos parámetros se le han pasado?`

Si miramos la firma de la funcion `append`

```go
func append(slice []Type, elems ...Type) []Type
```

Verás `elems ...Type`, lo que significa empaquetar todos los argumentos entrantes en un segmento de `elems` después del primer argumento.

> [!IMPORTANT]
> Una cosa importante a tener en cuenta es que **sólo el último argumento de una función puede ser `variadic`***..

Entonces, el primer argumento de la función `append` será un slice porque exige un `slice`, pero los argumentos posteriores se empaquetarán en elementos de un solo argumento.

## 1.2 Crear una variadic function

Como vimos anteriormente, la `variadic function` no es más que una función que acepta un número variable de argumentos. Para hacer que una función acepte un número variable de argumentos, necesitamos usar el `pack operator``...Type`.

> [!IMPORTANT]
> El `unpack operator` termina con `...` como un `slice...` mientras que el `pack operator` comienza con `...` como `...Type`.

Escribamos la función `getMultiples` cuyo primer argumento sea un `factor` de tipo `int`, que es un factor de multiplicación y los argumentos variables con el `unpack operator` aplicado al final (*por lo tanto, argumentos **variadic***) de tipo `int` se empaquetan en el `slice` `args`.

En esta función, estamos creando un `slice` vacío `multiples` usando la función `make` con una longitud igual a la longitud de los `args`, que es un `slice`. Usando `for range`, multiplicamos `factor` con elementos de `args` y los guardamos en `multiples`. Luego, devolvemos el `slice` `multiples`.

```go
func getMultiples(factor int, args ...int) []int {
	multiples := make([]int, len(args))

	for index, val := range args {
		multiples[index] = val * factor
	}

	return multiples
}
```

Veamos esto en un ejemplo en acción, podemos implementar esta función dentro de la función `main` como

```go
package main

import "fmt"

func main() {
	m := getMultiples(2, 2, 3)
	fmt.Println("Hello", m)
}

func getMultiples(factor int, args ...int) []int {
	multiples := make([]int, len(args))

	for index, arg := range args {
		multiples[index] = arg * factor
	}

	return multiples
}
```

```text
Hello [4 6]
```

[Ejemplo en vivo](https://go.dev/play/p/orA0CJtSnBs)

## 1.3  Como pasar un slice a una variadic function

Un `slice` es una referencia a un `array`, ¿qué sucede cuando pasa un segmento a una `variadic function` usando el `unpack operator`? ¿Go crea nuevos argumentos de `slice` o mantiene los mismos `slices`?

Dado que no tenemos ninguna herramienta para comparar, `args == s`, necesitamos mutar el `slice` `args` para verificar si el `slice` original `s` ha mutado.

```go
package main

import "fmt"

func getMultiples(factor int, args ...int) []int {
	for index, arg := range args {
		args[index] = arg * factor
	}

	return args
}

func main() {
	s := []int{10, 20, 30}

	mult1 := getMultiples(2, s...)

	fmt.Println(s, mult1)
}
```

```text
[20 40 60] [20 40 60]
```

[Ejemplo en vivo](https://go.dev/play/p/fbNgVGu6JZO)

En el programa anterior, modificamos ligeramente la `variadic function` `getMultiples` y, en lugar de crear un nuevo `slice`, asignamos valores de multiplicación a `args` reemplazando los elementos entrantes con elementos multiplicados.

> [!WARNING]
> En el resultado por consola del programa anterior, podemos ver que los valores de los `slices` cambiaron. Esto significa que, en caso de un `slice`, Go, cuando se pasa a una `variadic function` usando el `unpack operator`, usará el `array` al que referencia para construir un nuevo `slice`. Así que ten cuidado.

## 1.4 Slice arguments o Variadic arguments

Si los `variadic arguments` de una función se convierten en un `slice`. Entonces, ¿por qué necesitamos `variadic functions` cuando podemos lograr la misma funcionalidad con una funcion al uso a la que le pasamos `slices`?

Vemos un ejemplo de una función que acepta un  `num int` y un `slice` `nums []int` e imprime la posicion del numero `num` en dicho `slice` si se encuentra en el.

```go
package main

import (
    "fmt"
)

func find(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)

	found := false

	for i, n := range nums {
			if n == num {
					fmt.Println(num, "found at index", i, "in", nums)
					found = true
			}
	}

	if !found {
			fmt.Println(num, "not found in ", nums)
	}

	fmt.Printf("\n")
}

func main() {
	find(89, []int{89, 90, 95})
	find(45, []int{56, 67, 45, 90, 109})
	find(78, []int{38, 56, 98})
	find(87, []int{})
}
```

```text
type of nums is []int
89 found at index 0 in [89 90 95]

type of nums is []int
45 found at index 2 in [56 67 45 90 109]

type of nums is []int
78 not found in  [38 56 98]

type of nums is []int
87 not found in  []
```

[Ejemplo en vivo](https://go.dev/play/p/rG-XRL3yycJ)

A continuación se detallan las ventajas de utilizar `variadic arguments` en lugar de `slices`.

1. No es necesario crear un `slice` durante cada llamada a la función. Si observamos el programa anterior, hemos creado nuevos `slices` durante cada llamada de función.

	```go
	find(89, []int{89, 90, 95})
	find(45, []int{56, 67, 45, 90, 109})
	find(78, []int{38, 56, 98})
	find(87, []int{})
	```

	Esta creación de `slices` adicionales se puede evitar cuando se utilizan `variadic arguments`.

2. En la línea siguiente

	```go
	find(87, []int{})
	```

 	estamos creando un `slice` vacío solo para satisfacer la firma de la función `find`. Esto no es totalmente necesario en el caso de  `variadic functions`. En este ultimo caso podriamos haber escrito esta linea como

	```go
	find(87)
	```

3. Personalmente creo que el caso de `variadic functions` es más legible que el caso de funciones con `slices` como argumentos.

# 2. Referencias

[Function types](https://go.dev/ref/spec#Function_types)
[Variadic functions in go](https://medium.com/rungo/variadic-function-in-go)
[Variadic functions](https://golangbot.com/variadic-functions/)
[Go by examples: Variadic functions](https://gobyexample.com/variadic-functions)
