# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalación

Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/eliasjurado/greetings
```

## Uso

Ejemplo de cómo utilizar el paquete en tu código:

```go
package main

import (
	"fmt"
	"github.com/eliasjurado/greetings"
)

func main() {
	message, err := greetings.Hello("Elias")
	if err != nil {
		fmt.Println("Ocurrió un error:", err)
		return
	}
	fmt.Println(message)
}
```

Este ejemplo importa el paquete `github.com/eliasjurado/greetings` y llama a la función `Hello` para generar un saludo personalizado. Si ocurre un error, se imprime un mensaje de error.
