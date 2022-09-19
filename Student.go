package main

import (
    _ "github.com/microsoft/go-mssqldb"
    "database/sql"
    "context"
    "log"
    "fmt"
)

var db *sql.DB
var server = "studentsql.database.windows.net"
var port = 1433
var user = "muthyala"
var password = "Workhard@7777"
var database = "sqldb"

func main() {
    // Build connection string
    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
        server, user, password, port, database)

    var err error
    // Create connection pool
    db, err = sql.Open("sqlserver", connString)
    if err != nil {
        log.Fatal("Error creating connection pool: ", err.Error())
    }
    ctx := context.Background()
    err = db.PingContext(ctx)
    if err != nil {
        log.Fatal(err.Error())
    }
    fmt.Printf("Connected!\n")

    // Create employee
// CreateEmployee("Jake", "United States")
  // fmt.Printf("Inserted Data Successfully...") 

     // Read employees
    count, err := ReadEmployees()
    if err != nil {
        log.Fatal("Error reading Employees: ", err.Error())
    }
    fmt.Printf("Read %d row(s) successfully.\n", count)

    // Update from database
    updatedRows, err := UpdateEmployee("1.00", "NAGA")
    if err != nil {
        log.Fatal("Error updating Employee: ", err.Error())
    }
    fmt.Printf("Updated %d row(s) successfully.\n", updatedRows)

    // Delete from database
    deletedRows, err := DeleteEmployee("JAVA")
    if err != nil {
        log.Fatal("Error deleting Employee: ", err.Error())
    }
    fmt.Printf("Deleted %d row(s) successfully.\n", deletedRows)
   
}

// CreateEmployee inserts an employee record
func CreateEmployee(name,location string) {
     db.Query(`INSERT INTO Employees (name,Location)
    VALUES ($1,$2);`)
    fmt.Print("data inserted successfully...")
}

// ReadEmployees reads all employee records
func ReadEmployees() (int, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        return -1, err
    }

    tsql := fmt.Sprintf("SELECT name, Location FROM Employees;")

    // Execute query
    rows, err := db.QueryContext(ctx, tsql)
    if err != nil {
        return -1, err
    }
    defer rows.Close()
    var count int
    // Iterate through the result set.
    for rows.Next() {
        var name, location string
        // Get values from row.
        err := rows.Scan(&name, &location)
        if err != nil {
            return -1, err
        }
        fmt.Printf("Name: %s, Location: %s\n", name, location)
        count++
    }

    return count, nil
}
// UpdateEmployee updates an employee's information
func UpdateEmployee(name string, location string) (int64, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        return -1, err
    }

    tsql := fmt.Sprintf("UPDATE Employees SET Location = @location WHERE name = @name")

    // Execute non-query with named parameters
    result, err := db.ExecContext(
        ctx,
        tsql,
        sql.Named("Location", location),
        sql.Named("name", name))
    if err != nil {
        return -1, err
    }

    return result.RowsAffected()
}

// DeleteEmployee deletes an employee from the database
func DeleteEmployee(name string) (int64, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        return -1, err
    }

    tsql := fmt.Sprintf("DELETE FROM Employees WHERE name = @Name;")

    // Execute non-query with named parameters
    result, err := db.ExecContext(ctx, tsql, sql.Named("Name", name))
    if err != nil {
        return -1, err
    }

    return result.RowsAffected()
}