.PHONY: build
build:
	@echo "Building..."
	@truffle compile
	@abigen --abi=build/contracts/ArtistCollection.abi --pkg=contracts --out=artist_collection.go
