# Arrays

### Types

#### Array 
- Fixed length list of things
- Primitive data structure
- Rarely used directly
- `[5]int`

#### Slice 
- Array that can grow or shrink
- Every element must be of same type
- length, capacity and ptr to head that attaches to a fixed array - automatically has a pointer that updates stuff
- slice and its corresponding array stays at 2 different addresses
- does pass by value but still points to the original array
