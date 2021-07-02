package main

func main() {
	// injectionなServerの取得
	s := InitServer()

	// サーバーの起動
	s.Run()
}
