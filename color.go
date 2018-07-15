package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// HTTPリクエストからvalueNameで指定されたフォームデータの色情報を抽出する
// col[0] : R値
// col[1] : G値
// col[2] : B値
// col[3] : RGBの合計
func stripColor(r *http.Request, valueName string) [4]int64 {
	color := r.FormValue(valueName)
	var col = [4]int64{}
	col[0], _ = strconv.ParseInt(color[0:2], 16, 0)
	col[1], _ = strconv.ParseInt(color[2:4], 16, 0)
	col[2], _ = strconv.ParseInt(color[4:6], 16, 0)
	col[3] = col[0] + col[1] + col[2]
	return col
}

// /colorにきたリクエストを捌くハンドラ
func color(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		post(w, r)
		return
	}

	cols := [3][4]int64{
		stripColor(r, "color0"),
		{0, 0, 0, 0},
		stripColor(r, "color1"),
	}

	for i := 0; i < 3; i += 1 {
		cols[1][i] = int64(float64(cols[2][i]) / float64(cols[0][i]) * 255.0)
	}

	colstr := [3]string{}
	for i := 0; i < 3; i += 1 {
		for j := 0; j < 3; j += 1 {
			colstr[i] += fmt.Sprintf("%02x", cols[i][j])
		}
	}

	m := map[string]string{
		"color0": colstr[0],
		"color1": colstr[1],
		"color2": colstr[2],
	}

	t := template.Must(template.ParseFiles("templates/color.html"))

	if err := t.ExecuteTemplate(w, "color.html", m); err != nil {
		log.Fatal(err)
	}
}
