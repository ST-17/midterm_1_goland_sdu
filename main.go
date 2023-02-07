package main

import ("fmt"; "net/http"; "html/template")

type User struct{
	Name string
	Age uint16
	Money int16
	Avg_grades, Happiness float64
	Hobbies []string
}

func (u User) getAllInfo() string{
	return fmt.Sprintf("User name is: %s. He is %d years old and he has money equal to: %d!", u.Name,u.Age,u.Money)
}
func (u *User) setNewName(newName string){
	u.Name = newName
}
func home_page(w http.ResponseWriter, r *http.Request){
	bob := User{"Bob", 20, -50, 3.1, 0.8, []string{"SauBol", "Sing","Football"}}
	// fmt.Fprintf(w, `<h1>Main text</h1>
	// <b>Main text</b>`)
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}
func contact_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Contact page!")

}

func handleRequest(){
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contact_page)
	http.ListenAndServe(":8080", nil)
}

func main(){
	handleRequest()
}