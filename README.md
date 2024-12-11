# id

`id` is a Golang package that generates unique IDs based on Unix time and a sequencer. The generated ID is a signed number that can be verified using a validator. The package supports different validators and sequencers for flexible usage.

## Features

- **Unique ID Generation**: Combines Unix time and a sequencer to generate unique IDs.
- **Customizable Validators**: Supports various algorithms like Luhn for signing and verifying IDs.
- **Customizable Sequencers**: Allows different sequencer implementations, including range sequencers and database-backed sequencers.

## Installation

```bash
go get github.com/ceebydith/id
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/ceebydith/id"
)

func main() {
    sequencer := id.RangeSequencer(0, 9999)
    validator := id.LuhnValidator()
    generator := id.New(sequencer, validator)

    id, err := generator.Generate()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Generated ID: %d\n", id)
}
```

### Using PostgreSQL Sequence as Sequencer

To use a PostgreSQL sequence as the sequencer, you need a custom implementation of the `Sequencer` interface that interacts with your PostgreSQL database.

```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "github.com/ceebydith/id"
)

type pgSequencer struct {
    db        *sql.DB
    sequence  string
}

func (s *pgSequencer) Generate() (int64, error) {
    var value int64
    err := s.db.QueryRow(fmt.Sprintf("SELECT nextval('%s')", s.sequence)).Scan(&value)
    if err != nil {
        return 0, err
    }
    return value, nil
}

func NewPGSequencer(db *sql.DB, sequence string) id.Sequencer {
    return &pgSequencer{
        db:       db,
        sequence: sequence,
    }
}

func main() {
    connStr := "user=username dbname=mydb sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    sequencer := NewPGSequencer(db, "my_sequence")
    validator := id.LuhnValidator()
    generator := id.New(sequencer, validator)

    id, err := generator.Generate()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Generated ID: %d\n", id)
}
```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the [LICENSE](https://github.com/ceebydith/id/blob/main/LICENSE) file for details.