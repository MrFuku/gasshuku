package main

import (
	"testing"
)

func Test_Hoge(t *testing.T) {
	str := Hoge()
	if str == "Hello Hoge!!!" {
		t.Log("はじめのテストがパスしました")
	} else {
		t.Error("戻り値が「Hello Hoge!!!」ではありません")
	}
}
