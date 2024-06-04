- [Maps en go](#1-maps-en-go)
  - [Que es un map](#11-que-es-un-map)
  - [Crear un map vacio](#12-crear-un-map-vacio)
  - [Inicializar un map](#13-inicializar-un-map)
  - [Accediendo a los datos de un map](#14-accediendo-a-los-datos-de-un-map)
  - [Longitud de un map](#15-longitud-de-un-map)
  - [Eliminar un elemento de un map](#16-eliminar-un-elemento-de-un-map)
  - [Comparacion de maps](#17-comparacion-de-maps)
  - [Iteracion sobre un map](#18-iteracion-sobre-un-map)
  - [Map con otros tipos de datos](#19-map-con-otros-tipos-de-datos)
  - [Maps son tipos de referencia](#110-maps-son-tipos-de-referencia)
  - [Copiar un map](#111-copiar-un-map)
- [Referencias](#2-referencias)

# 1. Maps en go

## 1.1. Que es un map

Un `map` es como un array excepto que, en lugar de un índice o `index` entero, puede tener un `string` o cualquier otro tipo de datos como clave o `key`.

```go
{
  stringKey: intValue,
  stringKey: intValue
  ...
}
```

La sintaxis  para definir un map es la siguiente:

```go
var myMap map[keyType]valueType
```

Donde `keyType` es el tipo de datos de **map keys**, mientras que `valueType` es el tipo de datos de los **map values**. Un `map` es un tipo de datos compuesto *composite data type** porque está compuesto de tipos de datos primitivos.


Declaremos un simple mapa:

```go
package main

import "fmt"

func main() {
	var m map[string]int

	fmt.Println(m)
	fmt.Println("m == nil", m == nil)
}
```

```text
map[]
m == nil true
```

[Ejemplo en vivo](https://go.dev/play/p/Rv6dEDtbsoZ)

En el programa anterior, hemos declarado un mapa `m` que está vacío, porque el valor cero de un mapa es nulo. Pero lo que pasa con el map `zero value` es que no podemos agregarle valores porque, al igual que los slices, el `map` no contiene ningún dato, sino que hace referencia a la estructura de datos interna que contiene los datos.

Entonces, en el caso de un `map` nulo, falta la estructura de datos interna y asignarle cualquier dato provocará un error panic en tiempo de ejecución `panic: assignment to entry in nil map`. Puede utilizar un `map` nulo como variable para almacenar otro `map` no nulo.

## 1.2 Crear un map vacio

Un **`map` vacío** es como un slice vacío con una estructura de datos interna definida para que podamos usarlo para almacenar algunos datos. Al igual que el slice, podemos usar la función `make` para crear un `map` vacío.

```go
m := make(map[keyType]valueType)
```

Creemos un simple `map` `age` in el cual almacenaremos la edad de algunas personas.

```go
package main

import "fmt"

func main() {
	age := make(map[string]int)
	age["mina"] = 28
	age["john"] = 32
	age["mike"] = 55
	fmt.Println("age of john", age["john"])
}
```

```text
age of john 32
```

[Ejemplo en vivo](https://go.dev/play/p/mrQ_Hw0lZI0)

En el programa anterior, hemos creado un `map` vacío que contiene datos `int` y al que se hace referencia mediante claves o `keys` de `string`. Puede acceder o asignar un valor de un mapa usando una clave como `map[key]`.

Usando esa información, hemos asignado algunos datos de edad de `mina`, `john` y `mike`. Puede agregar tantos valores como desee, ya que un `map` tipo `slice` puede contener un número variable de elementos.

## 1.3 Inicializar un map

En lugar de crear un `map` vacío y asignar nuevos datos, podemos crear un `map` con algunos datos iniciales, como un `array` y un `slice`.

```go
package main

import "fmt"

func main() {
	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	fmt.Println(age)
}
```

```text
map[john:32 mike:55 mina:28]
```

[Ejemplo en vivo](https://go.dev/play/p/49br11mmkqx)

## 1.4 Accediendo a los datos de un map

En el caso de un `array` o `slice`, cuando intentamos acceder fuera del elemento de índice `index` (cuando el índice no existe), Go arrojará un error. Pero no en el caso de `map`.

Cuando intentas acceder al valor mediante una clave que no está en el mapa, Go no arrojará un error; en cambio, devolverá el `zero value` de `valueType`.

```go
package main

import "fmt"

func main() {
	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	fmt.Println(age["mina"])
	fmt.Println(age["jessy"])
}
```

```text
28
0
```

[Ejemplo en vivo](https://go.dev/play/p/fq6i39u8AeF)

`28` es correcto porque esa es la edad de `mina`, pero como `jessy` no está en el mapa, Go devolverá `0 `ya que el `zero value` del tipo de datos `int` es `0`.

Entonces, para verificar si existe una `key` en el mapa o no, Go proporciona otra sintaxis que devuelve 2 valores.

```go
value, ok := m[key]
```

veamos esta nueva sintaxis en un nuevo ejemplo

```go
package main

import "fmt"

func main() {
	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}
	minaAge, minaOk := age["mina"]
	jessyAge, jessyOk := age["jessy"]

	fmt.Println(minaAge, minaOk)
	fmt.Println(jessyAge, jessyOk)
}
```

```text
28 true
0 false
```

[Ejemplo en vivo](https://go.dev/play/p/YYYT2qdiGue)

Entonces, obtenemos información adicional sobre si existe una `key` o no. Si la `key` existe, el segundo parámetro será `true`; de lo contrario, será `false`.

## 1.5 Longitud de un map

Podemos averiguar cuántos elementos contiene un `map` usando la función `len`, que vimos en `array` y `slice`.

```go
package main

import "fmt"

func main() {
	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	fmt.Println("len(age) =", len(age))
}
```

```text
len(age) = 3
```

[Ejemplo en vivo](https://go.dev/play/p/94LOJeBXfq7)

> [!CAUTION]
> No hay nada como la capacidad del `slice` en el `map` porque Go toma el control completo de la estructura de datos interna del `map`. Por lo tanto, no intente utilizar la función `cap` en el `map`.

## 1.6 Eliminar un elemento de un map

A diferencia del `slice` donde necesita usar un truco para eliminar un elemento, Go proporciona una función de eliminación más sencilla para eliminar un elemento del `map`. La sintaxis de la función `delete` es la siguiente.

```go
func delete(m map[Type]Type1, key Type)
```

La función `delete` exige que el primer argumento sea un `map` y el segundo argumento sea un `key ` del `map`.

```go
package main

import "fmt"

func main() {
	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	delete(age, "john")
	delete(age, "jessy")

	fmt.Println(age)
}
```

```text
map[john:32 mike:55]
```

[Ejemplo en vivo](https://go.dev/play/p/i_M0xZ6jleP)

> [!IMPORTANT]
> Si la clave no existe en el `map`, como `jessy` en el ejemplo anterior, Go no generará un error al ejecutar la función de delete.

## 1.7 Comparacion de maps

Al igual que el `slice`, un `map` sólo se puede comparar con nulo o `nil`. Si estás pensando en iterar sobre un `map` y hacer coincidir cada elemento, estás en un gran problema. Pero si necesitas urgentemente comparar dos `maps`, utiliza la función `DeepEqual` del paquete [reflect](https://golang.org/pkg/reflect).

## 1.8 Iteracion sobre un map

Dado que no hay valores de índice `index` en el mapa, no puede usar un bucle `for` simple con un valor de índice `index` incremental hasta que llegue al final. Necesitas usar `for range` para hacerlo.

```go
package main

import "fmt"

func main() {
	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	for key, value := range age {
		fmt.Println(key, "=>", value)
	}
}
```

```text
mina => 28
john => 32
mike => 55
```

[Ejemplo en vivo](https://go.dev/play/p/MbLYk17Zu6C)

`range` en el bucle `for` devolverá la `key` y el valor del elemento del `map`. También puedes usar `_` (*blank identifier*) para ignorar la `key` o el valor en caso de que no lo necesites, al igual que un `array` y un `slice`.

> [!NOTE]
> El orden de recuperación de los elementos en el `map` es aleatorio cuando se utiliza para la iteración. Por lo tanto, no hay garantía de que siempre estén en orden. Eso también explica por qué no podemos comparar dos `maps`.

## 1.9 Map con otros tipos de datos

No es necesario que sólo los tipos `string` sean las `key` de un `map`. Todos los tipos comparables, como `boolean`, `int`, `float`, `complex`, `string`, etc., también pueden ser `key`. Esto debería ser muy obvio, pero `boolean` me deja con el culo torcio porque `boolean` solo puede representar 2 valores, `true` o `false`. Veamos qué pasa dónde podemos usarlo.

```go
package main

import "fmt"

func main() {
	age := map[bool]string{
		true:  "YES",
		false: "NO",
	}

	for key, value := range age {
		fmt.Println(key, "=>", value)
	}
}
```

```text
true => YES
false => NO
```

[Ejemplo en vivo](https://go.dev/play/p/NbeIj54VdFS)

Supongo que encontramos un caso de uso para valores clave `boolean`. Pero ¿qué pasa si añadimos claves duplicadas?

```go
package main

import "fmt"

func main() {
	age := map[bool]string{
		true:  "YES",
		false: "NO",
		true:  "YEAH",
	}

	for key, value := range age {
		fmt.Println(key, "=>", value)
	}
}
```

```text
./prog.go:9:3: duplicate key true in map literal
```

[Ejemplo en vivo](https://go.dev/play/p/sOC6TpettrK)

Esto prueba que no podemos agregar claves duplicadas en un mapa.

## 1.10 Maps son tipos de referencia

Al igual que el `slice`, el `map` hace referencia a una `struct` de datos interna. Cuando copia un `map` en un nuevo `map`, la `struct` de datos interna no se copia, solo se hace referencia.

```go
package main

import "fmt"

func main() {
	var ages map[string]int

	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	ages = age

	delete(ages, "john")

	fmt.Println("age", age)
	fmt.Println("ages", ages)
}
```

```text
age map[mike:55 mina:28]
ages map[mike:55 mina:28]
```

[Ejemplo en vivo](https://go.dev/play/p/n-U1AF4jy8M)

Como era de esperar, el `map` de `ages` ahora tiene dos elementos porque eliminamos uno. No sólo eso, sino que también obtuvimos el mismo cambio en el `map` de `age`. Esto demuestra que, al igual que el `slice` (pero a diferencia de una `array`), cuando asignas una variable con otra variable de `map`, comparten la misma estructura interna.

## 1.11 Copiar un map

Para copiar un `map`, debe utilizar el bucle `for`.

```go
package main

import "fmt"

func main() {
	ages := make(map[string]int)

	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	for key, value := range age {
		ages[key] = value
	}

	delete(ages, "john")

	fmt.Println("age", age)
	fmt.Println("ages", ages)
}
```

```text
age map[john:32 mike:55 mina:28]
ages map[mike:55 mina:28]
```

[Ejemplo en vivo](https://go.dev/play/p/uqwruAyVym4)

En el caso anterior, no copiamos el `map`, sino que utilizamos `key` y `value` del `map` para almacenarlos en un `map` diferente que implementa su propia estructura de datos subyacente.

> [!CAUTION]
> Dado que el `map` hace referencia a la `struct` de datos interna, el `map` pasado como parámetro de función comparte la misma `struct` de datos internos al igual que el slice. Por lo tanto, asegúrese de seguir las mismas pautas que se explican en la lección de [slices](../slices/slices.md).

# 2. Referencias

- [Maps type](https://go.dev/ref/spec#Map_types)
- [Making slices maps and channels](https://go.dev/ref/spec#Making_slices_maps_and_channels)
- [The anatomy of maps in go](https://medium.com/rungo/the-anatomy-of-maps-in-go)
- [Golang Maps Tutorial](https://golangbot.com/maps/)
- [Go by example: Maps](https://gobyexample.com/maps)
- [Learn go with tests: Maps](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps)
