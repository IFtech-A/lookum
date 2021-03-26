LOOKUM_PROJECT_NAME=backend
LOOKUM_TARGET_NAME=api
LOOKUM_SRC=$(SRC_DIR)/$(LOOKUM_PROJECT_NAME)
LOOKUM_BUILD=$(BUILD_DIR)/$(LOOKUM_PROJECT_NAME)
LOOKUM_PACKAGE=$(PACKAGE_DIR)/$(LOOKUM_PROJECT_NAME)/$(LOOKUM_TARGET_NAME)

.PHONY: api api-clean api-distclean api-package api-docker api-rebuild

api: prepare
	$(Q)$(MKDIR) $(LOOKUM_BUILD)
	$(Q)$(LNDIR) $(LOOKUM_SRC) $(LOOKUM_BUILD)
	$(Q)(cd $(LOOKUM_BUILD) && $(MAKE) $(LOOKUM_TARGET_NAME))
	$(Q)(echo "Build completed in directory: $(LOOKUM_BUILD)")

api-clean:
	$(Q)(cd $(LOOKUM_BUILD) && $(MAKE) clean)
	$(Q)(echo "Build directory cleaned: $(LOOKUM_BUILD)")

api-distclean:
	$(Q)$(RM)r $(LOOKUM_BUILD)
	$(Q)(echo "Build directory removed: $(LOOKUM_BUILD)")
	$(Q)$(RM)r $(LOOKUM_PACKAGE)
	$(Q)(echo "Package directory removed: $(LOOKUM_PACKAGE)")

api-rebuild: api-clean api

api-package: api
	$(Q)$(MKDIR) $(LOOKUM_PACKAGE)/bin
	$(Q)$(CP)r $(LOOKUM_BUILD)/bin $(LOOKUM_PACKAGE)/
	$(Q)$(CP)r $(LOOKUM_BUILD)/conf $(LOOKUM_PACKAGE)/
	$(Q)$(CP)r $(LOOKUM_BUILD)/web $(LOOKUM_PACKAGE)/
	$(Q)$(SCRIPT_DIR)/translate-links.sh $(LOOKUM_PACKAGE) $(LOOKUM_SRC)
	$(Q)(echo "Packaging completed in directory: $(LOOKUM_PACKAGE)")

api-docker: api-package
	$(Q)(cd $(DOCKER_DIR) && ./build.sh $(LOOKUM_TARGET_NAME))
