// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

library SafeConvert {
    function toUint256(int256 value) internal pure returns(uint256) {
        require(value >= 0, "Cannot convert negative int256 to uint256");

        return uint256(value);
    }

    function toUint256(int112 value) internal pure returns(uint256) {
        require(value >= 0, "Cannot convert negative int256 to uint256");

        // 有符号到无符号的扩大转换，在编译器看来可能是 危险操作。 不能直接 int112 -> uint256 
        // 哪怕是uint256 包含 uint112
        // 必须先进行类型提升 int112 -> int256 
        // 再进行类型转换 int256 -> uint256
        return uint256(int256(value));
    }

    /**
     * @dev 两种调用方式
     * @param x 加数
     * @param y 被加数
     * x.add(y)
     * SafeConvert.add(x, y)
     */
    function add(int256 x, int256 y) public pure returns(int256) {
        return x + y;
    }
}