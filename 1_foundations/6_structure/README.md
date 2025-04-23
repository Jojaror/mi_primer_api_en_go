# Estructura y organización de tu proyecto

Usa ```package nombre_paquete``` al inicio de cada archivo. Todos los archivos **.go** en un directorio = Mismo paquete. El compilador procesa todos los archivos del paquete juntos, comparten el mismo espacio de nombres. Las importaciones (import) solo se usan para traer código de otros paquetes (ya sean estándar como net/http o de terceros/propios en otros directorios).

Usa la siguiente instrucción para ejecutar nuestra API.

````bash
go run ./cmd/web
````
