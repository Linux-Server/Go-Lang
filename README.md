## Create a Module:
 - Create a folder
 - `mkdir main-module`
 - `cd main-module`

 ```go mod init sample.com/main-module```

- It will create a .mod file
- create a main file, which contain the main function
```
package main

func main() {
	fmt.Println("Welcome to main modile")

}

```

- To run the module, use the following command:
- ```go run .  OR  (file_name)```

## How to create and call package in same module (main-module):
- create a folder, to have the package
- Create a folder
 - `mkdir main-package_one`
 - `cd main-package_one`
 -  Create a file with any name (for ex: one.go) and add the following code
 ```
package one

import "fmt"

func One() {
	fmt.Println("One is active")
}
 
 ```
 - Use the following code to create a package inside the same module
 ```
package main

import (
	"fmt"
	"sample.com/main-module/main-package_one"
)

func main() {
	fmt.Println("Welcome to main modile")
	one.One() 

}
 ```

## How to create and call package in different module (second-module):

 - Second module can be exec or not exec

 - Create a folder
 - `mkdir second-module`
 - `cd second-module`

 ```go mod init sample.com/second-module```

 - Lets try the non-exec second module
 - create a file with any name
  - `touch greet.go`
  - Write a Greets function in it
  ```
package greet

import "fmt"

func Greets() {
	fmt.Println("welome Greets")
}

  ```

### Now call the second-module  function from main-module
-  Inorder to call another function from another module, we need to replace directive
- to do that , follow the command
`
go mod edit --replace {name of foreign module} = {path to that foreign module} 
`
```
go mod edit --replace sample.com/second-module=../second-module
```

```
package main

import (
	"fmt"
	"sample.com/second-module"
	"sample.com/main-module/one"
)

func main() {
	fmt.Println("Welcome to main modile")
	one.One()
    greet.Greets()


}

```
- After importing the path, use the tidy command
```
go mod tidy
````

# Workspace

- First create a folder for your workspace
- `mkdir my-workspace`
- `cd my-workspace`

- Next initialize the created folder as workspace , using the following command

```
go work init
```

- This above command will create a go.work file inside your workspace folder
- Now create a module inside your workspace
- `mkdir main-module`
- `cd main-module`
- Next initialize the module
- ```go mod init sample.com/main-module```
- create the entrypoint file (main.go)
```
package main

import (
	"fmt"

)

func main() {
	fmt.Println("This is workpace main module")

}


```

- Now connect this main-module in the workspace by using the following commaned

- `go work use . `  (if u are in main-module)
- `go work use ./main-module `  (if u are in workspace)

 - Now create a second module inside the worspace by following command

 - Create a folder
 - `mkdir second-module`
 - `cd second-module`

 ```go mod init sample.com/second-module```

 - Lets try the non-exec second module
 - create a file with any name
  - `touch second.go`
  - Write a Second function in it
  ```
package second

import "fmt"

func Second() {
	fmt.Println("Hello Second")
}

  ```
- Now connect the second-module with workspace using the following commad
- `go work use . `  (if u are in second-module)
- `go work use ./second-module `  (if u are in workspace)

### Call the Second function from the secon-module to main-module in then same workspace

- go to main-module and inside the main.go file, import like below code

```
package main

import (
	"fmt"

	"workspace.com/second-module"
)

func main() {
	fmt.Println("This is workpace main module")
	second.Second()

}

```

- Now call the entrypoint file of teh workspace using

` go run ./main-module`





