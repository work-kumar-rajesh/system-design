# Interface Segregation Principle (ISP) – Payment Processing Example

The Interface Segregation Principle (ISP) is one of the SOLID principles of object-oriented design. ISP states that no client should be forced to depend on methods it does not use. In practice, this means that large, general-purpose interfaces should be split into smaller, more specific ones so that classes only need to implement functionality that is relevant to them. This reduces coupling, simplifies testing, and improves maintainability.

---

## Detailed Theory

### What is ISP?
- **Definition:** ISP advocates that an interface should only include methods that are relevant to the clients that use it.
- **Goal:** Avoid "fat" interfaces by breaking them into smaller, more focused ones. This ensures that classes do not have to implement methods they don't need.

### Why is ISP Important?
- **Reduced Coupling:** Clients depend only on the methods they use, which minimizes the impact of changes.
- **Enhanced Maintainability:** Smaller interfaces are easier to understand, test, and modify.
- **Improved Reusability:** Focused interfaces allow classes to be reused in different contexts without carrying irrelevant functionality.
- **Clear Separation of Concerns:** Each interface has a distinct responsibility, making the overall system design cleaner.

---

## Real-Life Software Development Example: Payment Processing System

Imagine a payment processing system that supports multiple payment methods such as credit cards and digital wallets. Initially, a single broad interface might be defined to cover all possible operations for processing payments. However, not every payment method requires all operations. For example:
- **Credit Card Payment** requires validation of card details and processing the payment.
- **Digital Wallet Payment** only needs to process the payment without card validation or bank transfer operations.

### Without Applying ISP (Violation)

In the following design, a monolithic interface `PaymentMethod` forces all payment methods to implement methods they don't need.

```go
package main

import (
	"errors"
	"fmt"
)

// PaymentMethod defines a broad interface that includes methods for validating, processing, and initiating transfers.
type PaymentMethod interface {
	ValidateDetails(details string)
	ProcessPayment(amount float64) bool
	InitiateTransfer(account string, amount float64) bool
}

// CreditCardPayment implements PaymentMethod and supports all methods.
type CreditCardPayment struct {
	cardNumber string
}

func (cc *CreditCardPayment) ValidateDetails(details string) {
	cc.cardNumber = details
	fmt.Println("Credit card validated with details:", details)
}

func (cc *CreditCardPayment) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	return true
}

func (cc *CreditCardPayment) InitiateTransfer(account string, amount float64) bool {
	// For credit cards, initiating a transfer might be interpreted as a refund or a chargeback.
	fmt.Printf("Initiating refund of $%.2f to account %s\n", amount, account)
	return true
}

// DigitalWalletPayment also implements PaymentMethod,
// but it does not support bank transfer functionality.
type DigitalWalletPayment struct {
	walletID string
}

func (dw *DigitalWalletPayment) ValidateDetails(details string) {
	// Digital wallets might not require detailed validation.
	dw.walletID = details
	fmt.Println("Digital wallet set with ID:", details)
}

func (dw *DigitalWalletPayment) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing digital wallet payment of $%.2f\n", amount)
	return true
}

func (dw *DigitalWalletPayment) InitiateTransfer(account string, amount float64) bool {
	// Not applicable for digital wallets.
	fmt.Println("Error: Digital wallet does not support bank transfers.")
	return false
}

func main() {
	// Client code expects every PaymentMethod to support all operations.
	var payment PaymentMethod

	// Using CreditCardPayment
	payment = &CreditCardPayment{}
	payment.ValidateDetails("4111-1111-1111-1111")
	payment.ProcessPayment(150.0)
	payment.InitiateTransfer("RefundAccount", 20.0)

	fmt.Println()

	// Using DigitalWalletPayment - note the forced implementation of InitiateTransfer.
	payment = &DigitalWalletPayment{}
	payment.ValidateDetails("DW-12345")
	payment.ProcessPayment(75.0)
	if !payment.InitiateTransfer("RefundAccount", 10.0) {
		fmt.Println("DigitalWalletPayment: Transfer operation is not supported.")
	}
}
```

*Issues with this Design:*
- **Forced Implementation:** `DigitalWalletPayment` is forced to implement `InitiateTransfer` even though it is irrelevant, leading to errors or dummy behavior.
- **Unclear Contracts:** Client code must handle unexpected behavior because not all payment methods truly support the full interface.

---

## With ISP Applied (Refactored Design)

By applying ISP, we break the broad interface into smaller, more specific interfaces. This ensures that each payment method only implements the methods relevant to it.

### Refactored Interfaces:
- **PaymentProcessor:** For processing payments.
- **CardValidator:** For validating card details (only applicable to credit cards).
- **TransferInitiator:** For initiating transfers (only applicable to payment methods that support transfers).

```go
package main

import "fmt"

// PaymentProcessor defines the interface for processing payments.
type PaymentProcessor interface {
	ProcessPayment(amount float64) bool
}

// CardValidator defines the interface for validating card details.
type CardValidator interface {
	ValidateCard(details string)
}

// TransferInitiator defines the interface for initiating transfers.
type TransferInitiator interface {
	InitiateTransfer(account string, amount float64) bool
}

// CreditCardPayment implements PaymentProcessor, CardValidator, and TransferInitiator.
type CreditCardPayment struct {
	cardNumber string
}

func (cc *CreditCardPayment) ValidateCard(details string) {
	cc.cardNumber = details
	fmt.Println("Credit card validated with details:", details)
}

func (cc *CreditCardPayment) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	return true
}

func (cc *CreditCardPayment) InitiateTransfer(account string, amount float64) bool {
	fmt.Printf("Initiating refund of $%.2f to account %s\n", amount, account)
	return true
}

// DigitalWalletPayment implements only PaymentProcessor.
type DigitalWalletPayment struct {
	walletID string
}

func (dw *DigitalWalletPayment) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing digital wallet payment of $%.2f\n", amount)
	return true
}

func (dw *DigitalWalletPayment) SetWalletID(id string) {
	dw.walletID = id
	fmt.Println("Digital wallet set with ID:", id)
}

func main() {
	// Client code for credit card payments requires card validation and transfer capabilities.
	ccPayment := &CreditCardPayment{}
	ccPayment.ValidateCard("4111-1111-1111-1111")
	ccPayment.ProcessPayment(150.0)
	ccPayment.InitiateTransfer("RefundAccount", 20.0)

	fmt.Println()

	// Client code for digital wallet payments only requires payment processing.
	dwPayment := &DigitalWalletPayment{}
	dwPayment.SetWalletID("DW-12345")
	dwPayment.ProcessPayment(75.0)
	// Note: DigitalWalletPayment does not have a transfer capability, and the client is not forced to call it.
}
```

*Benefits of the Refactored Design:*
- **Focused Interfaces:**  
  - `CreditCardPayment` implements all relevant interfaces (`CardValidator`, `PaymentProcessor`, and `TransferInitiator`), since it supports all those operations.
  - `DigitalWalletPayment` implements only `PaymentProcessor`, since it does not support transfer operations.
- **Clear Contracts:**  
  - Client functions that require transfer capabilities can require a `TransferInitiator`, ensuring only suitable types are used.
  - Functions dealing with payment processing simply use `PaymentProcessor`, making substitution predictable.
- **Predictable Behavior:**  
  - There’s no forced implementation of irrelevant methods, so each class behaves according to its specific contract.
  
---

## Explanation of How ISP is Adhered to

### Common Contract Through Interfaces:
- The refactored design defines multiple focused interfaces:
  - **PaymentProcessor:** For processing payments.
  - **CardValidator:** For validating credit card details.
  - **TransferInitiator:** For initiating transfers.
- These interfaces form the “contract” for each payment method. Clients interact only with the interfaces that are relevant to the operation they need.

### Substitutability:
- **For Payment Processing:**  
  Both `CreditCardPayment` and `DigitalWalletPayment` implement `PaymentProcessor`. A client function that requires processing a payment can work with either type.
- **For Card Validation and Transfer Initiation:**  
  Only `CreditCardPayment` implements `CardValidator` and `TransferInitiator`. Thus, if a client requires these operations, it can rely on receiving an object that fully supports the functionality.

### No Unexpected Behavior:
- The client is never forced to call methods that are not applicable. A digital wallet payment object will never be treated as if it supports bank transfers.
- This segregation ensures that substituting one implementation for another (e.g., replacing a `CreditCardPayment` with another implementation of `CardValidator`) does not alter the expected behavior of the program.

### Correct Replacement:
- **Substitutability in Action:**  
  If a client function works with a `PaymentProcessor`, it can safely use any object that implements that interface—ensuring the behavior is consistent regardless of the underlying type.
- **No Unwanted Side Effects:**  
  The refactored design guarantees that each class fulfills the contract it advertises. There is no hidden or unexpected behavior because every method call is supported by the implementing class.

---

## Advantages of Applying ISP

1. **Improved Flexibility:** Classes only implement methods they require, making the system more adaptable.
2. **Simpler, Cleaner Code:** Avoids bloated interfaces and dummy method implementations.
3. **Enhanced Maintainability:** Smaller, focused interfaces are easier to test, understand, and modify.
4. **Decoupling:** Reduces dependencies between different parts of the system.

---

## Disadvantages of Applying ISP

1. **Increased Number of Interfaces:** The design may result in more interfaces, which could increase complexity.
2. **Potential Overhead:** Managing multiple interfaces can add some design overhead.
3. **Design Effort:** Careful planning is required to determine the correct granularity for interfaces.

---

## Conclusion

The Interface Segregation Principle is critical for designing flexible and maintainable systems. In our payment processing example, a broad interface forced digital wallet implementations to include irrelevant methods, causing unpredictable behavior. By splitting the interface into focused contracts—one for processing payments, one for validating cards, and one for initiating transfers—we ensure that each payment method only implements what it needs. This leads to a system where classes can be substituted for one another within the proper context without causing unwanted side effects, fully adhering to ISP.

