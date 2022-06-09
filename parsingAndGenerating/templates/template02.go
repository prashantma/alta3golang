/* Alta3 Research | RZFeeser
   Template - HTML template  */

package main


import (
    "html/template"
    "os"
)


type Todo struct {
    Title string
    Done  bool
}


type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}


func main() {

    // check our template for potential errors with Must
    tmpl := template.Must(template.ParseFiles("index.html")) //index02.html
    
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Laundry", Done: false},
                {Title: "Pull weeds in the garden", Done: true},
                {Title: "Sweep the stairs", Done: true},
            },
        }
        tmpl.Execute(os.Stdout, data)
}
// index02.html
/*
<h1>My TODO list</h1>
<ul>


            <li>Laundry</li>



            <li class="done">Pull weeds in the garden</li>



            <li class="done">Sweep the stairs</li>


</ul>
*/

// index.html
/*
<h1>My TODO list</h1>
<ul>
            <li>Laundry</li>
            <li class="done">Pull weeds in the garden</li>
            <li class="done">Sweep the stairs</li>
</ul>
*/
