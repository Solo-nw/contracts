// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.4;

import "@0xsequence/erc-1155/contracts/tokens/ERC1155/ERC1155MintBurn.sol";
import "@0xsequence/erc-1155/contracts/tokens/ERC1155/ERC1155Metadata.sol";
import "@0xsequence/erc-1155/contracts/tokens/ERC2981/ERC2981Global.sol";
import "@0xsequence/erc-1155/contracts/utils/SafeMath.sol";
import "./utils/Ownable.sol";
import "./utils/IERC20.sol";

contract ArtistCollection is ERC1155MintBurn, ERC1155Metadata, ERC2981Global, Ownable{
    using SafeMath for uint256;

    /***********************************|
    |             Variables             |
    |__________________________________*/
    IERC20 internal currency;
    mapping(uint256 => uint256) internal tokenPrices; // Stores the Price
    /***********************************|
    |             Constuctor            |
    |__________________________________*/
    constructor(string memory _name, string memory _baseURI, address _artistAddr, uint256 _royaltyBasisPoints, address _currency) 
                ERC1155Metadata(_name, _baseURI)
                Ownable(msg.sender, _artistAddr) 
                {
                    _setGlobalRoyaltyInfo(_artistAddr, _royaltyBasisPoints);
                    currency = IERC20(_currency);
                }

    /***********************************|
    |      Receiver Method Handler      |
    |__________________________________*/
    /**
    * @notice Prevents receiving Ether or calls to unsuported methods
    */
    fallback () external {
        revert("UNSUPPORTED_METHOD");
    }

    // sets token id base uri... 
    function setBaseURI(string calldata _newuri) public onlyCreator {
        _setBaseMetadataURI(_newuri);
    }

    function setContractName(string calldata _newName) public onlyCreator {
        _setContractName(_newName);
    }

    function transferOwnership(address _newOwner, uint256 _royaltyBasisPoints) public onlyOwner {
        _transferOwnership(_newOwner);
        _setGlobalRoyaltyInfo(_newOwner, _royaltyBasisPoints);
    }

    /***********************************|
    |         Getter Functions          |
    |__________________________________*/

    function getTokenPrice(uint256 _id) external view returns(uint256) {
        return tokenPrices[_id];
    }

    function getCurrency() external view returns(address) {
        return address(currency);
    }

    function setCurrency(address _currency) external onlyOwner {
        currency = IERC20(_currency);
    }
    
    function currenctyBalance() external view returns(uint256){
        return currency.balanceOf(msg.sender);
    }

    // sets price if it was not set... while minting it can be reset
    function setPrice(uint256 _id, uint256 _price) external onlyOwner {
        tokenPrices[_id] = _price;
    }
    /***********************************|
    |         Market Functions          |
    |__________________________________*/
    modifier tokenAvailable(uint256 _id) {
        require(balanceOf(owner, _id)>0);
        _;
    }

    function buy(uint256 _id, uint256 _amount, bytes memory _data) external 
        tokenAvailable(_id)
    {
        uint256 totalPrice = tokenPrices[_id] * _amount;
        bool _success = currency.transferFrom(msg.sender,owner,totalPrice);
        require(_success, "Buyer Couldn't Transfer Currency");
        _safeTransferFrom(owner, msg.sender, _id, _amount);
        _callonERC1155Received(owner, msg.sender, _id, _amount, gasleft(), _data);
    }


    /***********************************|
    |    Royalty Management Methods     |
    |__________________________________*/
    /**
    * @notice Will set the basis point and royalty recipient that is applied to all Skyweaver assets
    * @param _royaltyBasisPoints Basis points with 3 decimals representing the fee %
    *        e.g. a fee of 2% would be 20 (i.e. 20 / 1000 == 0.02, or 2%)
    */
    function setGlobalRoyaltyInfo(uint256 _royaltyBasisPoints) external onlyOwner() {
        _setGlobalRoyaltyInfo(owner, _royaltyBasisPoints);
    }

    /***********************************|
    |         Minting Methods           |
    |__________________________________*/
    function mint(uint256 id, uint256 amount, uint256 price, bytes memory data)
        public
        onlyOwner
    {
        _mint(owner, id, amount, data);
        tokenPrices[id] = price;
    }

    function mintBatch(uint256[] memory ids, uint256[] memory amounts, uint256[] memory prices, bytes memory data)
        public
        onlyOwner
    {
        _batchMint(owner, ids, amounts, data);
        uint i=0;
        for (i = 0; i < ids.length; i++) {
            tokenPrices[ids[i]] = prices[i];
        }
    }   

    /***********************************|
    |          Burning Functions        |
    |__________________________________*/

    /**
    * @notice Burn _amount of tokens of a given id from msg.sender
    * @dev This will not change the current issuance tracked in _supplyManagerAddr.
    * @param _id     Asset id to burn
    * @param _amount The amount to be burn
    */
    function burn(
        uint256 _id,
        uint256 _amount)
        external
    {
        _burn(msg.sender, _id, _amount);
    }

    /**
    * @notice Burn _amounts of tokens of given ids from msg.sender
    * @dev This will not change the current issuance tracked in _supplyManagerAddr.
    * @param _ids     Asset id to burn
    * @param _amounts The amount to be burn
    */
    function batchBurn(
        uint256[] calldata _ids,
        uint256[] calldata _amounts)
        external
    {
        _batchBurn(msg.sender, _ids, _amounts);
    }

    /***********************************|
    |          ERC165 Functions         |
    |__________________________________*/
    /**
    * @notice Query if a contract implements an interface
    * @param _interfaceID  The interface identifier, as specified in ERC-165
    * @return `true` if the contract implements `_interfaceID`
    */
    function supportsInterface(bytes4 _interfaceID) public override(ERC1155, ERC1155Metadata, ERC2981Global) virtual pure returns (bool) {
        return super.supportsInterface(_interfaceID);
    }
}
