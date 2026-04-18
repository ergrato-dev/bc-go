# Glosario — Semana 08: Métodos e Interfaces

Términos técnicos de la semana en orden alfabético.

---

## A

### `any`
Alias de `interface{}` introducido en Go 1.18. Representa la interface vacía que todos los tipos satisfacen. Una variable `any` puede almacenar cualquier valor, pero para operar sobre él se requiere una **type assertion**. Preferir interfaces concretas sobre `any` cuando sea posible.

---

## D

### Duck typing
Estrategia de tipado donde un tipo satisface una interface si tiene los métodos requeridos, sin necesidad de declararlo explícitamente. Nombre del concepto: "si camina como pato y grazna como pato, es un pato". En Go, la satisfacción de interfaces es siempre implícita.

---

## F

### `fmt.Stringer`
Interface de la librería estándar definida en el paquete `fmt`: `type Stringer interface { String() string }`. Si un tipo implementa `String() string`, `fmt.Println` y verbos `%v`/`%s` de `fmt.Printf` usan ese método automáticamente para representación textual.

---

## I

### Interface
Tipo en Go que define un conjunto de firmas de métodos. Ningún tipo declara que implementa una interface — la satisfacción es implícita. Una variable de tipo interface puede almacenar cualquier valor cuyo tipo concreto tenga los métodos requeridos.

### Interface vacía
`interface{}` o `any`. No requiere ningún método, por lo que todos los tipos la satisfacen. Útil cuando el tipo es verdaderamente desconocido; de lo contrario, usar una interface con métodos específicos.

---

## M

### Method set
El conjunto de métodos que puede llamar un tipo. El method set de `T` incluye solo métodos con receptor `T`. El method set de `*T` incluye métodos con receptor `T` y métodos con receptor `*T`. El method set determina qué interfaces puede satisfacer un tipo.

### Método (method)
Función con un receptor declarado entre `func` y el nombre. El receptor puede ser de tipo valor `(t T)` o puntero `(t *T)`. En Go se pueden definir métodos sobre cualquier tipo declarado en el mismo paquete.

---

## P

### Polimorfismo con interfaces
Capacidad de una función de operar sobre valores de tipos distintos siempre que todos satisfagan la misma interface. En Go se logra declarando el parámetro como tipo interface.

---

## R

### Receptor (receiver)
Parámetro especial de un método que aparece entre `func` y el nombre del método. Puede ser de valor `(r Tipo)` — trabaja sobre copia — o de puntero `(r *Tipo)` — accede al original.

---

## S

### Satisfacción de interface
Un tipo satisface una interface cuando implementa todos los métodos declarados en ella con las firmas exactas (nombre, parámetros y tipos de retorno). En Go esta verificación es implícita — el compilador la realiza automáticamente.

---

## T

### Type assertion
Operación para extraer el tipo concreto de una variable de interface: `valor, ok := v.(TipoConcreto)`. Si `ok` es `true`, `valor` tiene el tipo concreto y se pueden usar todos sus métodos. Si `ok` es `false`, `valor` es el zero value del tipo.

---

## V

### Verificación de compilación
Patrón `var _ Interface = Tipo{}` que fuerza al compilador a verificar que `Tipo` satisface `Interface`. Si algún método falta, el error aparece en compilación, no en runtime.
