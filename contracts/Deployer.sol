// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.4;

import "./Artist.sol";
import "./utils/Ownable.sol";

contract Deployer is Ownable {
    mapping(address => ArtistCollection) internal artistOwnerMap;

    constructor () Ownable(msg.sender, msg.sender) {}

    function deploy(string memory _collectionName, string memory _baseURI, address _artistAddr, uint256 _royaltyBasisPoints, address _currency) external onlyCreator {
        ArtistCollection _a = new ArtistCollection(_collectionName, _baseURI, _artistAddr, _royaltyBasisPoints, _currency);
        artistOwnerMap[_artistAddr] = _a;
    }

    function getArtistCollection(address _artistAddr) external view returns(address) {
        return address(artistOwnerMap[_artistAddr]);
    }
}
