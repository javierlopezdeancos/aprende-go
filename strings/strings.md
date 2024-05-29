- [Strings en Go](#1-strings-en-go)
  - [Longitud de una cadena](#11-longitud-de-una-cadena)
  - [Usando un loop for/range en un string](#12-usando-un-loop-forrange-en-un-string)
  - [Que es una runa](#13-que-es-una-runa)
  - [Strings son inmutables](#14-strings-son-inmutables)
  - [Strings usando backtick comillas invertidas](#15-strings-usando-backtick-comillas-invertidas)
  - [Comparacion de caracteres](#16-comparacion-de-caracteres)
- [Referencias](#2-referencias)

# 1. Strings en Go

Strings en Go merecen especial atención porque se implementan de manera muy diferente en Go en comparación con otros lenguajes.

Las cadenas **se definen entre comillas dobles "..." y no entre comillas simples, a diferencia de JavaScript**. Las cadenas en Go están codificadas en UTF-8 de forma predeterminada, lo que tiene más sentido en el siglo XXI.

Como UTF-8 admite el juego de caracteres ASCII, no necesita preocuparse por la codificación en la mayoría de los casos.

Escribamos un programa simple. Para definir una variable vacía de tipo `string`, utilice la palabra clave `string`.

```go
package main

import "fmt"

func main() {
  var s string

  s = "Hello World"

  fmt.Println(s)
}
```

```text
Hello World
```

[Ejemplo](https://go.dev/play/p/vMDoeaV3RCY)

## 1.1 Longitud de una cadena

Para encontrar la longitud de una cadena, puede usar la función `len`. La función `len` está disponible en el tiempo de ejecución de Go, por lo que no es necesario importarla desde ningún paquete.

```go
package main

import "fmt"

func main() {
  s := "Hello World"

  fmt.Println(len(s))
}
```

```text
11
```

[Ejemplo](https://go.dev/play/p/Kqj-TJMFyXP)

> [!TIP]
> `len` es una función universal para encontrar la longitud de cualquier tipo de datos, no es exclusiva de `strings`. Aprenderemos más sobre las funciones integradas de Go en próximos tutoriales.

En el programa anterior, `len(s)` imprimirá 11 en la consola ya que la cadena `s` tiene 11 caracteres, incluido un carácter de espacio.

Todos los caracteres de la cadena `Hello World` son caracteres ASCII válidos, por lo que esperamos que cada carácter ocupe solo un byte en la memoria (ya que los caracteres ASCII en UTF-8 ocupan 8 bits o 1 byte).

Verifiquemos eso usando un bucle `for` en la cadena `s`.

```go
package main

import "fmt"

func main() {
  s := "Hello World"

  for i := 0; i < len(s); i++ {
    fmt.Print(s[i], " ")
  }

  fmt.Println()
}
```

```text
72 101 108 108 111 32 87 111 114 108 100
```

[Ejemplo](https://go.dev/play/p/cE32NenaYmN)

¡Guau! Supongo que esperabas que `s[i]` fuera una letra en la cadena `s` donde `i` es el índice del carácter en la cadena que comienza en 0. Entonces, ¿qué es esto? Bueno, estos son el valor decimal de los caracteres `ASCII/UTF-8` en la cadena `Hello World` (consulte la [tabla](http://www.asciichart.com)).

**En Go, una cadena string es, de hecho, un slice de bytes de solo lectura**. Por ahora, imagina que un `slice` es como una matriz simple.

En el ejemplo anterior, estamos iterando sobre un slice de bytes (valores de la matriz uint8). Por lo tanto, `s[i]` imprime el valor decimal del byte que contiene el carácter. Pero para ver caracteres individuales, puede usar la cadena de formato `%c` en la declaración `Printf`. También puede usar la cadena de formato `%v` para ver el valor del byte y `%T` para ver el tipo de datos del valor.

```go
package main

import "fmt"

func main() {
  s := "Hello World"
  fmt.Println("len(s)", len(s))

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
}
```

```text
len(s) 11
H e l l o   W o r l d
72 101 108 108 111 32 87 111 114 108 100
48 65 6c 6c 6f 20 57 6f 72 6c 64
uint8 uint8 uint8 uint8 uint8 uint8 uint8 uint8 uint8 uint8 uint8
```

[Ejemplo](https://go.dev/play/p/wwqhgHcTeIU)

Como puedes ver, cada letra muestra un número decimal que contiene `8 bits` o `1 byte` de memoria en el tipo `uint8`.

Como sabemos, los caracteres `UTF-8` se pueden definir en un tamaño de memoria desde 1 byte (compatible con ASCII) hasta 4 bytes. Por lo tanto, en Go, todos los caracteres se representan en el tipo de datos `int32` (tamaño de 4 bytes). Una unidad de código es el número de bits que utiliza una codificación para una sola celda unitaria. Entonces, UTF-8 usa 8 bits y UTF-16 usa 16 bits para una unidad de código, eso significa que **UTF-8 necesita un mínimo de 8 bits o 1 byte para representar un carácter**.

Pero la pregunta más importante es, si todos los caracteres en `UTF-8` están representados en `int32`, entonces ¿por qué obtenemos el tipo `uint8` en el ejemplo anterior? Como se dijo anteriormente, en Go, una cadena es una porción de bytes de solo lectura. Cuando usamos la función len en una cadena, calcula la longitud de ese segmento.

Cuando usamos el bucle for, recorre el segmento y devuelve un byte a la vez o una unidad de código a la vez. Como hasta ahora, todos nuestros caracteres estaban en el conjunto de caracteres ASCII, el byte proporcionado por el bucle for era un carácter válido o una unidad de código era, de hecho, un `code point`.

Por lo tanto, `%c` en la declaración `Printf` podría imprimir un carácter válido de ese valor de byte. Pero como sabemos, el `code point` `UTF-8` o el valor de carácter se pueden representar mediante series de uno o más bytes (máximo 4 bytes). **¿Qué pasará en el bucle for que vimos antes si introducimos caracteres que no sean ASCII?**

Reemplacemos `o` en Hola por `õ` [letra o minúscula latina con tilde](http://www.utf8-chartable.de) que tiene una representación unicode `U+00F5` y está representado por 2 unidades de código (2 bytes) `c3` `b5` (representación hexadecimal). Entonces, en lugar de `6f` para el carácter `o`, deberíamos esperar `c3` `b5` para el carácter `õ`.

```go
package main

import (
 "fmt"
)

func main() {
  u := "Hellõ World"
  fmt.Println("len(u)", len(u))

  for i := 0; i < len(u); i++ {
  fmt.Printf("%c ", u[i])
  }

  fmt.Println("")

  for i := 0; i < len(u); i++ {
  fmt.Printf("%v ", u[i])
  }

  fmt.Println("")

  for i := 0; i < len(u); i++ {
  fmt.Printf("%x ", u[i])
  }

  fmt.Println("")
}
```

```text
len(s) 12
H e l l Ã µ   W o r l d
72 101 108 108 195 181 32 87 111 114 108 100
48 65 6c 6c c3 b5 20 57 6f 72 6c 64
```

[Ejemplo](https://go.dev/play/p/rhueGpn4pDc)

Del resultado anterior obtuvimos `c3 b5` en lugar de `6f` pero los caracteres de `Hellõ World` no se imprimieron muy bien. También vemos que `len(u)` devuelve 12 porque len cuenta el número de bytes en una cadena y eso causó este problema.

Como indexar una cadena (usando un bucle for en ella) se accede a bytes individuales, no a caracteres. Por lo tanto, c3 (195 decimal) en UTF-8 representa Ã y b5 (181 decimal) representa µ.

Para evitar el caos anterior, **Go introduce el tipo de datos runa** que es un alias de `int32` y les explicaba (pero aún no lo he demostrado) que Go representa un carácter en el tipo de datos `int32`.

> [!NOTE]
> Una respuesta interesante sobre por qué runa es int32 y no uint32 (ya que el valor del `code point` de carácter no puede ser negativo y el tipo de datos int32 puede contener valores tanto negativos como positivos) está [aquí](https://stackoverflow.com/questions/24714665/why-is-rune-in-golang-an-alias-for-int32-and-not-uint32).

Entonces, en lugar de slice de bytes, necesitamos convertir un string en un slice de runas.

```go
package main

import "fmt"

func main() {
  s := "Hellõ World"
  r := []rune(s)

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
}
```

```text
H e l l õ   W o r l d
72 101 108 108 245 32 87 111 114 108 100
48 65 6c 6c f5 20 57 6f 72 6c 64
int32 int32 int32 int32 int32 int32 int32 int32 int32 int32 int32
```

[Ejemplo](https://go.dev/play/p/ELgL-upVnz_r)

Convertimos una cadena en una porción de runas mediante **conversión de tipos**. Observe `f5` en el resultado anterior en lugar de `c3` `b5`.

Esto sucedió porque al convertir la cadena `s` en un slice de runas, `c3 b5` se convirtió a `f5` ya que `c3 b5` representa colectivamente el carácter `õ` y el `code point` de `õ` en la tabla UTF es f5 (por lo tanto, la representación del `code point` Unicode U+00F5) o decimal 245 ([consultar aquí](https://www.obliquity.com/computer/html/unicode0000.html)).

Además, obtuvimos la longitud 11 de la cadena `s`, lo cual es correcto, porque hay 11 runas en el segmento (o 11 `code point` o 11 caracteres). Y también demostramos que un `code point`  o un carácter en Go está representado por el tipo de datos `int32`.

## 1.2 Usando un loop for/range en un string

Si usa `range` dentro de un bucle for, `range` devolverá `runa` y el indice del byte del carácter.

```go
package main

import "fmt"

func main() {
  s := "Hellõ World"

  for index, char := range s {
    fmt.Printf("character at index %d is %c\n", index, char)
  }
}
```

```text
character at index 0 is H
character at index 1 is e
character at index 2 is l
character at index 3 is l
character at index 4 is õ
character at index 6 is
character at index 7 is W
character at index 8 is o
character at index 9 is r
character at index 10 is l
character at index 11 is d
```

[Ejemplo](https://go.dev/play/p/Xet2cJbywLH)

En el programa anterior, perdimos el índice 5 porque el quinto byte es la segunda `code unit` del carácter `õ`. Si no necesitas el valor del índice, puedes ignorarlo usando _ (`blank identifier`) en su lugar.

## 1.3 Que es una runa

Un string es una slice de bytes o enteros uint8, así de simple. Cuando usamos el bucle `for/range`, obtenemos runa porque cada carácter del string está representado por el tipo de datos de runa.

En Go, **un carácter se puede representar entre comillas simples**, también conocido como **carácter literal**. Por lo tanto, cualquier carácter `UTF-8` válido dentro de una comilla simple `(')` es una `runa` y su tipo es `int32`.

```go
package main

import "fmt"

func main() {
  r := 'õ'
  fmt.Printf("%x ", r)
  fmt.Printf("%v ", r)
  fmt.Printf("%T", r)

  fmt.Println()
}
```

```text
f5 245 int32
```

El programa anterior imprimirá `f5 245` `int32` que es un valor hexadecimal/decimal y un tipo de datos de valor de `code point` de `õ` en la tabla UTF.

[Ejemplo](https://go.dev/play/p/QNBsDunKTrJ)

## 1.4 Strings son inmutables

Como se ve en la definición anterior de strings, son un slice de bytes de solo lectura. Por lo tanto, si intentamos reemplazar cualquier byte en el segmento, el compilador arrojará un error.

```go
package main

import "fmt"

func main() {
  s := "Hello World"

  s[0] = 'F'

  fmt.Println(s)
}
```

```text
./prog.go:8:2: cannot assign to s[0] (neither addressable nor a map index expression)

Go build failed.
```

[Ejemplo](https://go.dev/play/p/9Uu5LqNqVkb)

El programa anterior no se compilará y el compilador arrojará un error, no se puede asignar a `s[0]` ya que la cadena `s` es un slice de bytes de solo lectura.

Sin embargo, puede crear un string a partir de un slice de bytes y no solo a partir de un string literal. Pero una vez realizada la conversión de slice a string, no podrá modificar el string como se explica en el ejemplo anterior.

```go
var1 := []uint8{72, 101, 108, 108, 111} // [72 101 108 108 111]
var2 := string(var1) // Hello
```

> [!NOTE]
> Recuerda que, un `byte` es un alias para `unit8` y `rune` es un alias para `int32`. Por lo tanto, puedes usarlos indistintamente.

## 1.5 Strings usando `backtick` (comillas invertidas)

En lugar de comillas dobles, también podemos usar el carácter de comilla invertida backtick (`) para representar una cadena en Go. Usando comillas dobles (“) debes escapar de nuevas líneas, tabulaciones y otros caracteres que no necesitan escaparse entre comillas invertidas.

Si pones un salto de línea en una cadena de acento grave, se interpreta como un carácter '\n', ver [string literals](https://golang.org/ref/spec#String_literals)

> [!NOTE]
> El valor de un string literal sin formato es el string formado por caracteres no interpretados (implícitamente codificados en UTF-8) entre las comillas invertidas; en particular, las barras invertidas no tienen un significado especial y el string puede contener nuevas líneas. Los caracteres de retorno de carro (\r) dentro de un string literal sin formato se descartan del valor de string sin formato.

Veamos un pequeño ejemplo

```go
package main

import "fmt"

func main() {
  s := `Hello,\n
  My Big Blue
  "World"!`

  fmt.Println(s)
}
```

```text
Hello,\n
  My Big Blue
 "World"!
```

[Ejemplo](https://go.dev/play/p/9Ir-0Lxx0u3)

Podemos ver que el formato original del string con una nueva línea, tabulación y las dobles comillas se mantuvieron en la salida y el carácter de nueva línea \n no afecto en nada mientras que se descartó el retorno de carro \r.

## 1.6 Comparacion de caracteres

Como el carácter representado entre comillas simples en Go es runa, la runa se puede comparar porque representan `code points` Unicode (valores `int32`). Por lo tanto, si un carácter tiene más valor decimal, será mayor que el carácter que tiene menor.

Veamos un ejemplo muy sencillo.

```go
package main

import (
 "fmt"
)

func main() {
  fmt.Printf("value of character a is %v of type %T\n", 'a', 'a')
  fmt.Printf("value of character b is %v of type %T\n", 'b', 'b')
  fmt.Println("hence 'b' > 'a' is", 'b' > 'a')
}
```

```text
value of character a is 97 of type int32
value of character b is 98 of type int32
hence 'b' > 'a' is true
```

[Ejemplo](https://go.dev/play/p/lxGiJzNeNWO)

Dado que el valor `int32` de `b` es mayor que `a`, la expresión `'b' > 'a'` será verdadera. Veamos otro ejemplo.

```go
package main

import (
 "fmt"
)

func main() {
  fmt.Printf("value of character a is %v of type %T\n", 'a', 'a')
  fmt.Printf("value of character A is %v of type %T\n", 'A', 'A')
  fmt.Println("hence 'A' > 'a' is", 'A' > 'a')

  fmt.Printf("\nvalue of character ℻ is %v of type %T\n", '℻', '℻')
  fmt.Printf("value of character ™ is %v of type %T\n", '™', '™')
  fmt.Println("hence '℻' > '™' is", '℻' > '™')
}
```

```text
value of character a is 97 of type int32
value of character A is 65 of type int32
hence 'A' > 'a' is false

value of character ℻ is 8507 of type int32
value of character ™ is 8482 of type int32
hence '℻' > '™' is true
```

[Ejemplo](https://go.dev/play/p/aw8Sv8Vto-c)

Como sabemos que los caracteres internamente no son más que `int32`, podemos hacer todo tipo de comparaciones con ellos. Por ejemplo, un bucle for entre dos rangos de valores de caracteres.

```go
package main

import (
  "fmt"
)

func main() {
  for i := 'a'; i < 'g'; i++ {
    fmt.Printf("character = '%c' with decimal value %v\n", i, i)
  }
}
```

```text
character = 'a' with decimal value 97
character = 'b' with decimal value 98
character = 'c' with decimal value 99
character = 'd' with decimal value 100
character = 'e' with decimal value 101
character = 'f' with decimal value 102
```

[Ejemplo](https://go.dev/play/p/kS4vxuSSmWg)

# 2. Referencias

- [Tipo de datos string en go](https://medium.com/rungo/string-data-type-in-go-8af2b639478)
- [String literals](https://golang.org/ref/spec#String_literals)
- [Strings types](https://go.dev/ref/spec#String_types)
