## Suma de Números en Go
Este proyecto es una aplicación que permite sumar dos números ingresados por el usuario mediante un formulario en un navegador web.

## Descripción
La aplicación utiliza el lenguaje de programación Go y la biblioteca estándar para:

* Crear un servidor HTTP.
* Renderizar una plantilla HTML para el formulario.
* Realizar la suma de dos números ingresados por el usuario.
* Al enviar los datos, el servidor procesa la solicitud, realiza la suma y muestra el resultado en la misma página.
## Estructura del Proyecto

    main.go       // Archivo principal con el código fuente de la aplicación
## Requisitos
    Tener instalado Go (versión 1.16 o superior).
    Navegador web.
    Editor de texto o IDE (opcional, para ver o editar el código).
    Git (si se desea clonar el repositorio).
    Instalación y Ejecución Local
1. Clonar el Repositorio
    Si el proyecto está alojado en un repositorio de Git, sigue estos pasos para clonarlo:

    Abre una terminal.

    
    git clone https://github.com/Cristina-Colcha/SumaGo.git 
    Navega al directorio del proyecto:

    cd <NOMBRE_DEL_REPOSITORIO>
2. Ejecutar el Proyecto
Asegúrate de estar en el directorio del proyecto.

Inicia el servidor con el comando:

    go run main.go
Si todo se ejecuta correctamente, deberías ver un mensaje en la terminal indicando que el servidor está corriendo:

    Servidor escuchando en http://localhost:8080
Abre un navegador web e ingresa la siguiente URL:

    http://localhost:8080
Interactúa con el formulario ingresando dos números y presionando el botón Sumar para ver el resultado.


## Licencia
Este proyecto es de código abierto y puede ser utilizado con fines educativos o personales.