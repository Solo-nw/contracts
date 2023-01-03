// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.4;
pragma experimental ABIEncoderV2;

import "./Artist.sol";
import "./utils/OwnableSimple.sol";

contract Deployer is OwnableSimple {
    mapping(address => ArtistCollection) internal artistOwnerMap;

    constructor () OwnableSimple(msg.sender) {}

    modifier isNotDeployed(address _artistAddr) {
        require(address(artistOwnerMap[_artistAddr])==address(0) || msg.sender == owner, "Artist Collection already deployed");
        _;
    }

    function deploy(string memory _collectionName, string memory _baseURI, address _artistAddr, uint256 _royaltyBasisPoints, address _currency) external isNotDeployed(_artistAddr){
        artistOwnerMap[_artistAddr] = new ArtistCollection(_collectionName, _baseURI, _artistAddr, _royaltyBasisPoints, _currency);
    }

    function getAllCollection(address[] memory _artistAddrs) external view returns(address[] memory) {
        address[] memory artistCollectionAddrs;
        for (uint i = 0; i < _artistAddrs.length; i++) {
            artistCollectionAddrs[i] = address(artistOwnerMap[_artistAddrs[i]]);
        }
        return artistCollectionAddrs;
    }

    function setBaseURIBatch(address[] memory _artistAddrs, string[] memory _baseURIs) external onlyOwner {
        require(_artistAddrs.length == _baseURIs.length, "Array lengths must match");
        for (uint i = 0; i < _artistAddrs.length; i++) {
            artistOwnerMap[_artistAddrs[i]].setBaseURI(_baseURIs[i]);
        }
    }

    function setCurrencyBatch(address[] memory _artistAddrs, address[] memory _currencies) external onlyOwner {
        require(_artistAddrs.length == _currencies.length, "Array lengths must match");
        for (uint i = 0; i < _artistAddrs.length; i++) {
            artistOwnerMap[_artistAddrs[i]].setCurrency(_currencies[i]);
        }
    }
}
