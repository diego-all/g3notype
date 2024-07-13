-- DDL statements
    CREATE TABLE IF NOT EXISTS products (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(64) NULL,
	description VARCHAR(200) NULL,
	price INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT DATETIME,
	updated_at TIMESTAMP NOT NULL
	);

-- DML statements [Dummy data]
INSERT INTO products (name, description, price, created_at, updated_at)
    VALUES ('Teléfono móvil', 'Smartphone de última generación', 799, DATETIME('now'), DATETIME('now'));

INSERT INTO products (name, description, price, created_at, updated_at)
    VALUES ('Camiseta', 'Camiseta de algodón', 20, DATETIME('now'), DATETIME('now'));

INSERT INTO products (name, description, price, created_at, updated_at)
    VALUES ('Sartén antiadherente', 'Sartén para cocinar', 35, DATETIME('now'), DATETIME('now'));

INSERT INTO products (name, description, price, created_at, updated_at)
    VALUES ('Balón de fútbol', 'Balón oficial de la FIFA', 50, DATETIME('now'), DATETIME('now'));

INSERT INTO products (name, description, price, created_at, updated_at)
    VALUES ('Muñeca', 'Muñeca de peluche para niños', 15, DATETIME('now'), DATETIME('now'));