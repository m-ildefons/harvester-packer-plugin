TARGETS := $(shell ls scripts)

.DEFAULT_GOAL := default

.dapper:
	@echo Downloading dapper
	@curl -sL https://releases.rancher.com/dapper/latest/dapper-$$(uname -s)-$$(uname -m) > .dapper.tmp
	@@chmod +x .dapper.tmp
	@./.dapper.tmp -v
	@mv .dapper.tmp .dapper

$(TARGETS): .dapper
	./.dapper $@

install:
	@packer plugins install --path bin/packer-plugin-harvester-amd64 packer.local/local/harvester
