# Go_logger

Logger function for golang

## Usage

### Normal Usage

Create the logger object with the base severity you need, and use the methods to call your prints

``` Golang
...
// Create the logger object
var log = logger.Logger{Level: 3, History: []string{}} // Level is the verbose level to use, sitory is a variable to store the logs for the panic handler

// print info
log.Info("Some info here")

// print Fatal error -> it exists the program with the error code specified
log.Fatal("error description", 123) // error code 123

// print debug
log.Debug("some debug info here") // Text only
log.Debug("some debug info here", someobject1, someobject2, someobjectN) // Text and object print

```

### Panic Handler

Panic handler serves to handle all unexpected errors and conveniently place them on a text file for reading and debug purposes. 

``` Golang
... 
// Create the logger object
var log = logger.Logger{Level: 3, History: []string{}}

func main() {

// Add panic hadler to this function
defer log.PanicHandler()

// print debug
log.Debug("some debug info here") // Text only
...
}
```

---

## V1.2.0

- Added capability to print logs on panic events 
  - On panic the code will create a panic.txt file on the executable location, this file contains the loggs history and the stack trace (does now show varibles)
- Fixed the missing Camel Case instances on upper objects and left snake case for variables (easier to read)
- Objects in debug and trace methods are now printed with their name / Value keys (before it was only values)

## V1.1.2

- Added capability to print objects on debug and trace method
- Added code to detect the caller function name 
  - Function names will be printed automatically on the debug, error and trace methods
