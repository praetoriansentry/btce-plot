build:
	go build -o plotter cmd/plotter/plotter.go
	go build -o loader cmd/loader/loader.go

clean:
	$(RM) plotter
	$(RM) loader
	$(RM) graph.png

neat:
	find . -type f -name '*.go' | xargs -I{} go fmt {}

install:
	go install cmd/plotter/plotter.go
	go install cmd/loader/loader.go
