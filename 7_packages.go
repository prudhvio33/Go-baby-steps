package main

/*
	Every executable go application must contain a main function. This function is the entry point for execution.
	The main function should reside in the main package.

	Source files belonging to a package should be placed in separate folders of their own. It is a convention in Go to
	name this folder with the same name of the package.

	It is illegal in Go to import a package and not to use it anywhere in the code. The compiler will complain if you do
	so. The reason for this is to avoid bloating of unused packages which will significantly increase the compilation time.
 */

 import "./rectangle_package"

 func main() {
 	rectangle_package.Area(10, 20)
 }