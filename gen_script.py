# read json file from build/contracts
# generate script for deployment

import json
import sys

with open("build/contracts/ArtistCollection.json", 'r') as f:
    data = json.load(f)

abi = data['abi']
bytecode = data['bytecode']
# write abi and bytecode to seperate files
with open("build/contracts/ArtistCollection.abi", 'w') as f:
    f.write(json.dumps(abi))

with open("build/contracts/ArtistCollection.bin", 'w') as f:
    f.write(bytecode)

