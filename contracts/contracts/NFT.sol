// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract NFT is ERC721URIStorage, Ownable {
    event MintNFT(address indexed account, string uri, uint256 tokenId);

    uint256 public tokenIdCounter;

    constructor() 
        Ownable(msg.sender)
        ERC721("AuctionNFT", "AUCNFT") {
    }

    function mint(address to, string memory uri) public onlyOwner returns(uint256) {
        uint256 newId = tokenIdCounter ++;
        _safeMint(to, newId);
        _setTokenURI(newId, uri);

        emit MintNFT(to, uri, newId);
        
        return newId;
    }
}

// 0x5FbDB2315678afecb367f032d93F642f64180aa3