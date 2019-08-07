package main

import (
   "net/http"
   "fmt"
   "time"
   "html/template"
)

type Welcome struct {
  Name string
  Time string
}

func main() {
  welcome := Welcome { "there", time.Now().Format(time.Stamp) }
  templates := template.Must( template.ParseFiles("frontend/index.html") )

  http.Handle("/frontend/", http.StripPrefix( "/frontend/", http.FileServer( http.Dir("frontend") ) ) )

  http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
    if name := r.FormValue("name"); name != "" {
       welcome.Name = name
    } else {
      welcome.Name = "there"
    }


    if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  fmt.Println(http.ListenAndServe(":8080", nil))
  fmt.println("Serving on port 8080")
}
