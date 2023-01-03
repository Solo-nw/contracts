// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.4;



contract OwnableSimple {
  address internal owner;

  event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

  constructor (address _firstOwner) {
    owner = _firstOwner;
    emit OwnershipTransferred(address(0), _firstOwner);
  }

  modifier onlyOwner() {
    require(msg.sender == owner, "Ownable#onlyOwner: SENDER_IS_NOT_OWNER");
    _;
  }

  function _transferOwnership(address _newOwner) internal {
    require(_newOwner != address(0), "Ownable#transferOwnership: INVALID_ADDRESS");
    owner = _newOwner;
    emit OwnershipTransferred(owner, _newOwner);
  }

  function getOwner() public view returns (address) {
    return owner;
  }
}