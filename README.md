# ExtendedMap in Go

`ExtendedMap` is a generic map implementation in Go that provides additional functionalities such as Push, Set, Pop,
Get, and Sort. It leverages Go's generics to provide a type-safe and flexible map structure.

## Features

- **Push**: Adds a new key-value pair to the map.
- **Set**: Updates the value of an existing key.
- **Get**: Retrieves the value of a given key.
- **Pop**: Removes a key-value pair from the map.
- **Sort**: Returns a sorted slice of key-value pairs.

## Usage

### Importing the Package

Make sure to include the `ExtendedMap` implementation in your project.

### Creating a New ExtendedMap

```go
m := NewExtendedMap[string, int]()
```

### Adding Key-Value Pairs

```go
m.Push("a", 10)
m.Push("b", 20)
```

### Updating a Value

```go
m.Set("a", 30)
```

### Retrieving a Value

```go
value, exists := m.Get("a")
```

### Removing a Key-Value Pair

```go
value, exists := m.Pop("a")
```

### Sorting Key-Value Pairs

```go
sortedPairs := m.Sort()
```

## Testing

Comprehensive unit tests are provided to ensure the correct behavior of the `ExtendedMap`. Refer to the test file for
detailed test cases.

