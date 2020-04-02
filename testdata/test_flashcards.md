### Explain how arrays in GO works differently then C ?
In GO Array works differently than it works in C

- Arrays are values, assigning one array to another copies all the elements
- If you pass an array to a function, it will receive a copy of the array, not a pointer to it
- The size of an array is part of its type. The types [10] int and [20] int are distinct

### How to copy map in Go?
You copy a map by traversing its keys. Unfortunately, this is the simplest way to copy a map in Go:
```go
a := map[string]bool{"A": true, "B": true}
b := make(map[string]bool)
for key, value := range a {
    b[key] = value
}
```

### How do you swap two values? Provide a few examples.
Two values are swapped as easy as this:
```go
a, b = b, a
```
```go
a, b, c = b, c, a
```
The swap operation in Go is guaranteed from side effects. The values to be assigned are guaranteed to be stored in temporary variables before starting the actual assigning, so the order of assignment does not matter.
