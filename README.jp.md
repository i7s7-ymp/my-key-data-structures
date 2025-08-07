[English](./README.md) | [日本語](./README.jp.md)

# My Data Structures

このリポジトリは、ソフトウェア開発でよく使用される 10 の重要なデータ構造を実装した Gist のコレクションです。もしお役に立てましたら、GitHub スターをいただけると嬉しいです！

## データ構造

- [リスト（List）](./docs/jp/list.jp.md)
- [スタック（Stack）](./docs/jp/stack.jp.md)
- [キュー（Queue）](./docs/jp/queue.jp.md)
- [ハッシュテーブル（Hash Table）](./docs/jp/hash-table.jp.md)
- [配列（Array）](./docs/jp/array.jp.md)
- [ヒープ（Heap）](./docs/jp/heap.jp.md)
- [木構造（Tree）](./docs/jp/tree.jp.md)
- [接尾辞木（Suffix Tree）](./docs/jp/suffix-tree.jp.md)
- [グラフ（Graph）](./docs/jp/graph.jp.md)
- [R 木（R-Tree）](./docs/jp/r-tree.jp.md)
- [頂点バッファ（Vertex Buffer）](./docs/jp/vertex-buffer.jp.md)

# データ構造の挙動確認

Docker を使用してローカル環境で挙動を確認できます。

## 必要なもの

- [Docker](https://www.docker.com/) がインストールされていること
- ターミナル環境

## 実行手順

1. **リポジトリをクローンまたは作業ディレクトリに移動**

   ```bash
   cd {リポジトリへのpath}/my-data-structures-go
   ```

2. **Docker イメージをビルド**
   以下のコマンドを実行して Docker イメージをビルドします。

   ```bash
   docker build -t go-code-runner -f docker/Dockerfile .
   ```

3. **コードを実行**
   ビルドしたイメージを使用してコードを実行します。
   ```bash
   docker run --rm go-code-runner <ファイル名>.go
   ```

# 参考資料

- [ByteByteGo - 10 Key Data Structures We Use Every Day](https://bytebytego.com/guides/10-key-data-structures-we-use-every-day/)
