# Makefile for common development tasks
menu:
	@echo 'build: Build iwd-frequency'
	@echo 'install: Install iwd-frequency to ~/.local/bin'

build:
	@echo 'Building iwd-frequency'
	@go build -o iwd-frequency

install:
	@echo 'Installing iwd-frequency to ~/.local/bin'
	@mv iwd-frequency /home/mcritchlow/.local/bin
