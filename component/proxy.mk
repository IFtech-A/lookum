PROXY_PROJECT_NAME=reverse_proxy
PROXY_TARGET_NAME=proxy
PROXY_SRC=$(SRC_DIR)/$(PROXY_PROJECT_NAME)
PROXY_BUILD=$(BUILD_DIR)/$(PROXY_PROJECT_NAME)
PROXY_PACKAGE=$(PACKAGE_DIR)/$(PROXY_PROJECT_NAME)/$(PROXY_TARGET_NAME)

.PHONY: proxy-distclean proxy-package proxy-docker proxy-prepare

proxy-prepare: prepare
	$(Q)$(MKDIR) $(PROXY_BUILD)
	$(Q)$(MKDIR) $(PROXY_BUILD)/ssl

	$(Q)$(MKDIR) $(PROXY_PACKAGE)
	
	$(Q)$(LNDIR) $(PROXY_SRC) $(PROXY_BUILD)
	$(Q)$(LNDIR) $(SRC_DIR)/cert $(PROXY_BUILD)/ssl

proxy-distclean:
	$(Q)$(RM)r $(PROXY_BUILD)
	$(Q)(echo "Build directory removed: $(PROXY_BUILD)")
	$(Q)$(RM)r $(PROXY_PACKAGE)
	$(Q)(echo "Package directory removed: $(PROXY_PACKAGE)")

proxy-package: proxy-prepare
	$(Q)$(CP)r $(PROXY_BUILD)/* $(PROXY_PACKAGE)
	$(Q)$(SCRIPT_DIR)/translate-links.sh $(PROXY_PACKAGE) $(PROXY_SRC)
	$(Q)$(SCRIPT_DIR)/translate-links.sh $(PROXY_PACKAGE) $(SRC_DIR)
	$(Q)(cat $(PROXY_PACKAGE)/ssl/server.crt $(PROXY_PACKAGE)/ssl/ca.crt > $(PROXY_PACKAGE)/ssl/tmp && mv $(PROXY_PACKAGE)/ssl/tmp $(PROXY_PACKAGE)/ssl/server.crt)
	$(Q)(echo "Packaging completed in directory: $(PROXY_PACKAGE)")

proxy-docker: proxy-package
	$(Q)(cd $(DOCKER_DIR) && ./build.sh $(PROXY_TARGET_NAME))
