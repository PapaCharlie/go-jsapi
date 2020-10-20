# Typed wrappers around `syscall/js.Value`

This project is intended to automatically generate code around the Javascript APIs to extract typed bindings for the
`js.Value` type that is returned by all calls made using the `syscall/js` package. I've found this to be of great help
with interacting with the DOM.

# Contributing
Contributions are very welcome, please create a corresponding file for the type you are trying to generate bindings for
in the `generator` package, run it with `make` and create a PR that includes the generated code as well.
