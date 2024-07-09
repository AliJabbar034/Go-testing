package main

import "testing"

func Test_isPrime (t *testing.T){

	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Negative number", -1, false},
		{"Zero", 0, false},
		{"One", 1, false},
		{"Two", 2, true},
		{"Three", 3, true},
		{"Four", 4, false},
		{"Five", 5, true},
		{"Six", 6, false},
		{"Prime number", 29, true},
		{"Non-prime number", 30, false},
	}

	for _,tt := range tests {

		
		t.Run(tt.name, func(t *testing.T) {
			result := IsPrime(tt.input)
			if result != tt.expected {
				t.Errorf("isPrime(%d) = %t; want %t", tt.input, result, tt.expected)
			}
		})
	
	}
 

}