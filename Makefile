libraries = var_dump

.PHONY: all/% format/% test/%

all: $(foreach library,$(libraries),all/$(library))

all/%: %
	cd "$*"; go build

test: $(foreach library,$(libraries),test/$(library))

test/%: %
	cd "$*"; go test

format: $(foreach library,$(libraries),format/$(library))

format/%: %
	cd "$*"; gofmt -d=true -tabs=false -tabwidth=2 -w=true .
