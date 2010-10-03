include $(GOROOT)/src/Make.inc

TARG=gobot
GOFILES=\
	net.go \
	main.go 

include $(GOROOT)/src/Make.cmd
