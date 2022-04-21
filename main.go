package mimi

import (
	"fmt"
	"math"
)

var gosa float64

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

//活性化関数
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

//損失関数
//出力と正解の誤差を定義する関数。誤差とは、あるべき状態との隔離の度合い
//誤差が大きくなれば、ニューラルネットワークが望ましい状態から離れていること

//二乗和誤差
//全ての出力層のニューロンで総和をとり、1/2を欠ける。二乗わ誤差を用いる
//ことで、ニューラルネットワークの出力がどの程度一致しているかを定量化できる。
//正解や出力が連続的なケースに向いている。回帰問題でよく利用される。
//yを出力値、ｔを正解値
func Square_sum(y, t []float64) float64 {
	var res float64
	for i := 0; i < len(y); i++ {
		res += (math.Pow(y[i]-t[i], 2))
	}
	return res / 2
}

//較差エントロピー誤差
/* 二つの分布の間のずれを表す尺度。分類問題でよく利用される。
正解に近づくほど小さくなり、正解から離れるにつれてどこまでも大きくなる
したがって、出力が正解から離れるほど、誤差が大きくなり、正解に近づく程
誤差が小さくなる。
　出力値と正解値の隔離が大きい時に学習速度が早くなるのが利点。
*/

func Cross_entropy(y, t []float64) float64 {
	var res float64
	gosa := 0.000000007
	for i := 0; i < len(y); i++ {
		res += (t[i] * math.Log(y[i]+gosa))
	}
	return res
}
