pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract METATRADE is Ownable {
    IERC20 public usdtToken;
    IERC20 public uncToken;
    ERC721 public nftToken;
    address public manager;
    event NFTTransfer(address indexed owner, uint256 tokenId);

    constructor(address _usdtToken,address _uncToken, address _nftToken,address _owner) Ownable(_owner){
        usdtToken = IERC20(_usdtToken);
        uncToken = IERC20(_uncToken);
        nftToken = ERC721(_nftToken);
    }

    function transferUSDTToNFT(uint256 _usdtAmount,uint256 tokenId) external {
        // Make sure caller has enough USDT allowance
        require(usdtToken.allowance(msg.sender, address(this)) >= _usdtAmount, "Allowance not enough");

        // Transfer USDT from sender to this contract
        require(usdtToken.transferFrom(msg.sender, address(this), _usdtAmount), "USDT transfer failed");

        // Mint new NFT to sender
        nftToken.safeTransferFrom(address(this),msg.sender, tokenId);

        emit NFTTransfer(msg.sender, tokenId);
    }
    function transferUNCToNFT(uint256 _uncAmount,uint256 tokenId) external {
        // Make sure caller has enough USDT allowance
        require(uncToken.allowance(msg.sender, address(this)) >= _uncAmount, "Allowance not enough");

        // Transfer USDT from sender to this contract
        require(uncToken.transferFrom(msg.sender, address(this), _uncAmount), "USDT transfer failed");

        // Mint new NFT to sender
        nftToken.safeTransferFrom(address(this),msg.sender, tokenId);

        emit NFTTransfer(msg.sender, tokenId);
    }

    // Owner function to withdraw USDT from contract
    function withdrawUNC(uint256 _amount) external onlyOwner {
        require(usdtToken.transfer(owner(), _amount), "USDT transfer failed");
    }
    function withdrawUSDT(uint256 _amount) external onlyOwner {
        require(uncToken.transfer(owner(), _amount), "USDT transfer failed");
    }


    // Owner function to withdraw NFTs from contract
    function withdrawNFT(uint256 _tokenId) external onlyOwner {
        nftToken.transferFrom(address(this), owner(), _tokenId);
    }
}
