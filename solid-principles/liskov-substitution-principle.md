# Liskov Substitution Principle (LSP) – Document Editing Example

The Liskov Substitution Principle (LSP) is a fundamental principle of object-oriented design stating that objects of a superclass should be replaceable with objects of its subclasses without affecting the correctness of the program. In essence, if class S is a subtype of class T, then objects of type T should be substitutable with objects of type S without altering the desirable properties of the program.

When LSP is violated, client code that relies on a base class may behave unpredictably when a subclass is substituted. In this example, we explore a document editing system where the base interface assumes the ability to edit and save documents. However, a read-only document does not support editing or saving, leading to a violation of LSP. We then refactor the design so that the read-only document is not forced into an interface it cannot fulfill.

---

## Detailed Theory

### What is LSP?
- **Definition:** If S is a subtype of T, then objects of type T should be replaceable with objects of type S without altering the correctness of the program.
- **Goal:** Ensure that subclasses extend base classes without changing their expected behavior, thereby guaranteeing that client code will function correctly regardless of which subclass is used.

### Why LSP is Important in Real-World Systems
- **Reliability:** Client code remains stable when new subtypes are introduced.
- **Maintainability:** Changes in subclass behavior do not cause unexpected issues in systems that depend on the base class.
- **Extensibility:** New types can be added without breaking existing functionality, which is crucial in evolving software systems.

---

## Real-World Scenario: Document Editing System

Imagine a software system that supports editing of documents. Initially, the system defines a single interface for documents that includes methods for editing and saving. A typical document (like a Word document) supports both editing and saving. However, there may also be documents that are read-only (e.g., legal documents or archived files) where editing and saving are not permitted.

### Problem: Violation of LSP

In the non-LSP design, a read-only document is forced to implement an interface that includes `Edit` and `Save` methods. When client code calls these methods, it either does nothing or returns an error, which violates the expectation that any document should support the complete functionality defined by the interface.

#### Non-LSP Implementation (Violation)

```go
package main

import (
	"errors"
	"fmt"
)

// Document defines an interface for editable documents.
type Document interface {
	Edit(content string)
	Save() error
	View() string
}

// WordDocument supports full editing and saving.
type WordDocument struct {
	content string
}

func (wd *WordDocument) Edit(content string) {
	wd.content = content
}

func (wd *WordDocument) Save() error {
	fmt.Println("WordDocument saved:", wd.content)
	return nil
}

func (wd *WordDocument) View() string {
	return wd.content
}

// ReadOnlyDocument is forced to implement Document but cannot truly edit or save.
type ReadOnlyDocument struct {
	content string
}

func (rod *ReadOnlyDocument) Edit(content string) {
	fmt.Println("Error: Cannot edit a read-only document.")
}

func (rod *ReadOnlyDocument) Save() error {
	return errors.New("error: cannot save a read-only document")
}

func (rod *ReadOnlyDocument) View() string {
	return rod.content
}

// Client function that expects any Document to support editing and saving.
func processDocument(doc Document, newContent string) {
	doc.Edit(newContent)
	err := doc.Save()
	if err != nil {
		fmt.Println("Processing error:", err)
	} else {
		fmt.Println("Document processed. Content:", doc.View())
	}
}

func main() {
	wd := &WordDocument{}
	rod := &ReadOnlyDocument{content: "Original Content"}

	fmt.Println("Processing WordDocument:")
	processDocument(wd, "Updated Word Content")

	fmt.Println("\nProcessing ReadOnlyDocument:")
	processDocument(rod, "Attempted Update")
}
```

**Expected Output:**
```
Processing WordDocument:
WordDocument saved: Updated Word Content
Document processed. Content: Updated Word Content

Processing ReadOnlyDocument:
Error: Cannot edit a read-only document.
Processing error: error: cannot save a read-only document
```

*Issue:* The client code treats both `WordDocument` and `ReadOnlyDocument` as fully editable documents. However, the read-only document cannot be edited or saved, leading to unexpected behavior.

---

## Refactored Design – Adhering to LSP

To adhere to LSP, we segregate the interfaces so that only editable documents expose methods for editing and saving. We introduce two interfaces:
- **ViewableDocument:** Contains methods for viewing the document.
- **EditableDocument:** Extends ViewableDocument with editing and saving capabilities.

This ensures that a read-only document implements only `ViewableDocument`, while a regular document implements `EditableDocument`. Client code that needs editing functionality will only operate on `EditableDocument`.

### Refactored Implementation

```go
package main

import "fmt"

// ViewableDocument defines an interface for documents that can be viewed.
type ViewableDocument interface {
	View() string
}

// EditableDocument extends ViewableDocument with editing and saving capabilities.
type EditableDocument interface {
	ViewableDocument
	Edit(content string)
	Save() error
}

// WordDocument implements EditableDocument.
type WordDocument struct {
	content string
}

func (wd *WordDocument) Edit(content string) {
	wd.content = content
}

func (wd *WordDocument) Save() error {
	fmt.Println("WordDocument saved:", wd.content)
	return nil
}

func (wd *WordDocument) View() string {
	return wd.content
}

// ReadOnlyDocument implements only ViewableDocument.
type ReadOnlyDocument struct {
	content string
}

func (rod *ReadOnlyDocument) View() string {
	return rod.content
}

// Client function that works with EditableDocument.
func processEditableDocument(doc EditableDocument, newContent string) {
	doc.Edit(newContent)
	err := doc.Save()
	if err != nil {
		fmt.Println("Error saving document:", err)
	} else {
		fmt.Println("Editable document processed. Content:", doc.View())
	}
}

// Client function that works with ViewableDocument.
func displayDocument(doc ViewableDocument) {
	fmt.Println("Document content:", doc.View())
}

func main() {
	// For editable documents.
	wordDoc := &WordDocument{content: "Initial Word Content"}
	processEditableDocument(wordDoc, "Updated Word Content")

	// For read-only documents.
	roDoc := &ReadOnlyDocument{content: "Read-Only Content"}
	displayDocument(roDoc)
}
```

**Expected Output:**
```
WordDocument saved: Updated Word Content
Editable document processed. Content: Updated Word Content
Document content: Read-Only Content
```

### How This Adheres to LSP

- **Substitutability:**  
  - The `EditableDocument` interface is only implemented by documents that truly support editing and saving (like `WordDocument`).
  - Client functions that require editing functionality only accept `EditableDocument`, ensuring that any passed object meets the expected contract.
- **Separation of Concerns:**  
  - `ReadOnlyDocument` does not implement methods it cannot support, thus it is not forced to fulfill a contract that doesn’t make sense for it.
- **Predictable Behavior:**  
  - When the client code calls `processEditableDocument`, it can be assured that the document can be edited and saved. There’s no risk of inadvertently passing a read-only document into a context where editing is expected.

---

# Explanation

## Common Contract Through Interfaces

In the refactored design, we define two interfaces:

- **ViewableDocument:**  
  Provides a method to view the document.

- **EditableDocument:**  
  Extends `ViewableDocument` by adding editing and saving functionalities.

These interfaces form the "contract" that any document type must adhere to. Client code interacts with these interfaces without needing to know the details of the underlying implementation.

---

## Substitutability

### For Viewing:
- **Any document**, whether it’s a `WordDocument` (editable) or a `ReadOnlyDocument`, implements the `ViewableDocument` interface.
- **Example:**  
  A function that accepts a `ViewableDocument` can work with both a `WordDocument` and a `ReadOnlyDocument` without any issues. This is because both provide a `View()` method, and the client code doesn't rely on editing functionality.

### For Editing:
- **Only documents that can be edited** implement the `EditableDocument` interface.
- **Example:**  
  A function that requires editing and saving will only accept an object that implements `EditableDocument` (like `WordDocument`). This guarantees that when client code calls `Edit()` or `Save()`, it will behave as expected.

---

## No Unexpected Behavior

- **In the original (LSP-violating) design:**  
  The `ReadOnlyDocument` was forced to implement `Edit()` and `Save()`, even though these operations weren’t applicable. When substituted into a context expecting a fully editable document, the behavior was unexpected (e.g., errors or no-ops).

- **In the refactored design:**  
  The responsibilities are clearly separated:
  - Editable documents (like `WordDocument`) adhere to the full contract of `EditableDocument`.
  - Read-only documents implement only `ViewableDocument`.

Thus, if a client expects a document that can be edited (an `EditableDocument`), it will never be given a read-only document, and vice versa. This ensures that objects of a subtype (e.g., `WordDocument` as an `EditableDocument`) can be used in any context where their supertype is expected without altering the program's correctness.

---

## Correct Replacement

### Substitutability Principle in Action:
- If we have a function that works with the `ViewableDocument` interface, you can replace a `WordDocument` (which is an `EditableDocument`, and therefore also a `ViewableDocument`) with a `ReadOnlyDocument` without any negative impact. This is because both provide the same `View()` functionality.

### No Unwanted Side Effects:
- The behavior expected by the client is preserved because the subclass (e.g., `WordDocument`) conforms fully to the expectations set by the interface (`EditableDocument`). There is no hidden behavior that could disrupt the system.



## Advantages of the Refactored Design

1. **Reliability:** Client code only uses classes that meet the required behavior, avoiding unexpected errors.
2. **Maintainability:** Each class has clear, distinct responsibilities, making the code easier to modify and extend.
3. **Extensibility:** New document types (e.g., hybrid documents that are editable but have some read-only sections) can be introduced without breaking existing functionality.
4. **Robustness:** The system remains robust as the invariants of each interface are strictly maintained.

---

## Conclusion

The Liskov Substitution Principle is essential for creating robust and maintainable systems. In our document editing example, forcing a read-only document into an editable interface led to unpredictable behavior. By segregating the interfaces into `EditableDocument` and `ViewableDocument`, we ensure that each document type only exposes the functionality it can support. This adherence to LSP not only prevents bugs but also promotes a design that is easier to extend and maintain in real-world software development.

