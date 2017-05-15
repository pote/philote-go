DEPS = $(firstword $(subst :, ,$(GOPATH)))/up-to-date
GPM ?= gpm

all: test

$(DEPS): Godeps | $(dir $(DEPS))
	$(GPM) get
	touch $@

$(dir $(DEPS)):
	mkdir -p $@

test:  $(DEPS)
	go test

##
# You're a PHONY! Just a big, fat PHONY.
##

.PHONY: test
