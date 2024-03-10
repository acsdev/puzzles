prepare-volumes:
	mkdir -p $$HOME/Dev/docker/volumes/go
	mkdir -p $$HOME/Dev/docker/volumes/.m2

##
## AFTER JAVA 8 IT IS POSSILE TO RUN WITHOUT COMPILE
##
run-java-codingami-closet-to-zero:
	cd code/codingame/closet_to_zero/java && java Main.java
	
run-java-freecodecamp-fibonnaci:
	cd code/freecodecamp/fibonnaci/java && java Main.java

run-java-freecodecamp-numbers_cansum:
	cd code/freecodecamp/numbers_cansum/java && java Main.java

run-java-gnirut-valid-words:
	cd code/gnirut/valid_words/java && java Main.java

run-java-gnirut-perfect-matches:
	cd code/gnirut/perfect_matches/java && java Main.java

run-java-spoj-aba12c:
	cd code/spoj/aba12c/java && java Main.java ../input

run-java-spoj-acido:
	cd code/spoj/acido/java && java Main.java ../input

run-java-spoj-chocopj09:
	cd code/spoj/chocopj09/java && java Main.java ../input

run-java-spoj-test:
	cd code/spoj/test/java/ && java Main.java ../input

run-go-freecodecamp-fibonnaci:
	cd code/freecodecamp/fibonnaci/go && go run main.go

run-go-gnirut-valid-words:
	cd code/gnirut/valid_words/go && go run main.go

run-go-gnirut-perfect-matches:
	cd code/gnirut/perfect_matches/go && go run main.go

run-go-gnirut-sort-arrays:
	cd code/gnirut/sort_arrays/go && go run main.go

run-go-kattis-3dprinter:
	cd code/kattis/3dprinter/go && go run main.go ../input

run-go-kattis-2048:
	cd code/kattis/2048/go && go run main.go ../inputs/input

run-go-random-binary-tree-mirror:
	cd code/random/binary_tree_mirror/go/mirror && go test -v -run "^TestMirror*" 

run-go-spoj-aba12c:
	cd code/spoj/aba12c/go && go run main.go ../input

run-go-spoj-acido:
	cd code/spoj/acido/go && go run main.go ../input

run-go-spoj-chocopj09:
	cd code/spoj/chocopj09/go && go run main.go ../input

run-go-spoj-obitetri:
	cd code/spoj/obitetri/go && go run main.go ../input

run-go-spoj-placar:
	cd code/spoj/placar/go && go run main.go ../input

run-go-spoj-rouba:
	cd code/spoj/rouba/go && go run main.go ../input

run-go-spoj-test:
	cd code/spoj/test/go && go run main.go ../input