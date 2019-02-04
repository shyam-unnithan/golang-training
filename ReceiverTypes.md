## Receiver Types
1. Value Receiver
2. Pointer Receiver

#### Value Receiver
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
func (p *Person) PrintName() {
    fmt.Printf(format: "\n%s %s\n", p.FirstName, p.LastName)
}

//A person method with value receiver
func (p *Person) PrintDetails() {
    fmt.Prtin(format: "[Date of Birth: %s, Email: %s, Location: %s]\n", p.Dob.String())
}

//A person method with value receiver
func (p *Person) changeLocation (newLocation string){
    p.Location = newLocation
}

func main() {
    var p Person
    p.FirstName = "Rob"
    p.LastName = "Pike"
    p.Dob = time.Date(year: 1957, time.February, day:17, hour:0, min:0, sec:0, nsec:0)
    p.Email ="pike@gmail.com"
    p.Location = "California"
}

// This is going to pass the address to p1
p1 := &Person {
    FirstName: "Shyam",
    LastName: "Unnithan",
    Dob: time.Date(year:1977, time.March, day:05, hour:0, min:0, sec:0, nsec: 0),
    Email: "shyam.unnithan@gmail.com",
    Location: "Bangalore",
}

p.ChangeLocation(newLocation: "Santa Clara")
p.PrintName()
p.PrintDetails() //This will still print the old details
p1.PrintName()
p1.PrintDetails()
```

#### Pointer Receiver
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
func (p Person) PrintName() {
    fmt.Printf(format: "\n%s %s\n", p.FirstName, p.LastName)
}

//A person method with value receiver
func (p Person) PrintDetails() {
    fmt.Prtin(format: "[Date of Birth: %s, Email: %s, Location: %s]\n", p.Dob.String())
}

//A person method with pointer receiver
func (p *Person) changeLocation (newLocation string){
    p.Location = newLocation
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

p.ChangeLocation(newLocation: "Santa Clara")
p.PrintName()
p.PrintDetails() //This will print thew location
p1.PrintName()
p1.PrintDetails()
```

### Escape Analysis
Go Compiler will go through the code to validate the correct use of reference and value types and in case it finds issues it will change it to proper types for the application to run

``` 
func something() int {
    i = &int
    return *i
}
```
Such cases will be evaluated by Escape Analysis to take a call on the variable types.
The focus of this system is to provide high performance