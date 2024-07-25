# QA

你今天要用golang寫一個車的模擬器，使用標準輸出來作唯一輸出介面，並使用github來繳交每一個版本

1. 你模擬了一台機車，機車上面有一具引擎，引擎具有啟動函數與一個加速度屬性
機車：此為一結構
啟動：此為函數，為機車的一個動作，函式只要印出"啟動了機車引擎"。
引擎：此為結構有實作字串函數，能印出加速度的數字，加速度為1。

下一題是腳踏車產品的進化

2. 你的腳踏車引擎非常受歡迎，即將有卡車、轎車、休旅車要採用這個引擎。
這些車都車的一種，都使用了你的引擎，其加速度也各自不同，請宣告這幾種車並且在啟動時印出，比如卡車："啟動了卡車引擎"。

# Build & Run

```bash
make && ./biker
```
