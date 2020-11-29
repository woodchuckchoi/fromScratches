package err

type CustomError struct {
	message string
}

func (c *CustomError) Error() string {
	return c.message
}

var (
	InvalidParameterError = &CustomError{
		message: "Invalid Parameters"
	}
	
	
)