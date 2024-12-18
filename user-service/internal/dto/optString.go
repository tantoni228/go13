package dto

type OptString struct {
	value string
	set   bool
}

func NewOptString(value string) OptString {
	return OptString{
		value: value,
		set:   true,
	}
}

func (os OptString) GetValue() string {
	return os.value
}

func (os OptString) IsSet() bool {
	return os.set
}
