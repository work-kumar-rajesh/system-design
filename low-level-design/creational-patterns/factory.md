# Factory Design Pattern 


The Factory Design Pattern is a creational pattern that provides a way to create objects without specifying the exact class 
of the object that will be created. The pattern defines an interface for creating objects but leaves the decision of which class
to instantiate to subclasses or concrete implementations.

## The factory pattern is useful when:
You have a family of related objects (with a common interface), and you want to instantiate different objects based on input or
configuration.The object creation is complex and needs to be abstracted away from the client code.
You need to control the instantiation process in a centralized manner (e.g., to ensure a singleton pattern or caching).

## Real-Life Example in Software Systems:
Consider a Payment Gateway system that supports multiple payment methods (like CreditCard, PayPal, and Crypto). You want to provide an interface to process payments, but the exact payment method implementation might vary based on user input or configuration.
Without Factory Design Pattern (Hardcoded Object Creation)

### Without the factory design pattern
```go
package main

import "fmt"

// PaymentProcessor is an interface for different payment methods
type PaymentProcessor interface {
    Pay(amount float64) string
}

// CreditCard is a struct that implements PaymentProcessor
type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f using Credit Card", amount)
}

// PayPal is a struct that implements PaymentProcessor
type PayPal struct{}

func (p *PayPal) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f using PayPal", amount)
}

// Main function demonstrating direct object creation
func main() {
    var processor PaymentProcessor

	// Assume we have some logic to decide which payment method to use
	method := "paypal" // This could come from a config or user input

	if method == "creditcard" {
		processor = &CreditCard{}
	} else if method == "paypal" {
		processor = &PayPal{}
	}

	// Use the selected payment method
	fmt.Println(processor.Pay(100.50))
}
```
Issues in the Above Example:
Tight Coupling: The main function is directly responsible for creating instances of the CreditCard or PayPal. This makes the code 
difficult to extend. Adding a new payment method would require modifying the client code. Scalability Problems: As the number of 
payment methods grows, the condition logic (e.g., if method == "creditcard") will keep increasing, making the code harder to maintain.

### With Factory Design Pattern
```go
package main

import "fmt"

// PaymentProcessor is an interface for different payment methods
type PaymentProcessor interface {
    Pay(amount float64) string
}

// CreditCard is a struct that implements PaymentProcessor
type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f using Credit Card", amount)
}

// PayPal is a struct that implements PaymentProcessor
type PayPal struct{}

func (p *PayPal) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f using PayPal", amount)
}

// PaymentFactory is a factory function that returns the appropriate PaymentProcessor based on the method
func PaymentFactory(method string) PaymentProcessor {
    switch method {
        case "creditcard":
        return &CreditCard{}
        case "paypal":
        return &PayPal{}
        default:
        return nil
    }
}

// Main function demonstrating the Factory pattern
func main() {
    // Get payment method from user input or config
    method := "paypal" // This could come from a config or user input

	// Create the appropriate PaymentProcessor using the factory
	processor := PaymentFactory(method)

	// Use the selected payment method
	if processor != nil {
		fmt.Println(processor.Pay(100.50))
	} else {
		fmt.Println("Invalid payment method")
	}
}
```
Benefits of Using the Factory Pattern:
1. Decoupling Object Creation: The client code (main) no longer directly creates instances of CreditCard or PayPal. It simply calls the factory to obtain the correct instance.
2. Extensibility: If a new payment method like CryptoPayment needs to be added, you can just modify the PaymentFactory function without changing any code in the main program.
3. Separation of Concerns: The responsibility of object creation is separated into the factory function, making the client code cleaner and more focused on business logic.
4. Code Maintainability: Since object creation is centralized, you can more easily update or refactor the object creation process (e.g., adding logging, error handling, or caching mechanisms).

## When to Use the Factory Design Pattern:
1. Object Creation Logic is Complex: If creating objects involves multiple steps, configurations, or parameters, using a factory
    pattern can centralize and simplify the logic.
2. You Want to Decouple Object Creation from Business Logic: The factory pattern allows you to separate the responsibility of creating 
   objects from the logic that uses them.
3. Family of Related Objects: When you have a family of objects with a common interface and the specific class to instantiate depends 
   on certain conditions or configurations.
4. Changing Class Instantiation Logic: If you anticipate changing how or which objects are created (e.g., switching from one class to 
   another based on configuration), the factory pattern provides an easy way to manage that.

## Real-Life Scenario Example: Database Connection Pooling

In a real-world application, you might need to connect to different types of databases like PostgreSQL, MySQL, or SQLite. You don't want to hardcode the connection logic for each type of database in your main code. Instead, you can use a factory pattern to manage the creation of database connections based on the configuration.
Example: Database Connection Factory Pattern

```go
package main

import "fmt"

// DatabaseConnection is an interface that all database connections implement
type DatabaseConnection interface {
    Connect() string
}

// PostgresConnection implements DatabaseConnection
type PostgresConnection struct{}

func (p *PostgresConnection) Connect() string {
    return "Connected to PostgreSQL"
}

// MySQLConnection implements DatabaseConnection
type MySQLConnection struct{}

func (m *MySQLConnection) Connect() string {
    return "Connected to MySQL"
}

// SQLiteConnection implements DatabaseConnection
type SQLiteConnection struct{}

func (s *SQLiteConnection) Connect() string {
    return "Connected to SQLite"
}

// DatabaseConnectionFactory is a factory function that returns a database connection
func DatabaseConnectionFactory(dbType string) DatabaseConnection {
    switch dbType {
        case "postgres":
        return &PostgresConnection{}
        case "mysql":
        return &MySQLConnection{}
        case "sqlite":
        return &SQLiteConnection{}
        default:
        return nil
    }
}

// Main function demonstrating the Factory pattern for database connections
func main() {
    dbType := "postgres" // This could come from a config or user input

	// Create the appropriate database connection using the factory
	conn := DatabaseConnectionFactory(dbType)

	// Use the connection
	if conn != nil {
		fmt.Println(conn.Connect())
	} else {
		fmt.Println("Invalid database type")
	}
}
```
Advantages of the Factory Pattern in the Database Example:
1. Scalability: You can easily add support for new databases by implementing a new DatabaseConnection type and updating the factory function.
2. Centralized Management: All database connection logic is encapsulated in the factory, which simplifies maintenance.
3. Configurable: You can easily switch between different database types by changing the configuration or input, without modifying the core application logic.

## Conclusion:
The Factory Design Pattern is a powerful tool for managing object creation in a way that decouples it from the rest of the 
application. It improves flexibility, scalability, and maintainability. By centralizing the creation logic, it allows for easier 
modifications and extensions of the application without changing the core business logic.

## Use the Factory pattern when:
1. You need to create objects from a family of related types.
2. The creation logic is complex or changes based on input.
3. You want to decouple object creation from application logic, leading to cleaner, more maintainable code.