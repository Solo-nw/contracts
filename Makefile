.PHONY: build
build:
	@echo "Building..."
	@truffle compile
	@python gen_script.py
	@abigen --abi=build/contracts/ArtistCollection.abi --pkg=contracts --out=artist_collection.go --bin=build/contracts/ArtistCollection.bin
