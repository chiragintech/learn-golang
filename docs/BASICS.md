# Go Fundamentals

### Installation

- Go runtime

### Benefits of Go

- Scalable backends
- Fast (Single native binary)
- Lightweight (Go Runtime in the executable binary cleans up unused memory at runtime using garbage collector)
- Good DevX
- Concurrent
- Statically typed (types checked at compile time before execution) (Unlike JS, Python - dynamically typed)
- Garbage collected
- Strongly typed (whenever an object is passed from a calling function to a called function, its type must be compatible with the type declared in the called function)

### Comparing Go with other languages

- **HIGHER Execution Speed** than Interpreted (run line by line using an "interpreter"): JS, Python, Ruby, PHP.
- **HIGHER Compilation Speed** (Not execution speed tho) than Compiled (build machine code/executable = compiling before execution): Rust, C, C++, Zig, Java, C#.
- Under compiled languages, most run directly in the CPU ("natively compiled") expect for Java and C# which run on top of a VM. (also Java allocates more on the heap)
- **Go runtime puts its execution speed similar to VM compiled** languages but needs much less memory since it does not run on top of a VM.
- C and Rust use low memory than Go since its gives more control to memory usage.
- Speed = CPU Cycles, Memory management in RAM.

### Basic Syntax

- `package main` tells the Go compiler this code is a standalone program, not a library for other programs to import.
- Package = Project
- Packages can be executable (main) and reusable helpers (other folders).
- `import "fmt"` pulls in the fmt package from the standard library so we can print to the console. Imports code and functions contained in that specific package.

### main and folders

- Program is run through `main.go`
- You can import other folders/"packages" through "repo-name/folder-name".
- Run using "folder-name.function-name".

### Basic Commands

- `go build`
- `go run`
- `go fmt` which formats the code in each file.
- `go install` compiles and installs a package.
- `go get` download raw source code of someone else's package.
- `go test` runs any tests in project.

### Build

- `go build` builds an executable of name of the repo.
- `./learn-golang` runs the executable.
- Compile errors stop build, runtime errors happening while program is running.
- Once you have an executable, you can just use that, no need of the compiler.
- But for JS, you need interpreter + code in production.
- Executable = Build + runtime (handles garbage collection + stack allocation than heap, goroutine scheduling)

### Types in detail (For performance and memory nerds)

- int  int8  int16  int32(rune)  int64 (Signed - postive and negative)
- uint uint8 uint16 uint32 uint64 uintptr (Unsigned - postive only)
- float32 float64
- complex64 complex128 (real + imaginary)
- Number is how many bits is taken to store that variable. 32/64 usually depending on user's env.
- JS just has a "Number" type.

### String Encoding

- In many languages like C, a character is a byte. 
- Using ASCII standard, we can represent 128 characters in 7 bits
- Strings are sequence of bytes.
- Rune = int32 can hold any unicode code point. Use with individual characters.
- Go uses UTF-8 (Unicode Transformation Format) Encoding that uses variable-length encoding - 1 to 4 one byte code units.
- Symbols like emoji can have multiple runes. Even sometimes a U+200D zero-width joiner combining 2 emojis - 1 glyph.
- Go handles wide variety of Unicode characters in strings.

### Variable Declaration

- You can declare outside a function, you cannot declare and assign it outside and use it inside a function.

### Functions

- Main does not need to be below all other helper functions.