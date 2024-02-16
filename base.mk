PATH_TO_FILES ?= ../testing
FILE_SUFFIX   ?= lab

KEY_FILE     ?= $(PATH_TO_FILES)/key_$(FILE_SUFFIX).txt
FILE_INPUT   ?= $(PATH_TO_FILES)/input.png
FILE_ENCODED ?= $(PATH_TO_FILES)/encoded_$(FILE_SUFFIX)
FILE_DECODED ?= $(PATH_TO_FILES)/decoded_$(FILE_SUFFIX).png

.PHONY: clear
clear:
	@rm -rf encdr
	@rm -rf $(PATH_TO_FILES)/*_$(FILE_SUFFIX)*

.PHONY: build
build:
	@go build -ldflags "-s -w" -o ./encdr main.go

.PHONY: text
text:
	@read -p "KEY: " KEY \
	&& ./encdr $${KEY} | ./encdr -d $${KEY} \
	&& echo $${KEY} > $(KEY_FILE)

.PHONY: file
file:
	@read -p "KEY: " KEY \
	&& ./encdr -i $(FILE_INPUT) -o $(FILE_ENCODED) $${KEY} \
	&& ./encdr -d -i $(FILE_ENCODED) -o $(FILE_DECODED) $${KEY} \
	&& echo $${KEY} > $(KEY_FILE) \
	&& echo "$$(sha1sum $(FILE_INPUT) | grep -o "^\S\+")  $(FILE_DECODED)" | sha1sum --check
