 package main
import "fmt"

func main() {
  //  representing string with `  `    
    message := `I love Go Programming`
    Table_Name := "Employees"
    name := "HI"
    Loc  := "HYD"
    tsql := fmt.Sprintf("INSERT INTO %s (name,Location) VALUES (%s,%s)",Table_Name,name,Loc)
    fmt.Printf("%s",tsql)
    fmt.Println(message)
}
