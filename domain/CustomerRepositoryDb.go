package domain

import (
	"database/sql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	client := d.client
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err := client.Query(findAllSql)
	if err != nil {
		log.Println("error al hacer la query" + err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("error al hacer la query" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, error) {
	byIdSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(byIdSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

	if err != nil {
		log.Fatal("error al ejecutar la query")
		return nil, err
	}

	return &c, nil
}
