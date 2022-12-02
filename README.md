# Snippetbox

## Routing request
| Method | URL Pattern | Handler    | Action    |
| :---:   | :---: | :---: | :---: |
| ANY | /   | home   | Display the home page   |
| ANY | /snippet/view?id=1    | snipperView   | Display a specific snippet   |
| POST | /snippet/create   | snippetCreate   | Create a new snippet   |
| ANY | /static   | http.FileServer   | Server a specific static file   |