import { describe, it } from "node:test"
import { expect } from "chai";
import hre from "hardhat";
const { ethers, networkHelpers } = await hre.network.connect();
import assert from 'assert';

describe("Hardhat Runtime Environment", function () {
  it("should have a config field", function () {
    assert.notEqual(hre.config, undefined);
  });
});


describe("Token contract", function() {
    it("Deployment should assign the total supply of tokens to the owner", async function() {
        /*
        ethers.getSigners() get a list of the Ethereum accounts.
        Account is used to send transactions to contracts and other accounts.
        */
        const[owner] = await ethers.getSigners();
        const hardhatToken = await ethers.deployContract("Token");
        const ownerBalance = await hardhatToken.balanceOf(owner.address);
        expect(await hardhatToken.totalSupply()).to.equal(ownerBalance);
    });
    it("Should transfer tokens between accounts", async function() {
        const [owner, addr1, addr2] = await ethers.getSigners();
        const hardhatToken = await ethers.deployContract("Token");
        await hardhatToken.transfer(addr1.address, 50);
        expect(await hardhatToken.balanceOf(addr1.address)).to.equal(50);
        await hardhatToken.connect(addr1).transfer(addr2.address, 50);
        expect(await hardhatToken.balanceOf(addr2.address)).to.equal(50);
    });
    /**
     * You can avoid code duplication and improve the performance of your test suite by using fixtures.
     * A fixture is a setup function that is run only the first time it's invoked.
     * On subsequent invocations, Hardhat will start the state of the network following
     * the fixture was initially executed.
     */
    async function deployTokenFixture() {      
      const [owner, addr1, addr2] = await ethers.getSigners();
      const hardhatToken = await ethers.deployContract("Token");
      await hardhatToken.waitForDeployment();
      return { hardhatToken, owner, addr1, addr2 };
    }
    describe("Deployment", function() {
      it("Should set the right owner", async function() {
        const { hardhatToken, owner } = await networkHelpers.loadFixture(deployTokenFixture);
        expect(await hardhatToken.owner()).to.equal(owner.address);
      });
      it("Should assign the total supply of tokens to the owner", async function () {
        const { hardhatToken, owner } = await networkHelpers.loadFixture(deployTokenFixture);
        const ownerBalance = await hardhatToken.balanceOf(owner.address);
        expect(await hardhatToken.totalSupply()).to.equal(ownerBalance);
      });
    });
    describe("Transactions", function () {
      it("Should transfer tokens between accounts", async function () {
        const { hardhatToken, owner, addr1, addr2 } = await networkHelpers.loadFixture(
          deployTokenFixture
        );
        await expect(
          hardhatToken.transfer(addr1.address, 50)
        ).to.changeTokenBalances(ethers, hardhatToken, [owner, addr1], [-50, 50]);
        await expect(
          hardhatToken.connect(addr1).transfer(addr2.address, 50)
        ).to.changeTokenBalances(ethers, hardhatToken, [addr1, addr2], [-50, 50])
      });
      it("Should emit Transfer events", async function() {
        const { hardhatToken, owner, addr1, addr2 } = await networkHelpers.loadFixture(
          deployTokenFixture
        );
        await expect(hardhatToken.transfer(addr1.address, 50))
          .to.emit(hardhatToken, "Transfer")
          .withArgs(owner.address, addr1.address, 50);
        await expect(hardhatToken.connect(addr1).transfer(addr2.address, 50))
          .to.emit(hardhatToken, "Transfer")
          .withArgs(addr1.address, addr2.address, 50);
      });
      it("Should fail if sender doesn't have enought tokens", async function() {
        const { hardhatToken, owner, addr1 } = await networkHelpers.loadFixture(
          deployTokenFixture
        );
        const initialOwnerBalance = await hardhatToken.balanceOf(owner.address);
        await expect(
          hardhatToken.connect(addr1).transfer(owner.address, 1)
        ).to.be.revertedWith("Not enough tokens");
        expect(await hardhatToken.balanceOf(owner.address)).to.equal(
          initialOwnerBalance
        );
      });
    });
})