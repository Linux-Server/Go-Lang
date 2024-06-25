package main

import (
	"fmt"
)

//  ASCCI  ==>  Ameerican starnda  = only 127  ||  a = 97
// Universal ==>  Unicode  ===>  "ഒ"  ==>    3346

//

//  Rune   ===>  no indivdual char in a any lang
// Lang   =>  1 char   ===>  1 ascii value /utf / unicode value
// malayalam => "ഒ"  ==>    3346

// Length   ==>  no of hex char
// Lang   ==>  1 char ==> an array of hex value
// malayalam => "ഒ"  ==>   [e0,b4,92]

func string_runes() {
	s := "sam" //Rune count = 6  , len = 18
	// name_one := "sam" // Rune and len = 3    =====> hex [73,61,6d]  ===> ascii [115,97,109]  ==> len and rune = 3
	// _ = name_one

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("the len is :", len(s))

	// find_hex(name)
	// find_ascii(name)
	// fmt.Println()
	// fmt.Println("Length:", len(name))

	// fmt.Println("Rune count:", utf8.RuneCountInString(name))

}

func find_hex(s string) {
	fmt.Println()
	fmt.Printf("Inside HEX : %x\n", s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

}

func find_ascii(s string) {
	fmt.Println()
	fmt.Printf("Inside ASCII : %v\n", s)
	for _, i := range s {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
	fmt.Println("ASCII LOOP OVER")

}
