Tengo la siguiente respuesta de una consulta a gemini:


INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
  VALUES ('Teléfono móvil', 'Smartphone de última generación', 799, 5, 'Modelo A', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
  VALUES ('Camiseta', 'Camiseta de algodón', 20, 1, 'Modelo B', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
  VALUES ('Sartén antiadherente', 'Sartén para cocinar', 35, 1, 'Modelo C', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
  VALUES ('Balón de fútbol', 'Balón oficial de la FIFA', 50, 1, 'Modelo D', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
  VALUES ('Muñeca', 'Muñeca de peluche para niños', 15, 1, 'Modelo E', DATETIME('now'), DATETIME('now'));
## Request JSON

### createBody:

"nombresito": "Teléfono móvil",
"descripcionsita": "Smartphone de última generación",
"precioaquel": 799,
"cantidadparce": 5,
"randomoelo": "Modelo A"
### updateBody:

"nombresito": "Camiseta",
"descripcionsita": "Camiseta de algodón",
"precioaquel": 20,
"cantidadparce": 1,
"randomoelo": "Modelo B"

Requiero la forma de extraer en 3 variables diferentes lo siguiente:

1. Los 5 inserts tal cual estan en un mismo string.
2. La parte de createBody completa y en el mismo orden que entrega gemini y sin tener en cuenta created_at ni updated_at.
3. La parte de updateBody completa y en el mismo orden que entrega gemini y sin tener en cuenta created_at ni updated_at. Garantizando que la data dummy de updateBody sea diferente a la de createBody, es decir se debe hacer esta validacion. En caso contrario el algoritmo esta erroneo por que gemini entrega informacion diferente.

Se debe garantizar que el algoritmo de extraccion sea generico, no siempre vendran las mismas columnas ni la misma cantidad de columnas. Pueden venir diferentes.

Podrias generar el script por favor, recuerda solo enfocarte en extraer.
Es decir: 
Extraer los 5 inserts utilizando una funcion.
Extraer lo que viene debajo de createBody en otra funcion (utilizar la etiqueta createBody:)
Extraer lo que viene debajo de updateBody en otra funcion (utilizar la etiqueta updateBody:)
