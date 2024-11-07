- [Maps en go](#1-maps-en-go)
  - [Crear un map vacio](#11-crear-un-map-vacio)
  - [Inicializar un map](#12-inicializar-un-map)
  - [Accediendo a los datos de un map](#13-accediendo-a-los-datos-de-un-map)
  - [Longitud de un map](#14-longitud-de-un-map)
  - [Eliminar un elemento de un map](#15-eliminar-un-elemento-de-un-map)
  - [Comparación de maps](#16-comparaci%C3%B3n-de-maps)
  - [Iteración sobre un map](#17-iteraci%C3%B3n-sobre-un-map)
  - [Map con otros tipos de datos](#18-map-con-otros-tipos-de-datos)
  - [Maps son tipos de referencia](#19-maps-son-tipos-de-referencia)
  - [Copiar un map](#110-copiar-un-map)
- [Aprender maps haciendo tests](#2aprender-maps-haciendo-tests)
  - [Buscar un item en un map por su key](#21-buscar-un-item-en-un-map-por-su-key)
    - [Escribiendo el primer test](#211-escribiendo-el-primer-test)
    - [Correr el test](#212-correr-el-test)
    - [Escribir el minimo codigo para correr el test y ver su output](#213-scribir-el-minimo-codigo-para-correr-el-test-y-ver-su-output)
    - [Escribir el codigo para que el test pase](#214-escribir-el-codigo-para-que-el-test-pase)
    - [Refactor](#215-refactor)
    - [Crear un helper para el test](#216-crear-un-helper-para-el-test)
    - [Usar un tipo personalizado para el diccionario](#217-usar-un-tipo-personalizado-para-el-diccionario)
    - [Escribir un test para el caso de que la word no este en el dictionary](#218-escribir-un-test-para-el-caso-de-que-la-word-no-este-en-el-dictionary)
    - [Escribir el codigo para que el test corra y poder ver su output](#219-escribir-el-codigo-para-que-el-test-corra-y-poder-ver-su-output)
    - [Escribir el codigo necesario para que el test pase](#2110-escribir-el-codigo-necesario-para-que-el-test-pase)
    - [Refactor](#2111-refactor)
  - [Agregar un item a un map](#22-agregar-un-item-a-un-map)
    - [Escribir el primer test para añadir un item a un map](#221-escribir-el-primer-test-para-a%C3%B1adir-un-item-a-un-map)
    - [Escribir el código necesario para que el test corra y poder ver su output](#222-escribir-el-c%C3%B3digo-necesario-para-que-el-test-corra-y-poder-ver-su-output)
    - [Escribir el código necesario para que el test pase](#223-escribir-el-c%C3%B3digo-necesario-para-que-el-test-pase)
    - [Punteros copias etc](#224-punteros-copias-etc)
    - [Refactor](#225-refactor)
    - [Escribir un test para el caso de que la word ya exista](#226-escribir-un-test-para-el-caso-de-que-la-word-ya-exista)
    - [Correr el test](#227-correr-el-test)
    - [Escribir el código necesario para que el test pase](#228-escribir-el-c%C3%B3digo-necesario-para-que-el-test-pase)
    - [Escribir el código necesario para que el test pase](#229-escribir-el-c%C3%B3digo-necesario-para-que-el-test-pase)
    - [Refactor](#2210-refactor)
  - [Actualizar el item de un map](#23-actualizar-el-item-de-un-map)
    - [Escribir el primer test](#231-escribir-el-primer-test)
    - [Correr el test](#232-correr-el-test)
    - [Escribir el código necesario para que el test corra y poder ver su output](#233-escribir-el-c%C3%B3digo-necesario-para-que-el-test-corra-y-poder-ver-su-output)
    - [Escribir el código necesario para que el test pase](#234-escribir-el-c%C3%B3digo-necesario-para-que-el-test-pase)
    - [Escribiendo un test para el caso de que la word que queremos actualizar sea nueva en el diccionario](#235-escribiendo-un-test-para-el-caso-de-que-la-word-que-queremos-actualizar-sea-nueva-en-el-diccionario)
    - [Correr el test](#236-correr-el-test)
    - [Escribir el código necesario para que el test corra y poder ver el test fallando en su output](#237-escribir-el-c%C3%B3digo-necesario-para-que-el-test-corra-y-poder-ver-el-test-fallando-en-su-output)
    - [Escribir el código necesario para que el test pase](#238-escribir-el-c%C3%B3digo-necesario-para-que-el-test-pase)
    - [Nota al declarar un nuevo error para Update](#239-nota-al-declarar-un-nuevo-error-para-update)
  - [Borrar un item de un map](#24-borrar-un-item-de-un-map)
    - [Escribir el primer test](#241-escribir-el-primer-test)
    - [Correr el test](#242-correr-el-test)
    - [Escribir el código necesario para que el test corra y poder ver el test fallando en su output](#243-escribir-el-c%C3%B3digo-necesario-para-que-el-test-corra-y-poder-ver-el-test-fallando-en-su-output)
    - [Escribir el código necesario para que el test pase](#244-escribir-el-c%C3%B3digo-necesario-para-que-el-test-pase)
- [Referencias](#3-referencias)

# 1. Maps en go

Un `map` es como un array excepto que, en lugar de un índice o `index` entero, **puede tener un `string` o cualquier otro tipo de datos siempre que sea un tipo de datos comparable como clave o `key`**.

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

## 1.7 Comparación de maps

Al igual que el `slice`, un `map` sólo se puede comparar con nulo o `nil`. Si estás pensando en iterar sobre un `map` y hacer coincidir cada elemento, estás en un gran problema. Pero si necesitas urgentemente comparar dos `maps`, utiliza la función `DeepEqual` del paquete [reflect](https://golang.org/pkg/reflect).

## 1.8 Iteración sobre un map

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

# 2. Aprender maps haciendo tests

Intentaremos cubrir con los siguientes ejemplos de tests

- Crear un map
- Buscar un item en un map
- Agregar un item a un map
- Actualizar un item en un map
- Eliminar un item de un map
- Aprender mas acerca de los errores
  - Como crear errores que son constantes
  - Crear wrappers de errores

## 2.1 Buscar un item en un map por su key

Buscaremos una forma de almacenar elementos mediante una `key` y buscarlos rápidamente.

Los `maps` te permiten almacenar elementos de forma similar a un diccionario. Puedes pensar en la `key` como la palabra y el valor como la definición. ¿Y qué mejor manera de aprender sobre Maps que crear nuestro propio diccionario?

### 2.1.1 Escribiendo el primer test

Primero, suponiendo que ya tenemos algunas palabras con sus definiciones en el diccionario, si buscamos una palabra, debería devolvernos la definición de la misma.

```go
package main

import "testing"

func TestSearch(t *testing.T) {
    dictionary := map[string]string{"test": "this is just a test"}

    got := Search(dictionary, "test")
    want := "this is just a test"

    if got != want {
        t.Errorf("got %q want %q given, %q", got, want, "test")
    }
}
```

Declarar un `map` es similar a un `array`. Excepto que comienza con la palabra clave `map` y requiere dos tipos. El primero es el tipo de la `key`, que está escrito dentro de `[]`. El segundo es el tipo del `value` correspondiente para esa `key`, que va justo después de `[]`.

El tipo de `key` es especial. **Solo puede ser un tipo comparable** porque sin la capacidad de saber si 2 claves son iguales, no tenemos forma de asegurarnos de que estamos obteniendo el valor correcto. Los tipos comparables se explican en profundidad en la [language spec](https://golang.org/ref/spec#Comparison_operators)

El tipo de valor, por otro lado, puede ser del tipo que desee. Incluso puede ser otro `map`.

### 2.1.2 Correr el test

Intentando ejecutar `go test` el compilador fallara `./prog_test.go:8:9: undefined: Search`.

[Ejemplo en vivo](https://go.dev/play/p/819gntjPNCS)

### 2.1.3 Escribir el minimo codigo para correr el test y ver su output

Tendremos que implementar la funcion `Search` para que el test pase.

```go
package main

func Search(dictionary map[string]string, word string) string {
    return ""
}
```

Esta funcion simplemente devuelve una cadena vacía. Ahora, si ejecutamos `go test`, debería lanzarnos el error que hemos definido cuando el valor devuelto por `Search` no es el esperado.

```text
got '' want 'this is just a test' given, 'test'.
```

### 2.1.4 Escribir el codigo para que el test pase

```go
func Search(dictionary map[string]string, word string) string {
    return dictionary[word]
}
```

Obtener un valor de un `map` es lo mismo que obtener un valor de un array de `map[key]`.

### 2.1.5 Refactor

### 2.1.6 Crear un helper para el test

```go
func TestSearch(t *testing.T) {
    dictionary := map[string]string{"test": "this is just a test"}

    got := Search(dictionary, "test")
    want := "this is just a test"

    assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

### 2.1.7 Usar un tipo personalizado para el diccionario

```diff
+ type Dictionary map[string]string

+ func (d Dictionary) Search(word string) string {
+        return d[word]
+ }
```

```go
package main

import "testing"

type Dictionary map[string]string

func (d Dictionary) Search(word string) string {
    return d[word]
}

func TestSearch(t *testing.T) {
    dictionary := Dictionary{"test": "this is just a test"}

    got := dictionary.Search("test")
    want := "this is just a test"

    assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

```text
=== RUN   TestSearch
--- PASS: TestSearch (0.00s)
PASS
```

Creamos un tipo de `Dictionary` que actúa como una wrapper alrededor del tipo `map` personalizando nuestro caso de uso. Con el tipo personalizado definido, podemos pasarlo como un `receiver` de la funcion `Search`, permitiendonos este diseno ejecutar `dictionary := Dictionary{"test": "this is just a test"}`.

### 2.1.8 Escribir un test para el caso de que la `word` no este en el `dictionary`

La búsqueda básica fue muy fácil de implementar, pero ¿qué pasará si proporcionamos un valor  `string` que no está en nuestro diccionario?

En realidad no recibimos nada a cambio. Esto es bueno porque el programa puede seguir ejecutándose, aunque es un error silencioso,  pero existe un enfoque mejor. La función podria informar que `word` no está en `dictionary`. De esta manera, el usuario no se pregunta si `word` no existe o si simplemente no hay una definición (esto puede no parecer muy útil para un `dictionary`, sin embargo, es un escenario que podría ser `key` en otros casos de uso).

Asi que escribamos nuestros dos casos de uso,

1. La función Search encuentra la `word` asociada a la `key`
2. La función Search no encuentra la `word` asociada a la `key`

```go
func TestSearch(t *testing.T) {
    dictionary := Dictionary{"test": "this is just a test"}

    t.Run("known word", func(t *testing.T) {
        got, _ := dictionary.Search("test")
        want := "this is just a test"

        assertStrings(t, got, want)
    })

    t.Run("unknown word", func(t *testing.T) {
        _, err := dictionary.Search("unknown")
        want := "could not find the word you were looking for"

        if err == nil {
            t.Fatal("expected to get an error.")
        }

        assertStrings(t, err.Error(), want)
    })
}
```

Pero como vemos, ahora nuestra función `Search` nos deberia devolver un error si no encuentra la `key` en el `map`, es decir, la firma de su return deberia ser `(string, error)`, devolviendonos un tipo `Error` en el caso  de no encontrar la `word` asociada a la `key` en el `dictionary`.

La forma de manejar este escenario en Go es devolver un segundo argumento que sea de tipo `Error`.

Tenga en cuenta que, para lanzar el mensaje de error, primero verificamos que el error no sea `nil` nulo y luego usamos el método `.Error()` para obtener el `string` que luego podemos pasar a la `assertion`.

### 2.1.9 Escribir el codigo para que el test corra y poder ver su output

```diff
func (d Dictionary) Search(word string) (string, error) {
-    return d[word]
+    return d[word], nil
}
```

```go
package main

import "testing"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    return d[word], nil
}

func TestSearch(t *testing.T) {
    dictionary := Dictionary{"test": "this is just a test"}

    t.Run("known word", func(t *testing.T) {
        got, _ := dictionary.Search("test")
        want := "this is just a test"

        assertStrings(t, got, want)
    })

    t.Run("unknown word", func(t *testing.T) {
        _, err := dictionary.Search("unknown")
        want := "could not find the word you were looking for"

        if err == nil {
            t.Fatal("expected to get an error.")
        }

        assertStrings(t, err.Error(), want)
    })
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

El test debería ahora fallar con un mensaje de error mucho mas claro.

```text
dictionary_test.go:22: expected to get an error.
```

### 2.1.10 Escribir el codigo necesario para que el test pase

```diff
- import "errors"
+ import (
+        "errors"
+         "testing"
+    )

func (d Dictionary) Search(word string) (string, error) {
-    return d[word], nil
+    definition, ok := d[word]
+    if !ok {
+        return "", errors.New("could not find the word you were looking for")
+    }

+    return definition, nil
}
```

```go
package main

import (
    "errors"
    "testing"
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", errors.New("could not find the word you were looking for")
    }

    return definition, nil
}

func TestSearch(t *testing.T) {
    dictionary := Dictionary{"test": "this is just a test"}

    t.Run("known word", func(t *testing.T) {
        got, _ := dictionary.Search("test")
        want := "this is just a test"

        assertStrings(t, got, want)
    })

    t.Run("unknown word", func(t *testing.T) {
        _, err := dictionary.Search("unknown")
        want := "could not find the word you were looking for"

        if err == nil {
            t.Fatal("expected to get an error.")
        }

        assertStrings(t, err.Error(), want)
    })
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

```text
=== RUN   TestSearch
=== RUN   TestSearch/known_word
=== RUN   TestSearch/unknown_word
--- PASS: TestSearch (0.00s)
    --- PASS: TestSearch/known_word (0.00s)
    --- PASS: TestSearch/unknown_word (0.00s)
PASS
```

[Ejemplo en vivo](https://go.dev/play/p/0-JllTc_JVF)

Para hacer que los test pasen, utilizamos una propiedad interesante de la búsqueda en el `map` mencionada anteriormente. El `map` puede devolver 2 valores. El segundo valor es un `boolean` que indica si la clave se encontró correctamente.

Esta propiedad nos permite diferenciar entre una `word` que no existe y una palabra que simplemente no tiene definición en el `dictionary`.

### 2.1.11 Refactor

Podemos deshacernos del magic error en nuestra función `Search` extrayéndolo en una variable. Esto también nos permitirá tener un mejor test.

```diff
+ var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
-        return "", errors.New("could not find the word you were looking for")
+        return "", ErrNotFound
    }

    return definition, nil
}
```

Añadimos una nueva función assertError

```diff
+ func assertError(t testing.TB, got, want error) {
+        t.Helper()

+        if got != want {
+            t.Errorf("got error %q want %q", got, want)
+     }
+ }
```

Al crear un nuevo helper `assertError`, podemos simplificar nuestra test y comenzar a usar nuestra variable `ErrNotFound` para que nuestra prueba no falle si cambiamos el texto de error en el futuro.

```diff
t.Run("unknown word", func(t *testing.T) {
-    _, err := dictionary.Search("unknown")
+    _, got := dictionary.Search("unknown")

-    want := "could not find the word you were looking for"

-    if err == nil {
-        t.Fatal("expected to get an error.")
-    }

-    assertStrings(t, err.Error(), want)

    assertError(t, got, ErrNotFound)
})
```

Quedando el ejemplo despues de ingresar los cambios del refactor

```go
package main

import (
    "errors"
    "testing"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}

func TestSearch(t *testing.T) {
    dictionary := Dictionary{"test": "this is just a test"}

    t.Run("known word", func(t *testing.T) {
        got, _ := dictionary.Search("test")
        want := "this is just a test"

        assertStrings(t, got, want)
    })

    t.Run("unknown word", func(t *testing.T) {
        _, got := dictionary.Search("unknown")

        assertError(t, got, ErrNotFound)
    })
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func assertError(t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error %q want %q", got, want)
    }
}
```

```text
=== RUN   TestSearch
=== RUN   TestSearch/known_word
=== RUN   TestSearch/unknown_word
--- PASS: TestSearch (0.00s)
    --- PASS: TestSearch/known_word (0.00s)
    --- PASS: TestSearch/unknown_word (0.00s)
PASS
```

[Ejemplo en vivo](https://go.dev/play/p/R-G7_PngXI5)

## 2.2 Agregar un item a un map

Hemos visto una excelente forma de buscar en el `dictionary`. Sin embargo, no tenemos forma de agregar nuevas `words` a nuestro `dictionary`, hagamoslo con tests.

### 2.2.1 Escribir el primer test para añadir un item a un `map`

```go
package main

import "testing"

type Dictionary map[string]string

func TestAdd(t *testing.T) {
    dictionary := Dictionary{}
    dictionary.Add("test", "this is just a test")

    want := "this is just a test"
    got, err := dictionary.Search("test")
    if err != nil {
        t.Fatal("should find added word:", err)
    }

    assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

En esta prueba, utilizamos nuestra función de `Search` para facilitar un poco la validación del `dictionary`.

### 2.2.2 Escribir el código necesario para que el test corra y poder ver su output

```diff
+ func (d Dictionary) Add(word, definition string) {
+ }
```

```go
package main

import "testing"

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) {
}

func TestAdd(t *testing.T) {
    dictionary := Dictionary{}
    dictionary.Add("test", "this is just a test")

    want := "this is just a test"
    got, err := dictionary.Search("test")
    if err != nil {
        t.Fatal("should find added word:", err)
    }

    assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

Ahora el test debería correr y fallar con el siguiente mensaje

```text
dictionary_test.go:31: should find added word: could not find the word you were looking for
```

### 2.2.3 Escribir el código necesario para que el test pase

```diff
func (d Dictionary) Add(word, definition string) {
+    d[word] = definition
}
```

```go
package main

import "testing"

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) {
    d[word] = definition
}

func TestAdd(t *testing.T) {
    dictionary := Dictionary{}
    dictionary.Add("test", "this is just a test")

    want := "this is just a test"
    got, err := dictionary.Search("test")
    if err != nil {
        t.Fatal("should find added word:", err)
    }

    assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

Añadir a un map es también similar a un array. Solo necesitas especificar una `key` y asignarle un valor.

### 2.2.4 Punteros copias etc

Una propiedad interesante de los `map` es que puedes modificarlos sin pasarles una dirección (por ejemplo, `&myMap`).

Entonces, cuando pasas un `map` a una función/método, de hecho lo estás copiando, pero solo la parte del puntero, no la estructura de datos subyacente que contiene los datos.

Un problema con los `map` es que pueden tener un valor nulo `nil`. Un mapa nulo se comporta como un mapa vacío cuando se lee, pero **intentar escribir en un mapa `nil` provocarán un `panic` en tiempo de ejecución**.

Por lo tanto, no es recomendable inicializar un `map` vacío:

```go
var m map[string]string
```

En su lugar, puedes inicializar un mapa vacío, o usar la palabra clave `make`:

```go
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```

Ambas formas son correctas, crean un empty `hash map` que apunta a `dictionary`. Lo que asegura que nunca se produzca un `panic` en tiempo de ejecución.

### 2.2.5 Refactor

No hay mucho que refactorizar en nuestra implementación, pero el test podría necesitar un poco de simplificación.

```diff
func TestAdd(t *testing.T) {
    dictionary := Dictionary{}
-    dictionary.Add("test", "this is just a test")
+    word := "test"
+    definition := "this is just a test"
+
    dictionary.Add(word, definition)

-    want := "this is just a test"
-    got, err := dictionary.Search("test")
-    if err != nil {
-        t.Fatal("should find added word:", err)
-    }

-    assertStrings(t, got, want)

    assertDefinition(t, dictionary, word, definition)
}
+
+ func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
+        t.Helper()
+
+        got, err := dictionary.Search(word)
+        if err != nil {
+            t.Fatal("should find added word:", err)
+        }
+
+        assertStrings(t, got, definition)
+}
```

```go
package main

import (
    "errors"
    "testing"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}

func (d Dictionary) Add(word, definition string) {
    d[word] = definition
}

func TestAdd(t *testing.T) {
    dictionary := Dictionary{}
    word := "test"
    definition := "this is just a test"

    dictionary.Add(word, definition)

    assertDefinition(t, dictionary, word, definition)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }
    assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

```text
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
PASS
```

[Ejemplo en vivo](https://go.dev/play/p/VbXgdhfsUYA)

Creamos variables para `word` y `definition`, y movimos la `assertion` de `definition` a su propia función auxiliar.

Nuestro `Add` se ve bien. ¡Excepto que no hemos tenido en cuenta el caso de lo que sucede cuando el valor que intentamos agregar ya existe!

`map` no arrojará un error si el valor ya existe. En su lugar, seguirán adelante y sobrescribirán el valor con el valor recién proporcionado. Esto puede ser conveniente en la práctica, pero hace que el nombre de nuestra función sea menos preciso. `Add` no deberia modificar los valores existentes. Sólo debería agregar nuevas `word` a nuestro `dictionary`.

### 2.2.6 Escribir un test para el caso de que la `word` ya exista

```diff
func TestAdd(t *testing.T) {
-    dictionary := Dictionary{}
-    word := "test"
-    definition := "this is just a test"
-
-    dictionary.Add(word, definition)
-
-    assertDefinition(t, dictionary, word, definition)
+
+    t.Run("new word", func(t *testing.T) {
+        dictionary := Dictionary{}
+        word := "test"
+        definition := "this is just a test"
+
+        err := dictionary.Add(word, definition)
+
+        assertError(t, err, nil)
+        assertDefinition(t, dictionary, word, definition)
+    })
+
+    t.Run("existing word", func(t *testing.T) {
+        word := "test"
+        definition := "this is just a test"
+        dictionary := Dictionary{word: definition}
+        err := dictionary.Add(word, "new test")
+
+        assertError(t, err, ErrWordExists)
+        assertDefinition(t, dictionary, word, definition)
+    })
+
+ func assertError(t testing.TB, got, want error) {
+     t.Helper()
+
+     if got != want {
+            t.Errorf("got error %q want %q", got, want)
+        }
+ }
}
```

```go
package main

import (
    "errors"
    "testing"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}

func (d Dictionary) Add(word, definition string) {
    d[word] = definition
}

func TestAdd(t *testing.T) {
    t.Run("new word", func(t *testing.T) {
        dictionary := Dictionary{}
        word := "test"
        definition := "this is just a test"

        err := dictionary.Add(word, definition)

        assertError(t, err, nil)
        assertDefinition(t, dictionary, word, definition)
    })

    t.Run("existing word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{word: definition}
        err := dictionary.Add(word, "new test")

        assertError(t, err, ErrWordExists)
        assertDefinition(t, dictionary, word, definition)
    })
}

func assertError(t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error %q want %q", got, want)
    }
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }
    assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

Para este test, modificamos `Add` para que devuelva un error, que estamos validando con una nueva variable de error, `ErrWordExists`. También modificamos el test anterior para comprobar si hay un error `nil`.

### 2.2.7 Correr el test

```text
./prog_test.go:34:10: dictionary.Add(word, definition) (no value) used as value
./prog_test.go:44:10: dictionary.Add(word, "new test") (no value) used as value
```

[Ejemplo en vivo](https://go.dev/play/p/iu9L7Tt4ajL)

### 2.2.8 Escribir el código necesario para que el test pase

```diff
- var ErrNotFound = errors.New("could not find the word you were looking for")

+ var (
+        ErrNotFound   = errors.New("could not find the word you were looking for")
+        ErrWordExists = errors.New("cannot add word because it already exists")
+ )

func (d Dictionary) Add(word, definition string) error {
    d[word] = definition
    return nil
}
```

ahora corriendo el test

```text
=== RUN   TestAdd
=== RUN   TestAdd/new_word
=== RUN   TestAdd/existing_word
    prog_test.go:47: got error %!q(<nil>) want "cannot add word because it already exists"
    prog_test.go:48: got "new test" want "this is just a test"
--- FAIL: TestAdd (0.00s)
    --- PASS: TestAdd/new_word (0.00s)
    --- FAIL: TestAdd/existing_word (0.00s)
FAIL
```

[Ejemplo en vivo](https://go.dev/play/p/4dfuvC1HIJs)

Ahora tenemos dos errores más. Todavía estamos modificando el valor y devolviendo un error `nill`.

### 2.2.8 Escribir el código necesario para que el test pase

```diff
func (d Dictionary) Add(word, definition string) error {
-    d[word] = definition
+    _, err := d.Search(word)
+
+    switch err {
+    case ErrNotFound:
+        d[word] = definition
+    case nil:
+        return ErrWordExists
+    default:
+        return err
+    }
+
    return nil
}
```

Aquí estamos usando un `switch` para hacer coincidir el error. Tener un `switch` como este proporciona una red de seguridad adicional, en caso de que `Search` devuelva un error distinto de `ErrNotFound`.

```go
package main

import (
    "errors"
    "testing"
)

var (
    ErrNotFound   = errors.New("could not find the word you were looking for")
    ErrWordExists = errors.New("cannot add word because it already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
    _, err := d.Search(word)

    switch err {
    case ErrNotFound:
        d[word] = definition
    case nil:
        return ErrWordExists
    default:
        return err
    }

    return nil
}

func TestAdd(t *testing.T) {
    t.Run("new word", func(t *testing.T) {
        dictionary := Dictionary{}
        word := "test"
        definition := "this is just a test"

        err := dictionary.Add(word, definition)

        assertError(t, err, nil)
        assertDefinition(t, dictionary, word, definition)
    })

    t.Run("existing word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{word: definition}
        err := dictionary.Add(word, "new test")

        assertError(t, err, ErrWordExists)
        assertDefinition(t, dictionary, word, definition)
    })
}

func assertError(t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error %q want %q", got, want)
    }
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }
    assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

```test
=== RUN   TestAdd
=== RUN   TestAdd/new_word
=== RUN   TestAdd/existing_word
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/new_word (0.00s)
    --- PASS: TestAdd/existing_word (0.00s)
PASS
```

[Ejemplo en vivo](https://go.dev/play/p/4_EGj8wbYMx)

### 2.2.9 Refactor

No tenemos mucho que refactorizar, pero a medida que nuestro uso de errores crece, podemos hacer algunas modificaciones.

```diff
+ const (
+        ErrNotFound   = DictionaryErr("could not find the word you were looking for")
+        ErrWordExists = DictionaryErr("cannot add word because it already exists")
+ )
+
+ type DictionaryErr string
+
+ func (e DictionaryErr) Error() string {
+        return string(e)
+ }
```

Hicimos que los errores fueran constantes, para esto necesitamos crear nuestro propio tipo `DictionaryErr `que implementa la interfaz de error. Puede leer más sobre los detalles en [este excelente artículo de Dave Cheney](https://dave.cheney.net/2016/04/07/constant-errors). En pocas palabras, hace que los errores sean más reutilizables e inmutables.

```go
package main

import (
    "testing"
)

const (
    ErrNotFound   = DictionaryErr("could not find the word you were looking for")
    ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
    _, err := d.Search(word)

    switch err {
    case ErrNotFound:
        d[word] = definition
    case nil:
        return ErrWordExists
    default:
        return err
    }

    return nil
}

func TestAdd(t *testing.T) {
    t.Run("new word", func(t *testing.T) {
        dictionary := Dictionary{}
        word := "test"
        definition := "this is just a test"

        err := dictionary.Add(word, definition)

        assertError(t, err, nil)
        assertDefinition(t, dictionary, word, definition)
    })

    t.Run("existing word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{word: definition}
        err := dictionary.Add(word, "new test")

        assertError(t, err, ErrWordExists)
        assertDefinition(t, dictionary, word, definition)
    })
}

func assertError(t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error %q want %q", got, want)
    }
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }
    assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

## 2.3 Actualizar el item de un map

### 2.3.1 Escribir el primer test

```diff
+ func TestUpdate(t *testing.T) {
+        word := "test"
+        definition := "this is just a test"
+        dictionary := Dictionary{word: definition}
+        newDefinition := "new definition"
+
+        dictionary.Update(word, newDefinition)
+
+        assertDefinition(t, dictionary, word, newDefinition)
+ }
```

```go
package main

import (
    "testing"
)

const (
    ErrNotFound   = DictionaryErr("could not find the word you were looking for")
    ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}

func TestUpdate(t *testing.T) {
    word := "test"
    definition := "this is just a test"
    dictionary := Dictionary{word: definition}
    newDefinition := "new definition"

    dictionary.Update(word, newDefinition)

    assertDefinition(t, dictionary, word, newDefinition)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }
    assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

`Update` esta muy relacionado con `Add`.

### 2.3.2 Correr el test

```text
./prog_test.go:35:13: dictionary.Update undefined (type Dictionary has no field or method Update)
```

[Ejemplo en vivo](https://go.dev/play/p/UgVzlz1Njr8)

### 2.3.3 Escribir el código necesario para que el test corra y poder ver su output

Ya sabemos cómo lidiar con un error de este tipo. Necesitamos definir nuestra función `Update` y pasarle el `dictionary` como `receiver`.

```go
func (d Dictionary) Update(word, definition string) {}
```

Una vez implementado esto, podemos ver que necesitamos cambiar la definición de `word`.

```text
=== RUN   TestUpdate
    prog_test.go:39: got "this is just a test" want "new definition"
--- FAIL: TestUpdate (0.00s)
FAIL
```

[Ejemplo en vivo](https://go.dev/play/p/Za8RJS2CXuj)

### 2.3.4 Escribir el código necesario para que el test pase

Ya vimos cómo resolver este problema cuando solucionamos el problema con `Add`. Entonces, implementemos algo realmente similar a `Add`.

```diff
func (d Dictionary) Update(word, definition string) {
+ d[word] = definition
}
```

No es necesario refactorizar esto, es cambio simple. Sin embargo, ahora tenemos el mismo problema que con `Add`. Si pasamos una palabra nueva, `Update` la agregará al diccionario.

### 2.3.4 Escribiendo un test para el caso de que la `word` que queremos actualizar sea nueva en el diccionario

```diff
+ t.Run("existing word", func(t *testing.T) {
+        word := "test"
+        definition := "this is just a test"
+        dictionary := Dictionary{word: definition}
+        newDefinition := "new definition"
+
+     err := dictionary.Update(word, newDefinition)
+
+        assertError(t, err, nil)
+        assertDefinition(t, dictionary, word, newDefinition)
+ })
+
+ t.Run("new word", func(t *testing.T) {
+     word := "test"
+        definition := "this is just a test"
+        dictionary := Dictionary{}
+
+        err := dictionary.Update(word, definition)
+
+        assertError(t, err, ErrWordDoesNotExist)
+ })
```

```go
package main

import (
    "testing"
)

const (
    ErrNotFound   = DictionaryErr("could not find the word you were looking for")
    ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}

func (d Dictionary) Update(word, definition string) {}

func TestUpdate(t *testing.T) {
    t.Run("existing word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{word: definition}
        newDefinition := "new definition"

        err := dictionary.Update(word, newDefinition)

        assertError(t, err, nil)
        assertDefinition(t, dictionary, word, newDefinition)
    })

    t.Run("new word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{}

        err := dictionary.Update(word, definition)

        assertError(t, err, ErrWordDoesNotExist)
    })
}

func assertError(t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error %q want %q", got, want)
    }
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)ings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

### 2.3.5 Correr el test

```text
./prog_test.go:38:10: dictionary.Update(word, newDefinition) (no value) used as value
./prog_test.go:49:10: dictionary.Update(word, definition) (no value) used as value
./prog_test.go:51:23: undefined: ErrWordDoesNotExist
```

Agregamos otro tipo de error más para cuando la `word` no existe en `dictionary`. También modificamos `Update` para devolver un valor de error.

Obtenemos 3 errores pero ya sabemos como resolverlos.

[Ejemplo en vivo](https://go.dev/play/p/zGxEgpTaKHY)

### 2.3.6 Escribir el código necesario para que el test corra y poder ver el test fallando en su output

```diff
const (
    ErrNotFound         = DictionaryErr("could not find the word you were looking for")
    ErrWordExists       = DictionaryErr("cannot add word because it already exists")
+ ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

+ func (d Dictionary) Update(word, definition string) error {
        d[word] = definition
+ return nil
}
```

Agregamos nuestro propio tipo de error y devolvemos un error `nill`.

Con estos cambios ahora deberíamos ver un error más claro.

```go
package main

import (
    "testing"
)

const (
    ErrNotFound         = DictionaryErr("could not find the word you were looking for")
    ErrWordExists       = DictionaryErr("cannot add word because it already exists")
    ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}

func (d Dictionary) Update(word, definition string) error {
    d[word] = definition
    return nil
}

func TestUpdate(t *testing.T) {
    t.Run("existing word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{word: definition}
        newDefinition := "new definition"

        err := dictionary.Update(word, newDefinition)

        assertError(t, err, nil)
        assertDefinition(t, dictionary, word, newDefinition)
    })

    t.Run("new word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{}

        err := dictionary.Update(word, definition)

        assertError(t, err, ErrWordDoesNotExist)
    })
}

func assertError(t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error %q want %q", got, want)
    }
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }
    assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

```text
=== RUN   TestUpdate
=== RUN   TestUpdate/existing_word
=== RUN   TestUpdate/new_word
    prog_test.go:55: got error %!q(<nil>) want "cannot update word because it does not exist"
--- FAIL: TestUpdate (0.00s)
    --- PASS: TestUpdate/existing_word (0.00s)
    --- FAIL: TestUpdate/new_word (0.00s)
FAIL
```

[Ejemplo en vivo](https://go.dev/play/p/wzxACQ4q-lW)

### 2.3.7 Escribir el código necesario para que el test pase

```diff
func (d Dictionary) Update(word, definition string) error {
-    d[word] = definition
+    _, err := d.Search(word)
+
+    switch err {
+        case ErrNotFound:
+        return ErrWordDoesNotExist
+    case nil:
+        d[word] = definition
+    default:
+        return err
+    }
+
    return nil
}
```

Esta función parece casi idéntica a `Add` excepto que cambiamos cuando actualizamos el `dictionary` y cuando devolvemos un error.

```go
package main

import (
    "testing"
)

const (
    ErrNotFound         = DictionaryErr("could not find the word you were looking for")
    ErrWordExists       = DictionaryErr("cannot add word because it already exists")
    ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}

func (d Dictionary) Update(word, definition string) error {
    _, err := d.Search(word)

    switch err {
    case ErrNotFound:
        return ErrWordDoesNotExist
    case nil:
        d[word] = definition
    default:
        return err
    }

    return nil
}

func TestUpdate(t *testing.T) {
    t.Run("existing word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{word: definition}
        newDefinition := "new definition"

        err := dictionary.Update(word, newDefinition)

        assertError(t, err, nil)
        assertDefinition(t, dictionary, word, newDefinition)
    })

    t.Run("new word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{}

        err := dictionary.Update(word, definition)

        assertError(t, err, ErrWordDoesNotExist)
    })
}

func assertError(t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error %q want %q", got, want)
    }
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }
    assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

```text
=== RUN   TestUpdate
=== RUN   TestUpdate/existing_word
=== RUN   TestUpdate/new_word
--- PASS: TestUpdate (0.00s)
    --- PASS: TestUpdate/existing_word (0.00s)
    --- PASS: TestUpdate/new_word (0.00s)
PASS
```

[Ejemplo en vivo](https://go.dev/play/p/Y3VQjg5dsO-)

### 2..3.8 Nota al declarar un nuevo error para Update

Podríamos reutilizar `ErrNotFound` y no agregar un nuevo error. Sin embargo, suele ser mejor tener un error preciso para cuando falla una `update`.

Tener errores específicos te brinda más información sobre lo que salió mal. A continuación se muestra un ejemplo en una aplicación web:

> Puede redirigir al usuario cuando se encuentre `ErrNotFound`, pero mostrar un mensaje de error cuando se encuentre `ErrWordDoesNotExist`.

## 2.4 Borrar un item de un map

### 2.4.1 Escribir el primer test

```diff
+ func TestDelete(t *testing.T) {
+        word := "test"
+        dictionary := Dictionary{word: "test definition"}
+
+        dictionary.Delete(word)
+
+        _, err := dictionary.Search(word)
+        if err != ErrNotFound {
+            t.Errorf("Expected %q to be deleted", word)
+        }
+ }
```

Nuestro test creara un `dictionary` con una `word` y luego borrara la `word` y luego chequeara si la `word` ha sido borrada.

```go
package main

import (
    "testing"
)

const (
    ErrNotFound         = DictionaryErr("could not find the word you were looking for")
    ErrWordExists       = DictionaryErr("cannot add word because it already exists")
    ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}

func (d Dictionary) Update(word, definition string) error {
    _, err := d.Search(word)

    switch err {
    case ErrNotFound:
        return ErrWordDoesNotExist
    case nil:
        d[word] = definition
    default:
        return err
    }

    return nil
}

func TestDelete(t *testing.T) {
    word := "test"
    dictionary := Dictionary{word: "test definition"}

    dictionary.Delete(word)

    _, err := dictionary.Search(word)
    if err != ErrNotFound {
        t.Errorf("Expected %q to be deleted", word)
    }
}
```

### 2.4.2 Correr el test

```text
./prog_test.go:49:13: dictionary.Delete undefined (type Dictionary has no field or method Delete)
```

[Ejemplo en vivo](https://go.dev/play/p/AaNMrY2c6sR)

### 2.4.3 Escribir el código necesario para que el test corra y poder ver el test fallando en su output

```diff
+ func (d Dictionary) Delete(word string) {}
```

Después de añadir esto los test deberían fallar con el siguiente mensaje

```text
=== RUN   TestDelete
    prog_test.go:55: Expected "test" to be deleted
--- FAIL: TestDelete (0.00s)
FAIL
```

[Ejemplo en vivo](https://go.dev/play/p/S6z86380JCB)

### 2.4.3 Escribir el código necesario para que el test pase

```diff
func (d Dictionary) Delete(word string) {
+ delete(d, word)
}
```

Go tiene una función de `delete` `built-in` que funciona en `maps`. Se necesitan dos argumentos. El primero es el `map` y el segundo es la `key` que hay que eliminar.

La función `delete` no devuelve nada y basamos nuestro método de eliminación en la misma noción. Dado que eliminar un valor que no existe no tiene ningún efecto, a diferencia de nuestros métodos `Add` y `Update`, no necesitamos complicar la API con errores.

# 3. Referencias

- [Maps type](https://go.dev/ref/spec#Map_types)
- [Making slices maps and channels](https://go.dev/ref/spec#Making_slices_maps_and_channels)
- [The anatomy of maps in go](https://medium.com/rungo/the-anatomy-of-maps-in-go)
- [Golang Maps Tutorial](https://golangbot.com/maps/)
- [Go by example: Maps](https://gobyexample.com/maps)
- [Learn go with tests: Maps](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps)
