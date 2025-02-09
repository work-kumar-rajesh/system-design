# Builder Design Pattern

The Builder Design Pattern is a creational pattern used to construct complex objects step by step. It allows you to create 
different representations of an object using the same construction process. The Builder pattern is especially useful when an 
object requires many parameters or has a complex configuration, and you want to avoid having large constructors or multiple 
constructor overloads. Key Concepts of the Builder Pattern:

1. Builder: Responsible for constructing the product.
2. Product: The final object that will be constructed.
3. Director: A class that defines the order in which to call the construction steps (this is optional but common).
4. Client: The class that uses the builder to construct the product.


## Structure of the Builder Pattern:
1. Product: The object to be constructed.
2. Builder: An abstract class or interface that defines the steps to build the product.
3. Concrete Builder: Implements the Builder interface and constructs the product step by step.
4. Director: Optional class responsible for managing the construction process.
5. Client: Uses the director and builder to create the object.


## When to Use the Builder Pattern:
1. When you need to create an object with many parameters, especially when some parameters are optional.
2. When the creation process involves multiple steps or has different configurations.
3. When you want to create different representations of the same type of object.


## Example in Go:

Let's say we're building a Pizza object, where a pizza can have different sizes, toppings, and crust types.


```go
package main

import "fmt"

// Product represents the final object
type Pizza struct {
    size    string
    crust   string
    toppings []string
}

func (p *Pizza) String() string {
    return fmt.Sprintf("Pizza [Size: %s, Crust: %s, Toppings: %v]", p.size, p.crust, p.toppings)
}

// Builder interface defines the steps to build the product
type PizzaBuilder interface {
    SetSize(size string) PizzaBuilder
    SetCrust(crust string) PizzaBuilder
    AddTopping(topping string) PizzaBuilder
    Build() *Pizza
}

// ConcreteBuilder is the actual implementation of the builder interface
type ConcretePizzaBuilder struct {
    pizza *Pizza
}

func NewConcretePizzaBuilder() *ConcretePizzaBuilder {
    return &ConcretePizzaBuilder{
        pizza: &Pizza{},
    }
}

func (b *ConcretePizzaBuilder) SetSize(size string) PizzaBuilder {
    b.pizza.size = size
    return b
}

func (b *ConcretePizzaBuilder) SetCrust(crust string) PizzaBuilder {
    b.pizza.crust = crust
    return b
}

func (b *ConcretePizzaBuilder) AddTopping(topping string) PizzaBuilder {
    b.pizza.toppings = append(b.pizza.toppings, topping)
    return b
}

func (b *ConcretePizzaBuilder) Build() *Pizza {
    return b.pizza
}


//The director helps to organize the construction process and ensures a product is built in a particular sequence of steps.
// Director is responsible for constructing a pizza using the builder
type PizzaDirector struct {
    builder PizzaBuilder
}

func NewPizzaDirector(builder PizzaBuilder) *PizzaDirector {
    return &PizzaDirector{builder: builder}
}

func (d *PizzaDirector) ConstructMargheritaPizza() *Pizza {
    return d.builder.SetSize("Medium").
    SetCrust("Thin").
    AddTopping("Cheese").
    AddTopping("Tomato").
    Build()
}

func (d *PizzaDirector) ConstructPepperoniPizza() *Pizza {
    return d.builder.SetSize("Large").
    SetCrust("Thick").
    AddTopping("Cheese").
    AddTopping("Pepperoni").
    Build()
}
```

```go
//Client Code

func main() {
    // Create a concrete builder
    pizzaBuilder := NewConcretePizzaBuilder()

    // Director to manage the construction
    director := NewPizzaDirector(pizzaBuilder)

    // Construct different types of pizzas
    margheritaPizza := director.ConstructMargheritaPizza()
    fmt.Println("Margherita Pizza:", margheritaPizza)

    pepperoniPizza := director.ConstructPepperoniPizza()
    fmt.Println("Pepperoni Pizza:", pepperoniPizza)
}
```
Output:

Margherita Pizza: Pizza [Size: Medium, Crust: Thin, Toppings: [Cheese Tomato]]
Pepperoni Pizza: Pizza [Size: Large, Crust: Thick, Toppings: [Cheese Pepperoni]]


## Explanation of the Example:
Pizza (Product): The Pizza struct represents the final product. It holds information such as size, crust type, and toppings.
PizzaBuilder (Builder Interface): The PizzaBuilder interface defines methods that must be implemented by a concrete builder. These methods set different parts of the pizza.
ConcretePizzaBuilder (Concrete Builder): The ConcretePizzaBuilder struct implements the PizzaBuilder interface. It is responsible for constructing a Pizza object step by step.
PizzaDirector (Director): The PizzaDirector helps to organize the construction process. It can create specific kinds of pizzas (like Margherita or Pepperoni) using the builder.
Client: The client uses the PizzaDirector and PizzaBuilder to create complex pizza objects. It does not need to know the details of how the pizza is constructed.


## Advantages of the Builder Pattern:
Separation of Concerns: The construction logic is separated from the product class. This improves readability and maintainability.
Avoids Constructor Overload: Instead of having multiple constructors with different combinations of parameters, the builder provides a fluent API to build objects step by step.
Immutability: The product (in this case, pizza) can be built step by step and is immutable once it is created.
Flexible Object Creation: You can create different types of objects with different configurations using the same builder.


## Disadvantages:
Increased Complexity: The Builder pattern introduces additional classes, which may be unnecessary for simpler objects.
Not Always Required: If the object creation is simple or doesn’t involve a large number of parameters, the Builder pattern may be overkill.

## Conclusion:
The Builder Design Pattern is a powerful tool for constructing complex objects, especially when the number of parameters grows,
and there are multiple configurations possible. It ensures a clean and readable API and prevents the need for constructor overloads, 
making the code easier to maintain and extend.


##  Important 

### Why we are using builder interface here and not directly concrete implementation ? 

#### Decoupling from Specific Implementation (Flexibility):

By using a Builder interface, you decouple the client code from the concrete builder implementation. The client 
(such as the PizzaDirector) does not need to know the specific type of builder being used. It only needs to know that it can 
interact with an object that implements the PizzaBuilder interface. This allows for flexibility, where you can easily swap 
different ConcreteBuilder implementations without affecting the client code.
For example, you could add a new VeganPizzaBuilder or GlutenFreePizzaBuilder that still adheres to the PizzaBuilder interface but
has different implementation details. The PizzaDirector would remain the same and could work with any builder that 
implements the PizzaBuilder interface.

```go
type VeganPizzaBuilder struct {
    pizza *Pizza
}

func (v *VeganPizzaBuilder) SetSize(size string) PizzaBuilder {
    v.pizza.size = size
    return v
}

func (v *VeganPizzaBuilder) SetCrust(crust string) PizzaBuilder {
    v.pizza.crust = crust
    return v
}

func (v *VeganPizzaBuilder) AddTopping(topping string) PizzaBuilder {
    if topping != "Cheese" {
        v.pizza.toppings = append(v.pizza.toppings, topping)
    }
    return v
}

func (v *VeganPizzaBuilder) Build() *Pizza {
    return v.pizza
}
```

#### Abstraction for Complex Builders:

In real-world scenarios, you may have multiple ConcreteBuilder classes for building objects in different ways. The interface helps 
abstract the construction process, so the client does not need to know how the object is being built internally. By interacting with 
the interface, the client can stay unaware of the complex internal logic or the details of the concrete builder. For example, 
one builder might construct the pizza with different default toppings or use different types of crusts. Using the interface allows you
to plug in different builder types without affecting the client logic.

#### Separation of Concerns:

The PizzaBuilder interface defines the essential steps for constructing a pizza. The ConcretePizzaBuilder provides the actual 
implementation of these steps. This separation of concerns ensures that the PizzaDirector is responsible for constructing the pizza 
in a specific sequence, while the ConcreteBuilder is responsible for the actual steps and logic of constructing the pizza.
Without the interface, the PizzaDirector would need to know the specifics of how a ConcretePizzaBuilder works, which tightly couples 
the classes together and violates the Single Responsibility Principle (SRP). The interface prevents this.

#### Testability:
When writing tests for the construction of an object, it's easier to mock or substitute the builder if it is based on an interface. 
For example, you could write unit tests for the PizzaDirector without worrying about the actual implementation of the builder. 
You just need to provide a mock or a simple test builder that implements the same interface. If the builder were concrete and not 
abstracted, it would be harder to mock or control the construction process in tests.

```go
type MockPizzaBuilder struct {
    pizza *Pizza
}

func (m *MockPizzaBuilder) SetSize(size string) PizzaBuilder {
    m.pizza.size = "Small" // Mock value
    return m
}

func (m *MockPizzaBuilder) SetCrust(crust string) PizzaBuilder {
    m.pizza.crust = "Thin" // Mock value
return m
}

func (m *MockPizzaBuilder) AddTopping(topping string) PizzaBuilder {
    return m // No-op for testing
}

func (m *MockPizzaBuilder) Build() *Pizza {
    return m.pizza
}
```

#### Consistency:

Using an interface guarantees that all builders adhere to the same contract, ensuring consistency in how objects are constructed. 
If each builder implements the same interface, you can be sure that any builder will be able to set the required properties and 
eventually return a built object, following the same steps. Without the interface, different ConcreteBuilder types could potentially 
diverge in their implementation details, which could introduce inconsistencies or lead to errors in the construction process.

#### Future Extension:
The interface-based design allows for future extensions without modifying existing code. If you want to add a new type of builder in the
future, all you need to do is implement the PizzaBuilder interface and add the logic for constructing that type of pizza.
This adheres to the Open-Closed Principle (OCP) from SOLID principles, which states that classes should be open for extension but 
closed for modification.

# REAL WORLD EXAMPLE
Here’s a concise comparison of using the Builder pattern versus not using the Builder pattern for constructing an HTTP request, 
along with their advantages and disadvantages.

## With Builder Pattern:

```go
package main

import (
    "bytes"
    "fmt"
    "net/http"
    "time"
)

type HTTPRequest struct {
    Method    string
    URL       string
    Headers   map[string]string
    Body      []byte
    Timeout   time.Duration
    AuthToken string
    RetryCount int
}

type HTTPRequestBuilder interface {
    SetMethod(method string) HTTPRequestBuilder
    SetURL(url string) HTTPRequestBuilder
    AddHeader(key, value string) HTTPRequestBuilder
    SetBody(body []byte) HTTPRequestBuilder
    SetTimeout(timeout time.Duration) HTTPRequestBuilder
    SetAuthToken(token string) HTTPRequestBuilder
    SetRetryCount(count int) HTTPRequestBuilder
    Build() *HTTPRequest
}

type ConcreteHTTPRequestBuilder struct {
    request *HTTPRequest
}

func NewHTTPRequestBuilder() *ConcreteHTTPRequestBuilder {
    return &ConcreteHTTPRequestBuilder{
        request: &HTTPRequest{
            Headers: make(map[string]string),
            Timeout: 30 * time.Second,
            RetryCount: 3,
        },
    }
}

func (b *ConcreteHTTPRequestBuilder) SetMethod(method string) HTTPRequestBuilder {
    b.request.Method = method
    return b
}

func (b *ConcreteHTTPRequestBuilder) SetURL(url string) HTTPRequestBuilder {
    b.request.URL = url
    return b
}

func (b *ConcreteHTTPRequestBuilder) AddHeader(key, value string) HTTPRequestBuilder {
    b.request.Headers[key] = value
    return b
}

func (b *ConcreteHTTPRequestBuilder) SetBody(body []byte) HTTPRequestBuilder {
    b.request.Body = body
    return b
}

func (b *ConcreteHTTPRequestBuilder) SetTimeout(timeout time.Duration) HTTPRequestBuilder {
    b.request.Timeout = timeout
    return b
}

func (b *ConcreteHTTPRequestBuilder) SetAuthToken(token string) HTTPRequestBuilder {
    b.request.AuthToken = token
    return b
}

func (b *ConcreteHTTPRequestBuilder) SetRetryCount(count int) HTTPRequestBuilder {
    b.request.RetryCount = count
    return b
}

func (b *ConcreteHTTPRequestBuilder) Build() *HTTPRequest {
    return b.request
}

func main() {
    builder := NewHTTPRequestBuilder()
    req := builder.SetMethod("POST").
    SetURL("https://example.com/api").
    AddHeader("Content-Type", "application/json").
    SetBody([]byte(`{"key": "value"}`)).
    Build()

	client := &http.Client{Timeout: req.Timeout}
	httpReq, _ := http.NewRequest(req.Method, req.URL, bytes.NewBuffer(req.Body))
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}
	fmt.Printf("Request: %+v\n", req)
	// client.Do(httpReq)  // Simulate sending the request
}
```

### Advantages of Using the Builder Pattern:

1. Readable & Flexible: Only set the fields you care about; optional fields are handled with defaults.
2. Avoids Error-Prone Constructors: Eliminates passing unnecessary or empty values for optional fields.
3. Method Chaining: Clear and intuitive method chaining to configure the request.
4. Easier Maintenance: Adding new fields or functionality doesn’t affect existing usage.

### Disadvantages of Using the Builder Pattern:

1. More Code: You need to create an extra builder struct and interfaces, leading to some initial overhead.
2. Learning Curve: For beginners, understanding the builder pattern and method chaining may take some time.

## Without Builder Pattern:
```go
package main

import (
    "bytes"
    "fmt"
    "net/http"
    "time"
)

type HTTPRequest struct {
    Method    string
    URL       string
    Headers   map[string]string
    Body      []byte
    Timeout   time.Duration
    AuthToken string
    RetryCount int
}

func NewHTTPRequest(method, url string, headers map[string]string, body []byte, timeout time.Duration, authToken string, retryCount int) *HTTPRequest {
    if timeout == 0 {
        timeout = 30 * time.Second  // Default timeout
    }
    if retryCount == 0 {
        retryCount = 3  // Default retry count
    }
    return &HTTPRequest{
        Method:    method,
        URL:       url,
        Headers:   headers,
        Body:      body,
        Timeout:   timeout,
        AuthToken: authToken,
        RetryCount: retryCount,
    }
}

func main() {
    headers := map[string]string{"Content-Type": "application/json"}
    req := NewHTTPRequest("POST", "https://example.com/api", headers, []byte(`{"key": "value"}`), 0, "", 0)

	client := &http.Client{Timeout: req.Timeout}
	httpReq, _ := http.NewRequest(req.Method, req.URL, bytes.NewBuffer(req.Body))
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}
	fmt.Printf("Request: %+v\n", req)
	// client.Do(httpReq)  // Simulate sending the request
}
```
### Advantages of Not Using the Builder Pattern:
1. Less Overhead: No extra code for builders or interfaces.
2. Simple and Straightforward: Direct approach to creating the object without additional abstraction.
3. Fewer Files: You don’t need extra structs or classes, keeping the code smaller.

### Disadvantages of Not Using the Builder Pattern:
1. Error-Prone: If you forget to set an optional field, or pass incorrect defaults (e.g., 0 for timeouts), the object may not behave as expected.
2. Cumbersome Constructors: The constructor with many parameters becomes hard to manage as the number of fields grows.
3. Less Extensible: Adding new parameters means updating the constructor and all places that use it, which could break existing code or increase complexity.

## Comparison Summary:
Aspect	                            With Builder Pattern	                            Without Builder Pattern
Code Complexity	             Higher (due to builder structure)	                    Lower (direct object creation)
Readability	                 High(method chaining, clear field customization)	    Moderate (manual handling of defaults and fields)
Flexibility	                       High (can easily add more fields)	            Low (adding fields requires function changes)
Error-Prone	                    Low (defaults are handled in the builder)	        High (possible to forget defaults, hard to manage)
Maintainability	                High (easy to add new fields)	                    Low (changing constructor impacts all uses)
Ease of Use	                     High (clear setter functions)	                    Low (function signatures with many parameters)
Performance	                 Slightly higher overhead (due to extra structs)	    Slightly lower overhead (direct manipulation)
Conclusion:

### Use the Builder Pattern when:
1. You have complex objects with many optional parameters.
2. You want clean, readable, and maintainable code. 
3. You expect to frequently extend the object with new fields.

Avoid the Builder Pattern when:
1. The object is simple with few parameters.
2. You want to keep your code minimal and avoid the overhead of builders. 
3. You don’t mind dealing with constructors with many parameters or managing defaults manually.

In real-world applications, Builder is particularly useful when constructing objects with many optional and 
required parameters (like HTTP requests, configurations, or complex data structures).
For simpler objects, direct construction might be sufficient.