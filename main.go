//// 2023/10/24 // 13:48 //

//// Packages

// Every Go program is made up of packages.

// You have probably noticed the mackage main at the top of all the programs you have
// been writing.

// A package names "main" has an entrypoint at the main() function. A main package is
// compiled into an executable program.

// // A package by any other name is a "library package". Libraries have no entry point.
// // Libraries simply export functionality that can be used by other packages. For example:

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	fmt.Println("My favorite number is", rand.Intn(10))
// }

// This program is an executable. It is a "main" package and imports from the fmt and
// math/rand library package.

// // Assignment

// // Fix the bug in the code

// // package mailio
// package main

// import (
// 	"fmt"
// )

// func test(text string) {
// 	fmt.Println(text)
// }

// func main() {
// 	test("starting Mailio server")
// 	test("stopping Mailio server")
// }

//// 13:52

//// Package Naming

// // Naming Convention

// // By comvention, a package's name is the same as the last element of its import path. For
// // instance, the main/rand package comprises files that begin with:
// package rand

// // That said, package names aren't required to match their import path. For example, I
// // could write a new package with the path github.com/mailio/rand and name the
// // package random:
// package random

// While in the above is possible, it is discouraged for the sake of consistency.

// One Package / Directory

// A directoy of Go code can have at most one package. All .go files in a single directory
// must all belong to the same package. If they don't an error will be thrown by the
// compiler. This is true for main and library packages alike.

//// 13:55

//// Modules

// Go programs are organized into packages. A package is a directory of Go code that's all
// compiled together. Functions, types, variables, and constants defined in one source file
// are visible to all other source files within the same package (directory).

// A repository contains one or more modules. A module is a collection of Go packages
// that are released together.

// A Go repository typically contains only one module,
// located at the root of the repository.

// A file names go.mod at the root of a project declares the module. It contains:
// * The module path
// * The version of the Go language your project requires
// * Optionally, any external package dependencies your project has

// // The module path is just the import path prefix for all packages within the module.
// // Here's an example of a go.mod file:

// module github.com/bootdotdev/exampleproject

// go 1.20

// require github.com/google/examplepackage v1.3.0

// Each module's path not only serves as an import path prefix for the packages within but
// also indicates where the go coammnd should look to download it. For example, to
// download the module golang.org/x/tools, the go command would consult the
// repository located at https://golang.org/x/tools.

// An "import path" is a string used to import a package. A package's import path is
// its module path joined with its subdirectory within the module. For example, the
// module github.com/google/go-cmp contains a package in the directory cmp/ . That
// package's import path is github.com/google/go-cmp/cmp . Packages in the standard
// library do not have a module path prefix.

// Do I need to put my package on GitHub?

// You don't need to publish your code to a remote repository before you can build it. A
// module can be defined locally without belonging to a repository. However, it's a good
// habit to keep a copy of all your projects on a remote server, like GitHub.

//// 14:03
//// 5-gopath
//// Setting up your machine

// You machine will contain many version control repositories (managed by Git, for
// example).

// Each repository contains one or more packages, but will typically be a single module.

// Each package consists of one or more Go source files in a single directory.

// The path to a package's directory determines its import path and where it can be
// downloaded form if you decide to host it on a remote version control system like
// Github or Gitlab.

//// A note on GOPATH

// The %GOPATH environment variable will be set by default somewhere on your machine
// (typically in the home directoy, ~/go). Since we will be working in the new
// "Go modules" setup, you don't need to worry about that. If you read something online about
// setting up your GOPATH, that documentation is probably out of date.

// These days you should avoid working in the $GOPATH/src directory. Again, that's the
// old way of doing things and can cause unexpected issues, so better to just avoid it.

// // Get into your workspace

// // Navigate to a location on your machine where you want to store some code. For
// // example, I store all my code in ~/workspace, then organize it into subfolders based on
// // the remote location. For example,
// ~/workspace/github.com/wagslane/go-password-validator
// https://github.com/wagslane/go-password-validator
// // That said, you can put your code wherever you want

// //// First Local Program

// // Once inside your personal workspace, create a new directory and enter it:
// mkdir hellogo
// cd hellogo

// // Inside the directory declare your module's name:
// go mod init {REMOTE}/{USERNAME}/hellogo
// // Where {REMOTE} is your preferred remote source provider (i.e. github.com) and
// // {USERNAME} is your Git username. If you don't use a remote provider yet, just use
// // example.com/username/hellogo

// // Print your go.mod file:
// cat go.mod

//// 14:14
//// Go Run

// // Iniside hellogo, create a new file called main.go.

// // Conventionally, the file in the main package that contains the main() function is
// // called main.go.

// // Paste the following code into your file:
// package main

// import "fmt"

// func main() {
// 	fmt.Println("hello world")
// }

// // Run the code
// go run main.go

// // The go run command is used to quickly compile and run a Go package. The compiled
// // binary is not saved in your working directory. Use go build instead to compile
// // production executables.

// // I rarely use go run other than to quickly do some testing or debugging.

// // Further reading
// // Execute go help run in your shell and read the instructions

//// 14:17
//// Go Build

// // go build compiles go code into an executable program

// // Build an executable

// // Ensure you are in your hellogo repo, then run:
// go build

// // Run the new program:
// ./hellogo

//// 14:18
//// Go Install

// // Build an executable

// // Ensure you are in your hellogo repo, then run:
// go install

// // Navigate out of your project directory:
// cd ../

// // Go has instaleld the hellogo program globally. Run it with:
// hellogo

// Tip about "not found"
// If you get an error regarding "hellogo not found" it means you probably don't have
// your Go environment setup properly. Specifically, go install is adding your binary to
// your GOBIN directory, but that may not be in your PATH.

// You can read more about that here in the go install docs.
// https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies

//// 14:24
//// Custom Package

// // Let's write a package to import and use in hellogo.

// // Create a sibling directory at the same level as the hellogo directory:
// mkdir mystrings
// cd mystrings

// // Initialize a module:
// go mod init {REMOTE}/{USERNAME}/mystring

// // Then create a new file mystring.go in that directory and paste the following code:

// // by convention, we nave our package the same as the directory
// package mystrings

// // Reverse reverses a string left to right
// // Notice that we need to capitalize the first letter of the function
// // If we don't then we won't be able access this function outside of the
// // mystrings package
// func Reverse(s string) string {
// 	result := ""
// 	for _, v := range s {
// 		result = string(v) + result
// 	}
// 	return result
// }

// // Note that there is no main.go or func main() in this package.

// // go build won't build an executable from a library package. However, go build will
// // still compile the package and save it to our local build cache. It's userful for checking for
// // compiler errors.

// // Run:
// go build

//// 14:27
//// Custom Package Continued

// // Let's use our new mystrings package in hellogo

// // Modify hellogo's main.go file:
// package main

// import (
// 	"fmt"

// 	"{REMOTE}/{USERNAME}/mystrings"
// )

// func main() {
// 	fmt.Println(mystring.Reverse("hello world"))
// }

// // Don't forget to replace {REMOTE} and {USERNAME} with the values you used before.
// // Then edit hellogo's go.mod file to contian the following:

// module example.com/username/hellogo

// go 1.20

// replace example.com/username/mystrings v0.0.0 => ../mystrings

// require (
// 	example.com/username/mystrings v0.0.0
// )

// // Now build and run the new program:
// go build
// ./hellogo

//// 14:31
//// Remote Packages

// Let's learn how to use an open-source package that's available online.

// A note on how you should publish modules

// Be aware that using the "replace=" keyword like we did in the last assignment isn't
// advised, but can be useful to get up and running quickly. The proper way to create and
// depend on modules is to publish them to a remote repository. When you do that, the
// "replace keyword can be dropped from the go.mod:

// // Bad
// // This works for local-only development
// module github.com/wagslane/hellogo

// fo 1.20

// replace github.com/wagslane/mystrings v0.0.0 => ../mystrings

// require (
// 	github.com/wagslane/mystrings v0.0.0
// )

// // Good
// // This works if we publish our modules to a remote location like Github as we should.
// module github.com/wagslane/hellogo

// go 1.20

// require (
// 	github.com/wagslane/mystrins v0.0.0
// )

// //// Assignment

// // First, create a new directory in the same parent directory as hellogo and mystrings
// // called datetest.

// // Create main.go in datetest and add the following code:
// package main

// import (
// 	"fmt"
// 	"time"

// 	tinytime "github.com/wagslane/go-tinytime"
// )

// func main() {
// 	tt := tinytime.New(1585750374)

// 	tt = tt.Add(time.Hour * 48)
// 	fmt.Println(tt)
// }

// // Initialize a module:
// go mod init {REMOTE}/{USERNAME}/datetest

// // Download and install the remote go-tinydate package using go get:
// go get github.com/wagslane/go-tinytime

// // Print the contents of your go.mod file to see the changes:
// cat go.mod

// // Compile and run your program:
// go build
// ./datetest

//// 14:40
//// Clean Packages

// I've often seen, and have been responsible for, throwing code into packages without
// much thought. I've quickly drawn a line in the sand and started putting code into
// different folders (which in Go are different packages by definition) just for the sake of
// findability. Learning to properly build small and reusable packages can take your Go
// career to the next level.

// Rules Of Thumb

// // 1. Hide internal logic

// // If you're familiar with the pillars of OOP, this is a practice in encapsulation.

// // Oftentimes an application will have complex logic that requires a lot of code. In almost
// // every case the logic that the application cares about can be exposed via an API, and
// // most of the dirty work can be kept within a package. For example, imagine we are
// // building an application that needs to classify images. We could build a package:

// package classifier

// // ClassifyImage classifies images as "hotdog" or "not hotdog"
// func ClassifyImage(image []byte) (imageType string) {
// 	return hasHotdogColors(image) && hasHotdogShape(image)
// }

// func hasHotdogShape(image []byte) bool {
// 	// internal logic that the application doesn't need to know about
// 	return true
// }

// func hasHotdogColors(image []byte) bool {
// 	// internal logic that the application doesn't need to know about
// 	return true
// }

// // We create an API by only exposing the function(s) that the application-level needs to
// // know about. All other logic is unexported to keep a clean separation of concerns. The
// // application doesn't need to know how to classify an image, just the result of the
// // classification.

// // 2. Don't change APIs

// // The unexported functions within a package can and should change often for testing,
// // refactoring, and bug fixing.

// // A well-designed library will have a stable API so that users aren't receiving breaking
// // changes each time they update the package version. In Go, this means not changing
// // exported function's signatures.

// // 3. Don't export functions from the main package

// // A main package isn't a library, there's no need to export functions from it.

// // 4. Packages shouldn't know about dependents

// // Perhaps one of the most important and most broken rules is that a package shouldn't
// // know anything about its dependents. In other words, a package should never have
// // specific knowledge about a particular application that uses it.

// // Further Reading
// // You can optionally read more here if you're interested.
// https://blog.boot.dev/golang/how-to-separate-library-packages-in-go/


