package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tsawler/bookings-app/pkg/config"
	"github.com/tsawler/bookings-app/pkg/models"
	"github.com/tsawler/bookings-app/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

type allEmps []*models.Employee

var emps = allEmps{
	{
		ID:      "1",
		Name:    "Lam",
		Postion: "Intern",
	},
	{
		ID:      "2",
		Name:    "Long",
		Postion: "Intern",
	},
	{
		ID:      "3",
		Name:    "Ha",
		Postion: "Developer",
	},
	{
		ID:      "4",
		Name:    "Luc",
		Postion: "Project Manager",
	},
}

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability handles post
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end is %s", start, end)))
}

type jsonRes struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// GET JSON
func (m *Repository) GetAvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonRes{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")

	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Employee
func (m *Repository) GetEmps(w http.ResponseWriter, r *http.Request) {
	out, err := json.MarshalIndent(&emps, "", "     ")

	if err != nil {
		log.Println(err)
	}
	w.Write(out)
	// json.NewEncoder(w).Encode(emps)
}

func (m *Repository) GetEmp(w http.ResponseWriter, r *http.Request) {
	empID := chi.URLParam(r, "id")

	for _, emp := range emps {
		if emp.ID == empID {
			log.Println(&emp)
			out, err := json.MarshalIndent(emp, "", "     ")

			if err != nil {
				log.Println(err)
			}
			w.Write(out)
			return
		}
	}
	fmt.Fprintln(w, "Not exist")
}

func (m *Repository) AddEmp(w http.ResponseWriter, r *http.Request) {
	var newEmp models.Employee
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEmp)
	emps = append(emps, &newEmp)
	out, err := json.MarshalIndent(&newEmp, "", "     ")

	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(out)

	// json.NewEncoder(w).Encode(newEmp)
}

func (m *Repository) UpdateEmp(w http.ResponseWriter, r *http.Request) {
	empID := chi.URLParam(r, "id")

	var updateEmp models.Employee

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Update fail !")
	}

	json.Unmarshal(reqBody, &updateEmp)

	for _, emp := range emps {
		if emp.ID == empID {
			if updateEmp.Postion != "" || updateEmp.Name != "" {
				if updateEmp.Name == "" {
					emp.Postion = updateEmp.Postion
				} else if updateEmp.Postion == "" {
					emp.Name = updateEmp.Name
				} else {
					emp.Name = updateEmp.Name
					emp.Postion = updateEmp.Postion
				}
			}

			// emp.ID = updateEmp.ID
			// emps = append(emps[:i], emp)

			out, err := json.MarshalIndent(emp, "", "     ")

			if err != nil {
				log.Println(err)
			}
			w.WriteHeader(http.StatusCreated)
			w.Write(out)
			return
		}
	}

}

func (m *Repository) DeleteEmp(w http.ResponseWriter, r *http.Request) {
	empID := chi.URLParam(r, "id")

	for i, emp := range emps {
		if emp.ID == empID {
			emps = append(emps[:i], emps[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", empID)
		}
	}
}
