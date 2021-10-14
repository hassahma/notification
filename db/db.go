package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"github.com/lib/pq"
	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

/*
CREATE TABLE customer (
id SERIAL,
customer_id TEXT,
notification_url TEXT,
active BOOL,
creation_date timestamp,
PRIMARY KEY(customer_id, notification_url),
CONSTRAINT no_duplicate_records UNIQUE (customer_id, notification_url),
CONSTRAINT no_duplicate_id_for_foreign_key_refs UNIQUE (id)
);

CREATE TABLE notification (
id SERIAL PRIMARY KEY,
customer INT,
customer_key text,
payload TEXT,
status TEXT,
incoming_order INT,
creation_date timestamp,
processsed_date timestamp,
Foreign KEY(customer) references customer(id)
);

*/
func SaveCustomer(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	vars :=  r.URL.Query()

	customerId := vars["customerId"]
	url := vars["url"]
	fmt.Println("url:", url)
	fmt.Println("customer id:", customerId)
	fmt.Println("request body:", r.Body)

	sqlStatement := `
					INSERT INTO customer (customer_id, notification_url, active)
					VALUES ($1, $2, $3)
					RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, customerId, url, false).Scan(&id)
	if err != nil {
		log.Println("DB error: ", err)
		pqErr := err.(*pq.Error)
		log.Println(pqErr.Code)
		if pqErr.Code == "23505" {
			http.Error(w, "Bad request! - The notification of same type already exists for this customer.", 400)
		} else {
			http.Error(w, "Bad request! - Please check the logs for details.", 400)
		}
	}
	fmt.Println("New record ID is:", id)
}
