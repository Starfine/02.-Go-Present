package math

import "testing"

func TestAdd(t *testing.T) {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if s := Add(i, j); s != i+j {
				t.Error("加法函數測試失敗！")
			} else {
				t.Log("加法函數測試成功！")
			}
		}

	}
}
