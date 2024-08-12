package models

import (
	"context"
	"time"
)

type Product struct {
 	Id	int	`json:"id"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Price	int	`json:"price"`
	Quantity	int	`json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
 	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) Insert(product Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int

		stmt := `insert into products (name, description, price, quantity, created_at, updated_at)
 	values ($1, $2, $3, $4, $5, $6) returning  id`

    err := db.QueryRowContext(ctx, stmt,
	product.Name,
	product.Description,
	product.Price,
	product.Quantity,
	time.Now(),
	time.Now(),
).Scan(&newID)


	if err != nil {
		return 0, err
	}

	return newID, nil
}



func (p *Product) GetOneById(id int) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

    	query := `select id, name, description, price, quantity, created_at, updated_at from products where id = $1`

	var product Product
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
	&product.Id,	&product.Name,
	&product.Description,
	&product.Price,
	&product.Quantity,

 	&product.CreatedAt,
	&product.UpdatedAt,
)

	if err != nil {
		return nil, err
	}

	return &product, nil
}


func (p *Product) Update(product Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel() // resource leaks

	stmt := `update products set
 	Name = $1,
	Description = $2,
	Price = $3,
	Quantity = $4,
	updated_at = $6
 	where id = $7`

	_, err := db.ExecContext(ctx, stmt,
	product.Name,
	product.Description,
	product.Price,
	product.Quantity,
	time.Now(),
 	product.Id,
)

	if err != nil {
		return 0, err
	}

	return 0, nil
}



func (p *Product) GetAll() ([]*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, price, quantity, created_at, updated_at from products order by name`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product

	for rows.Next() {
		var product Product
		err := rows.Scan(
	&product.Id,	&product.Name,
	&product.Description,
	&product.Price,
	&product.Quantity,

 	&product.CreatedAt,
	&product.UpdatedAt,
)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}



func (p *Product) DeleteByID(id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from products where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}