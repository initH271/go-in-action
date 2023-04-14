/*
Popcount(x uint64) 返回数字x中含二进制1bit的个数
*/
package popcount

/*
pc表存储每8bit数字存储的二进制bit个数

	0 0
	1 1
*/
var PC [256]byte

func init() {
	for i := range PC {
		PC[i] = PC[i/2] + byte(i&1)
	}
}

func Popcount(x uint64) int {

	return int(PC[byte(x>>(0*8))] +
		PC[byte(x>>(1*8))] +
		PC[byte(x>>(2*8))] +
		PC[byte(x>>(3*8))] +
		PC[byte(x>>(4*8))] +
		PC[byte(x>>(5*8))] +
		PC[byte(x>>(6*8))] +
		PC[byte(x>>(7*8))])

}
