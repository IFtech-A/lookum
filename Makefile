Q:=@
MKDIR:=mkdir -p
LNDIR:=lndir -silent
CP:=cp -fP
RM:=rm -f

ifeq ($(V),1)
Q:=
MKDIR:=$(MKDIR)v
LNDIR:=lndir
CP:=$(CP)v
RM:=$(RM)v
endif

BASE_DIR:=$(CURDIR)
COMPONENT_DIR:=$(BASE_DIR)/component
SCRIPT_DIR:=$(BASE_DIR)/script
DOCKER_DIR:=$(BASE_DIR)/docker
SRC_DIR:=$(BASE_DIR)/src
BUILD_DIR:=/home/src/lookum
PACKAGE_DIR:=/package/lookum

include $(COMPONENT_DIR)/*.mk

.DEFAULT_GOAL:=help

.PHONY: all build prepare clean distclean docker package

help:
		@echo "Generic targets"
		@echo "---------------"
		@echo "  help                     Display this help"
		@echo "  all                      Build All"
		@echo "  package                  Package All"
		@echo "  docker                   Build docker image All"
		@echo "  clean                    Clean All"
		@echo "  distclean                Distclean All"
		@echo ""
		@echo "  API TARGETS"
		@echo "  ------------------"
		@echo "  api                      Build api locally"
		@echo "  api-clean                Clean api"
		@echo "  api-distclean            Remove api build data"
		@echo "  api-package              Make api package"
		@echo "  api-docker               Build api docker image"
		@echo "  ------------------"
		@echo "  DB TARGETS"
		@echo "  ------------------"
		@echo "  db-distclean            Remove db package data"
		@echo "  db-package              Make db package"
		@echo "  db-docker               Build db docker image"
		@echo "  ------------------"
		@echo "  REVERSE PROXY TARGETS"
		@echo "  ------------------"
		@echo "  proxy-distclean            Remove nginx proxy package data"
		@echo "  proxy-package              Make nginx proxy package"
		@echo "  proxy-docker               Build nginx proxy docker image"
		@echo ""


prepare:
	$(Q)$(MKDIR) $(BUILD_DIR) $(PACKAGE_DIR)

clean: api-clean

distclean: api-distclean db-distclean proxy-distclean

all: api

docker: api-docker db-docker proxy-docker

package: api-package db-package proxy-docker

