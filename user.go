package main

import (
    "net/http"
    "html/template"
    "github.com/lcabrini/npk-common"
)

var userList = `
{{template "base" .}}

{{define "main"}}
<p>Test</p>
{{end}}
`

func ListUsers(w http.ResponseWriter, r *http.Request) {
    t, _ := template.New("base").Parse(npk.BaseTemplate)
    t.New("navbar").Parse(npk.Navbar)
    t.New("users").Parse(userList)
    t.ExecuteTemplate(w, "users", nil)
}
