package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Bill struct {
	Name       string `json:"name"`
	DayOfMonth int    `json:"dayOfMonth"`
	Paid       bool   `json:"paid"`
}

func SaveBill(db *sql.DB, bill Bill) error {
	b, err := GetByName(db, bill.Name)
	if err != nil {
		return err
	}

	if b.Name != "" {
		return fmt.Errorf("Bill with the name \"%s\" already exists", bill.Name)
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
	if errors.Is(err, sql.ErrNoRows) {
		return Bill{}, nil
	}

	if err != nil {
		return Bill{}, err
	}

	return bill, nil
}

func ListBills(db *sql.DB) ([]Bill, error) {
	var bills []Bill

	rows, err := db.Query("select name, day_of_month, paid from bills order by paid")
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

func ListBillsSoonToBePaid(db *sql.DB, maxInterval int) ([]Bill, error) {
	var bills []Bill
	rows, err := db.Query("select name, day_of_month, paid from bills where paid = false and day_of_month >= ? and day_of_month <= ? ",
		time.Now().Day(),
		time.Now().Day()+maxInterval,
	)
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

func RemoveBill(db *sql.DB, name string) (int, error) {
	res, err := db.Exec("delete from bills where name = ?", name)
	if err != nil {
		return 0, err
	}
	rows, err := res.RowsAffected()
	return int(rows), err
}

func Paid(db *sql.DB, name string) (int, error) {
	res, err := db.Exec("update bills set paid = true where name = ?", name)
	if err != nil {
		return 0, err
	}
	rows, err := res.RowsAffected()
	return int(rows), err
}

func Update(db *sql.DB, bill Bill) error {
	_, err := db.Exec("update bills set name = ?, day_of_month = ? where name = ?", bill.Name, bill.DayOfMonth, bill.Name)
	return err
}

func Reset(db *sql.DB) error {
	_, err := db.Exec("update bills set paid = false")
	return err
}
