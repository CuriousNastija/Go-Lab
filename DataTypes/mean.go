package sprint

/* Write a function that 
takes three float32 values as input 
and returns their mean as a float32 value. */

func Mean(a,b,c float32) float32 {
	res := (a + b + c) / 3
	return res
}