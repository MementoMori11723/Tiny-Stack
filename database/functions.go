package database

// test function

type ReturnType struct {
	Data interface{}
}

type contacts struct {
	Name  string
	Email string
}

func Fetch() ReturnType {
	return ReturnType{
		Data: []contacts{
			{"John Doe", "johndoe@none.com"},
			{"Monkey D. Luffy", "monkeydluffy@chad.com"},
		},
	}
}
