// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./Auction.sol";

// @title 拍卖工厂合约
contract AuctionFactory is Ownable {
    address[] public auctions;
    address public auctionImpl; 

    event AuctionCreated(address indexed auction);

    constructor(address _auctionImpl) 
        Ownable(msg.sender) 
        {
        auctionImpl = _auctionImpl;
    }

    // @dev 创建拍卖
    function createAuction(
        address nft, 
        uint256 tokenId, 
        uint256 duration, 
        address oracle) external returns(address) {
            // 判断nft 是否处于创建者
            require(IERC721(nft).ownerOf(tokenId) == msg.sender, "Not NFT Owner");

            // Auction auction = new Auction();
            // auction.initialize(nft, tokenId, duration, oracle);
            ERC1967Proxy proxy = new ERC1967Proxy(
                auctionImpl,
                abi.encodeWithSelector(
                    Auction.initialize.selector,
                    nft,
                    tokenId,
                    duration,
                    oracle,
                    msg.sender
                )
            );

            auctions.push(address(proxy));

            bytes memory bytecode = type(Auction).creationCode;
            console.log("================", bytecode.length);
            // console.logBytes(bytecode);

            emit AuctionCreated(address(proxy));

            return address(proxy);
    }

    // @dev 升级老拍卖的合约
    function upgradeAuction(address proxy, address newImplementation) external onlyOwner {
        // 判断是不是工厂创建的拍卖合约
        bool isKnown = false;
        for (uint i = 0; i < auctions.length; i ++) {
            if (auctions[i] == proxy) {
                // 是工厂创建的拍卖合约
                isKnown = true;
                break;
            }
        }
        require(isKnown, "Unknown proxy");

        // 调用upgradeTo() 升级
        (bool success, ) = proxy.call(abi.encodeWithSignature("upgradeTo(address)", newImplementation));
        require(success, "Upgrade auction failed");
    }

    // @dev 设置升级的logic合约地址，新的拍卖合约会走新的
    function setLogicAuction(address logicImplementation) external onlyOwner {
        auctionImpl = logicImplementation;
    }
}