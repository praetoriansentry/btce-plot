set terminal pngcairo enhanced font "arial,8" fontscale 1.0 size {{.X}}, {{.Y}}
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

set title "BTC-E ETH Prices" textcolor rgbcolor "white"


set xdata time
set timefmt "%Y-%m-%dT%H:%M:%S"
set xrange["{{.XMin}}":"{{.XMax}}"]
set bars 4.0
set style fill empty

plot '{{.DatName}}' using 1:2:3:4:5:($5 < $2 ? -1 : 1) with candlesticks palette

#title 'Quartiles'


