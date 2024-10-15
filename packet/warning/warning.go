package warning

type Warning string

const (
	CustomWarning                  Warning = "customWarning"
	PlayerAlreadyMadeActionWarning Warning = "playerAlreadyMadeActionWarning"
	ActionIgnoredDueToDeadWarning  Warning = "actionIgnoredDueToDeadWarning"
	SlowResponseWarning            Warning = "slowResponseWarning"
)
