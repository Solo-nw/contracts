.PHONY: build
build:
	@echo "Building..."
	@truffle compile
	@abigen --abi=build/contracts/ArtistCollection.abi --pkg=artist --out=artist_collection.go
	@mv artist_collection.go ./build/go/artist_contract/artist_collection.go
