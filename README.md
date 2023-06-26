# showGoroutinesInfo

一行書くだけで現在の goroutine の数とそれぞれの状態、作成場所を出力するライブラリです。
print debug のように気軽に使用することができます。

## Features

- runtime.Stack で出力される goroutine のスタックトレースをパースし、必要な情報だけ整形し出力しています。
- go tool trace のような高機能なものではなく、初学者が簡単に使用できるようなツールを目指しました。

## Note

このライブラリは2023年6月現在最新版の Go 1.21 の runtime.Stack の挙動に依存しています。よって実行する際は go1.21RC2 を使用してください。

[Go 1.21 Release Notes](https://tip.golang.org/doc/go1.21#runtime)

### go1.21RC2 インストール方法

```
go install golang.org/dl/go1.21rc2@latest
go1.21rc2 download
go1.21rc2 version
```

### 実行

```
go1.21rc2 run <your file>
```

## Usage

以下のように呼び出して使用します。

```go
package main

import "github.com/naka-c1024/showGoroutinesInfo"

func main() {
    showGoroutinesInfo.Do("region name")
}
```

## example

コード

```go
package main

import (
	"sync"

	"github.com/naka-c1024/showGoroutinesInfo"
)

func main() {
	showGoroutinesInfo.Do("init main")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		goroutineFirst()
	}()

	wg.Wait()

	showGoroutinesInfo.Do("last main")
}

func goroutineFirst() {
	showGoroutinesInfo.Do("goroutineFirst")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		goroutineSecond()
	}()

	wg.Wait()
}

func goroutineSecond() {
	showGoroutinesInfo.Do("goroutineSecond")
}

```

出力

```
=== goroutines info: init main ===

num goroutines -> 1

Goroutine ID: 1
State: running


=== goroutines info: goroutineFirst ===

num goroutines -> 2

Goroutine ID: 6
State: running
Created at: main.main in goroutine 1

Goroutine ID: 1
State: semacquire


=== goroutines info: goroutineSecond ===

num goroutines -> 3

Goroutine ID: 7
State: running
Created at: main.goroutineFirst in goroutine 6

Goroutine ID: 1
State: semacquire

Goroutine ID: 6
State: semacquire
Created at: main.main in goroutine 1


=== goroutines info: last main ===

num goroutines -> 1

Goroutine ID: 1
State: running
```
