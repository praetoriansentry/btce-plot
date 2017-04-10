set terminal pngcairo  transparent enhanced font "arial,8" fontscale 1.0 size {{.X}}, {{.Y}}
set output '{{.OutName}}'

set bar 1.000000 front
set style circle radius graph 0.02, first 0.00000, 0.00000
set style ellipse size graph 0.05, 0.03, first 0.00000 angle 0 units xy
set style textbox transparent margins  1.0,  1.0 border
unset logscale
set ytics  norangelimit
set title "Demo of plotting financial data"

set paxis 1 range [ * : * ] noreverse nowriteback
set paxis 2 range [ * : * ] noreverse nowriteback
set paxis 3 range [ * : * ] noreverse nowriteback
set paxis 4 range [ * : * ] noreverse nowriteback
set paxis 5 range [ * : * ] noreverse nowriteback
set paxis 6 range [ * : * ] noreverse nowriteback
set paxis 7 range [ * : * ] noreverse nowriteback
set lmargin  9
set rmargin  2
set colorbox vertical origin screen 0.9, 0.2, 0 size screen 0.05, 0.6, 0 front
x = 0.0
## Last datafile plotted: "finance.dat"
plot '{{.DatName}}' using 0:5 notitle with lines