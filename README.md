[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

# Optional

Optional is a simple package which provides Optional[T] type and helper functions for sql.Null* types.

# Usage example

```go
	val := None[int]()
	val = Of(2)
	slc, err := val.MarshalJSON()
	fmt.Println(string(slc), err)

	ptr := val.Ptr()
	fmt.Println(*ptr)
	// Output:
	// 2 <nil>
	// 2
```

# Inspiration

- [gomoni/null](https://github.com/gomoni/null)
- [guregu/null](https://github.com/guregu/null)
