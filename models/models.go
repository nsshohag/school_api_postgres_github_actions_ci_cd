package models

// Student ... that holds data and key = `json:"id"` and so on if not provided then it would be ID
/*
In Go, even if two structs have identical fields
they are considered different types if
they are defined in separate packages.
This is because Go treats types as unique based on their package namespace.
*/
// That's why using a common shared package and creating student data in different package having same types
// Student ... that holds data and key = `json:"id"` and so on if not provided then it would be ID
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class int    `json:"class"`
}
