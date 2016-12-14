package kitty

// Command defines what is considered a "command"
type Command interface {
	Checks() Checks
	Process(Context)
}

// Checks defines what to use when checking if for example the command is owner only, admin only, etc.
// Currently it's has "owner only"
type Checks struct {
	OwnerOnly bool
}
