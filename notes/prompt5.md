En un modelo de datos como el siguiente:

        id | int
		nombre | string
		descripcion| string
		precio | int
		cantidad | int
		random| int
		created_at|DATETIME('now')
		updated_at|DATETIME('now')

Como se conoce a la columna que define la entidad, es decir en este caso el nombre?

Podrias darme la respuesta en español





En un modelo de datos, la columna que define la entidad, proporcionando una descripción única o distintiva, como en este caso "nombre", se conoce como "atributo identificador" o "identificador natural". Aunque no necesariamente es la clave primaria, es un campo que describe o nombra la entidad de manera significativa.


Segun blancarte: 
Table ID vs Natural ID
Libro tiene ISBN, Persona tiene DNI, cedula etc. (Identificador del mundo real, natural ID o identificador natural)
Pueden coexistir en una misma tabla. Hay tablas que tienen el ID de siempre y una columna que puede funsionar con natural id. Si bien esta columna se puede llamar como sea, ya que dependera de cada tabla, la idea es que debe estar marcada como única, para impedir que existe otro registro con el mismo natural id, de la misma forma, deberemos crear un índice que  ayude a realizar búsquedas más rápida sobre ese campo.


DUDA DANIEL (segundo campo) En el dominio de las APIs es normal.
TODO DEPENDE!!
Cree pertinente en este punto manejar la segunda columna 8Idnetificador natural como unico) ?
Ej. no pueden haber libros con el mismo nombre, ene ste caso no seria identificadores unicos solo identificadores naturales,
Entonces es mas por tener claro el concepto y nombrarlo.

Daniel es practico o no? Sirve para algo?

Clave natural: Clave de dominio o clave de negocio. (En el modelo relacional) 






Relacion con diagrama de clases vs diagrama de datos

https://blog.codmind.com/table-id-vs-natural-id/






