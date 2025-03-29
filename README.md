## Goretry

The **goretry** package provides a flexible and configurable mechanism for retrying operations in Go. It supports different backoff strategies and allows customization of retry behavior through options.  
Objective
The main objective of the **goretry** package is to provide a simple and extensible way to retry operations that may fail intermittently. It supports various backoff strategies such as constant, linear, and exponential backoff, and allows users to define custom retry conditions.  
Installation
To install the goretry package, use the following command:

```bash
go get github.com/masilvasql/goretry
```
_______

### Usage

_**There are examples inside the examples folder.**_

<br>

#### Basic Usage
Here is a basic example of how to use the **goretry** package to retry an operation with default settings:

```go
package main

import (
	"context"
	"fmt"
	"github.com/masilvasql/goretry"
)

func main() {
	ctx := context.Background()
	retryFunc := func(ctx context.Context) (string, error) {
		// Your operation logic here
		return "success", nil
	}

	result, err := goretry.Do(ctx, retryFunc)
	if err != nil {
		fmt.Println("Operation failed:", err)
	} else {
		fmt.Println("Operation succeeded:", result)
	}
}
```
<br>

##### Custom Backoff Strategy

You can customize the backoff strategy using the provided options. Here is an example using exponential backoff:

```go
package main

import (
	"context"
	"fmt"
	"github.com/masilvasql/goretry"
	"time"
)

func main() {
	ctx := context.Background()
	retryFunc := func(ctx context.Context) (string, error) {
		// Your operation logic here
		return "success", nil
	}

	result, err := goretry.Do(ctx, retryFunc,
		goretry.WithBackoffStrategy(goretry.ExponentialBackoff(500*time.Millisecond, 2.0)),
		goretry.WithMaxRetries(5),
	)
	if err != nil {
		fmt.Println("Operation failed:", err)
	} else {
		fmt.Println("Operation succeeded:", result)
	}
}
```

#### Custom Retry Condition

You can also define custom conditions for retrying the operation. Here is an example:

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/masilvasql/goretry"
)

func main() {
	ctx := context.Background()
	retryFunc := func(ctx context.Context) (string, error) {
		// Your operation logic here
		return "", errors.New("temporary error")
	}

	result, err := goretry.Do(ctx, retryFunc,
		goretry.WithShouldRetry(func(err error) bool {
			// Custom retry condition
			return err.Error() == "temporary error"
		}),
	)
	if err != nil {
		fmt.Println("Operation failed:", err)
	} else {
		fmt.Println("Operation succeeded:", result)
	}
}
```

______

#### Backoff Strategies
The **goretry** package provides several built-in backoff strategies:

* Constant Backoff: Retries the operation after a fixed duration.
* Linear Backoff: Retries the operation with a delay that increases linearly with each attempt.
* Exponential Backoff: Retries the operation with a delay that increases exponentially with each attempt.
Example of Constant Backoff

```go
backoff := goretry.ConstantBackoff(500 * time.Millisecond)
```

Example of Linear Backoff

```go
backoff := goretry.LinearBackoff(500 * time.Millisecond)
```

Example of Exponential Backoff

```go
backoff := goretry.ExponentialBackoff(500 * time.Millisecond, 2.0)
```

______

#### Options
The **goretry** package allows customization through options:

* WithMaxRetries: Sets the maximum number of retry attempts.
* WithBackoffStrategy: Sets the backoff strategy.
* WithShouldRetry: Sets the custom retry condition.

_________
**Conclusion**
The goretry package is a powerful tool for handling retry logic in Go applications. It provides flexibility through various backoff strategies and customizable options, making it suitable for a wide range of use cases.