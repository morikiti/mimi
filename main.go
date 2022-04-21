package mimi

import (
	"fmt"
	"math"
)

func Hello() {
	fmt.Println("モジュールから日っぱてきたぽよ")
}

func Add(a, b int) int {
	return a + b
}

func Sum(a *[]float64) float64 {
	res := 0.
	c := *a
	for i := 0; i < len(c); i++ {
		res += c[i]
	}
	return res
}

//ステップ関数
func Step_function(x float64) float64 {
	var ans float64
	if x <= 0 {
		ans = 0.
	} else if 0 < x {
		ans = 1.
	}
	return ans
}

//シグモイド関数
func Sigmoid_func(x float64) float64 {
	return 1 / (1 + math.Exp(x))
}

//tanh
func Tsnh_func(x float64) float64 {
	return (math.Exp(x) - math.Exp(-x)) / (math.Exp(x) + math.Exp(-x))
}

// ReLu
func ReLU(x float64) float64 {
	var res float64
	if x <= 0 {
		res = 0.
	} else {
		res = x
	}
	return res
}

//Leaky ReLU
func LeakyReLU(x float64) float64 {
	var res float64
	if x <= 0 {
		res = x * 0.01
	} else {
		res = x
	}
	return res
}

//ソフトマックス関数
/*
func SoftMax(x float64) float64 {
	return math.Exp(x) / Sum()
} */
