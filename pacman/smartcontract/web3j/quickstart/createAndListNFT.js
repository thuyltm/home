import { Web3 } from 'web3';
const web3 = new Web3('http://localhost:7545');
const networkId = await web3.eth.net.getId();
const account = web3.eth.accounts.wallet.add("0xac7aef3f0e3b34d0051fd9f3d71d5a02434d3d6c1ac4d3ae931cdc5984df41a1");

import Marketplace from './contracts/Marketplace.json' with { type: 'json' };
import BoredPetsNFT from './contracts/BoredPetsNFT.json'  with { type: 'json' };

const boredPetsContractAddress = BoredPetsNFT.networks[networkId].address
const boredPetsContract = new web3.eth.Contract(BoredPetsNFT.abi, boredPetsContractAddress)

const marketPlaceContract = new web3.eth.Contract(Marketplace.abi, Marketplace.networks[networkId].address)
    
let listingFee = await marketPlaceContract.methods.LISTING_FEE().call()
listingFee = listingFee.toString()

let txn1 = boredPetsContract.methods.mint("URI1");
txn1.send({ from: account[0].address }).on('receipt', function (receipt) {
    console.log('minted');
    const tokenId = receipt.events.NFTMinted.returnValues[0];
    console.log(tokenId);
    marketPlaceContract.methods.listNft(boredPetsContractAddress, tokenId, Web3.utils.toWei(1, "ether"))
    .send({ from: account[0].address, value: listingFee }).on('receipt', function (receipt) {
            console.log('listed')
            console.log(receipt.events.NFTListed.returnValues[0]);
        });
})