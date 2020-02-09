package main

import (
    "time"
    "log"
    "net/http"
    "html/template"
    "github.com/google/uuid"
    "github.com/lcabrini/npk-common"
)

type Branch struct {
    Id uuid.UUID
    Name string
    CreatedAt time.Time
    Status string
}

func branchList() ([]Branch, error) {
    sql := `
    select *
    from branches
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

    var branches []Branch
    var branch Branch
    for rows.Next() {
        err = rows.Scan(&branch.Id, &branch.Name, &branch.CreatedAt,
            &branch.Status)
        branches = append(branches, branch)
    }

    return branches, nil
}

var branchListTpl = `
{{template "base" .}}

{{define "toolbar"}}
<div class="container has-text-right">
  <a href="/branches/add">
    <i class="fa fa-plus"></i>
  </a>
</div>
{{end}

{{define "main"}}
<table class="table is-fullwidth">
  <thead>
    <tr>
      <th>Name</th>
      <th>Created</th>
      <th>Status</th>
    </tr>
  </thead>

  <tbody>
    {{range .}}
    <tr>
      <td>{{.Name}}</td>
      <td>{{.CreatedAt}}</td>
      <td>{{.Status}}</td>
      <td>
        <a href="/branches/edit/{{.Id}}">
          <i class="fa fa-edit"></i>
        </a>
        <a href="/branches/delete/{{.Id}}">
          <i calss="fa fa-trash-alt"></i>
        </a>
      </td>
    </tr>
    {{end}}
  </tbody>
</table>
{{end}}
`

func ListBranches(w http.ResponseWriter, r *http.Request) {
    branches, err := branchList()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    t, _ := template.New("base").Parse(npk.BaseTemplate)
    t.New("navbar").Parse(npk.Navbar)
    t.New("branches").Parse(branchListTpl)
    err = t.ExecuteTemplate(w, "branches", branches)
    if err != nil {
        log.Printf("error: %v", err)
    }
}
