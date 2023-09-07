package utils

// ================================================
type Command interface {
  Declare()
  Run(options []string)
}
