# State Design Pattern

The State Design Pattern is a behavioral pattern that allows an object to alter its behavior when its internal state changes. Instead of implementing numerous conditional statements to change behavior based on state, the object delegates state-specific behavior to separate state objects. This results in cleaner, more maintainable code by encapsulating state-specific logic into dedicated classes.

---

## Key Concepts of the State Pattern

1. **Context:** The object whose behavior varies based on its internal state. It maintains a reference to a state object.
2. **State Interface:** Declares the methods that each concrete state must implement.
3. **Concrete States:** Classes that implement the state interface. Each state encapsulates behavior specific to that state.
4. **Client:** Interacts with the Context and, indirectly, with the concrete state classes without needing to know their specific implementations.

---

## Structure of the State Pattern

- **Context Class:** Maintains an instance of a Concrete State that defines the current state.
- **State Interface:** Declares common methods for handling requests.
- **Concrete State Classes:** Implement the state interface and define behavior for each state.
- **State Transition:** The Context can switch between different Concrete States as its internal state changes.

---

## When to Use the State Pattern

- When an object's behavior depends on its internal state, and it must change its behavior at runtime based on that state.
- To avoid large conditional statements (if-else or switch-case) that check for state.
- When state-specific behavior should be isolated and encapsulated within separate classes.
- To improve maintainability by making the state transitions explicit and easy to manage.

---

## Real-Life Example in Software Systems: Turnstile Gate

Imagine a turnstile gate at a subway station that can be either **Locked** or **Unlocked**. When the turnstile is locked, inserting a coin unlocks it, and pushing through then locks it again. The behavior changes based on the current state of the turnstile.

### Example in Go

```go
package main

import "fmt"

// State is the interface that declares actions for the context.
type State interface {
    Coin(context *Turnstile)
    Push(context *Turnstile)
}

// Turnstile is the context that holds a state.
type Turnstile struct {
    state State
}

// NewTurnstile initializes a turnstile in the Locked state.
func NewTurnstile() *Turnstile {
    t := &Turnstile{}
    t.state = &LockedState{}
    return t
}

// SetState allows the turnstile to change its state.
func (t *Turnstile) SetState(s State) {
    t.state = s
}

// Coin delegates the coin insertion action to the current state.
func (t *Turnstile) Coin() {
    t.state.Coin(t)
}

// Push delegates the push action to the current state.
func (t *Turnstile) Push() {
    t.state.Push(t)
}

// LockedState represents the state when the turnstile is locked.
type LockedState struct{}

func (ls *LockedState) Coin(context *Turnstile) {
    fmt.Println("Coin inserted. Turnstile is now unlocked.")
    context.SetState(&UnlockedState{})
}

func (ls *LockedState) Push(context *Turnstile) {
    fmt.Println("Turnstile is locked. Cannot push.")
}

// UnlockedState represents the state when the turnstile is unlocked.
type UnlockedState struct{}

func (us *UnlockedState) Coin(context *Turnstile) {
    fmt.Println("Coin inserted, but turnstile is already unlocked. Returning coin.")
}

func (us *UnlockedState) Push(context *Turnstile) {
    fmt.Println("Pushed through. Turnstile is now locked.")
    context.SetState(&LockedState{})
}

func main() {
    turnstile := NewTurnstile()

    // Attempt to push while locked.
    turnstile.Push() // Expected: "Turnstile is locked. Cannot push."

    // Insert coin to unlock.
    turnstile.Coin() // Expected: "Coin inserted. Turnstile is now unlocked."

    // Attempt to insert another coin.
    turnstile.Coin() // Expected: "Coin inserted, but turnstile is already unlocked. Returning coin."

    // Push through to lock again.
    turnstile.Push() // Expected: "Pushed through. Turnstile is now locked."
}
```

**Expected Output:**
```
Turnstile is locked. Cannot push.
Coin inserted. Turnstile is now unlocked.
Coin inserted, but turnstile is already unlocked. Returning coin.
Pushed through. Turnstile is now locked.
```

---

## Advantages of the State Pattern

1. **Simplifies Conditional Logic:** Eliminates large conditional statements by delegating state-specific behavior to separate classes.
2. **Encapsulation:** Keeps state-specific behavior isolated, making code easier to maintain and extend.
3. **Flexibility:** Allows dynamic change of behavior at runtime by switching state objects.
4. **Improved Readability:** Clearly separates the different behaviors based on state.

---

## Disadvantages of the State Pattern

1. **Increased Number of Classes:** Each state is implemented as a separate class, which can increase the number of classes in the system.
2. **Complexity in State Management:** Managing state transitions and ensuring all states are handled properly can add complexity.
3. **Overhead:** Additional objects and state transitions might introduce slight performance overhead in some cases.

---

## Real-World Scenario Example: Vending Machine

A vending machine can be in several states: waiting for coin insertion, coin inserted, dispensing product, or out of order. The State pattern can encapsulate the behavior for each of these states, making the code easier to manage and modify as new states or behaviors are introduced.

---

## Conclusion

The State Design Pattern provides a robust method for managing an object's behavior based on its internal state. By encapsulating state-specific behavior in separate classes, it eliminates the need for complex conditional logic and enhances code maintainability. While it introduces additional classes and complexity, its benefits in flexibility, readability, and dynamic behavior make it an invaluable pattern for state-driven systems.
