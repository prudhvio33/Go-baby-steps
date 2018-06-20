package computers

/*
	If a struct type starts with a capital letter, then it is a exported type and it can be accessed from other packages.
	Similarly if the fields of a structure start with caps, they can be accessed from other packages.
*/

type Spec struct { //exported struct
	Maker string //exported field
	model string //unexported field
	Price int //exported field
}