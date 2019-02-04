## Concurrency

Two primitives are used for implementing concurrency

1. goroutines
2. channels

### goroutines
* A goroutine is a function that is running concurrently

* Goroutines can be launched by placing the go keyword before a function

* Created and managed by Go runtime
    ```
    go foo()
    ```

* Executing based on a technique called m:n Scheduler (Maps m goroutines (goroutines are kind of green threads) to nOS threads)

 * Go scheduler is a m:n scheduler Multiplexes m go routines on nOS threads
Scheduler is responsible to schedule goroutines on limited os thread

### channels
* Channels provide a simple mechanism for goroutines to communicated and a powerful construct to build sophisticated concurrency patterns.

* #### Types of Channels
    1. UnBuffered Channel
       - Syncrhonous communication
       - Blocking Sending and receiving
    2. Buffered Channel
        - Asynchronous communication
        - Stores upto capacity elements supports FIFO

Declaring channels
```
var ch chan <Datatype to be used to communicate>
```
e.g
```
var ch chan string
```
Channel operator <- is used for send operation or -> receive operation. 
```
fmt.Println(<-counter)
val, ok := <-counter
if ok {
    fmt.Println(val)
}
```
If the value ok is false which means channel is closed. You will not be able to send data on a closed channel, but receive can happen until the data on the channel is consumed.

The close function closes channel for sending any further data


```
close (channel)


```

### WaitGroup
* Used to wait for goroutines to finish
    ```
    wg.Add(delta: 2)
    ...
    ...
    wg.Wait()
    ```
* Call Waitgroups Done to tell goroutines is completed
    ```
    defer wg.Done()
    ```
* Defer statements are always executed at the end even if program crashes

