ECHO=echo
GIT=git
GO=go
MKDIR=mkdir
MV=mv
RM=rm

#

help:
	@echo "make <command>"
	@echo "  install"
	@echo "    Install dependencies"
	@echo "  clean"
	@echo "    Remove dependencies"

install: backup clone get restore

clean: backup clear restore

#

backup:
	-$(MKDIR) .temp
	$(MV) function/main.go .temp/main.go

restore:
	-$(MKDIR) function
	$(MV) .temp/main.go function/main.go
	$(RM) -rf .temp

get:
	-$(GO) get -u github.com/GoogleCloudPlatform/cloud-functions-go

clone: clear
	$(GIT) clone https://github.com/GoogleCloudPlatform/cloud-functions-go function
	$(RM) -rf function/.git

clear:
	$(RM) -rf function