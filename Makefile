default:
	@export GOPATH=$$GOPATH:$$(pwd)  && go install dnslookup
run: default
	@bin/dnslookup
	@echo ""
push:
	git add --all && git commit -am 'update' && git push origin master
clean:
	@rm -rf bin
edit:
	@export GOPATH=$$GOPATH:$$(pwd)  && atom .
