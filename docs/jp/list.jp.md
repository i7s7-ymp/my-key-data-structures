# 連結リスト（Linked List）とは

連結リストは、データ要素（ノード）がポインタやリファレンスによって線形に接続されたデータ構造です。各ノードは、データ部分と次のノードへのポインタで構成されています。

## 基本的な特徴

- **動的サイズ**: 実行時にサイズを自由に変更可能
- **非連続メモリ**: 要素が連続したメモリ領域に格納されない
- **ポインタベース**: 各ノードが次のノードへのポインタを持つ
- **メモリ効率**: 必要な分だけメモリを確保（オーバーヘッドあり）
- **順次アクセス**: 先頭から順番にアクセスする必要がある

## 連結リストの種類

### 単方向連結リスト

- 各ノードが次のノードへのポインタのみを持つ
- 一方向にのみ移動可能

### 双方向連結リスト

- 各ノードが前後のノードへのポインタを持つ
- 双方向に移動可能

### 循環連結リスト

- 最後のノードが最初のノードを指す
- リング状の構造

# 行える処理

## 基本操作

### 挿入操作

- **先頭への挿入**: 新しいノードを先頭に追加
- **末尾への挿入**: 新しいノードを末尾に追加
- **中間への挿入**: 指定位置に新しいノードを追加

### 削除操作

- **先頭からの削除**: 先頭ノードを削除
- **末尾からの削除**: 末尾ノードを削除
- **中間からの削除**: 指定ノードを削除

### アクセス操作

- **検索**: 特定の値を持つノードを探索
- **取得**: 指定位置のノードを取得
- **走査**: 全ノードを順次処理

# 得意な処理（効率的な操作）

## 先頭への挿入・削除

- **計算量**: **O(1)**
- 先頭ノードのポインタが既知なら即座に実行可能
- 配列と比べて大幅に高速

## 動的なサイズ変更

- **計算量**: **O(1)** (挿入・削除時)
- 事前にサイズを決める必要がない
- メモリの無駄がない

## 任意位置への挿入・削除（ノードの参照がある場合）

- **計算量**: **O(1)**
- 対象ノードのポインタが既知なら高速
- ポインタの付け替えのみで完了

# 苦手な処理（非効率的な操作）

## ランダムアクセス

- **計算量**: **O(n)**
- インデックスでの直接アクセス不可
- 先頭から順次辿る必要がある

## 中間位置での操作（位置指定）

- **計算量**: **O(n)**
- 対象位置まで順次移動が必要
- 配列のような直接アクセス不可

## 逆方向の走査（単方向の場合）

- **計算量**: **O(n)**
- 単方向連結リストでは後ろから前への移動不可
- 全体を走査する必要がある

## 要素の検索

- **計算量**: **O(n)**
- 線形探索のみ可能
- ソートされていても二分探索は困難

# Go 言語での実装

## 基本的なノード構造

```go
// 単方向連結リストのノード
type ListNode struct {
    Data int       // データ部分
    Next *ListNode // 次のノードへのポインタ
}

// 双方向連結リストのノード
type DoublyListNode struct {
    Data int              // データ部分
    Next *DoublyListNode  // 次のノードへのポインタ
    Prev *DoublyListNode  // 前のノードへのポインタ
}
```

## 連結リスト構造体

```go
type LinkedList struct {
    Head *ListNode // 先頭ノードへのポインタ
    Size int       // リストのサイズ
}
```

## 基本操作の実装例

### 先頭への挿入 O(1)

```go
func (list *LinkedList) InsertAtHead(data int) {
    newNode := &ListNode{Data: data, Next: list.Head}
    list.Head = newNode
    list.Size++
}
```

### 末尾への挿入 O(n)

```go
func (list *LinkedList) InsertAtTail(data int) {
    newNode := &ListNode{Data: data, Next: nil}

    if list.Head == nil {
        list.Head = newNode
    } else {
        current := list.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
    list.Size++
}
```

### 先頭からの削除 O(1)

```go
func (list *LinkedList) DeleteFromHead() bool {
    if list.Head == nil {
        return false
    }

    list.Head = list.Head.Next
    list.Size--
    return true
}
```

### 検索 O(n)

```go
func (list *LinkedList) Search(target int) *ListNode {
    current := list.Head
    for current != nil {
        if current.Data == target {
            return current
        }
        current = current.Next
    }
    return nil
}
```

### 全要素の表示 O(n)

```go
func (list *LinkedList) Display() {
    current := list.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Data)
        current = current.Next
    }
    fmt.Println("nil")
}
```

## 使用例

```go
func main() {
    list := &LinkedList{}

    // 要素の挿入
    list.InsertAtHead(1)
    list.InsertAtHead(2)
    list.InsertAtTail(3)

    // リストの表示: 2 -> 1 -> 3 -> nil
    list.Display()

    // 要素の検索
    node := list.Search(1)
    if node != nil {
        fmt.Printf("Found: %d\n", node.Data)
    }

    // 先頭から削除
    list.DeleteFromHead()
    list.Display() // 1 -> 3 -> nil
}
```

## 配列との比較

| 操作                 | 連結リスト         | 配列     |
| -------------------- | ------------------ | -------- |
| **ランダムアクセス** | O(n)               | O(1)     |
| **先頭への挿入**     | O(1)               | O(n)     |
| **末尾への挿入**     | O(n)               | O(1)平均 |
| **中間への挿入**     | O(n)               | O(n)     |
| **先頭からの削除**   | O(1)               | O(n)     |
| **末尾からの削除**   | O(n)               | O(1)     |
| **検索**             | O(n)               | O(n)     |
| **メモリ使用量**     | 多い（ポインタ分） | 少ない   |

連結リストは、頻繁な挿入・削除が発生し、ランダムアクセスが少ない場合に効果的なデータ構造です。
