.DEFAULT_GOAL = compile

compile:
	@pushd dist; maintainer go vanity build -f ../modules.yml; popd
.PHONY: compile
