# Chain of Responsibility Design Pattern

The Chain of Responsibility (CoR) Design Pattern is a behavioral design pattern that allows an object to send a request without knowing which object in the chain will handle it. Instead of coupling the sender to a specific receiver, the request is passed along a chain of potential handlers until one of them processes it. This pattern promotes loose coupling and flexibility in assigning responsibilities.

---

## Key Concepts of the Chain of Responsibility Pattern

1. **Handler Interface:** Defines a method for handling requests and for setting the next handler in the chain.
2. **Concrete Handlers:** Classes that implement the Handler interface. Each handler processes the request if it meets certain criteria; otherwise, it forwards the request to the next handler.
3. **Client:** The code that sends the request into the chain without needing to know which handler will process it.
4. **Chain Setup:** The client (or an external coordinator) is responsible for linking the handlers in a specific order.

---

## Structure of the Chain of Responsibility Pattern

- **Handler Interface:** Declares methods for handling a request and for setting the next handler.
- **Base Handler (Optional):** A helper class that implements common behavior (like forwarding the request).
- **Concrete Handlers:** Specific classes that decide whether to process the request or pass it on.
- **Client:** Initiates the request and may configure the chain.

---

## When to Use the Chain of Responsibility Pattern

- **Multiple Handlers:** When you have several objects that can handle a request, and the handler isn't known a priori.
- **Decoupling:** To decouple the sender of a request from its receivers.
- **Dynamic Assignment:** When the set of potential handlers should be dynamically changed or extended.
- **Request Flexibility:** When requests need to be processed by multiple handlers, each adding its own processing.

---

## Real-Life Example in Software Systems: Support Ticket Handling

Imagine a support system where a customer service request can be handled by different levels of support:
- **Front Desk:** Handles simple requests (e.g., password resets).
- **Manager:** Handles moderate issues (e.g., software troubleshooting).
- **Director:** Handles critical issues (e.g., system outages).

Each handler in the chain examines the request severity and either processes it or forwards it to the next level.

---

## Example in Go

```go
package main

import "fmt"

// Request represents a support request with a problem description and a severity level.
type Request struct {
	Problem  string
	Severity int
}

// Handler interface defines methods for setting the next handler and processing the request.
type Handler interface {
	SetNext(handler Handler)
	Handle(request Request)
}

// BaseHandler provides default behavior for setting and forwarding to the next handler.
type BaseHandler struct {
	next Handler
}

func (bh *BaseHandler) SetNext(handler Handler) {
	bh.next = handler
}

func (bh *BaseHandler) Handle(request Request) {
	if bh.next != nil {
		bh.next.Handle(request)
	}
}

// FrontDesk handles low-severity requests.
type FrontDesk struct {
	BaseHandler
}

func (fd *FrontDesk) Handle(request Request) {
	if request.Severity <= 1 {
		fmt.Println("FrontDesk handling request:", request.Problem)
	} else {
		fmt.Println("FrontDesk forwarding request")
		fd.BaseHandler.Handle(request)
	}
}

// Manager handles moderate-severity requests.
type Manager struct {
	BaseHandler
}

func (m *Manager) Handle(request Request) {
	if request.Severity <= 3 {
		fmt.Println("Manager handling request:", request.Problem)
	} else {
		fmt.Println("Manager forwarding request")
		m.BaseHandler.Handle(request)
	}
}

// Director handles high-severity requests.
type Director struct {
	BaseHandler
}

func (d *Director) Handle(request Request) {
	// Director handles any remaining requests.
	fmt.Println("Director handling request:", request.Problem)
}

func main() {
	// Create the handlers.
	frontDesk := &FrontDesk{}
	manager := &Manager{}
	director := &Director{}

	// Set up the chain: FrontDesk -> Manager -> Director.
	frontDesk.SetNext(manager)
	manager.SetNext(director)

	// Create several support requests with varying severity.
	req1 := Request{Problem: "Forgot password", Severity: 1}
	req2 := Request{Problem: "Software not responding", Severity: 2}
	req3 := Request{Problem: "System outage", Severity: 5}

	// Process requests through the chain.
	frontDesk.Handle(req1)
	frontDesk.Handle(req2)
	frontDesk.Handle(req3)
}
```

**Expected Output:**
```
FrontDesk handling request: Forgot password
FrontDesk forwarding request
Manager handling request: Software not responding
FrontDesk forwarding request
Manager forwarding request
Director handling request: System outage
```

---

## Advantages of the Chain of Responsibility Pattern

1. **Loose Coupling:** The sender is decoupled from the receivers.
2. **Flexibility:** You can dynamically change or extend the chain by adding or removing handlers.
3. **Simplified Code:** Reduces the need for conditional statements to decide which object should handle the request.
4. **Responsibility Sharing:** Multiple handlers can cooperate to handle a request.

---

## Disadvantages of the Chain of Responsibility Pattern

1. **Uncertain Handling:** There's no guarantee that a request will be handled if no handler in the chain processes it.
2. **Performance Overhead:** Traversing the chain can be less efficient if the chain is long.
3. **Debugging Complexity:** Tracing the flow of a request through multiple handlers can be challenging.
4. **Maintenance Overhead:** Changing the order or behavior of handlers requires careful management.

---

## Conclusion

The Chain of Responsibility Design Pattern provides an elegant solution for handling requests by passing them along a chain of potential handlers. It promotes loose coupling between the sender and receiver and allows responsibilities to be dynamically assigned. While it may introduce some performance overhead and complexity in debugging, its benefits in flexibility and code maintainability make it a valuable pattern in many real-world applications, such as support ticket systems, event handling, and more.
