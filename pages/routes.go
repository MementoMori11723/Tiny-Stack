package pages

type Route map[string]func() string

var Routes = Route{
  "/": func() string {
    return "Hello, World!"
  },
}
