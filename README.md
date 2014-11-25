Bound
=====

Overview
--------

This is a ridiculously simple way to bind a function to a specific context for
use as an `http.Handler`.

Example
-------

```go
type app struct {
  content string
}

func (a app) handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%s", a.content)
}

a := app{
  content: "this is something",
}

http.ListenAndServe(":8080", bound.Handler{a, a.handler})
```

License
-------

3-clause BSD. A copy is included with the source.
