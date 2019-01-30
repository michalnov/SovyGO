package core

import (
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
	if err != nil {
		return core, err
	}
	return core, nil
}

func (c *Core) loadTemplates() error {
	var err error
	swap := c.Templates

	swap["index"], err = laodTemplate("index.html")
	if err != nil {
		return err
	}
	swap["login"], err = laodTemplate("login.html")
	if err != nil {
		return err
	}
	swap["register"], err = laodTemplate("register.html")
	if err != nil {
		return err
	}
	swap["register"], err = laodTemplate("register.html")
	if err != nil {
		return err
	}
	c.Templates = swap

	return nil
}

func laodTemplate(path string) (*template.Template, error) {
	absPath, _ := filepath.Abs("build/web_files/" + path)
	t, err := template.ParseFiles(absPath)
	if err != nil {
		return nil, err
	}
	tmpl := template.Must(t, err)
	if err != nil {
		return nil, err
	}
	return tmpl, err
}

//HomeHandler serve main htm page
func (c *Core) HomeHandler(w http.ResponseWriter, r *http.Request) {
	c.Templates["index"].Execute(w, nil)
}
