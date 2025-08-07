# ハッシュテーブル（Hash Table）とは

ハッシュテーブルは、キーと値のペアを効率的に格納・検索するデータ構造です。ハッシュ関数を使用してキーを配列のインデックスにマッピングし、高速なデータアクセスを実現します。

## 基本的な特徴

- **キー・バリュー型**: キーと値のペアでデータを管理
- **ハッシュ関数**: キーを配列のインデックスに変換する関数
- **高速アクセス**: 平均的に O(1)でのデータアクセス
- **動的サイズ**: 必要に応じてサイズを拡張可能
- **衝突処理**: 異なるキーが同じインデックスにマッピングされる問題への対処

## ハッシュテーブルの構成要素

### ハッシュ関数

- キーを配列のインデックスに変換する関数
- 良いハッシュ関数の条件：
  - 高速に計算できる
  - 値が均等に分散される
  - 決定的（同じキーは常に同じ値を返す）

### バケット配列

- 実際にデータを格納する配列
- 各要素（バケット）にキー・バリューペアを格納

### 衝突処理機構

- **チェイン法**: 同じインデックスの要素を連結リストで管理
- **オープンアドレス法**: 別の空いているインデックスを探す

# 行える処理

## 基本操作

### 挿入操作（Insert/Put）

- **新しいキー・バリューペアを追加**
- キーが既に存在する場合は値を更新

### 検索操作（Search/Get）

- **指定したキーに対応する値を取得**
- キーが存在しない場合は「見つからない」を返す

### 削除操作（Delete/Remove）

- **指定したキーとその値を削除**
- キーが存在しない場合は何もしない

### 存在確認操作（Contains/Has）

- **指定したキーが存在するかチェック**
- 真偽値を返す

## 補助操作

### サイズ取得

- **格納されている要素数を取得**

### 空判定

- **ハッシュテーブルが空かどうかを確認**

### 全要素の走査

- **すべてのキー・バリューペアを順次処理**

# 得意な処理（効率的な操作）

## データの検索

- **計算量**: **O(1)** (平均)
- ハッシュ関数により直接インデックスを計算
- 配列や連結リストと比べて圧倒的に高速

## データの挿入

- **計算量**: **O(1)** (平均)
- 既存データの移動が不要
- 動的にサイズを拡張可能

## データの削除

- **計算量**: **O(1)** (平均)
- 他の要素への影響が最小限
- 削除後の再配置が不要

## キーの存在確認

- **計算量**: **O(1)** (平均)
- 値を取得せずに存在だけを高速確認

# 苦手な処理（非効率的な操作）

## ソートされた順序での取得

- **計算量**: **O(n log n)**
- ハッシュテーブルは順序を保持しない
- ソートが必要な場合は別途処理が必要

## 最小値・最大値の検索

- **計算量**: **O(n)**
- 全要素を走査する必要がある
- 専用のデータ構造（ヒープなど）の方が効率的

## 範囲検索

- **計算量**: **O(n)**
- 特定の範囲内の値を効率的に取得できない
- B 木などの方が適している

## メモリ使用量

- **空間効率**: 負荷率に依存
- 衝突を避けるため、使用率を低く保つ必要がある
- メモリの無駄が発生しやすい

## 最悪ケースの性能

- **計算量**: **O(n)** (最悪の場合)
- すべてのキーが同じインデックスにマッピングされる場合
- ハッシュ関数の品質に依存

# Go 言語での実装

## 基本的な構造体

```go
// キー・バリューペア
type KeyValue struct {
    Key   string
    Value interface{}
    Next  *KeyValue // チェイン法用のポインタ
}

// ハッシュテーブル構造体
type HashTable struct {
    Buckets []*KeyValue // バケット配列
    Size    int         // 現在の要素数
    Capacity int        // バケットの容量
}
```

## ハッシュ関数の実装

```go
// シンプルなハッシュ関数（djb2アルゴリズム）
func (ht *HashTable) hash(key string) int {
    hash := 5381
    for _, char := range key {
        hash = ((hash << 5) + hash) + int(char)
    }
    return hash % ht.Capacity
}
```

## 基本操作の実装

### 挿入操作 O(1)平均

```go
func (ht *HashTable) Put(key string, value interface{}) {
    index := ht.hash(key)

    // 既存のキーをチェック
    current := ht.Buckets[index]
    for current != nil {
        if current.Key == key {
            current.Value = value // 値を更新
            return
        }
        current = current.Next
    }

    // 新しいキー・バリューペアを先頭に挿入
    newNode := &KeyValue{
        Key:   key,
        Value: value,
        Next:  ht.Buckets[index],
    }
    ht.Buckets[index] = newNode
    ht.Size++
}
```

### 検索操作 O(1)平均

```go
func (ht *HashTable) Get(key string) (interface{}, bool) {
    index := ht.hash(key)
    current := ht.Buckets[index]

    for current != nil {
        if current.Key == key {
            return current.Value, true
        }
        current = current.Next
    }

    return nil, false // キーが見つからない
}
```

### 削除操作 O(1)平均

```go
func (ht *HashTable) Delete(key string) bool {
    index := ht.hash(key)
    current := ht.Buckets[index]
    var prev *KeyValue

    for current != nil {
        if current.Key == key {
            if prev == nil {
                // 先頭要素の削除
                ht.Buckets[index] = current.Next
            } else {
                // 中間要素の削除
                prev.Next = current.Next
            }
            ht.Size--
            return true
        }
        prev = current
        current = current.Next
    }

    return false // キーが見つからない
}
```

### 存在確認 O(1)平均

```go
func (ht *HashTable) Contains(key string) bool {
    _, exists := ht.Get(key)
    return exists
}
```

## 補助メソッド

### ハッシュテーブルの作成

```go
func NewHashTable(capacity int) *HashTable {
    return &HashTable{
        Buckets:  make([]*KeyValue, capacity),
        Size:     0,
        Capacity: capacity,
    }
}
```

### 全要素の表示

```go
func (ht *HashTable) Display() {
    fmt.Printf("ハッシュテーブル (サイズ: %d, 容量: %d)\n", ht.Size, ht.Capacity)
    for i, bucket := range ht.Buckets {
        if bucket != nil {
            fmt.Printf("バケット[%d]: ", i)
            current := bucket
            for current != nil {
                fmt.Printf("(%s: %v)", current.Key, current.Value)
                if current.Next != nil {
                    fmt.Print(" -> ")
                }
                current = current.Next
            }
            fmt.Println()
        }
    }
}
```

### 負荷率の計算

```go
func (ht *HashTable) LoadFactor() float64 {
    return float64(ht.Size) / float64(ht.Capacity)
}
```

## 使用例

```go
func main() {
    // ハッシュテーブルを作成
    ht := NewHashTable(10)

    // データの挿入
    ht.Put("apple", 100)
    ht.Put("banana", 200)
    ht.Put("orange", 150)
    ht.Put("grape", 300)

    // データの検索
    if value, found := ht.Get("apple"); found {
        fmt.Printf("apple: %v\n", value)
    }

    // データの存在確認
    if ht.Contains("banana") {
        fmt.Println("banana が存在します")
    }

    // データの削除
    if ht.Delete("orange") {
        fmt.Println("orange を削除しました")
    }

    // ハッシュテーブルの表示
    ht.Display()

    // 負荷率の確認
    fmt.Printf("負荷率: %.2f\n", ht.LoadFactor())
}
```

## 他のデータ構造との比較

| 操作           | ハッシュテーブル | 配列     | 連結リスト | 二分探索木 |
| -------------- | ---------------- | -------- | ---------- | ---------- |
| **検索**       | O(1)平均         | O(n)     | O(n)       | O(log n)   |
| **挿入**       | O(1)平均         | O(1)末尾 | O(1)先頭   | O(log n)   |
| **削除**       | O(1)平均         | O(n)     | O(1)先頭   | O(log n)   |
| **順序保持**   | なし             | あり     | あり       | あり       |
| **メモリ効率** | 中程度           | 高い     | 中程度     | 中程度     |

ハッシュテーブルは、キーによる高速なデータアクセスが必要で、順序が重要でない場合に最適なデータ構造です。
