# TEMPLATE FOR GENRATION OF REST API WITH MYSQL DB

The Project Aims to Devlop Genric Template for Creation of REST API with MYSQL DB Integration. The project includes MYSQL DB integration. The Binary File in the Project Genrates the Request ,Response ,Error & CRUD Handlers You just need to Define a struct type for your API and Run the Reform Command in order to Build Your Complete API.

## Steps you need to follow
For Getting Started You Need to have Go 1.10+. Install or update the following Packages
* [Reform](https://github.com/go-reform/reform) - A better ORM for Go, based on non-empty interfaces and code generation
* [go-chi](https://github.com/go-chi/render) - easily manage HTTP request / response payloads of Go HTTP services 
* [packr](https://github.com/gobuffalo/packr) - The simple and easy way to embed static files into Go binaries.



1.  Install or update `reform` package, `reform` and `reform-db` commands
    ```
    go get -u gopkg.in/reform.v1/...
    
2.  Create a Table in your DataBase with all the Required Fields and a Field named id which will also be the pk.

3(a). Use `reform-db` command to generate models for your existing database schema.This command will Auto generate the Struct       
      corresponding to your Table Fields. In Case For example:
    ```
    reform-db -db-driver=mysql -db-source="root:root@tcp(localhost:3306)/DBNAME" init -gofmt=false    ```

3(b). In Case You Want to write  your own `struct` representing a table or view row.Rename the directory `expenses` to your Struct name .For eg for employee struct your package should be employees. and store it in file ` struct.go`:
      ```
    import (
	      "time"
        )

    //go:generate reform
    //reform:ExpenseTable
    type Expense struct {
	  ID          int32     `reform:"id,pk"`
	  Description string    `reform:"description"`
	  Type        string    `reform:"type"`
	  Amount      float64   `reform:"amount"`
	  CreatedOn   time.Time `reform:"created_on"`
	  UpdatedOn   time.Time `reform:"updated_on"`
    }  ```
  `  
   Magic comment `//reform:ExpenseTable` links this model to `ExpenseTable` table in SQL database which you created in step2.  
   The first value in field's `reform` tag is a column name. `pk` marks primary key.
   Use value `-` or omit tag completely to skip a field.
   Use pointers (recommended) or `sql.NullXXX` types for nullable fields.

4. Run `reform [package or directory]` or `go generate [package or file]`. This will create 'struct_reform.go`
   in the same package with type `ExpenseTable` and methods on `Expense`.

5. In the `crud-op.gotpl` file you need to provide the DBNAME                                                                                                                                                           
   `d,err:=sql.Open("mysql","root:root@tcp(localhost:3306)/DBNAME?charset=utf8&parseTime=True")` .

6. Now Execute the Main file with the command. `./main -S="Expense" ` The S flag here is used for passing the name of struct.This
   Command will generate `crud.go request.go response.go ` in the expenses directory.

7. Now you just need to call `expenses.Init()` . You can do this by rebuilding the main file.

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why


Give an example
123456yul/

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Reform](https://github.com/go-reform/reform) - A better ORM for Go, based on non-empty interfaces and code generation
* [go-chi](https://github.com/go-chi/render) - easily manage HTTP request / response payloads of Go HTTP services 
* [packr](https://github.com/gobuffalo/packr) - The simple and easy way to embed static files into Go binaries.

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

   * **Ayush Chauhan**   <https://github.com/arc1999>
   * **Shrikar Vaitala** <https://github.com/shrikar007>
   * **Pawas Seth**      <https://github.com/sethpawas>



## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* reform
* go-chi
* testing
