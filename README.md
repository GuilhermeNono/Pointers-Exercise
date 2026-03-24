# Go Pointers Study 🚀

A small Go study project focused on pointers, structs, and method receivers.

## What is this project? 📚

This repository contains a set of exercises (`Exercicios/exercicio1.go` to `Exercicios/exercicio9.go`) executed from `main.go`.

Main topics:
- 🔁 Swapping values with pointers
- ➕ Incrementing values through references
- 🧱 Struct initialization and reset
- 🛡️ Safe pointer handling (`nil` checks)
- 📈 Updating values conditionally by reference
- 🧩 Slices with pointer elements
- 🎯 Value receiver vs pointer receiver methods
- 🧠 Double pointers (`**int`)

## Exercises Overview 📝

1. `Swap(a, b *int)` - swap two integers using pointers.
2. `Inc(n *int)` - increment a value through a pointer.
3. `Reset(c *Counter)` - reset struct field values.
4. `NewCounterA()` / `NewCounterB()` - create structs with `new` and `&Type{}`.
5. `SafeInc(n *int)` - increment only when pointer is not `nil`.
6. `UpdateMax(max *int, v int)` - update max value conditionally.
7. `SliceAndPointers()` - use pointers to mutate slice elements.
8. `MoveVal` vs `MovePtr` - compare value and pointer receivers.
9. `AllocIfNil(n **int)` - allocate memory through a double pointer.

## Project Structure 🗂️

```text
Estudo/
  go.mod
  main.go
  Exercicios/
    exercicio1.go
    exercicio2.go
    exercicio3.go
    exercicio4.go
    exercicio5.go
    exercicio6.go
    exercicio7.go
    exercicio8.go
    exercicio9.go
```

## How to run ▶️

From the project root (`Estudo`):

```bash
go run .
```

If you want to refresh dependencies first:

```bash
go mod tidy
go run .
```

## Notes 💡

- This is a learning-oriented project.
- Some exercises are intentionally simple to highlight pointer behavior.
- Feel free to edit `main.go` and run one exercise at a time while studying.

