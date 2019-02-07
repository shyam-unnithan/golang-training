## Testing

1. Create a test suite which ends with _test.go
2. Write a test function with func TestXXX(t *testing.T)
3. Run the test using
    ```
    go test
    ```
4. Run test using verbose
    ```
    go test -v
    ```
5. Run with coverage details
    ```
    go test -v -cover
    ```
6. Example Test function should start with name Example e.g. ExampleSum()

7. Run with benchmark
    ```
    go test -v -bench
    ```

8. Package net/http/httptest can be used to unit test http applications
* ResponseRecorder: An implementaiton of ResponseWriter that records the mutations made on the Response so that you can verify the behavior
* Server:  A test web server only for running unit test

### Behavior Driven Development

#### Packages
* Ginkgo
* Gomega

Generating the bootstrap to run the specs for the directory and then creating a spec file. Then run generate for each of the .go file in the directory
```
$ cd path/to/books
$ ginkgo bootstrap
$ ginkgo generate book
```

### Mock Object libraries
Officially supported by the GoLang team gomock is a mock object library/tool which can be used to generate dummy functions for your interfaces by traversing through your code
* [gomock](https://github.com/golang/mock)
