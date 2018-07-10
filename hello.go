package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func main() {
	fmt.Printf("Hello World2\n")

	vm := otto.New()
	vm.Run(`
	a={}
	a['b']={}
	a['b']['c']=800
	a['b']['d']=900
    abc = 2 + 200;
	console.log("The value of abc is " + abc); // 4
	console.log("inputs="+a.b.c)
	console.log("inputs="+a.b.d)
`)
	vm.Set("def", 11222)
	vm.Run(`
    console.log("The value of def is " + def);
    // The value of def is 11
`)
	testhash := make(map[string]int)
	testhash["a"] = 1
	testhash["b"] = 2

	vm.Set("testhash", testhash)
	vm.Run(`
	def=def+1;
    console.log("The value of def is " + def);
    console.log("The value of def is " + testhash.a);
    console.log("The value of def is " + def);
	// The value of def is 11
	
`)
	v1, _ := vm.Get("def")
	fmt.Printf("v1=[%d]\n", v1)
	fmt.Println(v1)
	vm.Run(`
    console.log("The value of def is " + def);
	
`)
	object, _ := vm.Object(`xyzzy = {}`)
	object2, _ := vm.Object(`input = {}`)
	object.Set("volume", 11)
	object.Set("strVal", "StringValue")
	object.Set("input", object2)
	object2.Set("aaaa", "bbb")
	object2.Set("bbbb", 5)
	vm.Run(`
    console.log("The value of def is " + xyzzy.volume);
    console.log("The value of def is " + xyzzy.strVal);
    console.log("The value of def is " + xyzzy.input);
    console.log("The value of def is " + xyzzy.input.aaaa);
    console.log("The value of def is " + xyzzy.input.bbbb);
`)
}
