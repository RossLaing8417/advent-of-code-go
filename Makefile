year?=0
day?=0
part?=0

default: help

##@
##@ Usage commands
##@

.PHONY: gen
gen:	##@ Generate the solution files if they don't exist, including downloading the input data if the file .aoc_session exists and contians your session cookie: make gen [year=value] [day=value]
	@go run main.go -generate -year $(year) -day $(day) -part $(part)

.PHONY: test
test:	##@ Test the solution(s): make test [year=value] [day=value] [part=value]
	@go run main.go -test -year $(year) -day $(day) -part $(part) -cpuprofile="$(cpuprofile)" -memprofile="$(memprofile)"

.PHONY: run
run:	##@ Run the solution: make run [year=value] [day=value] [part=value]
	@go run main.go -run -year $(year) -day $(day) -part $(part) -cpuprofile="$(cpuprofile)" -memprofile="$(memprofile)"

##@
##@ Misc commands
##@

.PHONY: help
help: ##@ (Default) Print listing of key targets with their descriptions
	@printf "\nUsage: make <command>\n"
	@grep -F -h "##@" $(MAKEFILE_LIST) | grep -F -v grep -F | sed -e 's/\\$$//' | awk 'BEGIN {FS = ":*[[:space:]]*##@[[:space:]]*"}; \
	{ \
		if($$2 == "") \
			pass; \
		else if($$0 ~ /^#/) \
			printf "\n%s\n", $$2; \
		else if($$1 == "") \
			printf "     %-20s%s\n", "", $$2; \
		else \
			printf "\n    \033[34m%-20s\033[0m %s\n", $$1, $$2; \
	}'
