include $(GOROOT)/src/Make.inc

TARG=gobot
GOFILES=\
	net.go \
	irc.go \
	bot.go \
	main.go 

include $(GOROOT)/src/Make.cmd
