pragma solidity  ^0.8.24;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC721/ERC721.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/IERC20.sol";

// @author yueliyangzi
contract METANFT is ERC721{
    using SafeMathCell for uint256;
    IERC20 public usdtToken;
    IERC20 public uncToken;
    address public owner;
    mapping(uint256 => uint256) usdtPrice;
    mapping(uint256 => uint256) uncPrice;

    constructor(address _usdtTokenAddress,address _uncTokenAddress,address holder) ERC721("NFTContract", "NFTC") {
        usdtToken = IERC20(_usdtTokenAddress);
        uncToken = IERC20(_uncTokenAddress);
        owner = holder;
    }

    function mintNFT(address recipient, uint256 tokenId) external onlyManager{
        require(msg.sender == owner, "Only owner can mint NFTs");
        _safeMint(recipient, tokenId);
        usdtPrice[tokenId] = 0;
        uncPrice[tokenId] = 0;
    }
    function mintNFTs(address recipient, uint256[] memory tokenIds) external onlyManager{
        require(msg.sender == owner, "Only owner can mint NFTs");
        for (uint256 i = 0; i < tokenIds.length; i++) {
            _safeMint(recipient, tokenIds[i]);
            usdtPrice[tokenIds[i]] = 0;
            uncPrice[tokenIds[i]] = 0;
        }
    }
    function setUsdtPrice(uint256 tokenId,uint256 price) external onlyManager{
        usdtPrice[tokenId] = price;
    }
    function setUncPrice(uint256 tokenId,uint256 price) external onlyManager{
        uncPrice[tokenId] = price;
    }
    function transferOwnership(address newOwner) external {
        require(msg.sender == owner, "Only owner can transfer ownership");
        owner = newOwner;
    }

    function triggerOwnershipTransferUsdt(uint256 tokenId) external {
        //require(usdtPrice(tokenId) != 0, "nft has not set price");
        require(usdtToken.balanceOf(msg.sender) >= usdtPrice[tokenId], "Insufficient allowance");
        // Transfer USDT to the contract
        if (usdtPrice[tokenId] != 0 ){
            require(usdtToken.transferFrom(msg.sender, owner, usdtPrice[tokenId]), "USDT transfer failed");
            // Transfer ownership of NFT
            _transfer(owner,msg.sender, tokenId);
        }

    }
    function triggerOwnershipTransferUnc(uint256 tokenId) external  {
        //require(uncPrice(tokenId) != 0, "nft has not set price");
        require(usdtToken.balanceOf(msg.sender) >= uncPrice[tokenId], "Insufficient allowance");
        // Transfer USDT to the contract
        if (uncPrice[tokenId] != 0 ){
            require(uncToken.transferFrom(msg.sender, owner, uncPrice[tokenId]), "USDT transfer failed");
            // Transfer ownership of NFT
            _transfer(owner,msg.sender, tokenId);
        }
    }

    modifier onlyManager() {
        require(
            msg.sender == owner,
            "Only owner can call this."
        );
        _;
    }

}

// @title cell library
// @author yueliyangzi
library SafeMathCell {
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }

        uint256 c = a * b;
        require(c / a == b, "SafeMath:multiplication overflow");
        return c;
    }


    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b > 0, "SafeMath:division overflow");
        uint256 c = a / b;
        return c;
    }


    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a, "SafeMath: subtraction overflow");
        uint256 c = a - b;

        return c;
    }
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a, "SafeMath: addition overflow");

        return c;
    }

    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b != 0, "SafeMath: mod overflow");
        return a % b;
    }
    // _type 1.买入 2.卖出
    function recommender_radio(uint256 _generation,uint256 _type) internal pure returns(uint256 ratio){
        if(_type == 1){
            if(_generation == 1){
                return 30;
            }
            if(_generation == 2 ){
                return 20;
            }
            if(_generation >= 3 && _generation <= 8){
                return 5;
            }

        }
        if(_type == 2){
            if(_generation == 1){
                return 20;
            }
            if(_generation == 2 ){
                return 10;
            }
        }


    }
}