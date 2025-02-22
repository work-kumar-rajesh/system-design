# Proxy Design Pattern

The Proxy Design Pattern is a structural pattern that provides a surrogate or placeholder for another object to control access to it. It enables you to add functionality such as lazy initialization, caching, logging, or access control without modifying the original object's code.

## Key Concepts of the Proxy Pattern

1. **Subject:** The common interface that both the Real Subject and the Proxy implement.
2. **Real Subject:** The original object that the proxy represents. This is the object that performs the actual work.
3. **Proxy:** The object that controls access to the Real Subject. It can add additional behavior (e.g., logging, lazy initialization) before or after delegating the request to the Real Subject.
4. **Client:** The code that interacts with the Subject interface without knowing whether it is dealing with the Real Subject or the Proxy.

---

## When to Use the Proxy Pattern

- **Resource Management:** When the Real Subject is resource-intensive to create and you want to defer its instantiation (lazy initialization).
- **Access Control:** When you need to control access to the Real Subject, for example by checking user permissions or logging requests.
- **Remote Access:** When the Real Subject is located in a different address space (e.g., on a remote server) and you want to provide a local surrogate.
- **Caching:** When repeated requests can be optimized by caching responses.
- **Logging/Monitoring:** When you want to log requests or monitor usage without modifying the Real Subject.

---

## Real-Life Example in Software Systems: Remote Service Access

Consider a scenario where a client needs to access a remote service. The service is expensive to initialize and should only be created when needed. Additionally, you may want to log each request for monitoring purposes.

### Without the Proxy Pattern (Direct Access)
```go
package main

import "fmt"

// RemoteService represents a service that is expensive to initialize.
type RemoteService struct{}

func (rs *RemoteService) Request(data string) string {
    return "Response from remote service for: " + data
}

func main() {
    // Direct creation and usage of the remote service.
    service := &RemoteService{}
    fmt.Println(service.Request("Test Data"))
}
```

**Issues:**
- The client directly creates the RemoteService, which might be resource-intensive.
- No control is provided for lazy initialization, access control, or logging.

### With the Proxy Design Pattern
```go
package main

import "fmt"

// Subject defines the common interface.
type Subject interface {
    Request(data string) string
}

// RemoteService is the real subject.
type RemoteService struct{}

func (rs *RemoteService) Request(data string) string {
    return "Response from remote service for: " + data
}

// Proxy controls access to RemoteService.
type Proxy struct {
    realService *RemoteService
}

func (p *Proxy) Request(data string) string {
    // Lazy initialization: Create the RemoteService only when needed.
    if p.realService == nil {
        fmt.Println("Initializing remote service...")
        p.realService = &RemoteService{}
    }
    // Logging the request.
    fmt.Println("Proxy: Request received for data:", data)
    return p.realService.Request(data)
}

func main() {
    // The client interacts with the proxy through the Subject interface.
    var service Subject = &Proxy{}
    response := service.Request("Test Data")
    fmt.Println(response)
}
```

---

## Advantages of the Proxy Pattern

1. **Controlled Access:** The proxy controls how and when the client interacts with the Real Subject.
2. **Lazy Initialization:** The Real Subject can be created only when it is actually needed.
3. **Additional Functionality:** Additional behavior (e.g., logging, caching, security checks) can be added without modifying the Real Subject.
4. **Encapsulation:** The proxy hides the complexities and resource usage of the Real Subject from the client.

## Disadvantages of the Proxy Pattern

1. **Increased Complexity:** Adding a proxy introduces an extra layer of abstraction.
2. **Performance Overhead:** There is a slight performance cost due to additional method calls and checks.
3. **Maintenance:** Changes to the Real Subject's interface may require updates to the proxy.

---

## Real-World Scenario Example: Access Control

Imagine a system where sensitive data is stored on a remote server. Access to this data must be logged and controlled. A proxy can check user permissions before granting access, log each request for auditing, and only instantiate the remote connection when needed. This ensures that security and performance are managed effectively without burdening the client code.

---

## Conclusion

The Proxy Design Pattern provides a robust solution for managing access to objects. It introduces an intermediary layer that can add functionalities such as lazy initialization, logging, caching, and access control. While it increases the complexity of the system, its benefits in scenarios that require controlled and efficient access to resources often outweigh the drawbacks.
