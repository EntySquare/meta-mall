pragma solidity  ^0.8.24;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC721/ERC721.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/IERC20.sol";

// @author yueliyangzi
contract METANFT is ERC721{
    using SafeMathCell for uint256;
    IERC20 public usdtToken;
    IERC20 public uncToken;
    address public owner;

    constructor(address _usdtTokenAddress,address _uncTokenAddress,address holder) ERC721("NFTContract", "NFTC") {
        usdtToken = IERC20(_usdtTokenAddress);
        uncToken = IERC20(_uncTokenAddress);
        owner = holder;
    }

    function mintNFT(address recipient, uint256 tokenId) external {
        require(msg.sender == owner, "Only owner can mint NFTs");
        _safeMint(recipient, tokenId);
    }
    function mintNFTs(address[] memory recipients, uint256[] memory tokenIds) external {
        require(msg.sender == owner, "Only owner can mint NFTs");
        require(recipients.length == tokenIds.length, "Array lengths mismatch");

        for (uint256 i = 0; i < recipients.length; i++) {
            _safeMint(recipients[i], tokenIds[i]);
        }
    }
    function transferOwnership(address newOwner) external {
        require(msg.sender == owner, "Only owner can transfer ownership");
        owner = newOwner;
    }

    function triggerOwnershipTransferUsdt(uint256 tokenId, uint256 amount) external {
        require(ownerOf(tokenId) == msg.sender, "You are not the owner of this NFT");
        require(usdtToken.allowance(msg.sender, address(this)) >= amount, "Insufficient allowance");
        require(usdtToken.balanceOf(msg.sender) >= amount, "Insufficient balance");

        // Transfer USDT to the contract
        require(usdtToken.transferFrom(msg.sender, address(this), amount), "USDT transfer failed");

        // Transfer ownership of NFT
        _transfer(msg.sender, owner, tokenId);
    }
    function triggerOwnershipTransferUnc(uint256 tokenId, uint256 amount) external {
        require(ownerOf(tokenId) == msg.sender, "You are not the owner of this NFT");
        require(uncToken.allowance(msg.sender, address(this)) >= amount, "Insufficient allowance");
        require(uncToken.balanceOf(msg.sender) >= amount, "Insufficient balance");

        // Transfer USDT to the contract
        require(uncToken.transferFrom(msg.sender, address(this), amount), "USDT transfer failed");

        // Transfer ownership of NFT
        _transfer(msg.sender, owner, tokenId);
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