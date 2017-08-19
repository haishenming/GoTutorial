package first

// 判断101到200之间有多少个素数，并输出所有素数
import "fmt"

func JudgePrime() {
	// 判断101到200之间有多少个素数，并输出所有素数
	num := 0
	for i := 101; i <= 200; i++ {
		is_prime := true
		for j:=2;j<i;j++{
			if i % j == 0 {
				num ++
				is_prime = false
				break
			}
		}
		if is_prime{
			fmt.Println(i)
		}
	}
	fmt.Println("素数的个数是：",num)
}
