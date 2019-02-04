## Type System:

    1. Primitive Types
    2. Composite Types: Array, Slice, Map
    3. User-defined Types: struct, interface

### Struct

``` 
package main
import (
    "fmt"
    "time"
)

type Person struct {
    FirstName, LastName string
    Dob                 time.Time
    Email, Location     string
}

//A person method with value receiver
funct (p Person) PrintName() {
    fmt.Printf(format: "\n%s %s\n", p.FirstName, p.LastName)
}

//A person method with value receiver
func (p Person) PrintDetails() {
    fmt.Prtin(format: "[Date of Birth: %s, Email: %s, Location: %s]\n", p.Dob.String())
}

func main() {
    var p Person
    p.FirstName = "Rob"
    p.LastName = "Pike"
    p.Dob = time.Date(year: 1957, time.February, day:17, hour:0, min:0, sec:0, nsec:0)
    p.Email ="pike@gmail.com"
    p.Location = "California"
}

p1 := Person {
    FirstName: "Shyam",
    LastName: "Unnithan",
    Dob: time.Date(year:1977, time.March, day:05, hour:0, min:0, sec:0, nsec: 0),
    Email: "shyam.unnithan@gmail.com",
    Location: "Bangalore",
}

p.PrintName()
p.PrintDetails()
p1.PrintName()
p1.PrintDetails()
```
####  Type Embedding
```
type Developer struct {
    Employee //type embedding
    Skills []string
}
```