package domain

type Movie struct {
	Name     string
	Category string
	Cast     map[string]interface{}
}
