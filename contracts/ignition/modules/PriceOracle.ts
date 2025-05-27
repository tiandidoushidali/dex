import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

// npx hardhat ignition deploy ./ignition/modules/PriceOracle.ts
const PriceOracle = buildModule("PriceOracleV1", (m) => {
    // ETH/USD 价格地址
    const ethUsdAddr = "0x694AA1769357215DE4FAC081bf1f309aDC325306"
    const priceOracle = m.contract("PriceOracle", [ethUsdAddr]);
    return { priceOracle }
})

export default PriceOracle;

// 0x5FbDB2315678afecb367f032d93F642f64180aa3