package core

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	auth "github.com/michalnov/SovyGo/bin/server/modules/authentication"
	conf "github.com/michalnov/SovyGo/bin/server/modules/configuration"
)

//Core --
type Core struct {
	Config    conf.Config
	clients   []session
	Templates map[string]*template.Template
}

type session struct {
	sessionID string
	login     bool
	token     auth.Token
}

//NewCore ---
func NewCore() (Core, error) {
	var core Core
	var err error
	core.Config, err = conf.ReadConfig()
	core.loadTemplates()
	if err != nil {
		fmt.Printf("error while loading")
		panic(err)
		//return core, err
	}
	return core, nil
}

func (c *Core) loadTemplates() error {
	var err error
	c.Templates = make(map[string]*template.Template, 0)
	var swap = make(map[string]*template.Template, 0)

	swap["index"] = laodTemplate("index.html")
	if err != nil {
		return err
	}
	swap["login"] = laodTemplate("login.html")
	if err != nil {
		return err
	}
	swap["register"] = laodTemplate("register.html")
	if err != nil {
		return err
	}
	swap["test"] = laodTemplate("test.html")
	if err != nil {
		return err
	}
	c.Templates = swap

	return nil
}

func laodTemplate(path string) *template.Template {
	absPath, err := filepath.Abs("build/web_files/" + path)
	//fmt.Println(absPath)
	tmpl := template.Must(template.ParseFiles(absPath))
	if err != nil {
		fmt.Println("EEE 22")
		return nil
	}
	//fmt.Println("iii")
	return tmpl
}

//HomeHandler serve main htm page
func (c *Core) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EEE 110")
	//http.ServeFile(w, r, "web_files/test.html")
	c.Templates["index"].Execute(w, "")
}

//LoginHandler serve main htm page
func (c *Core) LoginHandler(w http.ResponseWriter, r *http.Request) {
	c.Templates["login"].Execute(w, nil)
}

//RegisterHandler serve main htm page
func (c *Core) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	c.Templates["Register"].Execute(w, nil)
}

//TestHandler serve main htm page
func (c *Core) TestHandler(w http.ResponseWriter, r *http.Request) {
	c.Templates["test"].Execute(w, nil)
}
