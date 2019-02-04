### Packages

#### 1. Package main

it compiles into executables 

#### 2. Library Packages

Any other name than main to reuse code. If you are looking for reusable code then use library packages

if directory structure is
src -> anz -> banking -> sqlstre -> 10 .go files

The name of package should be

    package sqlstore
###### Note: Typically use small case characters to define package name

Use the following command to compile the library and install it in the binary location
```
go install
```

This will compile the code and create a binary sqlstore under

pkg -> anz/banking -> sqlstore

###### Note: You don't have to do this step if you have referenced it in a main and you are building that main

#### Importing Package
You can import this package by using

```
import (
    "anz/banking/sqlstore"
)
```

Using this package
```
func main(){
    sqlstore.<object to be referenced>
}

```
You can also import with an alias name.

```
import (
    store "anz/banking/sqlstore"
)
```

If you use alias name . then you don't need any alias name.

``` 
func main(){
    Connect()
}
```

Sometimes you also use _ as the alias. Usually this is done with Vendor packages where you want the init() of the packages to be triggered

e.g.
```
database/sql
_ github.com/lib/pq //Initialize SQL Driver for PostGres
```

###### Note: this is only used mostly while writing unit test

You can use this package by using alias

```
func main() {
    store.<object to be referenced>
}
```

Any function, struct, variable named starting with Capital Letters then it is exported and available to be used any other package

Any function, struct, variable named starting with Lower Case Letters are not exported and will only be available in the same package

#### Init function
Before main runs the init function will run at the package level

```
func init() {
    // This will be automatically invoked before running main
}
```

###### Note: Not recommended to rely too much on init. May be required in special cases

