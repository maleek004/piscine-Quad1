package checkpoint

// we first create a helper function that checks if a number is prime, if its 1 or can be divided by 2 or 3 - false
// if its 2 or 3 then true 
// then we loop from 5 up to the square root of the number , to see if we can find any number that will divide it  

func IsPrime(n int) bool{
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3{
		return true
	}
	if n%2 ==0 || n%3 == 0{
		return false
	}

	for i:= 5; i*i <=n; i+=6{
		if n%i==0 || n%(i+2)==0{
			return false
		} 
	}
	return true
}

func FindPrevPrime(nb int) int {
	if nb <=2{
		return 0
	}
	for i := nb; i>=2; i--{
		if IsPrime(i){
		return i}
	}
	return 0
}