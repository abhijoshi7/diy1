package main

import (
	"database/sql"
	"fmt"
)

func getList(db *sql.DB, id int) ([]int, error) {
	rows, err := db.Query("SELECT product_id from available where store_id = $1", id)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	var pro []int
	fmt.Println(rows)
	for rows.Next() {
		var p temp
		if err := rows.Scan(&p.product_id); err != nil {
			return nil, err
		}
		pro = append(pro, p.product_id)
	}

	return pro, nil
}

func (p store) addProduct(db *sql.DB, id int) error {
	for _, pi := range p.Product_id {
		db.QueryRow(
			"INSERT INTO available(store_id, product_id, is_available) VALUES($1, $2, $3) RETURNING store_id",
			id, pi, p.Is_available).Scan(pi)
	}
	return nil
}
