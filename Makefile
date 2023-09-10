install:
	@mkdir -p $(HOME)/.goblaq 2>&1 > /dev/null
	go install $(PWD)/cmd/goblaq