- [Punteros en Go](#1-punteros-en-go)
  - [Cómo acceder a la dirección de memoria de una variable?](#11-c%C3%B3mo-acceder-a-la-direcci%C3%B3n-de-memoria-de-una-variable)
  - [Que es un puntero?](#12-que-es-un-puntero)
  - [Desreferenciar un puntero](#13-desreferenciar-un-puntero)
  - [Cambiar el valor de la variable usando un puntero](#14-cambiar-el-valor-de-la-variable-usando-un-puntero)
  - [La función new](#15-la-funci%C3%B3n-new)
  - [Pasar un puntero a una función](#16-pasar-un-puntero-a-una-funci%C3%B3n)
  - [Aritmética de punteros](#17-aritm%C3%A9tica-de-punteros)
- [Referencias](#2-referencias)

# 1. Punteros en Go

Un puntero es una variable que almacena los datos de la dirección de memoria a los que hace referencia otra variable. Los punteros tienen el poder de mutar los datos a los que apuntan.

Antes de comenzar a hablar sobre los punteros, aprendamos un par de cosas sobre los números hexadecimales. Un número hexadecimal es un número con base 16. Si eres un desarrollador web, entonces los estás usando durante mucho tiempo, porque en su mayoría; los colores se representan en formato hexadecimal. Por ejemplo, el blanco se representa como `#FFFFFF` y el negro como `#000000`.

En Go, puede guardar un número hexadecimal en una variable y Go proporciona una expresión literal para eso. Si un número comienza con 0x, entonces es un número hexadecimal.

**Code**

```go
package main

import "fmt"

func main() {
 a := 0x00
 b := 0x0A
 c := 0xFF

 fmt.Printf("variable a of type %T with value %v in hex is %X\n", a, a, a)
 fmt.Printf("variable b of type %T with value %v in hex is %X\n", b, b, b)
 fmt.Printf("variable c of type %T with value %v in hex is %X\n", c, c, c)
}
```

**Output**

```
variable a of type int with value 0 in hex is 0
variable b of type int with value 10 in hex is A
variable c of type int with value 255 in hex is FF
```

[Ejemplo](https://go.dev/play/p/NnAJ5go4dfz)

Del ejemplo anterior, podemos ver que los valores representados en el sistema **hexadecimal** se guardan en el sistema **decimal** con tipo de datos `int`.

Pero, ¿por qué estamos aprendiendo sobre números hexadecimales cuando hablamos de punteros? Bueno, primero hablemos de la **dirección de memoria**.

Cuando declara una variable y proporciona algún valor (*datos*), en el tiempo de ejecución de Go se asignará algo de memoria para el valor en la RAM y, según el tipo de datos, asignará un tamaño específico de memoria para almacenar ese valor.

Esa memoria tendrá alguna dirección de memoria (*como una dirección postal*) para que Go pueda encontrar el valor de esa variable cuando se le solicite. *Estas direcciones de memoria se representan en valores hexadecimales*.

## 1.1 Cómo acceder a la dirección de memoria de una variable?

Para acceder al valor de la dirección (*dato*) representado por una variable, Go proporciona el operador `&` (*ampersand*) que se usa delante del nombre de la variable. Al hacer esto, la expresión `&variable_name` devuelve la dirección de memoria del valor (*dato*) al que hace referencia la variable `variable_name`.

**Code**

```go
package main

import "fmt"

func main() {
 a := 0x00
 b := 0x0A
 c := 0xFF

 fmt.Println("&a =", &a)
 fmt.Println("&b =", &b)
 fmt.Println("&c =", &c)
}
```

**Output**

```
&a = 0xc00001c030
&b = 0xc00001c038
&c = 0xc00001c040
```

[Ejemplo](https://go.dev/play/p/15P6SHOBavA)

Vimos esto en la lección de [slices](../example-slices/slices.md) cuando intentábamos demostrar que dos slices pueden hacer referencia a valores del mismo array. En el ejemplo anterior, usando el operador `&`, encontramos la dirección de memoria de la variable `a`, `b` y `c`.

## 1.2 Que es un puntero?

**Un puntero es una variable que apunta a la ubicación de memoria de otra variable** (*en realidad, al valor al que hace referencia la variable*).

Como vimos anteriormente, podemos guardar un valor hexadecimal en una variable, pero under the hood, se guarda como un valor decimal de tipo `int`. Pero hay una trampa.

Aunque puedes guardar una dirección de memoria (*número hexadecimal*) en una variable, no es un puntero o no apunta a la ubicación de memoria de otra variable. Es solo un valor y no tiene idea de lo que significa ese valor.

Ahí es donde entra en juego el **puntero**. Un puntero es solo una variable pero de tipo especial y tipo de datos especial.

Un puntero también guarda la dirección de memoria, pero sabe dónde se encuentra esa memoria en la RAM y cómo recuperar el valor almacenado en esa dirección de memoria. **Puede realizar varios tipos de operaciones con él**, como **leer el valor** almacenado en la dirección de memoria o **escribir un nuevo valor**.

A diferencia de guardar el valor hexadecimal en una variable que tiene tipo `int`, el puntero tiene el tipo de datos `*int` si apunta a la dirección de memoria de datos `int` y `*string` si apunta a la dirección de memoria de datos `string`.

La sintaxis para crear o definir un puntero es `var p *Type` donde `Type` es un tipo de datos, el valor (*dato*) al que apuntará. Vamos a crear un puntero simple.

**Code**

```go
package main

import "fmt"

func main() {
 var pa *int

 fmt.Printf("pointer pa of type %T with value %v\n", pa, pa)
}
```

**Output**

```
pointer pa of type *int with value <nil>
```

[Example](https://go.dev/play/p/77BocN1gVXu)

En el ejemplo anterior, hemos creado el puntero `pa` que apunta a los datos de tipo `int` pero como no le estamos asignando ningún valor inicial, su valor cero es `nil`. Un puntero tiene un valor `nil` porque no apunta a ningún dato (valor) en la RAM en este momento.

Así que vamos a crear una variable de tipo `int` y hacer que `pa` apunte a ella.

**Code**

```go
package main

import "fmt"

func main() {
 a := 1
 var pa *int
 pa = &a

 fmt.Printf("pointer pa of type %T with value %v\n", pa, pa)
}
```

```
pointer pa of type *int with value 0xc00001c030
```

[Ejemplo](https://go.dev/play/p/LHlGgDwSpH6)

En el ejemplo anterior, hemos creado una variable `a` y le hemos asignado un valor inicial de `1`. Go guardará un número entero 1 en algún lugar de la RAM. Luego hemos creado el puntero `pa` que puede apuntar a un valor `int`.

Posteriormente, hemos asignado la dirección de memoria de la variable a (*su valor en realidad*) al puntero `pa` usando la expresión `pa = &a`. El programa anterior también se puede escribir con un formato abreviado de asignación de variables.

**Code**

```go
package main

import "fmt"

func main() {
 a := 1
 pa := &a

 fmt.Printf("pointer pa of type %T with value %v\n", pa, pa)
}
```

**Output**

```
pointer pa of type *int with value 0xc00001c030
```

[Ejemplo](https://go.dev/play/p/la95cObwviy)

En el formato abreviado, Go interpretará que estamos tratando de crear un puntero porque estamos asignando la dirección de memoria de una variable (*usando el operador &*) a la variable que estamos tratando de crear.

Cuando imprime el valor de `pa`, devuelve la dirección de memoria a la que apunta. Además, el tipo de datos de `pa` es `*int`, lo que significa que es un puntero que apunta a los datos de tipo `int`.

Sin embargo, no menciona explícitamente a qué variable o datos apunta. Pero puede encontrar los datos en esa dirección de memoria.

## 1.3 Desreferenciar un puntero

Para averiguar el valor (*dato*) al que apunta un puntero, necesitamos usar el operador `*`, también llamado operador de `desreferenciación` (*dereferencing operator*) que, si se coloca antes de una variable de puntero (*como el operador `&` para obtener la dirección de la memoria*), devuelve los datos en esa memoria.

**Code**

```go
package main

import "fmt"

func main() {
 a := 1
 pa := &a

 fmt.Printf("data at %v is %v\n", pa, *pa)
}
```

**Output**

```
data at 0xc000120000 is 1
```

[Ejemplo](https://go.dev/play/p/TnS5HySFnA9)

## 1.4 Cambiar el valor de la variable usando un puntero

Como vimos en el ejemplo anterior, podemos leer los datos en la ubicación de memoria a la que apunta un puntero, pero también podemos cambiar (escribir) el valor en esa ubicación de memoria.

**Code**

```go
package main

import "fmt"

func main() {
 a := 1
 pa := &a
 *pa = 2

 fmt.Printf("a = %v\n", a)
 fmt.Printf("data at %v is %v\n", pa, *pa)
}
```

**Output**

```
a = 2
data at 0xc0000b2000 is 2
```

[Ejemplo](https://go.dev/play/p/yv9QmgmVQCK)

Como puede ver en el ejemplo anterior, la sintaxis `*pa` lee el valor de la ubicación de memoria señalada por el puntero, pero se puede usar la misma expresión para escribir un nuevo valor en la misma ubicación de memoria.

Si se pregunta, ¿por qué cambió el valor de la variable? Esto se debe a que se ha escrito un nuevo valor en la dirección de memoria a la que hace referencia la variable `a`. Esto prueba que los punteros son bastante poderosos.

> La diferencia entre una variable y un puntero es que una variable almacena el valor en una dirección de memoria y el puntero apunta a una dirección de memoria.

## 1.5 La función new

Go proporciona la función `new` integrada que asigna memoria y devuelve un puntero a esa memoria. La sintaxis de la nueva función es la siguiente.

```go
func new(Type) *Type
```

El primer argumento de la nueva función es el tipo de datos y el valor devuelto de esta función es el puntero de ese tipo de datos. Esta función asignará algo de memoria, escribirá un valor cero del Tipo en esa ubicación de memoria y devolverá un puntero a esa ubicación de memoria.

**Code**

```go
package main

import "fmt"

func main() {
 pa := new(int)

 fmt.Printf("data at %v is %v\n", pa, *pa)
}
```

**Output**

```
data at 0xc00001c030 is 0
```

[Ejemplo](https://go.dev/play/p/WCVmwQCS6cG)

¿Esperaba que el valor (*dato*) en la ubicación de la memoria devuelto por la nueva función fuera nulo? Bueno, *el zero value de un puntero es nulo*, lo que significa que el puntero no apunta a ninguna memoria, pero cuando el puntero lo apunta a una ubicación de memoria, la memoria no puede estar vacía, debe contener algunos datos.

Go almacena el `zero value` del tipo de datos pasado a la nueva función y devuelve la dirección de memoria de la misma. Por lo tanto, si solo está interesado en un puntero, puede usar la nueva función en lugar de crear una nueva variable y luego un puntero que apunte al valor de la variable.

> Por lo tanto, la definición "**Un puntero es una variable que apunta a la dirección de memoria de otra variable**" no es estrictamente cierta. "**Un puntero es una variable que apunta a una dirección de memoria**" es más preciso.

## 1.6 Pasar un puntero a una función

Al igual que una variable, puede pasar un puntero a una función. hay dos maneras de hacer esto. Cree un puntero y luego páselo a la función o simplemente pase una dirección de una variable.

**Code**

```go
package main

import "fmt"

func changeValue(p *int) {
 *p = 2
}

func main() {
 a := 1
 pa := &a
 changeValue(pa)

 fmt.Printf("a = %v\n", a)
}
```

**Output**

```
a = 2
```

[Ejemplo](https://go.dev/play/p/CZof2zNT9Jo)

En el programa anterior, pasamos el puntero `pa` como argumento a la función `changeValue`. Este puntero apunta al valor de la variable `a`. Por lo tanto, dentro de la función, podemos escribir un nuevo valor en la dirección de memoria señalada por el puntero pa que también muta el valor de la variable `a`.

En lugar de adoptar este enfoque largo, podemos acortar el ejemplo anterior pasando la dirección de la variable `a` como argumento. El parámetro `p` de la función `changeValue` ahora es el puntero.

> Dos punteros que apuntan al mismo valor son iguales.

**Code**

```go
package main

import "fmt"

func changeValue(p *int) {
 *p = 2
}

func main() {
 a := 1
 changeValue(&a)

 fmt.Printf("a = %v\n", a)
}
```

**Output**

```
a = 2
```

[Ejemplo](https://go.dev/play/p/kr-LTGyjZs0)

En el programa anterior, la sintaxis del argumento de la función `changeValue` indica a Go que estamos esperando un puntero, especialmente la parte `*int` (*declaración de tipo*) del argumento `p`.

Puede pasar un puntero de un tipo de datos compuesto como un `array` a la función.

**Code**

```go
package main

import "fmt"

func changeValue(p *[3]int) {
 //*p == original array `a`
 // *p[0] != (*p)[0]
 (*p)[0] *= 2
 (*p)[1] *= 3
 (*p)[2] *= 4
}

func main() {
 a := [3]int{1, 2, 3}
 changeValue(&a)

 fmt.Printf("a = %v\n", a)
}
```

**Output**

```
a = [2 6 12]
```

[Ejemplo](https://go.dev/play/p/wp2QxYOSSBC)

También podríamos escribir el programa anterior usando la sintaxis abreviada proporcionada por Go para acceder a los datos desde un puntero de array.

**Code**

```go
package main

import "fmt"

func changeValue(p *[3]int) {
 // (*p)[0] == p[0]
 p[0] *= 2
 p[1] *= 3
 p[2] *= 4
}

func main() {
 a := [3]int{1, 2, 3}
 changeValue(&a)

 fmt.Printf("a = %v\n", a)
}
```

**Output**

```
a = [2 6 12]
```

**Pero pasar el puntero de matriz como parámetro de función no es idiomático para Go**. Deberíamos preferir `slices` en su lugar para esta funcionalidad. Como vimos en la lección de `slices`, podemos pasar un `slice` como argumento a una función y esa función puede mutar los valores dentro del `slice`

## 1.7 Aritmética de punteros

A diferencia del lenguaje C, donde un puntero puede incrementarse o disminuirse, Go no permite la [aritmética de punteros](https://www.tutorialspoint.com/cprogramming/c_pointer_arithmetic.htm).

# Referencias

[Punteros en Go (pointers)](https://medium.com/rungo/pointers-in-go-a789eafccd53)
