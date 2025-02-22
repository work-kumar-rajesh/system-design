# Mediator Design Pattern

The Mediator Design Pattern is a behavioral pattern that encapsulates how a set of objects interact. Instead of objects referring to and communicating directly with each other, they communicate through a mediator object. This reduces the dependencies between communicating objects, promoting loose coupling and making the system easier to maintain and extend.

---

## Key Concepts of the Mediator Pattern

1. **Mediator:** An interface that defines how colleagues (components) communicate.
2. **Concrete Mediator:** Implements the mediator interface and coordinates the interaction between colleague objects.
3. **Colleague:** Components that interact with each other indirectly through the mediator.
4. **Loose Coupling:** Colleague objects do not refer to each other explicitly, but instead, communicate via the mediator.

---

## Structure of the Mediator Pattern

- **Mediator Interface:** Declares a method for communicating with colleague objects.
- **Concrete Mediator:** Implements the mediator interface and facilitates the communication between colleagues.
- **Colleague Classes:** Each colleague has a reference to the mediator and uses it to communicate with other colleagues.
- **Client:** Configures the mediator and colleagues, initiating interactions between them.

---

## When to Use the Mediator Pattern

- **Complex Communication:** When you have a complex set of interactions between objects that are difficult to manage directly.
- **Reducing Dependencies:** To reduce the tight coupling between multiple components by centralizing their communication in a mediator.
- **Simplifying Maintenance:** When adding or modifying communication between components, changes need to be made in only one place (the mediator) rather than across all interacting objects.
- **Encapsulating Control Logic:** To encapsulate the control logic for object interactions in a single mediator class.

---

## Real-Life Example in Software Systems: Chat Room

Imagine a chat room where multiple users (colleagues) communicate with each other. Instead of each user sending messages directly to every other user, they send messages to a chat room mediator, which then distributes the messages to all users. This decouples the user objects from each other.

### Example in Go

```go
package main

import "fmt"

// Mediator interface defines the communication method.
type Mediator interface {
    SendMessage(sender string, message string)
    Register(user *User)
}

// ConcreteMediator: ChatRoom mediates communication between users.
type ChatRoom struct {
    users map[string]*User
}

func NewChatRoom() *ChatRoom {
    return &ChatRoom{users: make(map[string]*User)}
}

func (c *ChatRoom) Register(user *User) {
    c.users[user.name] = user
    user.chatRoom = c
}

func (c *ChatRoom) SendMessage(sender string, message string) {
    for name, user := range c.users {
        if name != sender {
            user.Receive(sender, message)
        }
    }
}

// Colleague: User that interacts via the mediator.
type User struct {
    name     string
    chatRoom Mediator
}

func NewUser(name string) *User {
    return &User{name: name}
}

func (u *User) Send(message string) {
    fmt.Printf("%s sends: %s\n", u.name, message)
    u.chatRoom.SendMessage(u.name, message)
}

func (u *User) Receive(sender string, message string) {
    fmt.Printf("%s receives from %s: %s\n", u.name, sender, message)
}

func main() {
    chatRoom := NewChatRoom()

    alice := NewUser("Alice")
    bob := NewUser("Bob")
    charlie := NewUser("Charlie")

    chatRoom.Register(alice)
    chatRoom.Register(bob)
    chatRoom.Register(charlie)

    alice.Send("Hello, everyone!")
    bob.Send("Hi, Alice!")
}
```

**Expected Output:**
```
Alice sends: Hello, everyone!
Bob receives from Alice: Hello, everyone!
Charlie receives from Alice: Hello, everyone!
Bob sends: Hi, Alice!
Alice receives from Bob: Hi, Alice!
Charlie receives from Bob: Hi, Alice!
```

---

## Advantages of the Mediator Pattern

1. **Loose Coupling:** Colleague objects are decoupled, as they communicate only with the mediator.
2. **Centralized Control:** The mediator encapsulates the interaction logic, making it easier to manage complex communication.
3. **Simplified Object Interaction:** Reduces the number of direct connections between objects, leading to more maintainable code.
4. **Flexibility:** Changing the interaction between objects requires modifications only in the mediator.

## Disadvantages of the Mediator Pattern

1. **Centralized Complexity:** The mediator can become overly complex if it handles too many responsibilities.
2. **Performance Overhead:** Communication through the mediator may introduce additional overhead.
3. **Single Point of Failure:** If the mediator fails, the communication between all colleague objects can be disrupted.
4. **Increased Indirection:** The added layer of abstraction may complicate understanding the flow of communication.

---

## Real-World Scenario Example: Air Traffic Control System

In an air traffic control system, multiple airplanes (colleagues) communicate with each other through a control tower (mediator). The control tower manages takeoffs, landings, and traffic flow, ensuring safe and efficient operations without airplanes having direct communication with each other.

---

## Conclusion

The Mediator Design Pattern offers a structured way to manage communication between objects by centralizing the interaction logic within a mediator. This promotes loose coupling and simplifies maintenance, especially in systems with complex communication needs. While it introduces an additional layer of abstraction, the benefits of centralized control and reduced dependencies often outweigh the drawbacks in large-scale systems.
