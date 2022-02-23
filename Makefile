.DEFAULT_GOAL = compile

compile:
	@rm -rf dist/*; pushd dist; maintainer go vanity build -f ../modules.yml; popd
.PHONY: compile
