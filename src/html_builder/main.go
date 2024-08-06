package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tmpl := template.Must(template.ParseGlob("source/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("source/group/*.html"))

	// 出力ファイルを作成
	f, err := os.Create("target/output.html")
	if err != nil {
		log.Fatalf("ファイル作成に失敗しました: %v", err)
	}
	defer f.Close()

	// テンプレートをファイルに書き出す
	data := ""
	err = tmpl.ExecuteTemplate(f, "template.html", data)
	if err != nil {
		log.Fatalf("テンプレートの実行に失敗しました: %v", err)
	}

	log.Println("HTMLファイルが作成されました: target/output.html")
}
