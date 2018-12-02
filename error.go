package first

// Error returns the first non-nil argument or nil.
func Error(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}
