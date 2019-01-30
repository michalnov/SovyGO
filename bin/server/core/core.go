package core

import (
	"html/template"

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
	t, err := template.ParseFiles("index.html")
	checkErr(err)
	tmpl := template.Must(t, err)
	checkErr(err)
	swap["index"] = tmpl
	t, err = template.ParseFiles("login.html")
	checkErr(err)
	tmpl = template.Must(t, err)
	checkErr(err)
	swap["login"] = tmpl
	if err != nil {
		return err
	}
	c.Templates = swap

	return nil
}
