

CREATE TABLE IF NOT EXISTS books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    descripcion TEXT NOT NULL,
    precio INTEGER NOT NULL,
    cantidad INTEGER NOT NULL,
    random INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- DML statements [Dummy data]

-- GENERADO CON SCRIPT 
INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
     VALUES ('El Hobbit', 'Aventura épica de Tolkien', 15, 10, 12345, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
    VALUES ('Cien años de soledad', 'Novela maestra de García Márquez', 12, 15, 67890, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
    VALUES ('El principito', 'Cuento clásico para todas las edades', 8, 20, 24680, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
    VALUES ('1984', 'Distopía de Orwell', 10, 18, 13579, DATETIME('now'), DATETIME('now'));

INSERT INTO Books (nombre, descripcion, precio, cantidad, random, created_at, updated_at)
    VALUES ('El Señor de los Anillos', 'Trilogía de fantasía épica', 25, 12, 56789, DATETIME('now'), DATETIME('now'));
