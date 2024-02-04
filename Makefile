prepare-volumes:
	mkdir -p $$HOME/Dev/docker/volumes/go
	mkdir -p $$HOME/Dev/docker/volumes/.m2

run-fibonacci-go:
	go run go/freecodecamp/fibonnaci/fibonacci.go 

run-fibonacci-java:
	java java/freecodecamp/fibonacci/Fibonacc.java