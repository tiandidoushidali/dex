import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

// npx hardhat ignition deploy ./ignition/modules/AuctionFactory.ts
const AuctionFactoryModule = buildModule("AuctionFactoryModuleV1", (m) => {
    const auctionAddr = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
    const auctionFactory = m.contract("AuctionFactory", [auctionAddr]);

    // m.encodeFunctionCall(auctionFactory, "initialize", [])

    return { auctionFactory };
});

export default AuctionFactoryModule;

// 0x5FbDB2315678afecb367f032d93F642f64180aa3
