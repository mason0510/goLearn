package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// 构建训练数据
	x := mat.NewDense(5, 2, []float64{1, 2, 1, 3, 1, 4, 1, 5, 1, 6})
	y := mat.NewDense(5, 1, []float64{3, 4, 5, 6, 7})

	// 求解最小二乘法
	var xTx mat.Dense
	xTx.Mul(x.T(), x)
	xTxInv := mat.NewDense(2, 2, nil)
	xTxInv.Inverse(&xTx)
	var xTy mat.Dense
	xTy.Mul(x.T(), y)
	var b mat.Dense
	b.Mul(xTxInv, &xTy)

	// 输出结果
	fmt.Println("线性回归系数：")
	fmt.Printf("%.3f\n", mat.Formatted(&b, mat.Prefix("    ")))
}
