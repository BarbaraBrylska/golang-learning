/*************************************************************************
	> File Name: bubblesort.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Thu 25 Oct 2012 06:54:05 PM CST
 ************************************************************************/

package bubblesort

func BubbleSort(value []int){
	for i:=0; i<len(value)-1; i++{
		for j:=i+1; j <len(value); j++{
			if value[i] > value[j]{
				value[i],value[j] = value[j],value[i]
			}
		}
	}
}
