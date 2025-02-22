# Single Responsibility Principle (SRP)

The Single Responsibility Principle (SRP) is one of the five SOLID principles of object-oriented design. It states that a class should have only one reason to change, meaning that it should have only one responsibility. By ensuring that a class focuses on a single task or functionality, SRP improves code maintainability, readability, testability, and reusability. When a class adheres to SRP, modifications in one part of the system (such as business logic or logging) do not affect other parts.

---

## Detailed Theory

### What is a Responsibility?
- A **responsibility** is a specific functionality or behavior that a class must encapsulate. For example, in an e-commerce application, responsibilities could include processing payments, updating inventory, sending notifications, or logging events.
  
### Why SRP is Important:
- **Separation of Concerns:** By isolating responsibilities, you ensure that changes in one aspect of the system do not ripple through unrelated areas.
- **Maintainability:** Smaller, focused classes are easier to understand, modify, and debug.
- **Testability:** Unit tests become more straightforward when a class has a single purpose.
- **Reusability:** A class with a focused responsibility is more likely to be reused across different parts of an application or in different projects.
- **Reduced Complexity:** SRP helps avoid large classes with tangled logic, making the system easier to navigate.

### Theoretical Example:
Imagine a class that handles both user authentication and logging. If the logging requirements change (e.g., to support multiple log levels or destinations), it might inadvertently affect the authentication logic if they are combined in a single class. SRP would suggest splitting this into separate classes:
- **AuthenticationService:** Handles verifying user credentials.
- **Logger:** Manages logging of events and errors.

---

## Real-Life Example: Order Processing System in an E-Commerce Application

In many real-world applications, such as e-commerce platforms, processing an order can involve multiple responsibilities:
- **Validation:** Ensuring the order data is correct.
- **Payment Processing:** Handling the payment transaction.
- **Inventory Update:** Adjusting inventory levels after a purchase.
- **Notification:** Sending a confirmation email or message.
- **Logging:** Recording the order process for audit purposes.

### Without Applying SRP

In the non-SRP version, we combine all these responsibilities into one large class (`OrderProcessor`). This makes the class complex, hard to maintain, and difficult to test independently.

```go
package main

import (
	"fmt"
	"time"
)

// Order represents an e-commerce order.
type Order struct {
	ID          int
	Customer    string
	Items       []string
	TotalAmount float64
	CreatedAt   time.Time
}

// OrderProcessor handles order processing, including validation, payment, inventory, logging, and notification.
type OrderProcessor struct {
	orders []Order
}

// ProcessOrder processes the given order by validating, processing payment, updating inventory, logging, and sending notifications.
func (op *OrderProcessor) ProcessOrder(order Order) bool {
	// Validation logic
	if order.TotalAmount <= 0 || len(order.Items) == 0 {
		fmt.Println("Order validation failed.")
		return false
	}
	fmt.Println("Order validated.")

	// Payment processing (simulated)
	if !op.processPayment(order.TotalAmount) {
		fmt.Println("Payment processing failed.")
		return false
	}
	fmt.Println("Payment processed.")

	// Inventory update (simulated)
	if !op.updateInventory(order.Items) {
		fmt.Println("Inventory update failed.")
		return false
	}
	fmt.Println("Inventory updated.")

	// Logging (simulated)
	op.logOrder(order)

	// Notification (simulated)
	op.sendNotification(order.Customer, order.ID)

	// Save order to system (simulated)
	op.orders = append(op.orders, order)
	fmt.Println("Order processed successfully.")
	return true
}

func (op *OrderProcessor) processPayment(amount float64) bool {
	fmt.Printf("Processing payment of $%.2f...\n", amount)
	// Simulate payment processing delay
	time.Sleep(500 * time.Millisecond)
	return true
}

func (op *OrderProcessor) updateInventory(items []string) bool {
	fmt.Println("Updating inventory for items:", items)
	// Simulate inventory update delay
	time.Sleep(300 * time.Millisecond)
	return true
}

func (op *OrderProcessor) logOrder(order Order) {
	fmt.Printf("Logging order #%d for customer %s at %s\n", order.ID, order.Customer, order.CreatedAt.Format(time.RFC822))
}

func (op *OrderProcessor) sendNotification(customer string, orderID int) {
	fmt.Printf("Sending notification to %s for order #%d\n", customer, orderID)
}

func main() {
	// Create a sample order
	order := Order{
		ID:          101,
		Customer:    "Alice",
		Items:       []string{"Laptop", "Mouse"},
		TotalAmount: 1200.50,
		CreatedAt:   time.Now(),
	}

	// Process the order using the non-SRP OrderProcessor
	processor := &OrderProcessor{}
	processor.ProcessOrder(order)
}
```

*Issues with this Approach:*
- **High Coupling:** All functionalities (validation, payment, inventory, logging, notification) are in one class.
- **Difficult Maintenance:** A change in one area (e.g., updating the payment logic) requires modifying the same class, risking unintended side effects.
- **Poor Testability:** Unit testing individual responsibilities becomes challenging.

---

### With SRP Applied

By refactoring the code according to SRP, we separate each responsibility into its own class or service. The `OrderProcessor` now delegates tasks to specialized services, making the system modular and easier to maintain.

```go
package main

import (
	"fmt"
	"time"
)

// Order represents an e-commerce order.
type Order struct {
	ID          int
	Customer    string
	Items       []string
	TotalAmount float64
	CreatedAt   time.Time
}

// PaymentService handles payment processing.
type PaymentService struct{}

func (ps *PaymentService) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing payment of $%.2f...\n", amount)
	time.Sleep(500 * time.Millisecond)
	return true
}

// InventoryService handles inventory updates.
type InventoryService struct{}

func (is *InventoryService) UpdateInventory(items []string) bool {
	fmt.Println("Updating inventory for items:", items)
	time.Sleep(300 * time.Millisecond)
	return true
}

// Logger handles logging operations.
type Logger struct{}

func (l *Logger) LogOrder(order Order) {
	fmt.Printf("Logging order #%d for customer %s at %s\n", order.ID, order.Customer, order.CreatedAt.Format(time.RFC822))
}

// NotificationService handles sending notifications.
type NotificationService struct{}

func (ns *NotificationService) SendNotification(customer string, orderID int) {
	fmt.Printf("Sending notification to %s for order #%d\n", customer, orderID)
}

// OrderValidator handles order validation.
type OrderValidator struct{}

func (ov *OrderValidator) Validate(order Order) bool {
	if order.TotalAmount <= 0 || len(order.Items) == 0 {
		fmt.Println("Order validation failed.")
		return false
	}
	fmt.Println("Order validated.")
	return true
}

// OrderProcessor now focuses solely on coordinating the order processing workflow.
type OrderProcessor struct {
	validator           *OrderValidator
	paymentService      *PaymentService
	inventoryService    *InventoryService
	logger              *Logger
	notificationService *NotificationService
	orders              []Order
}

func NewOrderProcessor() *OrderProcessor {
	return &OrderProcessor{
		validator:           &OrderValidator{},
		paymentService:      &PaymentService{},
		inventoryService:    &InventoryService{},
		logger:              &Logger{},
		notificationService: &NotificationService{},
	}
}

// ProcessOrder coordinates the workflow by delegating responsibilities.
func (op *OrderProcessor) ProcessOrder(order Order) bool {
	// Validate order
	if !op.validator.Validate(order) {
		return false
	}

	// Process payment
	if !op.paymentService.ProcessPayment(order.TotalAmount) {
		fmt.Println("Payment processing failed.")
		return false
	}
	fmt.Println("Payment processed.")

	// Update inventory
	if !op.inventoryService.UpdateInventory(order.Items) {
		fmt.Println("Inventory update failed.")
		return false
	}
	fmt.Println("Inventory updated.")

	// Log order
	op.logger.LogOrder(order)

	// Send notification
	op.notificationService.SendNotification(order.Customer, order.ID)

	// Save order to system
	op.orders = append(op.orders, order)
	fmt.Println("Order processed successfully.")
	return true
}

func main() {
	// Create a sample order
	order := Order{
		ID:          101,
		Customer:    "Alice",
		Items:       []string{"Laptop", "Mouse"},
		TotalAmount: 1200.50,
		CreatedAt:   time.Now(),
	}

	// Process the order using the refactored OrderProcessor that follows SRP
	processor := NewOrderProcessor()
	processor.ProcessOrder(order)
}
```

*Benefits of the SRP Approach:*
- **Separation of Concerns:** Each service (validation, payment, inventory, logging, notification) has a single responsibility.
- **Maintainability:** Changes in one service do not affect others.
- **Testability:** Individual services can be tested in isolation.
- **Reusability:** Services like `PaymentService` and `Logger` can be reused in other parts of the application.

---

## Advantages of Applying SRP

1. **Improved Readability:** Each class or service is focused on a single task, making the code easier to understand.
2. **Easier Maintenance:** Isolated responsibilities mean that changes in one area are less likely to affect other areas.
3. **Enhanced Testability:** Smaller, focused components are easier to test with unit tests.
4. **Greater Reusability:** Modular components can be reused across different parts of the application.
5. **Reduced Complexity:** Splitting responsibilities leads to simpler classes with clear purposes.

---

## Disadvantages of Applying SRP

1. **Increased Number of Classes:** Following SRP strictly can lead to a larger number of smaller classes, which might make the overall system more complex to navigate.
2. **Potential Overhead:** The added abstraction layers can sometimes introduce additional overhead in terms of code organization and management.
3. **Integration Complexity:** Combining multiple services to achieve a single workflow may require careful coordination.

---

## Conclusion

The Single Responsibility Principle is a fundamental design guideline that advocates for a clear separation of concerns within your code. By ensuring that each class or module has only one responsibility, developers can create systems that are easier to maintain, extend, and test. The real-life example of an order processing system illustrates how a monolithic approach can be refactored into a modular design where each service handles a distinct aspect of the order process. Although applying SRP can lead to an increased number of classes, the benefits in maintainability, testability, and clarity typically outweigh the drawbacks.

