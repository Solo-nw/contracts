// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.4;

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
        ArtistCollection _a = new ArtistCollection(_collectionName, _baseURI, _artistAddr, _royaltyBasisPoints, _currency);
        artistOwnerMap[_artistAddr] = _a;
    }

    function getCollection(address _artistAddr) external view returns(address) {
        return address(artistOwnerMap[_artistAddr]);
    }

    function setBaseURIBatch(address[] memory _artistAddrs, string memory _baseURI) external onlyOwner {
        for (uint i = 0; i < _artistAddrs.length; i++) {
            artistOwnerMap[_artistAddrs[i]].setBaseURI(string(abi.encodePacked(_baseURI, _artistAddrs[i])));
        }
    }

    function setCurrencyBatch(address[] memory _artistAddrs, address[] memory _currencies) external onlyOwner {
        require(_artistAddrs.length == _currencies.length, "Array lengths must match");
        for (uint i = 0; i < _artistAddrs.length; i++) {
            artistOwnerMap[_artistAddrs[i]].setCurrency(_currencies[i]);
        }
    }
}
