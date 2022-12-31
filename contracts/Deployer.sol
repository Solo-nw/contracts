// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.4;

import "./Artist.sol";
import "./utils/Ownable.sol";

contract Deployer is Ownable {
    mapping(address => ArtistCollection) internal artistOwnerMap;

    constructor () Ownable(msg.sender, msg.sender) {}

    modifier isNotDeployed(address _artistAddr) {
        require(address(artistOwnerMap[_artistAddr])==address(0), "Artist Collection already deployed");
        _;
    }

    function deploy(string memory _collectionName, string memory _baseURI, address _artistAddr, uint256 _royaltyBasisPoints, address _currency) external isNotDeployed(_artistAddr){
        ArtistCollection _a = new ArtistCollection(_collectionName, _baseURI, _artistAddr, _royaltyBasisPoints, _currency);
        artistOwnerMap[_artistAddr] = _a;
    }

    // ONLY USED TO REPLACE OLD CONTRACTS
    function forceDeploy(string memory _collectionName, string memory _baseURI, address _artistAddr, uint256 _royaltyBasisPoints, address _currency) external onlyCreator{
        ArtistCollection _a = new ArtistCollection(_collectionName, _baseURI, _artistAddr, _royaltyBasisPoints, _currency);
        artistOwnerMap[_artistAddr] = _a;
    }

    function getAllCollection(address[] memory _artistAddrs) external view returns(address[] memory) {
        address[] memory artistCollectionAddrs;
        for (uint i = 0; i < _artistAddrs.length; i++) {
            artistCollectionAddrs[i] = address(artistOwnerMap[_artistAddrs[i]]);
        }
        return artistCollectionAddrs;
    }
    
    function getArtistCollection(address _artistAddr) external view returns(address) {
        return address(artistOwnerMap[_artistAddr]);
    }
}
