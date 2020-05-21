all = build

clean:
	rm -rf build
	rm -rf dist
	rm -rf *.egg-info
build:
	python3 setup.py sdist bdist_wheel
publish:
	python3 -m twine upload  dist/*