.PHONY: install

PREFIX ?= /usr/local

install:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -m 0755 coolops $(DESTDIR)$(PREFIX)/bin/coolops
