PREFIX  ?= /usr/local
BINDIR  ?= $(PREFIX)/bin
MANDIR  ?= $(PREFIX)/share/man/man1
DESTDIR ?=

GO      := go
BINS    := lsdesktop desklaunch
MANPAGES := doc/lsdesktop.1 doc/desklaunch.1

.PHONY: all build install uninstall clean test

all: build

build:
	$(GO) build -o lsdesktop ./cmd/lsdesktop
	$(GO) build -o desklaunch ./cmd/desklaunch

install: build
	install -Dm755 lsdesktop $(DESTDIR)$(BINDIR)/lsdesktop
	install -Dm755 desklaunch $(DESTDIR)$(BINDIR)/desklaunch
	install -Dm644 doc/lsdesktop.1 $(DESTDIR)$(MANDIR)/lsdesktop.1
	install -Dm644 doc/desklaunch.1 $(DESTDIR)$(MANDIR)/desklaunch.1

uninstall:
	rm -f $(DESTDIR)$(BINDIR)/lsdesktop
	rm -f $(DESTDIR)$(BINDIR)/desklaunch
	rm -f $(DESTDIR)$(MANDIR)/lsdesktop.1
	rm -f $(DESTDIR)$(MANDIR)/desklaunch.1

clean:
	rm -f $(BINS)

test:
	$(GO) test ./...
