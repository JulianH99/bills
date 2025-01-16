package services

import (
	"database/sql"
	"fmt"
)

type Bill struct {
	Name       string
	DayOfMonth int
	Paid       bool
}

func SaveBill(db *sql.DB, bill Bill) error {
	b, err := GetByName(db, bill.Name)
	if err != nil {
		return err
	}

	if b.Name != "" {
		return fmt.Errorf("Bill with the name \"%s\" already exists\n", bill.Name)
	}

	_, err = db.Exec(`insert into bills (name, day_of_month, paid) values (?, ?, ?)`, bill.Name, bill.DayOfMonth, bill.Paid)
	if err != nil {
		return err
	}

	return nil
}

func GetByName(db *sql.DB, name string) (Bill, error) {
	var bill Bill

	res := db.QueryRow("select name, day_of_month, paid from bills where name = ?", name)
	err := res.Scan(&bill.Name, &bill.DayOfMonth, &bill.Paid)
	if err != nil {
		return Bill{}, err
	}

	return bill, nil
}

func ListBills(db *sql.DB) ([]Bill, error) {
	var bills []Bill

	rows, err := db.Query("select name, day_of_month, paid from bills")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bill Bill
		err := rows.Scan(&bill.Name, &bill.DayOfMonth, &bill.Paid)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}

	return bills, nil
}
