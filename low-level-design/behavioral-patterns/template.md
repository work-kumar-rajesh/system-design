# Template Design Pattern

The Template Design Pattern (or Template Method Pattern) is a behavioral design pattern that defines the skeleton of an algorithm in a method, deferring some steps to subclasses. This allows subclasses to redefine certain steps of an algorithm without changing its overall structure, promoting code reuse and consistency.

---

## Key Concepts of the Template Pattern

1. **Template Method:** Defines the skeleton of an algorithm. It calls primitive operations (steps) that may be implemented in subclasses.
2. **Abstract Class/Interface:** Declares abstract operations (hooks) that subclasses must implement.
3. **Concrete Class:** Implements the abstract operations and may override default behavior provided by the abstract class.
4. **Hook Methods:** Optional methods with default behavior that can be overridden by subclasses.

---

## Structure of the Template Pattern

- **Abstract Class/Interface:** Defines the template method and the abstract steps.
- **Concrete Classes:** Provide specific implementations for the abstract steps.
- **Client:** Uses the concrete classes without needing to know the underlying algorithm details.

---

## When to Use the Template Pattern

- When you have an algorithm that can be decomposed into a series of steps and you want to allow subclasses to provide their own implementations for one or more of these steps.
- To enforce a common workflow across different subclasses while allowing them to vary certain steps.
- When you want to avoid code duplication by placing common code in a single template method.

---

## Real-Life Example in Software Systems: Data Parsing

Consider a scenario where different types of data files (e.g., CSV, JSON) need to be parsed. While the overall process remains the same (read data, process data, write data), the specifics of each step can vary depending on the file format. The Template Pattern allows you to define the workflow once and let subclasses handle the specific details.

### Example in Go

```go
package main

import "fmt"

// DataParser defines the abstract operations for parsing data.
type DataParser interface {
    readData() string
    processData(data string) string
    writeData(data string)
}

// Parse is the template method that defines the skeleton of the algorithm.
func Parse(parser DataParser) {
    data := parser.readData()
    processed := parser.processData(data)
    parser.writeData(processed)
}

// CSVParser is a concrete implementation for parsing CSV data.
type CSVParser struct{}

func (p *CSVParser) readData() string {
    return "raw, csv, data"
}

func (p *CSVParser) processData(data string) string {
    return "processed csv data"
}

func (p *CSVParser) writeData(data string) {
    fmt.Println("CSVParser writes:", data)
}

// JSONParser is another concrete implementation for parsing JSON data.
type JSONParser struct{}

func (p *JSONParser) readData() string {
    return `{"data": "raw json data"}`
}

func (p *JSONParser) processData(data string) string {
    return "processed json data"
}

func (p *JSONParser) writeData(data string) {
    fmt.Println("JSONParser writes:", data)
}

func main() {
    fmt.Println("Using CSV Parser:")
    var csvParser DataParser = &CSVParser{}
    Parse(csvParser)

    fmt.Println("\nUsing JSON Parser:")
    var jsonParser DataParser = &JSONParser{}
    Parse(jsonParser)
}
```

**Expected Output:**
```
Using CSV Parser:
CSVParser writes: processed csv data

Using JSON Parser:
JSONParser writes: processed json data
```

---

## Advantages of the Template Pattern

1. **Code Reuse:** Common parts of the algorithm are implemented in the template method, reducing code duplication.
2. **Consistency:** Enforces a consistent algorithm structure across multiple implementations.
3. **Flexibility:** Allows subclasses to vary the behavior of individual steps without altering the overall algorithm.
4. **Ease of Maintenance:** Changes to the common workflow need to be made only in the template method.

---

## Disadvantages of the Template Pattern

1. **Inflexibility in the Algorithm:** The overall structure is fixed by the template method, limiting how much the algorithm can be altered.
2. **Subclass Dependency:** Subclasses must adhere to the structure defined by the abstract class, which may lead to tightly coupled designs.
3. **Difficulty in Understanding Flow:** New developers may find it challenging to trace the flow of control through the template and its overridden methods.

---

## Real-World Scenario Example: Data Processing Pipelines

Imagine a data processing pipeline where data needs to be ingested, transformed, and then stored. The steps for reading and writing data may be common, but the transformation logic can vary significantly between different data formats (e.g., CSV vs. JSON). The Template Pattern lets you define the overall pipeline in one place while allowing specialized transformation logic to be implemented in separate classes.

---

## Conclusion

The Template Design Pattern provides a structured way to define the skeleton of an algorithm, allowing specific steps to be deferred to subclasses. It promotes code reuse and consistency while giving subclasses the flexibility to implement or override parts of the process. This pattern is particularly valuable in scenarios where the overall process is fixed but the details may vary, such as in data parsing, processing pipelines, and similar workflows.
