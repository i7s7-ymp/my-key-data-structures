package main

import (
    "fmt"
    "time"
)

func main() {
    // 配列の基本操作
    fmt.Println("=== 配列の操作 ===")
    array := [5]int{1, 2, 3, 4, 5} // 固定サイズの配列
    fmt.Println("配列:", array)

    // インデックスアクセス (O(1))
    fmt.Println("インデックス 2 の要素:", array[2])

    // 配列のサイズは固定のため、挿入や削除はできない
    // array[5] = 6 // コンパイルエラー: 配列のサイズを超える操作は不可

    // スライスの基本操作
    fmt.Println("\n=== スライスの操作 ===")
    slice := []int{1, 2, 3, 4, 5} // 動的サイズのスライス
    fmt.Println("スライス:", slice)

    // インデックスアクセス (O(1))
    fmt.Println("インデックス 2 の要素:", slice[2])

    // 末尾への要素追加 (平均 O(1))
    start := time.Now()
    slice = append(slice, 6)
    fmt.Println("要素追加後のスライス:", slice)
    fmt.Println("追加処理時間:", time.Since(start))

    // 中間への挿入 (O(n))
    start = time.Now()
    index := 2
    slice = append(slice[:index], append([]int{99}, slice[index:]...)...)
    fmt.Println("中間挿入後のスライス:", slice)
    fmt.Println("中間挿入処理時間:", time.Since(start))

    // 中間からの削除 (O(n))
    start = time.Now()
    slice = append(slice[:index], slice[index+1:]...)
    fmt.Println("中間削除後のスライス:", slice)
    fmt.Println("中間削除処理時間:", time.Since(start))

    // スライスの容量確認
    fmt.Println("\n=== スライスの容量確認 ===")
    fmt.Printf("スライスの長さ: %d, 容量: %d\n", len(slice), cap(slice))

    // スライスの容量が増える挙動
    fmt.Println("\n=== スライスの容量が増える挙動 ===")
    for i := 0; i < 10; i++ {
        slice = append(slice, i)
        fmt.Printf("追加後: 長さ: %d, 容量: %d\n", len(slice), cap(slice))
    }
}