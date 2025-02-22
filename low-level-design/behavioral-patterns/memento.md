# Memento Design Pattern

The Memento Design Pattern is a behavioral pattern that captures and externalizes an object's internal state without violating its encapsulation. This pattern enables you to restore an object to a previous state (undo functionality) without exposing its implementation details.

---

## Key Concepts of the Memento Pattern

1. **Originator:** The object whose state needs to be saved and restored. It creates a memento containing a snapshot of its current state.
2. **Memento:** The object that stores the internal state of the Originator. It is opaque to other objects, preserving encapsulation.
3. **Caretaker:** Manages the mementos. It requests a memento from the Originator, stores it, and can later provide it back to restore the state.
4. **Client:** The code that initiates requests to save and restore the state using the Caretaker and Originator.

---

## Structure of the Memento Pattern

- **Originator:** Contains the business logic and state that needs to be saved. It provides methods to create and restore mementos.
- **Memento:** A lightweight object that holds the state of the Originator. Its structure is usually hidden from other objects.
- **Caretaker:** Keeps track of the mementos and decides when to restore a previous state. It does not, however, modify or inspect the memento's contents.
- **Client:** Coordinates the saving and restoring process by interacting with the Originator and the Caretaker.

---

## When to Use the Memento Pattern

- When you need to implement undo/redo functionality.
- When you want to capture and restore an object's state without exposing its implementation.
- When maintaining a history of state changes is important.
- When an object’s state is complex and should not be directly manipulated by external objects.

---

## Real-Life Example in Software Systems: Text Editor Undo Feature

Consider a text editor that allows users to undo their changes. Each time the user makes a change, the current state of the document is saved as a memento. If the user chooses to undo, the document is restored to a previous state using one of the saved mementos.

---

## Example in Go

```go
package main

import "fmt"

// Memento stores the state of the Originator.
type Memento struct {
	state string
}

// Originator holds the current state and can create and restore mementos.
type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

// CreateMemento creates a memento capturing the current state.
func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

// RestoreState restores the originator's state from the memento.
func (o *Originator) RestoreState(m *Memento) {
	o.state = m.state
}

// Caretaker manages the mementos.
type Caretaker struct {
	mementos []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	if index < 0 || index >= len(c.mementos) {
		return nil
	}
	return c.mementos[index]
}

func main() {
	originator := &Originator{}
	caretaker := &Caretaker{}

	// Initial state
	originator.SetState("State #1")
	caretaker.AddMemento(originator.CreateMemento())

	// Change state
	originator.SetState("State #2")
	caretaker.AddMemento(originator.CreateMemento())

	// Change state again
	originator.SetState("State #3")
	fmt.Println("Current State:", originator.GetState())

	// Restore to previous state
	originator.RestoreState(caretaker.GetMemento(1))
	fmt.Println("Restored State:", originator.GetState())
}
```

**Expected Output:**
```
Current State: State #3
Restored State: State #2
```

---

## Advantages of the Memento Pattern

1. **Encapsulation:** The internal state of the Originator is not exposed to external objects.
2. **Undo/Redo Functionality:** Enables the implementation of powerful undo/redo mechanisms.
3. **Separation of Concerns:** The Originator handles state management while the Caretaker manages memento storage.
4. **Simplicity:** Mementos are typically lightweight, making it an efficient way to capture state snapshots.

## Disadvantages of the Memento Pattern

1. **Memory Overhead:** Storing mementos for complex objects can consume significant memory.
2. **Performance Impact:** Frequent creation and storage of mementos may affect performance.
3. **Limited Scope:** Mementos capture only a snapshot of the state; they may not cover all dynamic aspects of an object's behavior.

---

## Real-World Scenario Example: Game State Saving

In a video game, the Memento pattern can be used to save the state of the game (e.g., player's progress, position, score). When the player chooses to load a saved game, the system restores the state from the saved memento, enabling an "undo" of the gameplay up to that point.

---

## Conclusion

The Memento Design Pattern offers an elegant solution for saving and restoring an object's state without violating encapsulation. It is particularly useful for implementing undo functionality, managing state history, and ensuring that the internal representation of an object remains hidden from external manipulation. While it may introduce memory and performance overhead, its benefits in maintaining a clean separation of concerns and enabling state rollback make it an invaluable pattern in many applications.
