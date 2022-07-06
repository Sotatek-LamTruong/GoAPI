package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

type Employee struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Postion string `json:"position"`
}

// var employees = allEmp{
// 	{
// 		ID:      "1",
// 		Name:    "Lam",
// 		Postion: "Intern",
// 	},
// 	{
// 		ID:      "2",
// 		Name:    "Long",
// 		Postion: "Intern",
// 	},
// 	{
// 		ID:      "3",
// 		Name:    "Ha",
// 		Postion: "Developer",
// 	},
// 	{
// 		ID:      "4",
// 		Name:    "Luc",
// 		Postion: "Project Manager",
// 	},
// }
