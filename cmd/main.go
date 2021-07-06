package main

func main() {
	// injectionなServerの取得
	s := NewServer()

	// サーバーの起動
	s.Run()
}
