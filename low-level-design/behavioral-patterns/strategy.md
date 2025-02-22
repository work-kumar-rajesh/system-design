# Strategy Design Pattern

The Strategy Design Pattern is a behavioral design pattern that defines a family of algorithms, encapsulates each one, and makes them interchangeable. This pattern lets the algorithm vary independently from the clients that use it, promoting flexibility and reusability.

## Key Concepts of the Strategy Pattern

1. **Strategy:** An interface that defines a common method for performing an algorithm.
2. **Concrete Strategies:** Classes that implement the Strategy interface, each providing a specific algorithm.
3. **Context:** Maintains a reference to a Strategy object and uses it to perform the operation.
4. **Client:** The code that configures the Context with a specific Strategy and invokes its behavior.

---

## Structure of the Strategy Pattern

- **Strategy Interface:** Declares the method(s) that all concrete strategies must implement.
- **Concrete Strategies:** Implement the Strategy interface with different algorithms or behaviors.
- **Context:** Contains a reference to a Strategy and delegates the algorithm execution to the current Strategy.
- **Client:** Chooses the appropriate Strategy based on the context and sets it in the Context.

---

## When to Use the Strategy Pattern

- When you have multiple algorithms for a specific task and want to switch between them at runtime.
- To avoid complex conditional statements that select different behaviors.
- When you need to decouple the algorithm from the context in which it is used.
- To promote code reuse by encapsulating each algorithm in its own class.

---

## Real-Life Example in Software Systems: Tax Calculation

Consider an order processing system where the tax calculation may vary based on the region. For instance, the United States might have a 7% tax rate while the European Union might have a 20% tax rate. The Strategy Pattern allows the order to use different tax calculation strategies dynamically.

### Example in Go

```go
package main

import "fmt"

// Strategy interface for tax calculation.
type TaxStrategy interface {
    Calculate(amount float64) float64
}

// Concrete Strategy: USTaxStrategy calculates tax at 7%
type USTaxStrategy struct{}

func (s *USTaxStrategy) Calculate(amount float64) float64 {
    return amount * 0.07
}

// Concrete Strategy: EUTaxStrategy calculates tax at 20%
type EUTaxStrategy struct{}

func (s *EUTaxStrategy) Calculate(amount float64) float64 {
    return amount * 0.20
}

// Context: Order uses a TaxStrategy to calculate total price.
type Order struct {
    amount      float64
    taxStrategy TaxStrategy
}

// SetTaxStrategy allows changing the tax calculation strategy.
func (o *Order) SetTaxStrategy(strategy TaxStrategy) {
    o.taxStrategy = strategy
}

// CalculateTotal computes the total amount including tax.
func (o *Order) CalculateTotal() float64 {
    tax := o.taxStrategy.Calculate(o.amount)
    return o.amount + tax
}

func main() {
    order := &Order{amount: 100.0}
    
    // Using US Tax Strategy.
    order.SetTaxStrategy(&USTaxStrategy{})
    fmt.Printf("Total with US Tax: %.2f\n", order.CalculateTotal())
    
    // Switching to EU Tax Strategy.
    order.SetTaxStrategy(&EUTaxStrategy{})
    fmt.Printf("Total with EU Tax: %.2f\n", order.CalculateTotal())
}
```

**Output:**
```
Total with US Tax: 107.00
Total with EU Tax: 120.00
```

---

## Advantages of the Strategy Pattern

1. **Flexibility:** Algorithms can be changed at runtime by switching strategies.
2. **Encapsulation:** Each algorithm is encapsulated in its own class, leading to cleaner code.
3. **Elimination of Conditional Statements:** Reduces the need for complex conditional logic to select an algorithm.
4. **Reusability:** Strategies can be reused across different contexts.

## Disadvantages of the Strategy Pattern

1. **Increased Number of Classes:** Each strategy is a separate class, which might increase the number of classes in the system.
2. **Overhead:** Switching between strategies may introduce slight overhead.
3. **Client Awareness:** The client must be aware of the different strategies available to select the appropriate one.

---

## Conclusion

The Strategy Design Pattern offers a robust method for managing multiple algorithms by encapsulating them into separate classes. It promotes a clean separation of concerns, improves maintainability, and provides flexibility by allowing the algorithm to change at runtime. Whether used for tax calculation, sorting, or any varying behavior, the Strategy Pattern helps in writing modular and adaptable code.
