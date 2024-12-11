// Package id provides functionality to generate unique IDs based on Unix time and a sequencer.
// The generated ID is a signed number that can be verified using a validator.
// This package supports various validators and sequencers for flexible usage.
//
// # Overview
//
// The id package offers:
// - Unique ID generation using Unix time and sequencers.
// - Customizable validators for signing and verifying IDs.
// - Flexible sequencer implementations, including in-memory and database-backed sequencers.
//
// # Installation
//
// To install the package, use:
//
//	go get github.com/ceebydith/id
//
// # Usage
//
// Example of generating an ID with a range sequencer and Luhn validator:
//
//	package main
//
//	import (
//	    "fmt"
//	    "github.com/ceebydith/id"
//	)
//
//	func main() {
//	    sequencer := id.RangeSequencer(0, 9999)
//	    validator := id.LuhnValidator()
//	    generator := id.New(sequencer, validator)
//
//	    id, err := generator.Generate()
//	    if err != nil {
//	        panic(err)
//	    }
//
//	    fmt.Printf("Generated ID: %d\n", id)
//	}
//
// Example of using a PostgreSQL sequence as the sequencer:
//
//	package main
//
//	import (
//	    "database/sql"
//	    "fmt"
//	    _ "github.com/lib/pq"
//	    "github.com/ceebydith/id"
//	)
//
//	type pgSequencer struct {
//	    db        *sql.DB
//	    sequence  string
//	}
//
//	func (s *pgSequencer) Generate() (int64, error) {
//	    var value int64
//	    err := s.db.QueryRow(fmt.Sprintf("SELECT nextval('%s')", s.sequence)).Scan(&value)
//	    if err != nil {
//	        return 0, err
//	    }
//	    return value, nil
//	}
//
//	func NewPGSequencer(db *sql.DB, sequence string) id.Sequencer {
//	    return &pgSequencer{
//	        db:       db,
//	        sequence: sequence,
//	    }
//	}
//
//	func main() {
//	    connStr := "user=username dbname=mydb sslmode=disable"
//	    db, err := sql.Open("postgres", connStr)
//	    if err != nil {
//	        panic(err)
//	    }
//	    defer db.Close()
//
//	    sequencer := NewPGSequencer(db, "my_sequence")
//	    validator := id.LuhnValidator()
//	    generator := id.New(sequencer, validator)
//
//	    id, err := generator.Generate()
//	    if err != nil {
//	        panic(err)
//	    }
//
//	    fmt.Printf("Generated ID: %d\n", id)
//	}
package id
