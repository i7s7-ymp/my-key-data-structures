package main

import (
    "fmt"
    "time"
    "unsafe"
)

// 静的配列の実装
type StaticArray struct {
    data []int
    size int
}

// 静的配列のコンストラクタ
func NewStaticArray(size int) *StaticArray {
    return &StaticArray{
        data: make([]int, size),
        size: size,
    }
}

// 要素の取得 (O(1))
func (sa *StaticArray) Get(index int) (int, error) {
    if index < 0 || index >= sa.size {
        return 0, fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, sa.size-1)
    }
    return sa.data[index], nil
}

// 要素の設定 (O(1))
func (sa *StaticArray) Set(index int, value int) error {
    if index < 0 || index >= sa.size {
        return fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, sa.size-1)
    }
    sa.data[index] = value
    return nil
}

// サイズの取得
func (sa *StaticArray) Size() int {
    return sa.size
}

// 配列の表示
func (sa *StaticArray) String() string {
    return fmt.Sprintf("StaticArray%v", sa.data)
}

// 動的配列の実装
type DynamicArray struct {
    data     []int
    length   int
    capacity int
}

// 動的配列のコンストラクタ
func NewDynamicArray() *DynamicArray {
    initialCapacity := 2
    return &DynamicArray{
        data:     make([]int, initialCapacity),
        length:   0,
        capacity: initialCapacity,
    }
}

// 指定された容量で動的配列を作成
func NewDynamicArrayWithCapacity(capacity int) *DynamicArray {
    if capacity < 1 {
        capacity = 2
    }
    return &DynamicArray{
        data:     make([]int, capacity),
        length:   0,
        capacity: capacity,
    }
}

// 容量を拡張する内部メソッド
func (da *DynamicArray) resize() {
    oldCapacity := da.capacity
    da.capacity *= 2
    newData := make([]int, da.capacity)
    
    // 手動でデータをコピー（copy関数を使わない）
    for i := 0; i < da.length; i++ {
        newData[i] = da.data[i]
    }
    
    da.data = newData
    fmt.Printf("容量を %d から %d に拡張しました\n", oldCapacity, da.capacity)
}

// 要素の取得 (O(1))
func (da *DynamicArray) Get(index int) (int, error) {
    if index < 0 || index >= da.length {
        return 0, fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, da.length-1)
    }
    return da.data[index], nil
}

// 要素の設定 (O(1))
func (da *DynamicArray) Set(index int, value int) error {
    if index < 0 || index >= da.length {
        return fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, da.length-1)
    }
    da.data[index] = value
    return nil
}

// 末尾への要素追加 (平均 O(1))
func (da *DynamicArray) Append(value int) {
    if da.length >= da.capacity {
        da.resize()
    }
    da.data[da.length] = value
    da.length++
}

// 指定位置への要素挿入 (O(n))
func (da *DynamicArray) Insert(index int, value int) error {
    if index < 0 || index > da.length {
        return fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, da.length)
    }
    
    if da.length >= da.capacity {
        da.resize()
    }
    
    // 要素を右にシフト（手動実装）
    for i := da.length; i > index; i-- {
        da.data[i] = da.data[i-1]
    }
    
    da.data[index] = value
    da.length++
    return nil
}

// 指定位置の要素削除 (O(n))
func (da *DynamicArray) Delete(index int) error {
    if index < 0 || index >= da.length {
        return fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, da.length-1)
    }
    
    // 要素を左にシフト（手動実装）
    for i := index; i < da.length-1; i++ {
        da.data[i] = da.data[i+1]
    }
    
    da.length--
    return nil
}

// 長さの取得
func (da *DynamicArray) Length() int {
    return da.length
}

// 容量の取得
func (da *DynamicArray) Capacity() int {
    return da.capacity
}

// 配列の表示
func (da *DynamicArray) String() string {
    result := "DynamicArray["
    for i := 0; i < da.length; i++ {
        if i > 0 {
            result += ", "
        }
        result += fmt.Sprintf("%d", da.data[i])
    }
    result += "]"
    return result
}

// 配列を空にする
func (da *DynamicArray) Clear() {
    da.length = 0
}

// メモリ使用量の情報を表示
func (da *DynamicArray) MemoryInfo() {
    elementSize := int(unsafe.Sizeof(int(0)))
    usedMemory := da.length * elementSize
    allocatedMemory := da.capacity * elementSize
    fmt.Printf("メモリ情報: 使用中=%dバイト, 確保済み=%dバイト, 使用率=%.1f%%\n", 
        usedMemory, allocatedMemory, float64(usedMemory)/float64(allocatedMemory)*100)
}

func main() {
    // 組み込みの配列とスライスのデモンストレーション
    demonstrateBuiltinArrays()
    
    // カスタム実装のデモンストレーション
    demonstrateCustomArrays()
}

func demonstrateBuiltinArrays() {
    fmt.Println("========================================")
    fmt.Println("=== 組み込みの配列とスライスのデモ ===")
    fmt.Println("========================================")
    
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

func demonstrateCustomArrays() {
    fmt.Println("\n========================================")
    fmt.Println("=== カスタム実装の配列のデモ ===")
    fmt.Println("========================================")
    
    // 静的配列のテスト
    fmt.Println("=== 静的配列のテスト ===")
    staticArray := NewStaticArray(5)
    fmt.Println("初期状態:", staticArray)
    
    // 要素の設定
    for i := 0; i < staticArray.Size(); i++ {
        staticArray.Set(i, (i+1)*10)
    }
    fmt.Println("値設定後:", staticArray)
    
    // 要素の取得
    if val, err := staticArray.Get(2); err == nil {
        fmt.Printf("インデックス 2 の要素: %d\n", val)
    }
    
    // 範囲外アクセスのテスト
    if _, err := staticArray.Get(10); err != nil {
        fmt.Println("エラー:", err)
    }
    
    // 動的配列のテスト
    fmt.Println("\n=== 動的配列のテスト ===")
    dynamicArray := NewDynamicArray()
    fmt.Printf("初期状態: %s (長さ: %d, 容量: %d)\n", 
        dynamicArray, dynamicArray.Length(), dynamicArray.Capacity())
    
    // 要素の追加
    fmt.Println("\n--- 要素の追加テスト ---")
    for i := 1; i <= 8; i++ {
        start := time.Now()
        dynamicArray.Append(i * 10)
        fmt.Printf("追加後: %s (長さ: %d, 容量: %d) 処理時間: %v\n", 
            dynamicArray, dynamicArray.Length(), dynamicArray.Capacity(), time.Since(start))
    }
    
    // メモリ使用量の表示
    fmt.Println("\n--- メモリ使用量 ---")
    dynamicArray.MemoryInfo()
    
    // 中間への挿入
    fmt.Println("\n--- 中間への挿入テスト ---")
    start := time.Now()
    err := dynamicArray.Insert(3, 999)
    if err == nil {
        fmt.Printf("インデックス 3 に 999 を挿入: %s 処理時間: %v\n", 
            dynamicArray, time.Since(start))
    }
    
    // 要素の取得
    if val, err := dynamicArray.Get(3); err == nil {
        fmt.Printf("インデックス 3 の要素: %d\n", val)
    }
    
    // 要素の変更
    dynamicArray.Set(0, 1111)
    fmt.Printf("インデックス 0 を 1111 に変更: %s\n", dynamicArray)
    
    // 中間からの削除
    fmt.Println("\n--- 中間からの削除テスト ---")
    start = time.Now()
    err = dynamicArray.Delete(3)
    if err == nil {
        fmt.Printf("インデックス 3 を削除: %s 処理時間: %v\n", 
            dynamicArray, time.Since(start))
    }
    
    // 性能テスト
    fmt.Println("\n--- 性能テスト ---")
    performanceTest()
}

func performanceTest() {
    fmt.Println("大量データでの性能テスト...")
    
    // カスタム動的配列の性能テスト
    da := NewDynamicArray()
    
    // 10000個の要素を追加
    start := time.Now()
    for i := 0; i < 10000; i++ {
        da.Append(i)
    }
    appendTime := time.Since(start)
    
    // 先頭に挿入（最悪ケース）
    start = time.Now()
    da.Insert(0, -1)
    insertTime := time.Since(start)
    
    // 中間から削除
    start = time.Now()
    da.Delete(5000)
    deleteTime := time.Since(start)
    
    // ランダムアクセス
    start = time.Now()
    for i := 0; i < 1000; i++ {
        da.Get(i * 10)
    }
    accessTime := time.Since(start)
    
    fmt.Printf("性能結果:\n")
    fmt.Printf("- 10000要素の追加: %v\n", appendTime)
    fmt.Printf("- 先頭への挿入: %v\n", insertTime)
    fmt.Printf("- 中間からの削除: %v\n", deleteTime)
    fmt.Printf("- 1000回のランダムアクセス: %v\n", accessTime)
    fmt.Printf("最終状態: 長さ=%d, 容量=%d\n", da.Length(), da.Capacity())
    da.MemoryInfo()
}