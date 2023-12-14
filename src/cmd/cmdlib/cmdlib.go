package cmdlib

// ================================================
type ICommand interface {
  Declare()
  Run([]string)
}

// ================================================
type Command struct {}

func (this *Command) Declare() {/* @virtual */}

func (this *Command) Run(options []string) {/* @virtual */}
