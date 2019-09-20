package boeiegast 
        
import (
    "testing"
)

func TestUpperFunc(t *testing.Testing){

}

func ExampleUpperFunc(){

}

func BenchmarkUpperFunc(b *testing.B){
	for i := 0; i < b.N; i++ {
		UpperFunc()	//Enter the values that your function needs between the parentheses
	}
}