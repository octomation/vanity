.DEFAULT_GOAL = compile

compile:
	@ ./Taskfile compile
.PHONY: compile

list:
	@ ./Taskfile modules
.PHONY: list
