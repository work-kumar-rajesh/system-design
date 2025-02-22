# Adapter Design Pattern

The Adapter Design Pattern is a structural design pattern that allows objects with incompatible interfaces to collaborate. It works as a bridge between two incompatible interfaces by converting the interface of one class (the **Adaptee**) into an interface expected by the client (the **Target**). This enables legacy or third-party classes to work with new systems without modifying their source code.

## Key Concepts of the Adapter Pattern

1. **Target:** The interface that the client expects.
2. **Adaptee:** The existing class with an incompatible interface.
3. **Adapter:** The class that implements the Target interface and wraps an instance of the Adaptee, translating calls into the format that the Adaptee expects.
4. **Client:** The code that uses the Target interface without needing to know about the Adaptee.

---

## Structure of the Adapter Pattern:
1. **Target Interface:** Defines the domain-specific interface that the client uses.
2. **Adaptee:** A class that provides useful functionality but with a non-compatible interface.
3. **Adapter:** Implements the Target interface and internally calls the Adaptee's methods to perform the conversion.
4. **Client:** Interacts with objects through the Target interface.

---

## When to Use the Adapter Pattern:
- When you need to use an existing class but its interface doesn't match the one you require.
- When integrating legacy components or third-party libraries with your system.
- When you want to provide a unified interface for a set of classes with different interfaces.
- When you need to convert data or behavior from one interface to another without changing the original class.

---

## Real-Life Example in Software Systems: Audio Playback

Consider an audio playback system where the client expects a standard interface for playing audio files. However, you have a legacy audio player that uses a different method signature. The Adapter pattern allows you to wrap the legacy audio player and make it compatible with the expected interface.

### Without the Adapter Pattern (Direct Integration)
```go
package main

import "fmt"

// Target interface that the client expects
type AudioPlayer interface {
    Play(fileName string)
}

// ModernAudioPlayer implements the target interface
type ModernAudioPlayer struct{}

func (m *ModernAudioPlayer) Play(fileName string) {
    fmt.Println("Playing modern audio file:", fileName)
}

// LegacyAudioPlayer is an existing class with an incompatible interface
type LegacyAudioPlayer struct{}

func (l *LegacyAudioPlayer) PlayAudio(fileName string) {
    fmt.Println("Playing legacy audio file:", fileName)
}

func main() {
    // Client uses ModernAudioPlayer directly
    var player AudioPlayer = &ModernAudioPlayer{}
    player.Play("song_modern.mp3")
    
    // LegacyAudioPlayer cannot be used directly because its interface differs:
    legacyPlayer := &LegacyAudioPlayer{}
    legacyPlayer.PlayAudio("song_legacy.mp3") // Incompatible with AudioPlayer interface
}
```

**Issues:**
- The client must handle incompatible interfaces.
- LegacyAudioPlayer cannot be directly integrated with code expecting the AudioPlayer interface.

### With the Adapter Pattern
```go
package main

import "fmt"

// AudioPlayer is the target interface
type AudioPlayer interface {
    Play(fileName string)
}

// ModernAudioPlayer implements the AudioPlayer interface
type ModernAudioPlayer struct{}

func (m *ModernAudioPlayer) Play(fileName string) {
    fmt.Println("Playing modern audio file:", fileName)
}

// LegacyAudioPlayer is the existing class with a different interface
type LegacyAudioPlayer struct{}

func (l *LegacyAudioPlayer) PlayAudio(fileName string) {
    fmt.Println("Playing legacy audio file:", fileName)
}

// AudioAdapter adapts LegacyAudioPlayer to the AudioPlayer interface
type AudioAdapter struct {
    legacyPlayer *LegacyAudioPlayer
}

func (a *AudioAdapter) Play(fileName string) {
    // Delegate call to the legacy player's method
    a.legacyPlayer.PlayAudio(fileName)
}

func main() {
    // Client uses ModernAudioPlayer directly
    var player AudioPlayer = &ModernAudioPlayer{}
    player.Play("song_modern.mp3")
    
    // Client can now also use LegacyAudioPlayer via the AudioAdapter
    var adaptedPlayer AudioPlayer = &AudioAdapter{legacyPlayer: &LegacyAudioPlayer{}}
    adaptedPlayer.Play("song_legacy.mp3")
}
```

---

## Advantages of the Adapter Pattern:
1. **Interface Compatibility:** Bridges incompatible interfaces so that legacy and new code can work together.
2. **Reusability:** Allows the use of legacy or third-party classes without modifying their source code.
3. **Decoupling:** Separates the client code from the implementation details of the adaptee.
4. **Flexibility:** Supports multiple adapters to handle different incompatible interfaces.

## Disadvantages of the Adapter Pattern:
1. **Increased Complexity:** Introduces additional layers of abstraction.
2. **Performance Overhead:** May add a slight overhead due to extra method calls.
3. **Limited Exposure:** The adapter might not expose all functionalities of the adaptee.

---

## Conclusion:
The Adapter Design Pattern is an effective solution for integrating classes with incompatible interfaces. By wrapping an existing class with an adapter, you can reuse legacy components, integrate third-party libraries, and maintain a consistent interface for client code. This pattern promotes decoupling and flexibility in system design, making it easier to evolve your codebase without significant refactoring.
