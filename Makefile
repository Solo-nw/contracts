.PHONY: build deploy
build:
	@echo "Building..."
	@truffle compile
	@python gen_script.py
	@abigen --abi=build/contracts/ArtistCollection.abi --pkg=artist --out=./artist/artist_collection.go
	@abigen --abi=build/contracts/Deployer.abi --pkg=deployer --out=./deployer/deployer.go

deploy:
	@echo "Deploying..."
	@truffle migrate --network mumbai
	@truffle run verify Deployer --network=mumbai
