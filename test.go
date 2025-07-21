package main

import "fmt"

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	//Adds entry to the map so long as no entry already exists for the given key
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
	// panic("impossible bro")
}

// check adds an error message to the map only if a validation check is not 'ok'
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func main() {

	// error := map[string]string{"error": "test1"}
	// error2 := map[string]string{"error2": "test2"}

	validation:=New()

	// struc1 := Validator{
	// 	Errors: error,
	// }

	// fmt.Println(struc1.Valid())
	// fmt.Println(New())

	validation.Check(1 ==9,"error","test1")
	validation.Check(1==0,"error3","test1")
	fmt.Println(validation.Errors)

}
