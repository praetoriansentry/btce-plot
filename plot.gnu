set terminal pngcairo enhanced font "arial,8" fontscale 1.0 size {{.X}}, {{.Y}}
set output '{{.OutName}}'

set xdata time
set timefmt "%Y-%m-%dT%H:%M:%S"
set xrange["{{.XMin}}":"{{.XMax}}"]
set bars 4.0
set style fill empty

plot '{{.DatName}}' using 1:2:3:4:5 with candlesticks title 'Quartiles'


