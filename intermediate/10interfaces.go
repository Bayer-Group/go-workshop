package main

import "fmt"

type dbTransaction interface {
	query(string, ...interface{}) (string, error)
	commit() error
}

func main() {
	// TODO write a struct with methods that implement the `dbTransaction` interface.
	// 	 on `query()` return some string as if you called your database
	// TODO invoke `retrieveData` with an instance of your transaction and an id for your query
	// TODO print the result of `retrieveData` to console
}

func retrieveData(tx dbTransaction, id string) string {
	defer tx.commit()

	result, err := tx.query("select * from data where id = $1", id)
	if err != nil {
		fmt.Println("error encountered during query:", err.Error())
	}

	return result
}
