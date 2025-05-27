// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces/IPriceOracle.sol";
import "./library/SafeConvert.sol";
import "hardhat/console.sol";

contract Auction is Initializable, UUPSUpgradeable, OwnableUpgradeable {
    using SafeConvert for int256;
    struct Bid {
        address bidder; // 出价者
        uint256 amount; // 数量
        address currency; // 货币
        uint256 usdValue; // usd 值
    }
    
    event NFTDeposit(address indexed account, uint256 tokenId);
    event PlaceBit(address indexed account, address currency, uint256 amount, int256 price);
    event Settle(address indexed account, uint256 amount);
    
    address nft;
    uint256 tokenId;
    address seller;
    uint256 deadline;

    /**
     * @dev 以下两种赋值无效，只在constructor 示例话会生效
     * 但是此处使用代理必须放到initialize中
     */
    // bool notSettled = true; // 优化gas
    // bool notDeposit = true;
    bool notSettled;
    bool notDeposit;

    Bid public highestBid; // 最高出价者

    IPriceOracle oracle;

    function initialize(
        address _nft,
        uint256 _tokenId,
        uint256 _duration, 
        address _oracle,
        address account) external initializer {
            __Ownable_init();

            nft = _nft;
            tokenId = _tokenId;
            seller = account;
            deadline = block.timestamp + _duration;

            notDeposit = true; // 优化gas 默认为true
            notSettled = true;
            console.log("seller:", seller);

            oracle = IPriceOracle(_oracle);
    }

    function depositNFT() external {
        require(msg.sender == seller, "Only seller can deposit");
        console.log("notSettled:", notSettled);
        console.log("notDeposit:", notDeposit);
        console.log("seller:", msg.sender, address(this));
        require(notDeposit, "NFT already deposit");

        IERC721(nft).transferFrom(msg.sender, address(this), tokenId);

        notDeposit = true;
        emit NFTDeposit(msg.sender, tokenId);
    }

    // @dev 拍卖出价
    // @param token 代币token or eth
    // @param amount 出价
    function placeBid(address token, uint256 amount) external payable {
        // 判断是否在有效时间
        require(block.timestamp <= deadline, "Time is not reached!");

        if (token != address(0)) {
            require(IERC20(token).balanceOf(msg.sender) >= amount, "Balance is not enough");
            // 代币支付
            // 将eth 转移到 合约
            IERC20(token).transferFrom(msg.sender, address(this), amount);
        } else {
            // eth 支付
            require(msg.value == amount, "Eth amount not match");
        }
        // 确保oracle 不是零地址
        int256 tokenUsdPrice = oracle.getUsdPrice(token, amount);
        console.log("tokenUsdPrice:", uint256(tokenUsdPrice), uint256(tokenUsdPrice) * amount);
        console.log("highestBid.bidder:", highestBid.bidder);
        if (highestBid.bidder == address(0)) {
            // // 方式一
            // highestBid = Bid({bidder: msg.sender, amount: amount, currency: token, usdValue: 111});
            // 方式二
            highestBid = Bid(msg.sender, amount, token, tokenUsdPrice.toUint256() * amount);
        } else {
            if (highestBid.usdValue < tokenUsdPrice.toUint256() * amount) {
                highestBid.bidder = msg.sender;
                highestBid.currency = token;
                highestBid.usdValue = tokenUsdPrice.toUint256() * amount;
            }
        }

        emit PlaceBit(msg.sender, token, amount, tokenUsdPrice);
    }

    // @dev 拍卖结束
    function settle() external {
        require(block.timestamp >= deadline, "Time is not reached");
        require(notSettled, "Auction has settled");

        console.log("highestBid.bidder:", highestBid.bidder);

        if (highestBid.bidder != address(0)) {
            require(highestBid.bidder == msg.sender, "You are not highest bidder");
            // 将nft转送
            IERC721(nft).transferFrom(address(this), highestBid.bidder, tokenId);
            if (highestBid.currency == address(0)) {
                // 使用的eth支付
                payable(highestBid.bidder).transfer(highestBid.amount);
            } else {
                // 将代币token转给seller
                IERC20(highestBid.currency).transfer(seller, highestBid.amount);
            }
            emit Settle(msg.sender, highestBid.amount);
        } else {
            // 没有人出价 将nft 返回
            IERC721(nft).transferFrom(address(this), msg.sender, tokenId);

            emit Settle(msg.sender, 0);
        }
        notSettled = false;
    }

    function _authorizeUpgrade(address newImpletion) internal override onlyOwner{

    }
}