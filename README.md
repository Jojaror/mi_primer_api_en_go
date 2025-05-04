# Mi Primer API en Go
Una guía paso a paso para crear aplicaciones web con Go

## Inspirado en "Let's Go"

Bienvenido a esta guía simplificada para Go. En este proyecto aprenderás los cimientos para desarrollar aplicaciones web sólidas, combinando enfoques prácticos en código con descripciones detalladas.

### Que encontrarás en este proyecto?

1. **Foundations**: Inicio del proyecto con el "Hola Mundo" y conceptos básicos de Go.  
2. **Configuration and error handling**: Configuración inicial y manejo de errores de forma robusta.  
3. **Database-driven responses**: Respuestas dinámicas con integración a bases de datos.  
4. **Dynamic HTML templates**: Creación y uso de plantillas HTML flexibles.  
5. **Middleware**: Implementación de capas intermedias para el procesamiento de peticiones.  
6. **Advance routing**: Manejo avanzado de rutas para solicitudes complejas.  
7. **Stateful HTTP**: Gestión del estado en aplicaciones HTTP para interacciones continuas.  
8. **Security improvements**: Implementación de mejoras de seguridad para salvaguardar la aplicación.  
9. **User authentication**: Estrategias robustas para la autenticación de usuarios.  
10. **Using request context**: Uso efectivo de contextos en las peticiones para un control detallado.  
11. **Optional Go features**: Exploración de características opcionales que enriquecen la funcionalidad.  
12. **Testing**: Prácticas y herramientas para asegurar la calidad y confiabilidad del código.  
13. **Extras**: Finalización del proyecto.

### Puntos de acceso

| Método | Ruta               | Manejador       | Descripción                          |
|--------|--------------------|-----------------|--------------------------------------|
| GET    | /                  | home            | Muestra la página principal
| GET    | /snippet/view/:id  | view            | Muestra la página de un objeto en especifico
| GET    | /snippet/create    | create          | Muestra la página de formulario de creacion de objeto
| GET    | /user/signup       | userSignup      | Muestra la página de regristro de usuario
| GET    | /user/login        | userLogin       | Muestra la página de inicio de sesión
| GET    | /static/           | http.FileServer | Sirve archivos estaticos especificos
| POST   | /snippet/create    | createPost      | Crea un nuevo objeto
| POST   | /user/signup       | userSignupPost  | Crea un nuevo usuario
| POST   | /user/login        | userLoginPost   | Autentica al usuario
| POST   | /user/logout       | userLogoutPost  | Cierra la sesión del usuario


___________________________________________________________________________________
> **Nota:** Este proyecto se inspira en el libro **Let's Go 2nd edition** de Alex Edwards, y se enriquece con aportes y mejoras propias. Cada sección de esta guía representa un paso en el desarrollo progresivo de la API, que se documentará y acompañará de material explicativo para facilitar el aprendizaje. Este proyecto debe ser visto como complemento y nunca como reemplazo al mencionado libro.
