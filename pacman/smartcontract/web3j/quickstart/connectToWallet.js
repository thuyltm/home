import { Web3 } from 'web3';
const web3 = new Web3('http://localhost:7545');
const privateKey = "0x"+"ac7aef3f0e3b34d0051fd9f3d71d5a02434d3d6c1ac4d3ae931cdc5984df41a1";
const account = web3.eth.accounts.wallet.add(privateKey);

console.log("Wallet Address %s", account[0].address);

// create transaction object to send 1 eth to '0xa32...c94' address from the account[0]
const tx = {
	from: account[0].address,
	to: '0xa3286628134bad128faeef82f44e99aa64085c94',
	value: web3.utils.toWei('1', 'ether'),
};
// the "from" address must match the one previously added with wallet.add

// send the transaction
const txReceipt = await web3.eth.sendTransaction(tx);

console.log('Tx hash:', txReceipt.transactionHash);