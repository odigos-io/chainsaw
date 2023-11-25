package v1alpha1

// Script describes a script to run as a part of a test step.
type Script struct {
	// Content defines a shell script (run with "sh -c ...").
	// +optional
	Content string `json:"content,omitempty"`

	// SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.
	// +optional
	SkipLogOutput bool `json:"skipLogOutput,omitempty"`

	// Check is an assertion tree to validate the operation outcome.
	// +optional
	Check *Check `json:"check,omitempty"`
}
