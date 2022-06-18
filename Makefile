GO = go

# V is the verbosity flag. When 0, no commands will be output during execution
# of the makefile. When 1, all commands are output. Q is a verbosity prefix.
# When V is 1, the filter is non-empty, the if falls through to the true branch,
# and the whole Q is therefore blank. When V is 0, the whole Q is sell to @ and
# the prefix will silence any command it's prefixing.
V  = 0
Q  = $(if $(filter 1,$V),,@)

OS = $(shell uname -s | tr A-Z a-z)
CFLAGS ?= -g -O2
BUILDFLAGS = GOFLAGS=-mod=vendor
APPVERSION=$(or $(shell git rev-parse HEAD 2>/dev/null),"unknown")
PROGRAMS = server runner

.PHONY: all
all: build

.PHONY: build
build: $(PROGRAMS)

$(PROGRAMS):
	@echo "Building $@"
	$Q $(BUILDFLAGS) $(GO) build \
		-o bin/$@ ./cmd/$@

# Misc
.PHONY: clean
clean:
	$Q rm -f bin/$(PROGRAMS)
	$Q $(BUILDFLAGS) $(GO) clean ./...
