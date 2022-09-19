// Go connection Sample Code:
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
	fmt.Printf("Connected!")

    /* Create Table 
	val:=Create_Table("STUDENT_RAMA")
	if(val == 0){
        log.Fatal("Error Creating Table:")
    } else {
	fmt.Printf("Table Created Successfully...!") 
	}   */


	/* DROP  TABLE from database
    rows,err := drop_Table("STUDENT_RAMA")
    if err != nil {
        log.Fatal("Error deleting Table: ", err.Error())
    } else {
		 fmt.Printf("Dropped Table successfully.\n",rows)
	}  */

	// Insert_Data
    updatedRows, err := Insert_Data("Employees","MANI","PRASAD")
    if err != nil {
        log.Fatal("Error updating Employee: ", err.Error())
    } else {
		fmt.Printf("Updated %d row(s) successfully.\n", updatedRows)
	} 
    
}

// Drop Table
func drop_Table(name string) (int64, error) {
    ctx := context.Background()
    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        return -1, err
    }

    tsql := fmt.Sprintf("DROP TABLE IF EXISTS %s ;",name)
    // Execute non-query with named parameters
    result, err := db.ExecContext(ctx, tsql, sql.Named("Name", name))
    if err != nil {
        return -1, err
		}
	return result.RowsAffected()	
	}

func Create_Table(Table_Name string)int{
	ctx := context.Background()
    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        return 0
    }
    tsql := fmt.Sprintf("CREATE TABLE %s (NAME VARCHAR(50),RNO INT) ;",Table_Name)
    // Execute non-query with named parameters
    result, err := db.ExecContext(ctx, tsql, sql.Named("Name", Table_Name))
    if err != nil {
        return 0
		}
    fmt.Printf("res :",result)
	return 1
	}

func Insert_Data(Table_Name string, name string ,location string) (int64, error) {
    ctx := context.Background()
    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        return -1, err
    }
	//fmt.Printf("HI HE",Table_Name)
	tsql := `
      INSERT INTO Employees (Name, Location) VALUES (@Name, @Location);
      select isNull(SCOPE_IDENTITY(), -1);
    `

    //tsql := fmt.Sprintf("INSERT INTO %s (name, Location) VALUES(%s,%s) ;",Table_Name,name,location)
	//fmt.Printf("%s",tsql)

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

func Insert_Data1(Table_Name string,name string, location string) (int64, error) {
    ctx := context.Background()
    var err error

    // Check if database is alive.
    err = db.PingContext(ctx)
    if err != nil {
        return -1, err
    }

    tsql := `
      INSERT INTO Employees (Name, Location) VALUES (@Name, @Location);
      select isNull(SCOPE_IDENTITY(), -1);
    `
    fmt.Printf(Table_Name)
    stmt, err := db.Prepare(tsql)
    if err != nil {
       return -1, err
    }
    defer stmt.Close()

    row := stmt.QueryRowContext(
        ctx,
        sql.Named("Name", name),
        sql.Named("Location", location))
    var newID int64
    err = row.Scan(&newID)
    if err != nil {
        return -1, err
    }

    return newID, nil

}