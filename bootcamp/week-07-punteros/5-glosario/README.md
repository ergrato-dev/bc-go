# Glosario — Semana 07: Punteros

Términos técnicos de la semana en orden alfabético.

---

## A

### Auto-deref (auto-derreferencia)
Mecanismo de Go por el cual, al acceder a un campo de struct a través de un puntero con notación de punto (`p.Campo`), el compilador inserta la derreferencia automáticamente (`(*p).Campo`). También aplica al llamar métodos: `variable.Metodo()` se convierte en `(&variable).Metodo()` si el método tiene receptor `*T` y `variable` es direccionable.

---

## D

### Derreferencia (dereference)
Operación de acceder al valor almacenado en la dirección a la que apunta un puntero. En Go se realiza con el operador `*` aplicado al puntero: `*p` lee o escribe el valor en la dirección que guarda `p`.

### Dirección de memoria (memory address)
Número que identifica una posición específica en la RAM donde se almacena un dato. En Go se obtiene con el operador `&`. Se imprime con el verbo `%p` en `fmt.Printf`.

---

## E

### Escape analysis
Análisis que realiza el compilador de Go para decidir si una variable se almacena en el stack (acceso rápido, liberación automática al terminar la función) o en el heap (vida más larga, gestionada por el garbage collector). Cuando una variable "escapa" al heap, suele ser porque un puntero a ella puede sobrevivir a la función que la creó.

---

## H

### Heap
Región de memoria de vida dinámica gestionada por el garbage collector de Go. Las variables que deben sobrevivir al stack frame de su función (por ejemplo, si un puntero a ellas se retorna o se almacena externamente) se ubican en el heap. Más lento de acceder que el stack pero de tamaño flexible.

---

## M

### Método de puntero (pointer receiver method)
Método definido con receptor `*T`. Recibe la dirección del valor receptor y puede modificar sus campos. Ejemplo: `func (c *Contador) Incrementar()`. Necesario cuando el método cambia el estado del struct.

### Método de valor (value receiver method)
Método definido con receptor `T` (sin asterisco). Recibe una copia del valor receptor. No puede modificar el original. Adecuado para métodos de solo lectura en structs pequeños.

---

## N

### `new(T)`
Función builtin de Go que asigna memoria para un valor de tipo `T` inicializado en su zero value y retorna un `*T`. `new(int)` retorna `*int` con `*p == 0`. Equivalente a `&T{}` para structs.

### nil (puntero nil)
Zero value de cualquier tipo puntero en Go. Un puntero nil no apunta a ninguna dirección válida. Derreferir un puntero nil (`*p` cuando `p == nil`) produce un pánico en tiempo de ejecución: `panic: runtime error: invalid memory address or nil pointer dereference`.

---

## P

### Paso por puntero (pass by pointer)
Patrón en el que se pasa la dirección de una variable (`&x`) a una función o método. La función puede modificar el valor original a través de la derreferencia. En Go no existe "paso por referencia" real; se pasa un puntero, que es un valor que contiene una dirección.

### Paso por valor (pass by value)
Comportamiento predeterminado en Go: al pasar una variable a una función, se crea una copia independiente. La función trabaja sobre la copia y el original no se ve afectado. Aplica a todos los tipos: `int`, `string`, structs, arrays.

### Puntero (pointer)
Variable que almacena la dirección de memoria de otra variable. En Go el tipo de un puntero a `T` se escribe `*T`. El zero value de cualquier tipo puntero es `nil`.

---

## S

### Stack
Región de memoria de acceso rápido y gestión automática. Cada función tiene su propio stack frame donde se almacenan sus variables locales. Al retornar la función, el stack frame se descarta. Las variables en el stack no requieren garbage collection.

---

## Z

### Zero value de puntero
El valor inicial (zero value) de cualquier tipo puntero en Go es `nil`. Al declarar `var p *int`, `p` vale `nil` hasta que se le asigne una dirección con `p = &x` o `p = new(int)`.
