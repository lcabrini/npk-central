package main

import (
    "log"
    "net/http"
    "html/template"
    "github.com/lcabrini/npk-common"
)

func userList() ([]npk.User, error) {
    sql := `
    select *
    from users
    where status != 'deleted'
    `

    db, err := npk.DBConnection(config)
    if err != nil {
        return nil, err
    }

    rows, err := db.Query(sql)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []npk.User
    var user npk.User
    for rows.Next() {
        err = rows.Scan(&user.Id, &user.Username, &user.Password,
            &user.CreatedAt, &user.Status)
        users = append(users, user)

    }

    return users, nil
}

var userListTpl = `
{{template "base" .}}

{{define "main"}}
<table class="table is-fullwidth">
  <thead>
    <tr>
      <th>Username</th>
      <th>Created</th>
      <th>Status</th>
      <th>Actions</th>
    </tr>
  </thead>

  <tbody>
    {{range .}}
      <tr>
        <td>{{.Username}}</td>
        <td>{{.CreatedAt}}</td>
        <td>{{.Status}}</td>
        <td>
          <a href="#"><i class="fa fa-edit"></i></a>
          <a href="#"><i class="fa fa-trash-alt"></i></a>
        </td>
      </tr>
    {{end}}
  </tbody>
</table>
{{end}}
`

func ListUsers(w http.ResponseWriter, r *http.Request) {
    users, err := userList()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    t, _ := template.New("base").Parse(npk.BaseTemplate)
    t.New("navbar").Parse(npk.Navbar)
    t.New("users").Parse(userListTpl)
    err = t.ExecuteTemplate(w, "users", users)
    if err != nil {
        log.Printf("error: %v", err)
    }
}
