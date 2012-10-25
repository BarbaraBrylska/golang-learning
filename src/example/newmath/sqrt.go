/*************************************************************************
	> File Name: sqrt.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Thu 25 Oct 2012 07:48:43 PM CST
 ************************************************************************/

package newmath
func Sqrt(x float64) float64{
	z := 0.0
	for i := 0; i < 1000; i++{
		 z -= (z*z - x) / (2 * x)
	}
	return z
}
