# Open-Closed Principle (OCP)

The Open-Closed Principle (OCP) is one of the five SOLID principles of object-oriented design. It states that software entities (classes, modules, functions, etc.) should be **open for extension** but **closed for modification**. In other words, you should be able to add new functionality to a system without changing the existing code, which minimizes the risk of introducing new bugs and makes the system more maintainable.

---

## Detailed Theory

### What Does It Mean to Be Open for Extension?
- **Open for Extension:** A module or class can have its behavior extended without modifying its source code. This is typically achieved through techniques like inheritance, composition, and polymorphism.

### What Does It Mean to Be Closed for Modification?
- **Closed for Modification:** Once a class or module is written and tested, its source code should not be modified. Instead, new functionality should be added by extending its behavior through new code.

### Why Is OCP Important?
- **Maintainability:** Changes in requirements can be accommodated by adding new code rather than altering existing, stable code.
- **Reliability:** By not modifying tested code, you reduce the risk of introducing regressions.
- **Scalability:** Systems designed with OCP are easier to scale, as new features can be integrated without disturbing the existing architecture.

---

## Real-Life Example: Discount Calculation in an E-Commerce Application

Consider an e-commerce application where different discount strategies are applied to customer orders. In a non-OCP implementation, you might have a large function or class with conditional logic to determine which discount to apply. With OCP, you can design the system to be extensible so that new discount types can be added without modifying existing code.

### Without Applying OCP

In this example, a single `DiscountCalculator` handles multiple discount types using conditional statements. Adding a new discount requires modifying the class.

```go
package main

import "fmt"

// DiscountCalculator calculates discount based on customer type.
type DiscountCalculator struct{}

// CalculateDiscount applies different discount rates based on customer type.
func (dc *DiscountCalculator) CalculateDiscount(customerType string, amount float64) float64 {
    if customerType == "Regular" {
        return amount * 0.05 // 5% discount for regular customers
    } else if customerType == "Premium" {
        return amount * 0.10 // 10% discount for premium customers
    } else if customerType == "VIP" {
        return amount * 0.15 // 15% discount for VIP customers
    }
    return 0
}

func main() {
    calculator := &DiscountCalculator{}
    discount := calculator.CalculateDiscount("Premium", 200.0)
    fmt.Printf("Discount for Premium customer: $%.2f\n", discount)
}
```

*Issues:*
- **Modification Required:** To add a new discount strategy, you must modify the `CalculateDiscount` method.
- **Conditional Complexity:** As more customer types are added, the conditional logic becomes cumbersome and error-prone.

---

## With OCP Applied

By applying OCP, we can separate the discount calculation into distinct strategies. Each strategy implements a common interface. The system is now open for extension—new discount strategies can be added without modifying existing code.

```go
package main

import "fmt"

// DiscountStrategy defines the interface for discount calculation.
type DiscountStrategy interface {
    Calculate(amount float64) float64
}

// RegularDiscount provides a 5% discount.
type RegularDiscount struct{}

func (rd *RegularDiscount) Calculate(amount float64) float64 {
    return amount * 0.05
}

// PremiumDiscount provides a 10% discount.
type PremiumDiscount struct{}

func (pd *PremiumDiscount) Calculate(amount float64) float64 {
    return amount * 0.10
}

// VIPDiscount provides a 15% discount.
type VIPDiscount struct{}

func (vd *VIPDiscount) Calculate(amount float64) float64 {
    return amount * 0.15
}

// DiscountContext holds a reference to a discount strategy.
type DiscountContext struct {
    strategy DiscountStrategy
}

// SetStrategy allows changing the discount strategy.
func (dc *DiscountContext) SetStrategy(strategy DiscountStrategy) {
    dc.strategy = strategy
}

// GetDiscount calculates the discount using the current strategy.
func (dc *DiscountContext) GetDiscount(amount float64) float64 {
    if dc.strategy == nil {
        return 0
    }
    return dc.strategy.Calculate(amount)
}

func main() {
    context := &DiscountContext{}

    // Apply Regular discount
    context.SetStrategy(&RegularDiscount{})
    fmt.Printf("Regular Discount: $%.2f\n", context.GetDiscount(200.0))

    // Apply Premium discount
    context.SetStrategy(&PremiumDiscount{})
    fmt.Printf("Premium Discount: $%.2f\n", context.GetDiscount(200.0))

    // Apply VIP discount
    context.SetStrategy(&VIPDiscount{})
    fmt.Printf("VIP Discount: $%.2f\n", context.GetDiscount(200.0))
}
```

*Benefits:*
- **Extensibility:** Adding a new discount type requires creating a new strategy without changing existing classes.
- **Reduced Complexity:** Each discount calculation is isolated in its own class.
- **Maintainability:** Changes in discount logic affect only the relevant strategy class.

---

## Advantages of Applying OCP

1. **Enhances Maintainability:** New features can be added with minimal risk to existing functionality.
2. **Improves Reliability:** Well-tested, stable code remains unchanged, reducing the likelihood of bugs.
3. **Facilitates Reuse:** Modules can be reused in different contexts without modification.
4. **Promotes Clean Architecture:** Encourages separation of concerns and modular design.

---

## Disadvantages of Applying OCP

1. **Increased Complexity:** May lead to an increased number of classes or interfaces.
2. **Overhead:** Additional abstraction layers can add complexity to the system architecture.
3. **Design Challenge:** Requires careful planning to design extensible systems from the start.

---

## Conclusion

The Open-Closed Principle is a vital guideline for creating maintainable and scalable software systems. By ensuring that modules are open for extension but closed for modification, developers can add new functionality without risking regression in existing code. The provided example of a discount calculation system illustrates how conditional logic can be refactored into separate strategy classes, leading to a design that is easier to extend, maintain, and test.

