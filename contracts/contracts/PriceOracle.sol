// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import "hardhat/console.sol";

// @title ChainLink 价格预言机
contract PriceOracle {
    AggregatorV3Interface public ethUsdFeed;
    mapping(address=>AggregatorV3Interface) public tokenFeeds;

    constructor(address _ethUsedFeed) {
        ethUsdFeed = AggregatorV3Interface(_ethUsedFeed);
    }

    function addTokenFeed(address token, address feed) external {
        tokenFeeds[token] = AggregatorV3Interface(feed);
    }

    function getUsdPrice(address token, uint256 amount) public view returns(int256) {
        console.log("----enter in:", 2000);
        return 2000;
        // if (token == address(0)) {
        //     (
        //         /** uint80 roundId */,
        //         int256 price,
        //         /** uint256 startAt */,
        //         /** uint256 updatedAt */,
        //         /** uint80 answeredInRound */
        //     ) = ethUsdFeed.latestRoundData();
        //     return int256(amount) * price / 1e8;
        // } else {
        //     (
        //         /** uint80 roundId */,
        //         int256 tokenPrice,
        //         /** uint256 startAt */,
        //         /** uint256 updatedAt */,
        //         /** uint80 answeredInRound */
        //     ) = tokenFeeds[token].latestRoundData();
        //     return int256(amount) * tokenPrice / 1e8;
        // }
    }
}