# first

Hi.  I like to add fallback logic to my programs, but sometimes it's awkward.

I wrote this little library so we can write:

```go
use(first.String(os.Getenv("CONFIG_FIELD_NAME"), "some sort of default"))
```

as opposed to:

```go
configValue := os.Getenv("CONFIG_FIELD_NAME")
if configValue == "" {
    configValue = "some sort of default"
}
use(configValue)
```

First also supports:

 * `frist.String("arbitrary", "numbers", "of", "arguments")`
 * `s := "pointers"; first.String(&s)`
 * `first.String(func() { return "functions returning strings" })`
 * Interfaces with a `String() string` method

Don't `panic()`.
