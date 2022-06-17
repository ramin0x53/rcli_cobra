bin_path := ./bin/rcli
config_path := ~/.configrcli.json

build:
	go build -o $(bin_path)

rmconfig:
	if [ -f $(config_path) ]; then rm $(config_path); fi;

install: rmconfig
	go install