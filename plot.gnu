{{ if eq .Mode "canvas" }}
set terminal canvas enhanced font "arial,12" fontscale 1.0 size {{.X}}, {{.Y}} jsdir './js' mousing
{{ else }}
set terminal pngcairo enhanced font "arial,12" fontscale 1.0 size {{.X}}, {{.Y}}
{{ end }}
set output '{{.OutName}}'

set border linecolor rgbcolor "yellow"
set key textcolor rgbcolor "white"

set obj 1 rectangle behind from screen 0,0 to screen 1,1
set obj 1 fillstyle solid 1.0 fillcolor rgbcolor "black"

set palette defined (-1 'red', 1 'green')
set cbrange [-1:1]
unset colorbox

set style fill solid noborder
set boxwidth {{ .BucketSeconds }} absolute

set xdata time
set timefmt "%Y-%m-%dT%H:%M:%S"
set xrange["{{.XMin}}":"{{.XMax}}"]
set bars 4.0
set style fill empty

set format x "%m/%d %H:%M"

set title "BTC-E ETH Prices" textcolor rgbcolor "white"
set xlabel "Time in UTC" textcolor "yellow"
show xlabel
set ylabel "ETH Price in USD" textcolor "yellow"
show ylabel

plot '{{.DatName}}' using 1:2:3:4:5:($5 < $2 ? -1 : 1) with candlesticks palette title 'Candlesticks'


