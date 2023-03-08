DIR := ${CURDIR}

.DEFAULT_GOAL = help

E:=@
ifeq ($(V),1)
	E=
endif

.PHONY: default

default: build

############################################################################
# OS/ARCH detection
############################################################################
os1=$(shell uname -s)
os2=
ifeq ($(os1),Darwin)
os1=darwin
os2=osx
else ifeq ($(os1),Linux)
os1=linux
os2=linux
else ifeq (,$(findstring MYSYS_NT-10-0-, $(os1)))
os1=windows
os2=windows
else
$(error unsupported OS: $(os1))
endif

arch1=$(shell uname -m)
ifeq ($(arch1),x86_64)
arch2=amd64
else ifeq ($(arch1),aarch64)
arch2=arm64
else ifeq ($(arch1),arm64)
arch2=arm64
else
$(error unsupported ARCH: $(arch1))
endif


############################################################################
# Vars
############################################################################

build_dir := $(DIR)/.build/$(os1)-$(arch1)

go_version_full := $(shell cat .go-version)
go_version := $(go_version_full:.0=)
go_dir := $(build_dir)/go/$(go_version)

ifeq ($(os1),windows)
	go_bin_dir = $(go_dir)/go/bin
	go_url = https://storage.googleapis.com/golang/go$(go_version).$(os1)-$(arch2).zip
	exe=".exe"
else
	go_bin_dir = $(go_dir)/bin
	go_url = https://storage.googleapis.com/golang/go$(go_version).$(os1)-$(arch2).tar.gz
	exe=
endif

go_path := PATH="$(go_bin_dir):$(PATH)"

go-check:
ifeq (go$(go_version), $(shell $(go_path) go version 2>/dev/null | cut -f3 -d' '))
else
	@echo "Installing go $(go_version)..."
	$(E)rm -rf $(dir $(go_dir))
	$(E)mkdir -p $(go_dir)
	$(E)curl -sSfL $(go_url) | tar xz -C $(go_dir) --strip-components=1
endif

## Checks installed go version and prints the path it is installed.
go-bin-path: go-check
	@echo "$(go_bin_dir):${PATH}"

install-toolchain: go-check

.PHONY: build

## Compiles all Galadriel binaries.
build: credential-composer-add-dc

# This is the master template for compiling Go binaries
define binary_rule
.PHONY: $1
$1: | go-check
	@echo Building $1...
	$(E)$(go_path) go build -o $1 $2
endef

# This dynamically generates targets for each binary using
# the binary_rule template above
$(eval $(call binary_rule,credential-composer-add-dc,plugin.go))

CONTAINER_OPTIONS = docker podman
CONTAINER_EXEC := $(foreach exec,$(CONTAINER_OPTIONS),\
     $(if $(shell which $(exec)),$(exec)))

