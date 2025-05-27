import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

// npx hardhat ignition deploy ./ignition/modules/Auction.ts
const AuctionModule = buildModule("AuctionV1", (m) => {
    const auction = m.contract("Auction")
    // const nftAddr = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
    // const oracleAddr = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
    // const time = Date.now() / 1000;
    // const data = m.call(auction, "initialize", [nftAddr, 1, time, oracleAddr])

    return { auction };
});

export default AuctionModule;

// 0x5FbDB2315678afecb367f032d93F642f64180aa3