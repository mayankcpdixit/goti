# govim

A tool to copy base golang template of your choice when you try to vim a new file. Takes care of duplicate filename. Takes space separated file names and created a snake_case.go file name. 


## Install
1. Clone the repo. Run `make` and open new tab.
2. Run `govim test`

## Usage
1. `govim test` -> created `test.go` with default template
2. `govim test` again -> will open the existin file if filename already exists
3. `govim "This is The file Name"` -> will create `this_is_the_file_name.go` with default template
