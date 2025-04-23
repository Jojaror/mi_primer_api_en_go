# Enrutamiento de peticiones

El **servemux** de **Go** soporta dos tipos de patrones de URL:  
- **Rutas fijas (fixed paths):** No terminan con una barra inclinada final (/) y solo coinciden cuando la URL es exactamente la misma.  
- **Rutas de subárbol (subtree paths):** Terminan con una barra inclinada final y coinciden cuando el inicio de la ruta de la URL se ajusta al patrón definido.

En el **servemux** de **Go**, las rutas fijas se ejecutan únicamente cuando hay una coincidencia exacta, mientras que las rutas de subárbol se activan siempre que la URL comience con el patrón definido.

## Restringiendo el inicio de nuestro patrón de URLs

Para asegurarnos de que se atiendan únicamente las solicitudes a la raíz, podemos emplear el siguiente filtro:

````go
if r.URL.Path != "/" {
  http.NotFound(w, r)
  return
}
````
````go
if r.URL.Path != "/" { ... }
````
Esta es una estructura de control ```if``` (condicional). ```r.URL.Path``` Accede al campo ```URL``` dentro de ```r``` (contiene la información de la URL solicitada por el cliente) y luego al campo ```Path``` Dentro de ```URL```, que es una cadena de texto que representa la parte de la ruta de la *URL* (por ejemplo, /, /contacto, /usuario/123). ```!= "/"``` Aqui comparamos si es exactamente **"no igual a"** la ruta de la *URL* solicitada **"/"**.
````go
http.NotFound(w, r)
````
Si la condición del ```if``` es verdadera (la ruta no es /), se llama a esta función del paquete ```net/http```. Es una forma rápida y estándar de enviar una respuesta *HTTP 404 Not Found* al cliente. Toma el *ResponseWriter (w)* y el *Request (r)* como argumentos.
````go
return
````
```return``` Esta palabra clave detiene la ejecución de la función ```home``` en este punto. Es importante porque si enviamos una respuesta *404*, no queremos continuar y ejecutar el *w.Write(...)* que viene después.
