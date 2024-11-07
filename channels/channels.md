# 1. Channel

- [Channel](#1-channel)
  - [Que son los channels?](#11-que-son-los-channels)
  - [Declarar un channel](#12-declarar-un-channel)
  - [Leer y escribir datos en un channel](#13-leer-y-escribir-datos-en-un-channel)
  - [Channels en la practica](#14-channels-en-la-practica)
  - [Deadlock](#15-deadlock)
  - [Cerrando un channel](#16-cerrando-un-channel)
    - [Usando un bucle for para leer datos de un channel cerrado](#161-usando-un-bucle-for-para-leer-datos-de-un-channel-cerrado)
   	- [Tamaño del buffer o capacidad del channel](#162-tama%C3%B1o-del-buffer-o-capacidad-del-channel)
   	- [Longitud y capacidad de channel](#163-longitud-y-capacidad-de-channel)
    - [Leer desde un channel cerrado](#164-leer-desde-un-channel-cerrado)
  - [Trabajando con múltiples goroutines](#17-trabajando-con-m%C3%BAltiples-goroutines)
  - [Channels unidireccionales](#18-channels-unidireccionales)
  - [Goroutine anonima](#19-goroutine-anonima)
  - [Channel como tipo de datos de un channel](#110-channel-como-tipo-de-datos-de-un-channel)
  - [Select](#111-select)
    - [Select con default, default case](#1111-select-con-default-default-case)
    - [Deadlock](#1112-deadlock)
    - [Nil channel](#1113-nil-channel)
    - [Añadiendo un timeout](#1114-a%C3%B1adiendo-un-timeout)
    - [Select vacío](#1115-select-vac%C3%ADo)
    - [WaitGroup](#1116-waitgroup)
    - [Worker pool](#1117-worker-pool)

## 1.1 Que son los `channels`?

Un canal es un objeto de comunicación mediante el cual las [goroutines](../goroutines/goroutines.md) pueden comunicarse entre sí. Técnicamente, un canal es una tubería de transferencia de datos donde se pueden pasar o leer datos. Por lo tanto, una `goroutine` puede enviar datos a un canal, mientras que otras `goroutines` pueden leer esos datos desde el mismo canal.

## 1.2 Declarar un `channel`

Go proporciona la palabra clave `chan` para crear un `channel`.

```go
var c chan int
```

Un `channel` puede transportar datos de un solo tipo de datos. No se permite transportar ningún otro tipo de datos desde ese `channel`.

```go
package main

import "fmt"

func main() {
	var c chan int
	fmt.Print(c)
}
```

```text
<nil>
```

[Ejemplo en vivo](https://go.dev/play/p/iWOFLfcgfF-)

El programa anterior declaramos un `channel` `c` que puede transportar datos de tipo `int`. El programa anterior se imprime porque el `zero value` de un `channel` es nulo. Pero un `channel` nulo no es útil. No podemos pasar datos ni leer datos de un `channel` que sea nulo. Por lo tanto, tenemos que usar la función `make` para crear un `channel` listo para usar.

```go
package main

import "fmt"

func main() {
	c := make(chan int)

	fmt.Printf("type of `c` is %T\n", c)
	fmt.Printf("value of `c` is %v\n", c)
}
```

Hemos utilizado la sintaxis abreviada `:=` para crear un `channel` usando la función `make`. El programa anterior produce el siguiente resultado.

```text
type of `c` is chan int
value of `c` is 0xc000076060
```

[Ejemplo en vivo](https://go.dev/play/p/N4dU7Ql9bK7)

Veamos que el valor del `channel` `c` parece que es una dirección de memoria. **Los `channels` por defecto son `pointers` punteros**. Principalmente, cuando desea comunicarse con una `goroutine`, pasa el `channel` como argumento a la función o método. Por lo tanto, cuando la `goroutine` recibe ese `channel` como argumento, no es necesario eliminar la referencia para enviar o extraer datos de ese canal.

## 1.3 Leer y escribir datos en un `channel`

Go nos proporciona una sintaxis muy fácil de recordar `left arrow syntax` `<-` para leer y escribir datos de un `channel`.

```go
c <- data
```

La sintaxis anterior significa que queremos enviar o escribir datos en el `channel` `c`. Nota la dirección de la flecha, apunta desde datos al `channel` `c`. Por lo tanto, podemos imaginar que estamos intentando enviar datos a `c`.

```go
<- c
``

La sintaxis anterior significa que queremos leer algunos datos del channel `c`. Nota la dirección de la flecha, comienza en el canal `c`. Esta declaración no incluye datos, pero aun así es una declaración válida. con una variable podriamos almacenar los datos provenientes del `channel`:

```go
var data int
data = <- c
```

Ahora los datos que leamos del `channel` `c` que es de tipo `int `se pueden almacenar en la variable `data` de tipo `int`.

Podemos reescribir esta sintaxis en una sola línea.

```go
data := <- c
```

Go descubrirá el tipo de datos que se transportan en el `channel` `c` y les dará un tipo de datos válido.

**Todas las operaciones de `channel` anteriores están bloqueadas by default**. [En la sección de goroutines](../goroutines/goroutines.md), vimos `time.Sleep` para lograr el bloqueando de una `goroutine`. Las operaciones en el `channel` de lectura/escritura son de naturaleza bloqueante. Cuando se escriben datos en el `channel`, la `goroutine` se bloquea hasta que otra rutina los lea desde ese `channel`. Al mismo tiempo, como vimos en el capítulo de concurrencia, las `channel operations` le dicen al `scheduler` que `shedule` otra `goroutine`, es por eso que un programa no se bloquea para siempre en la misma `goroutine`. Estas características de un `channel` son muy útiles en la comunicación de `goroutines`, ya que nos impiden tener que escribir bloqueos con trucos para que su comunicacion funcione entre sí.


## 1.4 `Channels` en la practica

Hablemos con una `goroutine` sobre su código:

```go
package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "John"
	fmt.Println("main() stopped")
}
```

```text
main() started
Hello John!
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/OeYLKEz7qKi)

Hablemos de la ejecución en el programa anterior paso a paso.

1. Declaramos una función de `greet` que acepta un `channel` `c` de tipo de datos de transporte `string`. En esa función, estamos leyendo datos del `channel` `c` imprimiendo esos datos a la consola.
2. En la función `main`, el programa imprime main iniciado en la consola ya que es la primera declaración.
3. Luego creamos el `channel` `c` de tipo `string` usando la función `make`.
4. Pasamos el `channel` `c` a la función `greet` pero lo ejecutamos como una rutina usando la palabra clave go.
5. En este punto, el proceso tiene 2 `goroutines` mientras que la `goroutine` activa es la `goroutine` `main` (consulte la lección [goroutines](../goroutines/goroutines.md) para saber cuál es). Despues de este punto  el control pasa a la siguiente línea.
6. Enviamos un valor `string` `John` al `channel` `c`. En este punto, la `goroutine` se bloquea hasta que alguna `goroutine` la lea. Go `scheduler` schedule la `greet` `goroutine` y su ejecución comienza.
7. Después de esa ejecución, la `main goroutine` se activa y ejecuta la declaración final, imprimiendo `main() stopped`.

## 1.5 Deadlock

Como mencionamos, cuando escribimos o leemos datos de un `channel`, esa `goroutine` se bloquea y el control se pasa a las `goroutines` disponibles. ¿Qué pasa si no hay otras rutinas disponibles? Imagine que todas están durmiendo. Ahí es donde se produce un error de interbloqueo `deadblok` que bloquea todo el programa.

> [!NOTE]
> Si está intentando leer datos de un `channel` pero el `channel` no tiene un valor disponible, se bloquea la `goroutine` actual y se desbloquea otra con la esperanza de que alguna `goroutine` envíe un valor al `channel`. Por lo tanto, esta operación de lectura se bloqueará. De manera similar, si va a enviar datos a un `channel`, bloqueará la rutina actual y desbloqueará otras hasta que alguna `goroutine` lea los datos. Por lo tanto, esta operación de envío se bloqueará.

Un ejemplo simple de `deadlock` sería que solo la rutina principal realizara alguna `channel operation` de lectura escritura.

```go
package main

import "fmt"

func main() {
	fmt.Println("main() started")

	c := make(chan string)
	c <- "John"

	fmt.Println("main() stopped")
}
```

```go
main() started
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/tmp/sandbox2794472245/prog.go:9 +0x6a
exit status 2
```

fatal error: all goroutines are asleep — deadlock!. Para que todas las `goroutines` están dormidas o no hay otras `goroutines` disponibles to `schedule`.

## 1.6 Cerrando un `channel`

Se puede cerrar un `channel` para que no se puedan enviar más datos a través de él. La `goroutine` del receptor puede averiguar el estado del `channel` usando `val, ok := <- channel` sintaxis donde `ok` es `true` si el `channel` está abierto o se pueden realizar operaciones de lectura y `false` si el `channel` está cerrado y no se pueden realizar más operaciones de lectura. Un `channel` se puede cerrar usando la función integrada close con sintaxis `close(channel)`. Veamos un ejemplo sencillo.

```go
package main

import "fmt"

func greet(c chan string) {
	<-c // for John
	<-c // for Mike
}

func main() {
	fmt.Println("main() started")

	c := make(chan string, 1)

	go greet(c)
	c <- "John"

	close(c) // closing channel

	c <- "Mike"
	fmt.Println("main() stopped")
}
```

```text
main() started
panic: send on closed channel

goroutine 1 [running]:
main.main()
	/tmp/sandbox2641477502/prog.go:20 +0xd6
```

[Ejemplo en vivo](https://go.dev/play/p/LMmAq4sgm02)

> [!NOTE]
> Este ejemplo nos ayudara a ilústrate el concepto de bloqueo,
>
> 1. Primero enviemos la operación `c <- "John"` que manda el `stream` `Jhon` al `channel` `c` esto bloqueara la `goroutine` `main` y alguna `goroutine` tiene que leer datos del `channel`, por lo tanto, Go `scheduler` agenda la `goroutine` `greet`.
> 2. Luego, tras la primera operación de lectura `<-c` en la `goroutine` `greet` no es bloqueante porque los datos están presentes en el canal `c` para ser leídos.
> 3. La segunda operación de lectura `<-c` sera bloqueante para la el `channel` `c` porque el `channel` `c` no tiene ningún dato para leer.
> 4. Por lo tanto, Go `scheduler` activa la `goroutine` `main` y el programa comienza la ejecución desde la función `close(c)`.

Del error anterior, podemos ver que estamos intentando enviar datos en un `channel` cerrado. Para comprender mejor la usabilidad de los `channels` cerrados, veamos el bucle `for`.

### 1.6.1 Usando un bucle `for` para leer datos de un `channel` cerrado

Se puede utilizar una sintaxis infinita para el bucle `for{}` para leer múltiples valores enviados a través de un `channel`.

```go
package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c) // close channel
}

func main() {
	fmt.Println("main() started")
	c := make(chan int)

	go squares(c) // start goroutine

	// periodic block/unblock of main goroutine until chanel closes
	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "<-- loop broke!")
			break // exit break loop
		} else {
			fmt.Println(val, ok)
		}
	}

	fmt.Println("main() stopped")
}
```

En el ejemplo anterior, estamos mandando el cuadrado `squares` de números `int` del `0` al `9` uno por uno al `channel` `c` en la `goroutine` function `squares`. En la `goroutine` `main`, estamos leyendo esos números dentro de un bucle `for` infinito.

En bucle el `for` infinito de la `goroutine` `main` que va a estar encantado en recibir los valores de los cuadrados de los numeros del `0` al `9`, como necesitamos una condición para romper el bucle en algún punto, estamos leyendo el valor del canal con sintaxis `val, ok := <-c`. Aquí, `ok` nos dará información adicional cuando el `channel` esté cerrado. Por lo tanto, en la `goroutine` de `squares`, después de escribir todos los datos, cerramos el canal usando la sintaxis `close(c)`. Cuando `ok` es `true`, el programa imprime el valor con `val` y el estado del `channel` con `ok`. Cuando es `false`, salimos del ciclo usando la palabra clave `break`. Por lo tanto, el programa anterior produce el siguiente resultado.

```text
main() started
0 true
1 true
4 true
9 true
16 true
25 true
36 true
49 true
64 true
81 true
0 false <-- loop broke!
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/X58FTgSHhXi)

> [!NOTE]
> Cuando el `channel` está cerrado, el valor leído por la `goroutine` es el valor cero del tipo de datos del `channel`. En este caso, dado que el `channel` transporta el tipo de datos `int`, será `0` como podemos ver en el resultado. Cerrar el `channel` no bloquea la `goroutine` actual a diferencia de leer o escribir un valor en el canal.

Para evitar la molestia de verificar manualmente la condición de cierre del `channel`, Go proporciona el bucle `for range` más sencillo que se cerrará automáticamente cuando el `channel` esté cerrado. Modifiquemos nuestro programa anterior anterior con esta opcion.

```go
package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c) // close channel
}

func main() {
	fmt.Println("main() started")
	c := make(chan int)

	go squares(c) // start goroutine

	// periodic block/unblock of main goroutine until chanel closes
	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("main() stopped")
}
```

En este ejemplo usamos `for range` en vez de un bucle `for{}` infinito. El bucle `for range` leera el valor del `channel` `c`, uno por cada pasada hasta que, al quedarse sin valores en la pasada final, lo cierra. El programa anterior produce el siguiente resultado.

```text
main() started
0
1
4
9
16
25
36
49
64
81
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/ICCYbWO7ZvD)

> Si no cierras el `chanel` en el bucle `for range`, el programa generará un error fatal de `deadlock` en tiempo de ejecución.

### 1.6.2 Tamaño del `buffer` o capacidad del `channel`

Como vimos, cada operación de envío al `channel` bloquea la `goroutine` actual. Pero es que hasta ahora usamos la función `make` sin el segundo parámetro. Este segundo parámetro es la `capacity` de un `channel` o el `size` del `buffer` de este `channel`. **De forma predeterminada, el tamaño del `buffer` de un `channel` es `0`**, también denominado `unbuffered channel`. En este caso, **todo lo escrito en el `channel` está inmediatamente disponible para leer**.

Pero cuando el tamaño del `buffer` no es cero, la `goroutine` no se bloquea hasta que el `buffer` esté lleno. Cuando el `buffer` está lleno, cualquier valor enviado al `channel` se agrega al `buffer` descartando el último valor en el `buffer` que está disponible para leer (donde se bloqueará la `goroutine`). Pero hay un problema: **la operación de lectura en el buffer tiene sed**. Eso significa que **una vez que comienza la operación de lectura, continuará hasta que el búfer esté vacío**. Técnicamente, eso significa que la **`goroutine` que lee el `channel` del `buffer` no se bloqueará hasta que el `buffer` esté vacío**.

Podemos definir un `channel`  con `buffer` con la siguiente sintaxis.

```go
c := make(chan Type, n)
```

Esto crea un `channel` de un tipo de datos `Type` con tamaño de `buffer` `n`. Hasta que el canal reciba `n+1` operaciones de envío, no bloqueará la `goroutine` actual.

Demostremos que `goroutine` no se bloquea hasta que el búfer está lleno y se desborda por `overflow`.

```go
package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)

	go squares(c)

	c <- 1
	c <- 2
	c <- 3

	fmt.Println("main() stopped")
}
```

```text
main() started
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/k0usdYZfp3D)

En el programa anterior, el `channel` `c` de elementos enteros `int` tiene una capacidad de `buffer` de `3`. Eso significa que puede contener 3 valores, lo cual ocurre en la línea no. 20,

```go
	c <- 3
```

pero como el búfer no se desborda por **overflowing** (*ya que no **enviamos** ningún valor nuevo*), la `goroutine` `main` no se bloqueará y el programa existe.

Veamos que pasa si enviamos algunos valores extra.

```go
package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)

	go squares(c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocks here

	fmt.Println("main() stopped")
}
```


```text
main() started
1
4
9
16
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/KGyiskRj1Wi)

Como hemos visto, cuando enviamos el numero `4` al `channel` `c` con la instrucción `c <- 4`, ahora el `buffer` del `channel` `c` esta lleno, la `goroutine` `main` se bloquea y la `goroutine` `squares` imprimira todos los valores.

### 1.6.3 Longitud y capacidad de channel

De manera similar a un `slice`, un `buffered channel` tiene longitud y capacidad.

- La **longitud** de un `channel` es la **cantidad de valores en cola (no leídos) en el `buffer` del `channel`**.
- La **capacidad** de un `channel` es el **tamaño del `buffer`**.

Para calcular la longitud, usamos la función `len`, mientras que para averiguar la capacidad, usamos la función `cap`, como un `slice`.

```go
package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v", len(c), cap(c))
	fmt.Println()
}
```

```text
Length of channel c is 2 and capacity of channel c is 3
```

[Ejemplo en vivo](https://go.dev/play/p/qsDZu6pXLT7)

Si nos preguntamos por qué el programa anterior funciona bien y no se produjo un error de `deadlock`. Esto se debe a que, como la capacidad del `channel` es `3` y solo hay `2` valores disponibles en el `buffer`, Go no intentó programar otra `goroutine` bloqueando la ejecución de la `goroutine` `main`. Simplemente puede leer estos valores en la rutina principal si lo desea, porque incluso si el búfer no está lleno, eso no le impide leer valores del canal.

Aquí hay otro ejemplo.

```go
package main

import "fmt"

func sender(c chan int) {
	c <- 1 // len 1, cap 3
	c <- 2 // len 2, cap 3
	c <- 3 // len 3, cap 3
	c <- 4 // <- goroutine blocks here
	close(c)
}

func main() {
	c := make(chan int, 3)

	go sender(c)

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))

	// read values from c (blocked here)
	for val := range c {
		fmt.Printf("Length of channel c after value '%v' read is %v\n", val, len(c))
	}
}
```

```text
Length of channel c is 0 and capacity of channel c is 3
Length of channel c after value '1' read is 3
Length of channel c after value '2' read is 2
Length of channel c after value '3' read is 1
Length of channel c after value '4' read is 0
```

[Ejemplo en vivo](https://go.dev/play/p/-gGpm08-wzz)

```go
package main

import (
	"fmt"
	"runtime"
)

func squares(c chan int) {
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)
	go squares(c)

	fmt.Println("active goroutines", runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocks here

	fmt.Println("active goroutines", runtime.NumGoroutine())

	go squares(c)

	fmt.Println("active goroutines", runtime.NumGoroutine())

	c <- 5
	c <- 6
	c <- 7
	c <- 8 // blocks here

	fmt.Println("active goroutines", runtime.NumGoroutine())
	fmt.Println("main() stopped")
}
```

```text
main() started
active goroutines 2
1
4
9
16
active goroutines 2
active goroutines 2
25
36
49
64
active goroutines 2
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/sdHPDx64aor)

### 1.6.4 Leer desde un `channel` cerrado

Usando `buffered channels` y `for range`, podemos leer desde `channels` cerrados. **Dado que en el caso de los `channels` cerrados los datos residen en el `buffer`, aún podemos extraer esos datos**.

```go
package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	// iteration terminates after receving 3 values
	for elem := range c {
		fmt.Println(elem)
	}
}
```

```text
1
2
3
```

[Ejemplo en vivo](https://go.dev/play/p/vULFyWnpUoj)

## 1.7 Trabajando con múltiples `goroutines`

Escribamos 2 `goroutines`, una `square` para calcular el cuadrado de números enteros `int` y otra `cube` para el cubo de números enteros `int`. s

```go
package main

import "fmt"

func square(c chan int) {
	fmt.Println("[square] reading")
	num := <-c
	c <- num * num
}

func cube(c chan int) {
	fmt.Println("[cube] reading")
	num := <-c
	c <- num * num * num
}

func main() {
	fmt.Println("[main] main() started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	testNum := 3
	fmt.Println("[main] sent testNum to squareChan")

	squareChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] sent testNum to cubeChan")

	cubeChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] reading from channels")

	squareVal, cubeVal := <-squareChan, <-cubeChan
	sum := squareVal + cubeVal

	fmt.Println("[main] sum of square and cube of", testNum, " is", sum)
	fmt.Println("[main] main() stopped")
}
```

Hablemos del ejemplo anterior paso a paso.

1. Creamos 2 funciones `square` y `cube` que ejecutaremos como `goroutines`. Ambos reciben el `channel` de tipo `int` como argumento `c` y leemos datos del mismo en la variable `num`. Luego escribimos datos en el `channel` `c` en la siguiente línea.
2. En la `goroutine` `main`, creamos 2 `channels` `squareChan` y `cubeChan` de tipo `int` usando la función `make`.
3. Luego ejecutamos la `goroutine` `square` y `cube`.
4. Dado que el control todavía está dentro de la `goroutine` `main`, la variable `testNumb` obtiene el valor de `3`.
5. Luego enviamos datos a `squareChan` y `cubeChan`. La `goroutine` `main` se bloqueará hasta que estos canales lo lean. Una vez que lo lean, `main` `goroutine` continuará ejecutándose.
6. Cuando en la `goroutine` `main` intentamos leer datos de `channels` determinados, el control se bloqueará hasta que estos `channels` escriban algunos datos de sus `goroutines`. Aquí, hemos utilizado la sintaxis abreviada `:=` para recibir datos de múltiples `channels`.
7. Una vez que estas `goroutines` escriban algunos datos en el `channel`, la `goroutine` `main` se bloqueará.
8. Cuando finaliza la operación de escritura del `channel`, la `goroutine main` comienza a ejecutarse. Luego calculamos la suma y la imprimimos en la consola.

Por lo tanto, el ejemplo anterior producirá el siguiente resultado.

```text
[main] main() started
[main] sent testNum to squareChan
[square] reading
[main] resuming
[main] sent testNum to cubeChan
[cube] reading
[main] resuming
[main] reading from channels
[main] sum of square and cube of 3  is 36
[main] main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/6wdhWYpRfrX)

## 1.8 `Channels` unidireccionales

Hasta ahora, hemos visto `channels` que pueden transmitir datos desde ambos lados, canales en los que podemos realizar operaciones de lectura y escritura. Pero también podemos crear `channels` que sean de naturaleza unidireccional. Por ejemplo, `channels` de solo recepción que solo permiten operaciones de lectura en ellos y `channels` de solo envío que solo permiten operaciones de escritura en ellos.

El `channel` unidireccional también se crea usando la función `make` pero con una sintaxis de flecha adicional.

```go
roc := make(<-chan int) // receive-only channel
soc := make(chan<- int) // send-only channel
```

En el programa anterior, `roc` es un `channel` de solo recepción donde la dirección de flecha en la función `make` que apunta hacia afuera de la palabra clave `chan`. Mientras que `soc` es un `channel` de solo envío donde la dirección de la flecha en la función `make` apunta hacia la palabra clave `chan`. También tienen un tipo diferente.

```go
package main

import "fmt"

func main() {
	roc := make(<-chan int)
	soc := make(chan<- int)

	fmt.Printf("Data type of roc is `%T`\n", roc)
	fmt.Printf("Data type of soc is `%T\n", soc)
}
```

```text
Data type of roc is `<-chan int`
Data type of soc is `chan<- int
```

[Ejemplo en vivo](https://go.dev/play/p/JZO51IoaMg8)

Pero, **¿para qué sirve un `channel` unidireccional?** El uso de `channels` unidireccionales **aumenta el `type safety` de un programa**. Por tanto, el programa es menos propenso a errores.

Pero supongamos que tenemos una `goroutine` en la que solo necesita leer datos de un canal y la `goroutine` `main` necesita leer y escribir datos desde/hacia el mismo `channel`. ¿Cómo hariamos eso?

Afortunadamente, Go proporciona una sintaxis más sencilla para **convertir un canal bidireccional en un canal unidireccional**.

```go
package main

import "fmt"

func greet(roc <-chan string) {
	fmt.Println("Hello " + <-roc + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "John"
	fmt.Println("main() stopped")
}
```

```text
main() started
Hello John!
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/k3B3gCelrGv)

Modificamos la `goroutine` `greet` para convertir el canal bidireccional `c` en el canal `roc` de solo recepción o lectura en la función `greet`. Desde ese momento sólo podemos leer en ese canal. Cualquier operación de escritura en él resultará en un fatal error `"invalid operation: roc <- "some text" (send to receive-only type <-chan string)"`.

## 1.9 `Goroutine` anonima

En el capítulo de [goroutines](../goroutines/), aprendimos sobre las `goroutines` anónimas. También podemos implementar `channels` con ellas. Modifiquemos el ejemplo simple anterior para implementar el `channel` en una `goroutine` anónima.

Este fue nuestro ejemplo anterior.

```go
package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)
	go greet(c)

	c <- "John"
	fmt.Println("main() stopped")
}
```

```text
main() started
Hello John!
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/c5erdHX1gwR)

A continuación se muestra el ejemplo modificado en el que hacemos que la `goroutine` `greet` fuera una `goroutine` anónima.

```go
package main

import "fmt"

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	// launch anonymous goroutine
	go func(c chan string) {
		fmt.Println("Hello " + <-c + "!")
	}(c)

	c <- "John"
	fmt.Println("main() stopped")
}
```

```text
main() started
Hello John!
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/cM5nFgRha7c)

## 1.10 `Channel` como tipo de datos de un `channel`

```go
package main

import "fmt"

// gets a channel and prints the greeting by reading from channel
func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

// gets a channels and writes a channel to it
func greeter(cc chan chan string) {
	c := make(chan string)
	cc <- c
}

func main() {
	fmt.Println("main() started")

	// make a channel `cc` of data type channel of string data type
	cc := make(chan chan string)

	go greeter(cc) // start `greeter` goroutine using `cc` channel

	// receive a channel `c` from `greeter` goroutine
	c := <-cc

	go greet(c) // start `greet` goroutine using `c` channel

	// send data to `c` channel
	c <- "John"

	fmt.Println("main() stopped")
}
```

```text
main() started
Hello John!
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/xVQvvb8O4De)

## 1.11 Select

`select` es como un `switch` sin ningún argumento de entrada pero solo se usa para `channels operations`. La instrucción `select` se utiliza para realizar una operación en solo un `channel` entre muchos, seleccionado sobre que `channel` realizar la operación condicionalmente por un bloque de casos.

Primero veamos un ejemplo y luego analicemos cómo funciona.

```go
package main

import (
	"fmt"
	"time"
)

var start time.Time
func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
```

En el ejemplo anterior, podemos ver que la instrucción `select` es como la instruccion `switch`, pero en lugar de operaciones booleanas, agregamos `channel operations` como lectura o escritura o una combinación de lectura y escritura. La declaración de `select` está bloqueando la `goroutine` excepto cuando tiene un caso predeterminado o `default`. Una vez que se cumpla una de las condiciones del caso, se desbloqueará.

Si todas las declaraciones `case` (`channel operations`) están bloqueando, la declaración de `select` esperará hasta que una de las declaraciones `case` (su `channel opertation`) se desbloquee y en ese caso se ejecutará. Si algunas o todas las `channel operations` no son bloqueantes, entonces uno de los `cases` no bloqueantes se elegirá aleatoriamente y se ejecutará inmediatamente.

Para explicar el programa anterior,

1. Iniciamos `goroutines` `service1` y `service2` con canales independientes `chan ` `chan2`.
2. Luego intorducimos una declaración `select` con 2 `cases`. Un `case` lee un valor de `chan1` y otro de `chan2`. Dado que estos canales no tienen `buffer`, la operación de lectura bloquearán (por lo tanto, las operaciones de escritura). 3. Entonces ambos casos de `select` están bloqueando. Por lo tanto, `select` esperará hasta que uno de los `cases` se desbloquee.
4. Cuando el control está en la instrucción de `select`, la `goroutine` `main` se bloqueará y programará todas las `goroutines` presentes en la instrucción de `select` (una a la vez), que son `service1` y `service2`.
5. `service1` espera 3 segundos y luego se desbloquea escribiendo en `chan1`.
6. De manera similar, `service2` espera 5 segundos y luego se desbloquea escribiendo en `chan2`.
7. Luego, dado que el `service1` se desbloquea antes que el `service2`, el `case` 1 se desbloqueará primero y, por lo tanto, ese caso se ejecutará y los demás `cases` (aquí el `case` 2) se ignorarán.
8. Una vez terminada la ejecución del `case`, continua la ejecución de la función `main`.

```text
main() started 0s
Response from service 1 Hello from service 1 3s
main() stopped 3s
```

[Ejemplo en vivo](https://go.dev/play/p/ar5dZUQ2ArH)

> [!TIP]
> El programa anterior simula un servicio web del mundo real donde un balanceador de carga recibe millones de solicitudes y tiene que devolver una respuesta de uno de los servicios disponibles. Usando `gorutines`, `channels` y `select`, podemos solicitar una respuesta a múltiples servicios, y se puede usar uno que responda rápidamente.


Para simular que:

1. Todos los `cases` son bloqueantes.
2. La respuesta está disponible casi al mismo tiempo.

simplemente podemos eliminar la llamada de `Sleep`.

```go
package main

import (
	"fmt"
	"time"
)

var start time.Time
func init() {
	start = time.Now()
}

func service1(c chan string) {
	c <- "Hello from service 1"
}

func service2(c chan string) {
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
```

Este ejemplo podria lanzar la siguiente salida

```text
main() started 0s
Response from service 2 Hello from service 2 0s
main() stopped 0s
```

Pero algunas veces podria lanzar esta otra salida

```text
main() started 0s
service1() started 484.8µs
Response from service 1 Hello from service 1 984µs
main() stopped 984µs
```

[Ejemplo en vivo](https://go.dev/play/p/giSkkqt8XHb)


Esto sucede porque las operaciones sobre `chan1` y `chan2` ocurren casi al mismo tiempo, pero aún así, hay cierta diferencia de tiempo en la ejecución y programación.

Para simular que:

1. Todos los `cases` no son bloqueantes.
2. La respuesta está disponible al mismo tiempo.

Podemos usar un `buffered channel`.

```go
package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {
	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	chan1 <- "Value 1"
	chan1 <- "Value 2"
	chan2 <- "Value 1"
	chan2 <- "Value 2"

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
```


[Ejemplo en vivo](https://go.dev/play/p/RLRGEmFQP3f)


En este ejemplo nos encontramos que la salida puede ser la siguiente

```text
main() started 0s
Response from chan2 Value 1 0s
main() stopped 1.0012ms
```

Pero tambien nos encontramos que podria ser esta otra

```text
main() started 0s
Response from chan1 Value 1 0s
main() stopped 0s
```

En el programa anterior, ambos `channels` tienen `2` valores en su `buffer`. Dado que estamos enviando `2` valores a un `channel` con capacidad de `buffer` `2`, estas `channel operations` no se bloquearán y el control irá a la instrucción de `select`. Dado que la lectura del `buffered channel` es una operación no bloqueante hasta que todo el `buffer` está vacío y leemos solo un valor en la condición `case`, todas las `case operations` no son bloqueantes. Por lo tanto, `Go runtime` seleccionará cualquier `case operation` al azar.

### 1.11.1 Select con default, default case

Al igual que la declaración `switch`, la declaración `select` también tiene `default case`. **El caso predeterminado es no bloqueante**. Pero eso no es todo, el caso predeterminado hace que la instrucción `select` siempre sea no bloqueante. Eso significa que la operación de envío y recepción de cualquier `channel` (con o sin `buffer`) siempre es no bloqueante.

Si un valor está disponible en cualquier `channel`, `select` ejecutará ese `case`. De lo contrario, ejecutará inmediatamente el `default case`.

```go
package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	fmt.Println("service1() started", time.Since(start))
	c <- "Hello from service 1"
}

func service2(c chan string) {
	fmt.Println("service2() started", time.Since(start))
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	default:
		fmt.Println("No response received", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
```

```text
main() started 0s
No response received 0s
main() stopped 0s
```

[Ejemplo en vivo](https://go.dev/play/p/rFMpc80EuT3)

En el programa anterior, dado que los `channels` no son `buffered channels` y el valor no está disponible inmediatamente en las operaciones de ambos `channels`, se ejecutará el `default case`. Si la declaración de `select` anterior no tuviera `default case`, `select` habría estado bloqueando y la respuesta habría sido diferente.

Dado que, de forma predeterminada, `select` no bloquea, el `scheduler` no recibe una llamada de `main goroutine` para programar las rutinas disponibles. Pero podemos hacerlo manualmente llamando a `time.Sleep`. De esta manera, todas las `gorutines` se ejecutarán y morirán, devolviendo el control a `main goroutine`, que se reactivará después de un tiempo. Cuando `main goroutine` se active, los `channels` tendrán valores disponibles inmediatamente.

Veamoslo en el siguiente ejemplo.

```go
package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	fmt.Println("service1() started", time.Since(start))
	c <- "Hello from service 1"
}

func service2(c chan string) {
	fmt.Println("service2() started", time.Since(start))
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	time.Sleep(3 * time.Second)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	default:
		fmt.Println("No response received", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
```

En el cual podriamos tener la siguiente salida:

```text
main() started 0s
service1() started 0s
service2() started 0s
Response from service 1 Hello from service 1 3s
main() stopped 3s
```

Pero que algunas veces podria ser:

```text
main() started 0s
service1() started 0s
service2() started 0s
Response from service 2 Hello from service 2 3.0000957s
main() stopped 3.0000957s
```

[Ejemplo en vivo](https://go.dev/play/p/eD0NHxHm9hN)

### 1.11.2 Deadlock

El `default` `case` es útil cuando no hay `channels` disponibles para enviar o recibir datos. Para evitar un punto muerto `deadlock`, podemos usar el `default` `case`. Esto es posible porque todas las operaciones del `channel` debido al caso predeterminado no son de bloqueo, Go no programa ninguna otra `goroutine` para enviar datos a los `channels` si los datos no están disponibles de inmediato.

```go
package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	default:
		fmt.Println("No goroutines available to send data", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
```

```text
main() started 0s
No goroutines available to send data 0s
main() stopped 0s
```

[Ejemplo en vivo](https://go.dev/play/p/S3Wxuqb8lMF)

De manera similar a recibir, en la operación de envío, si otras `goroutines` están inactivas (no listas para recibir el valor), se ejecuta el caso predeterminado.

### 1.11.3 `nil` channel

Como sabemos, el valor predeterminado de un canal es `nil` nulo. Por lo tanto, no podemos realizar operaciones de envío o recepción en un canal nulo `nil`. Pero en el caso de que se utilice un canal nulo `nil` en la declaración de selección, se producirá uno de los siguientes errores o ambos.

```go
package main

import "fmt"

func service(c chan string) {
	c <- "response"
}

func main() {
	fmt.Println("main() started")

	var chan1 chan string

	go service(chan1)

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res)
	}

	fmt.Println("main() stopped")
}
```

```text
main() started
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive (nil chan)]:
main.main()
	/tmp/sandbox3177050049/prog.go:17 +0x85

goroutine 6 [chan send (nil chan)]:
main.service(...)
	/tmp/sandbox3177050049/prog.go:6
created by main.main in goroutine 1
	/tmp/sandbox3177050049/prog.go:14 +0x73
```

[Ejemplo en vivo](https://go.dev/play/p/uhraFubcF4S)

Del ejemplo anterior, podemos ver que `select` (sin `cases`) significa que la declaración de `select` está virtualmente vacía porque **se ignoran los `cases` con `nil` `channel`**. Pero como la instrucción `select{}` vacía bloquea la `main` `goroutine` y la `goroutine` `service` está programada `scheculed` en su lugar, la operación del `channel` en `nil` `channels` arroja un `chan send (nil chan)` error. Para evitar esto, utilizamos el `default` `case`.

```go
package main

import "fmt"

func service(c chan string) {
	c <- "response"
}

func main() {
	fmt.Println("main() started")

	var chan1 chan string

	go service(chan1)

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res)
	default:
		fmt.Println("No response")
	}

	fmt.Println("main() stopped")
}
```

```text
main() started
No response
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/upLsz52_CrE)

El programa anterior no solo ignora el bloque de `cases` sino que ejecuta el `default` `case` inmediatamente. Por lo tanto, el `scheduler` no tiene tiempo para programar `schedule` la `service` `goroutine`. Pero, en realidad,  este es un muy mal diseño. Siempre debes comprobar si un `channel` tiene un valor nulo `nill`.

### 1.11.4 Añadiendo un timeout

El programa anterior no es muy útil ya que solo se ejecuta el `default` `case`. Pero a veces, lo que queremos es que cualquier servicio disponible responda en el momento adecuado, si no es así, entonces se debe ejecutar el `default` `case`. Esto se puede hacer usando un `case` con una operación de `channel` que se desbloquea después de un tiempo definido. Esta operación de `channel` es proporcionada por la función `After` del paquete `time`.

Veamos un ejemplo:

```go
package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("No response received", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
```

[Ejemplo en vivo](https://go.dev/play/p/mda2t2IQK__X)

El programa anterior arroja el siguiente resultado después de 2 segundos.

```text
main() started 0s
No response received 2s
main() stopped 2s
```

En el programa anterior, `<-time.After(2 * time.Second)` se desbloquea después de `2` segundos y devuelve el tiempo en el que se desbloqueó, pero aquí no nos interesa su valor de retorno. Dado que también actúa como una `goroutine`, tenemos 3 `goroutines` de las cuales ésta se desbloquea primero. Por lo tanto, se ejecuta el caso correspondiente a esa operación de `goroutine`.

Esto es útil porque no desea esperar demasiado para obtener una respuesta de los servicios disponibles, donde el usuario tenga que esperar demasiado tiempo antes de obtener algo del servicio. Por lo tanto, si sumamos `10 * time.Second` en el ejemplo anterior, se imprimirá la respuesta del `service1`.

```go
package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	case <-time.After(10 * time.Second):
		fmt.Println("No response received", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
```

```text
main() started 0s
Response from service 1 Hello from service 1 3s
main() stopped 3s
```

[Ejemplo en vivo](https://go.dev/play/p/qoc-xuI9pUg)

### 1.11.5 `select` vacío

Al igual que el bucle vacío `for{}`, una sintaxis `select{}` vacía también es válida, pero hay un problema. Como sabemos, la declaración de `select` se bloquea hasta que uno de los `cases` se desbloquea y, dado que no hay declaraciones de `cases` disponibles para desbloquearlo, la `main` `goroutine` se bloqueará para siempre, lo que provocará un `deadlock`.

```go
package main

import "fmt"

func service() {
	fmt.Println("Hello from service!")
}

func main() {
	fmt.Println("main() started")

	go service()

	select {}

	fmt.Println("main() stopped")
}
```

En el programa anterior, como sabemos, `select` bloqueará la `main` `goroutine`, el `scheduler` hara `schedule` de otra `goroutine` disponible que es el `service`. Pero después de eso, morirá y el `scheduler` tendrá que programar otra `goroutine` disponible, pero como la `main` `goroutine` está bloqueada y no hay otras `goroutines` disponibles, se produce un `deadlock`.

```text
main() started
Hello from service!
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [select (no cases)]:
main.main()
        program.Go:16 +0xba
exit status 2
```

[Ejemplo en vivo](https://go.dev/play/p/-pBd-BLMFOu)

### 1.11.6 `WaitGroup`

Imaginemos una condición en la que necesitamos saber si todas las `goroutines` terminaron su trabajo. Esto es algo opuesto a `select` donde solo necesitabamos que una condición fuera `true`, pero aquí necesitamos que **todas las condiciones sean `true`** para desbloquear la `main` `goroutine`. Aquí la condición es que la `channel` opertaion funcione correctamente.

`WaitGroup` es una `struct` con un valor de contador `counter` que rastrea cuántas `goroutines` se generaron y cuántas completaron su trabajo. Este contador, cuando llega a cero, significa que todas las `goroutines` han hecho su trabajo.

Profundicemos en un ejemplo y veamos la terminología.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func service(wg *sync.WaitGroup, instance int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
	wg.Done() // decrement counter
}

func main() {
	fmt.Println("main() started")
	var wg sync.WaitGroup // create waitgroup (empty struct)

	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment counter
		go service(&wg, i)
	}

	wg.Wait() // blocks here
	fmt.Println("main() stopped")
}
```

En el programa anterior, creamos una `struct` vacía (con campos de valor cero `zero-value`) `wg` de tipo `sync.WaitGroup`. La `struct` `WaitGroup` tiene campos no exportados como `noCopy`, `state1` y `sema` cuya implementación interna no necesitamos saber. Esta `struct` tiene tres métodos a conocer. `Add`, `Wait` y `Done`.

El método `Add` espera un argumento `int` que es `delta` para el contador `WaitGroup`. El contador no es más que un número entero con valor predeterminado `0`. Contiene cuántas `goroutines` se están ejecutando. **Cuando se crea `WaitGroup`, su valor de contador  `counter` es `0`** y podemos incrementarlo pasando `delta` como parámetro usando el método `Add`.

> [!IMPORTANT]
> Recuerde, el contador `counter` no entiende de su incremento cuando se lanza una nueva `goroutine`, por lo que debemos incrementarlo manualmente.

El método `Wait` se utiliza para bloquear la `goroutine` actual desde donde fue llamada. Una vez que el contador `counter` llegue a `0`, esa `goroutine` se desbloqueará. Por lo tanto, necesitamos algo para disminuir el contador `counter`.

El método `Done` disminuye el contador `counter`. No acepta ningún argumento, por lo que sólo disminuye el contador en `1`.

En el programa anterior, después de crear `wg`, ejecutamos el bucle `for` 3 veces. En cada turno, lanzamos 1 `goroutine` e incrementamos el contador en `1`. Eso significa que ahora tenemos `3` `goroutines` esperando ser ejecutadas y el contador `counter` de `WaitGroup` es `3`.

> [!IMPORTANT]
> Observe que pasamos un puntero a `wg` a la `goroutine` `service`.

Esto se debe a que en la `goroutine` `service`, una vez que terminamos con lo que se suponía que debía hacer, necesitamos llamar al método `Done` para disminuir el contador. Si se pasara `wg` como valor, `wg` en `main` no disminuiría.

> [!TIP]
> Si queremos mutar un valor interno de una estructura, siempre pasamos un puntero a esa estructura.

Una vez que el bucle `for` terminó de ejecutarse, todavía no pasamos el control a otras `goroutines`. Esto se hace explicitamente llamando al método `Wait` en `wg` como `wg.Wait()`. Esto bloqueará la `main` `goroutine` hasta que el contador `counter` llegue a `0`. Una vez que el contador `counter` llegue a `0`, llamamos al método `Done` en `wg` 3 veces desde cada una de las 3 `goroutines` una vez terminen su trabajo, la `main` `goroutine` se desbloqueará y comenzará a ejecutar más código.

Por lo tanto, el programa anterior produce el resultado siguiente.

```text
main() started
Service called on instance 2
Service called on instance 3
Service called on instance 1
main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/8qrAD9ceOfJ)

El resultado anterior puede ser diferente para ti, ya que el orden de ejecución de las `goroutines` puede variar.

> [!NOTE]
> El método `Add` acepta el tipo de `int`, eso significa que `delta` también puede ser negativo. Para saber más sobre esto, visite la [documentación oficial](https://golang.org/pkg/sync/#WaitGroup.Add).

### 1.11.7 Worker pool

Como sugiere el nombre, un `worker pool` es una colección de `goroutines` que trabajan simultáneamente de forma concurrente para realizar un trabajo. En `WaitGroup`, vimos una colección de `goroutines` funcionando simultáneamente pero no tenían un trabajo específico. Una vez que les lanzas `channels`, tienen algo de trabajo que hacer y se convierten en un `worker pool`.

Entonces, el concepto detrás del `worker pool` es mantener un grupo de `goroutines` de `workers` que reciben alguna `task` y devuelve el resultado. Una vez que todos hayan terminado con su trabajo, recopilamos el resultado. Todas estas `goroutines` utilizan el mismo `channel` para fines individuales.

Veamos un ejemplo simple con dos `channels`, `tasks` y `results`.

```go
package main

import (
	"fmt"
	"time"
)

// worker than make squares
func sqrWorker(tasks <-chan int, results chan<- int, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond) // simulating blocking task
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		results <- num * num
	}
}

func main() {
	fmt.Println("[main] main() started")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		go sqrWorker(tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	// receving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped")
}
```

Vamos a explicar que esta pasando en el ejemplo anterior.

- `sqrWorker` es una función worker que toma el `channel` de `tasks`, el `channel` de `results` y un id `instance`. El trabajo de esta `goroutine` es enviar cuadrados del número recibido del `channel` de `tasks` al `channel` de `results`.

- En la función `main`, creamos `tasks` y `reults` como `buffered channels` con capacidad de `10` elementos `int`. Por lo tanto, **cualquier operación de envío no bloqueará hasta que el búfer esté lleno**.

- Luego generamos múltiples instancias de `sqrWorker` como `goroutines` con los dos `channels` anteriores y el parámetro id `instance` para obtener información sobre qué `worker` está ejecutando una `task`.

- Luego pasamos 5 `tasks` al `channel` de `tasks` que sin bloquearlo ya que `5` < `10` que era la capacidad de nuestros `buffered channels`.

- Como terminamos con el `channel` de `tasks`, lo cerramos con `close(tasks)`. Esto no es necesario, pero ahorrará mucho tiempo en el futuro si obtenemos algunos errores.

- Luego, usando el bucle `for`, con 5 iteraciones, extraemos datos del `channel` de `results`. Dado que la operación de lectura en un búfer vacío es bloqueante, el `scheduler` programara una `goroutina` desde el `worker pool`. Hasta que esa `goroutine` devuelva algún resultado, la `main` `goroutine` estará bloqueada.

- Dado que estamos simulando una operación de bloqueo en `worker` `goroutine`, se llamará al `scheduler` para programar otra `goroutine` disponible. Cuando `worker` `goroutine` este disponible, escribe en el `channel` de `results`. Como escribir en un `buffered channel` no bloquea hasta que el `buffer` esté lleno, escribir en el `channel` de `results` aquí no es bloqueante. Además, aunque el `worker` `goroutine` actual no estaba disponible, se ejecutaron muchas otras `worker` `goroutines` consumiendo valores del `buffer` de `tasks`. Después de que todas las `worker` `goroutines` consuman todas las `tasks`, el bucle `for range` finaliza cuando el `buffer` del `channel` de `tasks` este vacío. No arrojará un error de `deadlock` ya que el cerramos el `channel` de `tasks`.

- A veces, todas las `worker` `goroutines` pueden estar inactivas o `sleeping`, por lo que la `main` `goroutine` se activará y funcionará hasta que el `buffer` del `channel` de `results` vuelva a estar vacío.

- Después de que todas las `worker` `goroutines` hayan muerto o `died`, la `main` `goroutine` recuperará el control e imprimirá los resultados restantes del `channel` de `results` y continuará su ejecución.


```text
[main] main() started
[main] Wrote 5 tasks
[worker 1] Sending result by worker 1
[worker 2] Sending result by worker 2
[main] Result 0 : 16
[main] Result 1 : 0
[worker 0] Sending result by worker 0
[main] Result 2 : 4
[worker 1] Sending result by worker 1
[main] Result 3 : 36
[worker 2] Sending result by worker 2
[main] Result 4 : 64
[main] main() stopped
```

[Ejemplo en vivo](https://go.dev/play/p/IYiMV1I4lCj)

El ejemplo anterior es complejo pero explica muy bien cómo múltiples `goroutines` pueden leer del mismo canal y hacer el trabajo con elegancia. Las `goroutines` son poderosas cuando el trabajo del `worker` está bloqueado. Si elimina la llamada `time.Sleep()` del ejemplo anterior, entonces solo una `goroutine` realizará el trabajo, ya que no hay otras `goroutines` programadas hasta que finalice el bucle `for range` y la `goroutine` muera.

> [!IMPORTANT]
> Puede obtener resultados diferentes como en el ejemplo anterior dependiendo de qué tan rápido sea su sistema porque si todas las `worker` `goroutines` están bloqueadas, incluso por un microsegundo, la `main` `goroutine` se activará como explicamos.

Ahora, usemos el concepto de `WaitGroup` para sincronizar `goroutines`. Usando el ejemplo anterior con `WaitGroup`, podemos lograr los mismos resultados pero de manera más elegante.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// worker than make squares
func sqrWorker(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond)
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		results <- num * num
	}

	// done with worker
	wg.Done()
}

func main() {
	fmt.Println("[main] main() started")

	var wg sync.WaitGroup

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sqrWorker(&wg, tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	// wait until all workers done their job
	wg.Wait()

	// receving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // non-blocking because buffer is non-empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped")
}
```
