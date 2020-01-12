package main

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

func (p Page) saveData() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0666)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// func main() {
// 	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
// 	p1.saveData()
// 	p2, _ := loadPage("TestPage")
// 	fmt.Println(string(p2.Body))
// }
