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

{{define "toolbar"}}
<div class="container has-text-right">
  <a href="/users/add">
    <i class="fa fa-plus"></i>
  </a>
</div>
{{end}}

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

var userForm = `
{{template "base" .}}

{{define "toolbar"}}{{end}}

{{define "main"}}

<div class="container">
  {{if .}}
    <h1 class="title">Edit User</h1>
  {{else}}
    <h1 class="title">Add User</h1>
  {{end}}

  <form method="post">
    <div class="field">
      <label class="label">Username</label>
      <div class="control">
        <input class="input" name="username" type="text">
      </div>
    </div>

    <div class="field">
      <label class="label">Password</label>
      <div class="control">
        <input class="input" name="password1" type="password"
          placeholder="Password">
      </div>
    </div>

    <div class="field">
      <div class="control">
        <input class="input" name="password2" type="password"
          placeholder="Password confirmation">
      </div>
    </div>

    {{if .}}
    <div class="field">
      <label class="label">Status</label>
      <div class="control">
        <div class="select">
          <select name="status">
            <option value="new">New</option>
            <option value="active">Active</option>
            <option value="inactive">Inactive</option>
          </select>
        </div>
      </div>
    </div>
    {{end}}

    <div class="field">
      <p class="control has-text-right">
        <button class="button is-success">Submit</button>
      </p>
    </div>
  </form>
</div>
{{end}}
`

func AddUser(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        t, _ := template.New("base").Parse(npk.BaseTemplate)
        t.New("navbar").Parse(npk.Navbar)
        t.New("userForm").Parse(userForm)
        err := t.ExecuteTemplate(w, "userForm", nil)
        if err != nil {
            log.Printf("Error: %v", err)
        }

    case "POST":
        break

    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
