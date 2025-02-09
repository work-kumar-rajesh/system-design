# Singleton Design Pattern

The Singleton design pattern is a type of Creational Pattern that ensures a class has only one instance and provides a global
point of access to that instance. The key idea is that no matter how many times you request an instance of the class, you always
get the same one.

	
This is useful when you need a single point of control or coordination throughout the system, such as managinga single connection
to a database, logging, or configuration settings.Key Characteristics of Singleton Pattern:
1. Single Instance: Only one instance of the class exists.
2. Global Access: The instance is accessible globally throughout the application.
3. Lazy Initialization: The instance is created only when it is needed, which helps optimize performance.
4. Thread-Safety: In multi-threaded applications, ensuring that only one instance is created even when multiple
   threads attempt to create it simultaneously.

	
## Problem without Singleton :
Imagine a scenario where we need a logger. If each component or part of the application creates a new instance of the logger,
it could lead to problems like excessive resource usage, inconsistent logging output, and difficulty managing log states.A
Singleton can help by ensuring that only one logger instance is used across the application.
Basic Structure of Singleton Pattern


## The Singleton pattern typically consists of the following:
1. Private Constructor: The class’s constructor is made private so that instances can’t be created outside the class.
2. Static Instance: A static variable that holds the single instance of the class.
3. Public Access Method: A public method (usually GetInstance()) to provide access to the instance.

	
## Singleton Pattern in Go (Golang)
Golang does not have explicit support for access control modifiers (like private or protected), but we can achieve the desired 
effect using package-level variables and functions. In Go, a typical Singleton pattern can be implemented using a combination of 
sync.Once to ensure that the initialization happens only once and a global variable to hold the instance.
	
	
## Example Implementation of Singleton Pattern in Golang
Step 1: Create a Singleton Logger
Let’s create a simple Logger that ensures only one instance is created.

```go
package logger

import (
   "fmt"
   "sync"
)

//singelton logger instance
type Logger struct {
   level string
}

// Declare a private global variable to hold the single instance of Logger
var instance *Logger

// Use sync.Once to ensure the initialization is done only once
var once sync.Once

// GetInstance returns the singleton instance of the Logger
func GetInstance() *Logger {
   // The function inside the Once.Do will be executed only once
   once.Do(func() {
      instance = &Logger{
         level: "INFO", 
      }
   })
   return instance
}

// Log - a simple method to log messages
func (l *Logger) Log(message string) {
   fmt.Printf("[%s] %s\n", l.level, message)
}
```

Step 2: Using the Singleton in Main Function

Now, let's demonstrate how to use the Logger singleton in a program.
```go
package main

import (
   "example/logger" // Import the package where logger is defined
)

func main() {
// Get the singleton logger instance
   logger1 := logger.GetInstance()
   logger1.Log("This is the first log message.")

// Get another reference to the same logger instance
   logger2 := logger.GetInstance()
   logger2.Log("This is the second log message.")

// Check if both logger references point to the same instance
   if logger1 == logger2 { 
	   fmt.Println("Both logger1 and logger2 are the same instance.")
   }
}
```

Step 3: Output

[INFO] This is the first log message.
[INFO] This is the second log message.
Both logger1 and logger2 are the same instance.


## Explanation of the Code:
1. Global Instance Variable: instance holds the single instance of the Logger.
2. Once and Do: The sync.Once ensures that the initialization of instance happens only once. The Once.Do method guarantees that the 
   code inside it will only execute once, regardless of how many times GetInstance() is called.
3. Log Method: The Log method prints the message with a log level, in this case, "INFO".
4. global Access: By calling GetInstance(), you can access the same Logger instance.


## Advantages of Singleton in Go
1. Lazy Initialization: The instance is created only when it's first needed, not when the program starts.
2. Thread Safety: sync.Once ensures thread-safe initialization, so multiple goroutines can safely call GetInstance() without causing race 
   conditions.
3. Global Access Point: A single point of access to the instance of the logger (or any other resource).


## Variations of Singleton in Go

1. Thread-Safe Singleton with sync.Mutex  :
   In this version, the sync.Mutex lock prevents multiple goroutines from creating separate instances of the Logger.This
   approach is a bit more manual compared to sync.Once.

```go
package logger

import (
   "fmt"
   "sync"
)

type Logger struct {
	level string
}

var instance *Logger
var mu sync.Mutex

func GetInstance() *Logger {
   mu.Lock()
   defer mu.Unlock()
   if instance == nil {
   instance = &Logger{level: "INFO"}
   }
   return instance
}

func (l *Logger) Log(message string) {
   fmt.Printf("[%s] %s\n", l.level, message)
}
```
	

2. Eager Initialization:
   In this version, the Logger instance is created as soon as the package is imported, meaning the instance is available immediately,
   but still ensures there’s only one instance.

```go
package logger

import "fmt"

type Logger struct {
   level string
}

// Eagerly initialize the instance when the package is imported
var instance = &Logger{level: "INFO"}

func GetInstance() *Logger {
   return instance
}

func (l *Logger) Log(message string) {
   fmt.Printf("[%s] %s\n", l.level, message)
}
```


## Conclusion
The Singleton pattern is widely used in scenarios where you need a single point of control or access for an object, 
like in logging systems, configuration settings, and database connections. In Go, you can implement the Singleton pattern 
using sync.Once, sync.Mutex, or even eager initialization, depending on your needs.
In the example above, we have used sync.Once to ensure that the instance is created lazily and is thread-safe. 
You should choose your Singleton approach based on the specific requirements of your project,
such as lazy initialization vs. eager initialization and the need for concurrency handling.

## Simple Example
```go
package creational
type singleton struct{
   count int
}
var instance *singleton
func GetInstance() *singleton {
   if instance == nil {
      instance = new(singleton)
   }
   return instance
}
func (s *singleton) AddOne() int {
   s.count++
   return s.count
}
```
