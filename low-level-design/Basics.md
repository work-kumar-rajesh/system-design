# Design Patterns Overview

**Design patterns** are proven, reusable solutions to common problems in software design. They provide a shared language for developers and improve the design and maintenance of software systems. Design patterns are generally classified into three main categories:

- **Creational Patterns:**  
  Deal with object creation mechanisms. They help create objects in a manner suitable to the situation, enhancing flexibility and reuse.  
  **Examples:** Builder, Factory, Prototype, Singleton.

- **Structural Patterns:**  
  Concerned with how classes and objects are composed to form larger structures, ensuring that if one part changes, the entire structure does not need to be recompiled.  
  **Examples:** Adapter, Bridge, Composite, Decorator, Facade, Flyweight, Proxy.

- **Behavioral Patterns:**  
  Focus on communication between objects, detailing how objects interact and distribute responsibilities.  
  **Examples:** Chain of Responsibility, Command, Interpreter, Mediator, Memento, Observer, State, Strategy, Template Method, Visitor.

---

## Summary Table of Design Patterns

| **Design Pattern**             | **Category**  | **When to Use**                                                                                 | **Where to Use**                                                                                   | **Real-Life Example**                                                    |
|--------------------------------|---------------|-------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------|
| **Builder**                    | Creational    | Complex object creation<br>Many optional parameters<br>Step-by-step construction needed         | When object creation requires a flexible, multi-step process                                       | Building a pizza, constructing an HTTP request                           |
| **Factory**                    | Creational    | When object creation should be abstracted from the client<br>Choosing among related types         | When instantiating objects based on configuration or user input                                    | Payment systems selecting between CreditCard, PayPal, etc.               |
| **Prototype**                  | Creational    | When cloning is more efficient than creating from scratch<br>Resource-intensive instantiation     | In systems requiring many similar objects that share common properties                              | Cloning game characters, document templates                              |
| **Singleton**                  | Creational    | When only one instance is needed<br>Global access to a resource is required                      | For shared resources like loggers or configuration managers                                        | Logger, configuration manager                                            |
| **Adapter**                    | Structural    | When integrating incompatible interfaces<br>Adapting legacy components                          | When wrapping legacy or third-party classes to work with new interfaces                            | Adapting a legacy audio player for modern systems                        |
| **Bridge**                     | Structural    | When abstraction and implementation need to vary independently<br>Decoupling required            | In systems where you want to vary the interface independently of its implementation                | Remote control for different TV brands                                   |
| **Composite**                  | Structural    | When building part-whole hierarchies<br>Treating individual and composite objects uniformly       | In recursive structures like file systems, organizational charts, or GUIs                            | File systems, GUI component trees                                        |
| **Decorator**                  | Structural    | When adding responsibilities to objects dynamically<br>Avoiding subclass explosion              | When enhancing objects at runtime without altering the original code                               | Adding logging or encryption to notifications                            |
| **Facade**                     | Structural    | When a simplified interface is needed<br>Hiding complex subsystem details                        | In systems with multiple interdependent components                                                 | Home theater systems, computer startup sequences                         |
| **Flyweight**                  | Structural    | When many objects share common data<br>Memory usage is critical                                  | In systems where numerous objects share intrinsic state                                            | Text editor character rendering, particle systems                        |
| **Proxy**                      | Structural    | When controlling access is needed<br>Lazy initialization or additional functionality            | When adding a control layer (e.g., caching, security) without modifying the original object          | Virtual proxies for remote services                                      |
| **Chain of Responsibility**    | Behavioral    | When multiple objects can handle a request<br>Decoupling sender from receiver                    | In event processing or support systems where a request should be handled by one of many handlers    | Support ticket escalation (front desk, manager, director)                |
| **Command**                    | Behavioral    | When encapsulating a request as an object<br>Support for queuing or undo/redo operations          | In systems requiring decoupled execution of actions                                                | Remote controls, text editor command history                             |
| **Interpreter**                | Behavioral    | When defining and interpreting a simple language or grammar<br>For domain-specific languages       | In parsing and evaluating expressions or configuration files                                       | Mathematical expression evaluation, simple scripting languages             |
| **Mediator**                   | Behavioral    | When centralizing complex communication is needed<br>Reducing direct dependencies                 | In systems with many interacting objects that need a central coordinator                           | Chat room systems, air traffic control                                   |
| **Memento**                    | Behavioral    | When saving and restoring an object's state<br>Implementing undo/redo functionality                | In systems where maintaining a history of state changes is important                               | Text editor undo feature, game state saving                              |
| **Observer**                   | Behavioral    | When one object’s state change should trigger updates in many others<br>Implementing publish/subscribe | In event-driven systems or GUIs where multiple objects need to react to changes                   | News subscription systems, GUI event handling                            |
| **State**                      | Behavioral    | When an object’s behavior depends on its state<br>Avoiding complex conditional logic               | In systems where an object should change behavior at runtime based on its internal state             | Turnstile gates, vending machines                                        |
| **Strategy**                   | Behavioral    | When multiple interchangeable algorithms exist<br>Selecting an algorithm at runtime              | In systems requiring dynamic selection of algorithms                                               | Tax calculation systems, sorting algorithms                              |
| **Template Method**            | Behavioral    | When the overall algorithm is fixed but steps can be overridden<br>Enforcing a common workflow       | In frameworks where the process is defined but parts need customization                            | Data parsing pipelines, game loops                                       |
| **Visitor**                    | Behavioral    | When operations need to be performed on elements of a complex structure<br>Separating operations from data structure | In systems like compilers or document processing where new operations are added frequently            | Document processing, compiler AST traversal                              |

---


