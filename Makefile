VENV_PATH = ./venv
VENV = . $(VENV_PATH)/bin/activate;

.PHONY: run
run:
	cd adapter; go run main.go
configure:
	python3 -m venv $(VENV_PATH)
	$(VENV) pip install -r requirements.txt
clean:
	rm -rf venv