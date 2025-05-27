// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

interface IPriceOracle {
    function getUsdPrice(address token, uint256 amount) external view returns(int256);
}