import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

// npx hardhat ignition deploy ./ignition/modules/NFT.ts
const NftModule = buildModule("NftModuleV1", (m) => {
    const nftContract = m.contract("NFT", [])

    return { nftContract }
});

export default NftModule;