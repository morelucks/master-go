<h1 align="center">
  Master

  <img src="https://go.dev/images/go-logo-blue.svg" alt="Go Logo" width="120"/><br/>
</h1>

> Learn Go step by step, hands-on, modular, and fun.  
![Go Version](https://img.shields.io/badge/Go-1.23+-blue)
![License](https://img.shields.io/badge/license-MIT-green)

---


---

## 👋 Welcome to **Master Go**
Your **personal Go learning lab**, inspired by **Go-Ethereum** and **Cosmos SDK**.

This repository is designed for **learning Go step by step**, experimenting with concepts, and building small projects, all while gradually evolving toward a **professional Go project structure**.

---

## 🌱 What You’ll Learn

* Go basics: variables, functions, structs, and interfaces
* Advanced concepts: concurrency with goroutines, channels
* Building small projects: CLI apps, mini-tools
* Professional Go structure: `cmd/`, `internal/`, `pkg/`

---

## 📁 Folder Structure

```text
master-go/
├── learning/               # Learning lab for Go concepts
│   ├── getting-started/    # Setup, hello world
│   ├── basics/             # Variables, constants, types
│   ├── control-flow/       # If, for, switch
│   ├── arrays/             # Arrays
│   ├── slices/             # Slices
│   ├── strings/            # String handling
│   ├── functions/          # Functions & methods
│   ├── maps/               # Maps Maps (key-value collections) a.k.a dictionaries
│   ├── structs/            # Structs & interfaces
│   ├── file-io/           # File operations (read/write)
│   ├── pointers/           # Pointers & memory
│   ├── oop/                # Object-oriented patterns in Go(method and interfaces)
│   ├── concurrency/        # Goroutines, channels, sync
│   └── practices/          # Practice exercises & challenges
│
├── projects/               # Mini projects (CLI tools, small apps)
│
├── cmd/                    # Entry points for real applications
│   └── mastergod/          # Example main application
│
├── internal/               # Private business logic
│   ├── core/
│   └── services/
│
├── pkg/                    # Shared reusable libraries/utilities
│
└── go.mod                  # Go module file
```

---

## 🚀 How to Run

1. Clone the repo:

```bash
git clone https://github.com/morelucks/master-go.git
cd master-go
```

2. Run any example:

```bash
go run learning/getting-started/main.go
```


---

## 🧩 Examples

### Variables

```go
name := "Ethereum"
age := 10
fmt.Println(name, age)
```

### Functions

```go
func Add(a, b int) int {
    return a + b
}
```

### Concurrency

```go
go func() { fmt.Println("Hello from a goroutine!") }()
```

---

## 💡 Tips

* Keep practicing in `/getting-started`
* Build real apps in `/cmd` + `/internal` + `/pkg`
* Run `go mod tidy` to keep dependencies clean

---

## 🌟 Goal

 **Master Go** aims to help you **learn Go in depth**, experiment confidently, and progress toward **professional-level Go projects** — inspired by Go-Ethereum and Cosmos SDK

---

## 🤝 Contribute

Feel free to add:

* New examples
* Mini projects
* Improvements to folder organization

### 🪪 License

MIT License © 2025 [Kamshak Lucky]
See [LICENSE](LICENSE) for details.
