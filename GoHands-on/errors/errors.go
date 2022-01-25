//41.Custom error
package errors

//donot export as first letter of structure is small
type structure_err struct{
	err_text string
}

func New(err_text string) error{
	e := structure_err{err_text}
	return &e
	//return &structure_err{err_text}
}

func (e *structure_err)Error() string{
	return e.err_text
}