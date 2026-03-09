In Bitcoin, payments are realized in completely different way. There are:
1. No accounts
2. No balances
3. No addresses
4. No coin
5. No senders and receivers

Since blockchain is a public and open database, so

1, 2, 3> Not store sensitive information about wallet owners

4> Coins are not collected in accounts

5> Not transfer money from one address to another

What's inside a transaction?
```golang
type Transaction struct {
    ID []byte
    Vin []TXInput
    Vout []TXOutput
}
```
- Vin: output of a or multiple previous transaction
- Vout: where coins are actually stored

![Transaction Diagram](https://jeiwan.net/images/transactions-diagram.png)

```golang
type TXOutput struct {
    Value int
    ScriptPubKey string
}
```
the Value field store "coins". Bitcoin use __ScriptPubKey to define outputs locking and unlocking logic, storing means locking them with a puzzle__. This is made intentionally, to avoid possible hacks and misuses. By the way, having such scripting language means that Bitcoin can be used as a smart-contract platform as well

In Bitcoin, __the value field stores the number of satoshis, not the number of BTC__. A satoshi is a hundred millionth of a bitcoin

One important thing about outputs is that they are invisible, which means that you cannot reference a part of its value. __It is spent as a whole__. And if its value is greater than required, a change is generated and sent back to the sender.

From now on, every block must store at least one transaction and it's no more possible to mine blocks without transactions. This means that __we should remove the Data field of Block and store transactions instead__

#### The egg
When a miner starts mining a block, it adds a coinbase transaction to it. A coinbase transaction is a special type of transactions, which doesn't require previously existing outputs. It creates outputs out of nowhere. The egg without a chicken.

#### Hashing
We're using hashing as a mechanism of providing unique representation of data. We want __all transactions in a block to be uniquely identified__ by a single hash. To achieve this, we get hashes of each transaction, concatenate them, and get a hash of the concatenated combination.

Bitcoin uses a more elaborate technique: it represents __all transactions containing in a block as a Merkle tree__ and uses the __root hash of the tree in the Proof-of-Work__ system. This approach allows to quickly check if a block contains certain transaction, having only just the root hash and without downloading all the transactions.

#### Unspent Transaction Ouputs
Unspent Transaction outputs (UTXO). Unspent means that these outputs weren't reference in any inputs

#### Definition
```golang
//Vin: output of a or multiple previous transaction
//Vin holds the spent Vout index of a TXOutput list
//Vout: where coins are actually stored in a TXOutput list
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}
// Vout reference a specific index within a transaction's TXOutput list
// When a TXOutput is referenced in a TXInput, it's spent as a whole
type TXInput struct {
	Txid      []byte
	Vout      int *****
	ScriptSig string # who spent Vout
}

// TXOutput store "coins" (notice the Value field)
type TXOutput struct {
	Value        int *****
	ScriptPubKey string # the owner of Value
}
```

```golang
// NewCoinbaseTX creates a first coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	/* A coinbase transaction has only one input.
	   Txid is empty and Vout equal to -1.
	   Also, a coinbase transaction doesn't store a script in ScriptSig*/
	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to} //subsidy is the amount of reward, we'll store the reward as a constant
	/* In Bitcoin, this number is not stored anywhere and
	 calculated based only the total number of blocks */
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	return &tx
}
```

### Limitation of this approach
1. Getting balance requires scanning the whole blockchain, which can take very long time when there are many and many blocks. Also, it can take a lot of time if we want to validate later transactions.
2. Mempool: This is where transactions are stored before being packed in a block. In our current implementation, a block contains only one transaction, and this is quire inefficient