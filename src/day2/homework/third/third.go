package third

func Factorial(n uint64) uint64{
	var ret,i uint64 = 1,1
	for i=1;i<=n;i++ {
		ret = ret * i
	}
	return ret
}
