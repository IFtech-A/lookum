DB_PROJECT_NAME=backend
DB_TARGET_NAME=db
DB_SRC=$(SRC_DIR)/$(DB_PROJECT_NAME)
DB_BUILD=$(BUILD_DIR)/$(DB_PROJECT_NAME)
DB_PACKAGE=$(PACKAGE_DIR)/$(DB_PROJECT_NAME)/$(DB_TARGET_NAME)

.PHONY: db-distclean db-package db-docker db-prepare

db-prepare: prepare
	$(Q)$(MKDIR) $(DB_BUILD)
	$(Q)$(LNDIR) $(DB_SRC) $(DB_BUILD)

db-distclean:
	$(Q)$(RM)r $(DB_BUILD)
	$(Q)(echo "Build directory removed: $(DB_BUILD)")
	$(Q)$(RM)r $(DB_PACKAGE)
	$(Q)(echo "Package directory removed: $(DB_PACKAGE)")

db-package: db-prepare
	$(Q)$(MKDIR) $(DB_PACKAGE)/sql
	$(Q)$(CP)r $(DB_BUILD)/sql $(DB_PACKAGE)/
	$(Q)$(SCRIPT_DIR)/translate-links.sh $(DB_PACKAGE) $(DB_SRC)
	$(Q)(echo "Packaging completed in directory: $(DB_PACKAGE)")

db-docker: db-package
	$(Q)(cd $(DOCKER_DIR) && ./build.sh $(DB_TARGET_NAME))
