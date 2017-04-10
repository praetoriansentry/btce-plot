build:
	go build -o plotter cmd/plotter/plotter.go

clean:
	$(RM) plotter
	$(RM) graph.png

neat:
	find . -type f -name '*.go' | xargs -I{} go fmt {}

