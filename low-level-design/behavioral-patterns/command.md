# Command Design Pattern

The Command Design Pattern is a behavioral pattern that encapsulates a request as an object, thereby allowing you to parameterize clients with different requests, queue or log requests, and support undoable operations. It decouples the object that invokes the operation from the one that knows how to perform it.

---

## Key Concepts of the Command Pattern

1. **Command:** An interface or abstract class that declares an execution method.
2. **Concrete Command:** Implements the Command interface and binds a Receiver to an action.
3. **Receiver:** The object that performs the actual operation when the command's execute method is called.
4. **Invoker:** The object that holds and calls the command, triggering the action.
5. **Client:** The code that creates the command, sets its Receiver, and assigns it to the Invoker.

---

## Structure of the Command Pattern

- **Command Interface:** Declares the `Execute()` method that all concrete commands implement.
- **Concrete Commands:** Implement the command interface, often storing a reference to a Receiver.
- **Receiver:** Contains the business logic to perform the actual operation.
- **Invoker:** Maintains a command object and invokes its `Execute()` method.
- **Client:** Creates a concrete command and sets the receiver.

---

## When to Use the Command Pattern

- **Parameterization of Requests:** When you need to parameterize objects with operations.
- **Decoupling:** To decouple the sender of a request from the receiver.
- **Undo/Redo Operations:** When you want to support reversible operations.
- **Command Queuing:** When requests should be queued, logged, or executed asynchronously.

---

## Real-Life Example in Software Systems: Remote Control for Home Automation

Consider a home automation system where a remote control can issue commands to turn lights on or off. The Command pattern encapsulates these requests as objects, decoupling the remote control (Invoker) from the actual devices (Receivers).

### Example in Go

```go
package main

import "fmt"

// Command interface declares the Execute method.
type Command interface {
    Execute()
}

// Receiver: Light, which can be turned on or off.
type Light struct {
    location string
}

func (l *Light) On() {
    fmt.Println(l.location, "light is ON")
}

func (l *Light) Off() {
    fmt.Println(l.location, "light is OFF")
}

// Concrete Command: LightOnCommand encapsulates turning the light on.
type LightOnCommand struct {
    light *Light
}

func (c *LightOnCommand) Execute() {
    c.light.On()
}

// Concrete Command: LightOffCommand encapsulates turning the light off.
type LightOffCommand struct {
    light *Light
}

func (c *LightOffCommand) Execute() {
    c.light.Off()
}

// Invoker: RemoteControl holds a command and triggers its execution.
type RemoteControl struct {
    command Command
}

func (r *RemoteControl) SetCommand(c Command) {
    r.command = c
}

func (r *RemoteControl) PressButton() {
    if r.command != nil {
        r.command.Execute()
    }
}

func main() {
    // Receiver: a light in the living room.
    livingRoomLight := &Light{location: "Living Room"}

    // Concrete Commands.
    lightOn := &LightOnCommand{light: livingRoomLight}
    lightOff := &LightOffCommand{light: livingRoomLight}

    // Invoker: remote control.
    remote := &RemoteControl{}

    // Turn the light on.
    remote.SetCommand(lightOn)
    remote.PressButton()

    // Turn the light off.
    remote.SetCommand(lightOff)
    remote.PressButton()
}
```

**Expected Output:**
```
Living Room light is ON
Living Room light is OFF
```

---

## Advantages of the Command Pattern

1. **Decoupling:** Separates the sender (Invoker) from the receiver, allowing independent evolution.
2. **Flexibility:** New commands can be added without modifying existing code.
3. **Undo/Redo Support:** Commands can be stored and reversed to implement undo functionality.
4. **Command Queuing:** Requests can be queued or logged for later execution.
5. **Composite Commands:** Multiple commands can be combined into a macro command.

---

## Disadvantages of the Command Pattern

1. **Proliferation of Classes:** Each command is a separate class, which can lead to an increased number of classes.
2. **Increased Complexity:** The abstraction may be overkill for simple operations.
3. **Overhead:** Encapsulating each request as an object may add memory and processing overhead in high-frequency scenarios.

---

## Conclusion

The Command Design Pattern provides a robust way to decouple the requester of an action from the object that performs it. By encapsulating requests as objects, the pattern supports flexible command management, including queuing, logging, and undoable operations. While it may increase the number of classes and add some complexity, the benefits in decoupling and flexibility make it a valuable tool in designing scalable and maintainable systems.
