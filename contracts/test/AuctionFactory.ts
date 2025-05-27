import {
  time,
  loadFixture,
} from "@nomicfoundation/hardhat-toolbox/network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import hre from "hardhat";

describe("Test AuctionFactory", () => {
    async function deployAuctionFactory() {
        console.log("hre ethers version:", hre.ethers.version)
        // Contracts are deployed using the first signer/account by default
        const [owner, otherAccount] = await hre.ethers.getSigners();

        const Auction = await hre.ethers.getContractFactory("Auction")
        const auction = await Auction.connect(owner).deploy()
        await auction.waitForDeployment()
        const AuctionFactory = await hre.ethers.getContractFactory("AuctionFactory");
        const auctionFactory = await AuctionFactory.connect(owner).deploy(auction.getAddress())
        await auctionFactory.waitForDeployment()
        const NFT = await hre.ethers.getContractFactory("NFT")
        const nft = await NFT.connect(owner).deploy()
        await nft.waitForDeployment()
        const PriceOracle = await hre.ethers.getContractFactory("PriceOracle")
        const ethUsdAddr = "0x694AA1769357215DE4FAC081bf1f309aDC325306"
        const priceOracle = await PriceOracle.connect(owner).deploy(ethUsdAddr)
        await priceOracle.waitForDeployment()

        console.log("nft address:", await nft.getAddress())
        console.log("auction address:", await auction.getAddress())
        console.log("auctionFactory address:", await auctionFactory.getAddress())
        console.log("oracle address:", await priceOracle.getAddress())

        return { auction, auctionFactory, nft, priceOracle };
    }

    // 定义 sleep 函数
    const sleep = (ms: any) => new Promise(resolve => setTimeout(resolve, ms));
    it("Test CreateAuction", async () => {
        const { auction, auctionFactory, nft, priceOracle } = await loadFixture(deployAuctionFactory)
        const signers = await hre.ethers.getSigners();
        for (let i = 0; i < signers.length; i++) {
            console.log(`Signer ${i}: ${await signers[i].getAddress()}`);
        }
        // 铸造nft
        const nftTX = await nft.mint(signers[0], "http//loaclhost.com/token/1")
        const nftReceipt = await nftTX.wait()
        // console.log("nftTX:", nftTX)
        // console.log("nftReceipt:", nftReceipt)
        // 如果合约支持返回值，你可以使用 callStatic 方式获取返回值：
        // 查找特定事件
        if (nftReceipt != null) {
            for (const log of nftReceipt.logs) {
                const parsed = nft.interface.parseLog(log)
                // console.log("parsed:", parsed)
                if (parsed?.name == "MintNFT") {
                    const [to, uri, tokenId] = parsed.args
                    console.log("tokenId:", tokenId)
                    console.log("nft owner:", await nft.ownerOf(tokenId));
                    
                    const duration = 5n
                    const tx = await auctionFactory.connect(signers[0]).createAuction(nft.getAddress(), tokenId, duration, priceOracle.getAddress())
                    const receipt = await tx.wait()
                    // console.log("tx:", tx)

                    for (const alog of receipt!.logs) {
                        const aparsed = auctionFactory.interface.parseLog(alog)
                        if (aparsed?.name == "AuctionCreated") {
                            const [proxy] = aparsed.args
                            console.log("proxy address:", proxy)
                            console.log("AuctionCreated aparsed:", proxy)
                            // 异步的，注意await
                            // 否则还没有approve 成功调用depositNFT会失败
                            // await nft.connect(signers[0]).setApprovalForAll(proxy, true)
                            await nft.connect(signers[0]).approve(proxy, tokenId)
                            // 判断是否授权成功
                            const approvAddress = await nft.connect(signers[0]).getApproved(tokenId)
                            console.log("NFT approve success:", approvAddress)
                            // 将代理合约与auction绑定
                            const auctionProxy = auction.attach(proxy)
                            // 将owner nft 转移到合约
                            await auctionProxy.depositNFT()

                            console.log("B 玩家拍卖报价")
                            // B 玩家拍卖报价
                            const ethAmount = hre.ethers.parseEther("1");
                            const bidTx =await auctionProxy.connect(signers[1]).placeBid(hre.ethers.ZeroAddress, ethAmount, 
                                {value: ethAmount}
                            )
                            const bidReceipt = await bidTx.wait()
                            // C 玩家拍卖报价
                            const ethAmountC = hre.ethers.parseEther("2")
                            const bidTxC = await auctionProxy
                                .connect(signers[2])
                                .placeBid(hre.ethers.ZeroAddress, ethAmountC, 
                                    {value: ethAmountC}
                                )
                            const bidReceiptC = await bidTxC.wait()
                            console.log("sleep before:", Date.now())
                            await sleep(6000); // 异步等待
                            console.log("sleep after:", Date.now())

                            // C 结算
                            await auctionProxy.connect(signers[2]).settle()
                        }
                    }
                    // auctionFactory.interface.parseLog


                    // auctionFactory.connect(signers[0])
                    // 授权合约转移tokenId
                    // nft.connect(signers[0]).setApprovalForAll(auction, true)
                    // nft.connect(signers[0]).approve(auctionFactory!, tokenId!)
                }
            }
        }
    })
})