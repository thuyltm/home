import { Web3 } from 'web3';
const web3 = new Web3('http://localhost:7545');

const networkId = await web3.eth.net.getId();

import Marketplace from './contracts/Marketplace.json' with { type: 'json' };

const marketPlaceContract = new web3.eth.Contract(Marketplace.abi, Marketplace.networks[networkId].address)

let listedNfts = await marketPlaceContract.methods.getListedNfts().call();
listedNfts.map(i => {
    const nft = {
        price: i.price,
        tokenId: i.tokenId,
        seller: i.seller,
        owner: i.buyer
    }
    if (nft.tokenId > 0) {
        console.log(nft)
    }    
})