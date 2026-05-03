```sh
truffle test test/marketplace.js --config truffle-config.ovm.js --network ganache
#Contract: Marketplace
#marketplace 0x9B988D3Fb9A2B8E81581Fb1f41E1843538781E43
#token_owner 0x9653e0E3bB2942408B0d49B14dCa8Ab209182888
#buyer 0x7De63d38Ad58fE9953124E78C1DE7c7178A7bE11
#   ✔ should validate before listing (145ms)
#   1) should list nft

#   Events emitted during test:
#   ---------------------------

#    BoredPetsNFT.Transfer(
#      from: <indexed> 0x0000000000000000000000000000000000000000 of unknown class (type: address),
#      to: <indexed> 0x9653e0E3bB2942408B0d49B14dCa8Ab209182888 of unknown class (type: address),
#      tokenId: <indexed> 1 (type: uint256)
#    )

#    BoredPetsNFT.MetadataUpdate(
#      _tokenId: 1 (type: uint256)
#    )

#    BoredPetsNFT.ApprovalForAll(
#      owner: <indexed> 0x9653e0E3bB2942408B0d49B14dCa8Ab209182888 of unknown class (type: address),
#      operator: <indexed> 0x9B988D3Fb9A2B8E81581Fb1f41E1843538781E43 (Marketplace) (type: address),
#      approved: true (type: bool)
#    )

#    [object Object].NFTMinted(
#      1 (type: uint256)
#    )

#    BoredPetsNFT.Transfer(
#      from: <indexed> 0x9653e0E3bB2942408B0d49B14dCa8Ab209182888 of unknown class (type: address),
#      to: <indexed> 0x9B988D3Fb9A2B8E81581Fb1f41E1843538781E43 (Marketplace) (type: address),
#      tokenId: <indexed> 1 (type: uint256)
#    )

#    [object Object].NFTListed(
#      nftContract: 0x2dA783C0C473B67Dc06EbBdB038Bc2a8665E5502 (BoredPetsNFT) (type: address),
#      tokenId: 1 (type: uint256),
#      seller: 0x9653e0E3bB2942408B0d49B14dCa8Ab209182888 of unknown class (type: address),
#      owner: 0x9B988D3Fb9A2B8E81581Fb1f41E1843538781E43 (Marketplace) (type: address),
#      price: 5000000000000000 (type: uint256)
#   )


#    ---------------------------
#    ✔ should validate before buying
#    ✔ should modify listings when nft is bought (624ms)
#    2) should validate reselling
#    > No events were emitted
#    3) should resell nft

#    Events emitted during test:
#    ---------------------------

#    BoredPetsNFT.Transfer(
#      from: <indexed> 0x0000000000000000000000000000000000000000 of unknown class (type: address),
#      to: <indexed> 0x9653e0E3bB2942408B0d49B14dCa8Ab209182888 of unknown class (type: address),
#      tokenId: <indexed> 3 (type: uint256)
#    )

#    BoredPetsNFT.MetadataUpdate(
#      _tokenId: 3 (type: uint256)
#    )

#    BoredPetsNFT.ApprovalForAll(
#      owner: <indexed> 0x9653e0E3bB2942408B0d49B14dCa8Ab209182888 of unknown class (type: address),
#      operator: <indexed> 0x9B988D3Fb9A2B8E81581Fb1f41E1843538781E43 (Marketplace) (type: address),
#      approved: true (type: bool)
#    )

#    [object Object].NFTMinted(
#      3 (type: uint256)
#    )

#    BoredPetsNFT.Transfer(
#      from: <indexed> 0x9653e0E3bB2942408B0d49B14dCa8Ab209182888 of unknown class (type: address),
#      to: <indexed> 0x9B988D3Fb9A2B8E81581Fb1f41E1843538781E43 (Marketplace) (type: address),
#      tokenId: <indexed> 3 (type: uint256)
#    )

#    [object Object].NFTListed(
#      nftContract: 0x2dA783C0C473B67Dc06EbBdB038Bc2a8665E5502 (BoredPetsNFT) (type: address),
#      tokenId: 3 (type: uint256),
#      seller: 0x9653e0E3bB2942408B0d49B14dCa8Ab209182888 of unknown class (type: address),
#      owner: 0x9B988D3Fb9A2B8E81581Fb1f41E1843538781E43 (Marketplace) (type: address),
#      price: 5000000000000000 (type: uint256)
#    )

#    BoredPetsNFT.Transfer(
#      from: <indexed> 0x9B988D3Fb9A2B8E81581Fb1f41E1843538781E43 (Marketplace) (type: address),
#      to: <indexed> 0x7De63d38Ad58fE9953124E78C1DE7c7178A7bE11 of unknown class (type: address),
#      tokenId: <indexed> 3 (type: uint256)
#    )

#    [object Object].NFTSold(
#      nftContract: 0x2dA783C0C473B67Dc06EbBdB038Bc2a8665E5502 (BoredPetsNFT) (type: address),
#      tokenId: 3 (type: uint256),
#      seller: 0x9653e0E3bB2942408B0d49B14dCa8Ab209182888 of unknown class (type: address),
#      owner: 0x7De63d38Ad58fE9953124E78C1DE7c7178A7bE11 of unknown class (type: address),
#      price: 5000000000000000 (type: uint256)
#    )

#    BoredPetsNFT.Approval(
#      owner: <indexed> 0x7De63d38Ad58fE9953124E78C1DE7c7178A7bE11 of unknown class (type: address),
#      approved: <indexed> 0x9B988D3Fb9A2B8E81581Fb1f41E1843538781E43 (Marketplace) (type: address),
#      tokenId: <indexed> 3 (type: uint256)
#    )

#    BoredPetsNFT.Transfer(
#      from: <indexed> 0x7De63d38Ad58fE9953124E78C1DE7c7178A7bE11 of unknown class (type: address),
#      to: <indexed> 0x9B988D3Fb9A2B8E81581Fb1f41E1843538781E43 (Marketplace) (type: address),
#      tokenId: <indexed> 3 (type: uint256)
#    )

#    [object Object].NFTListed(
#      nftContract: 0x2dA783C0C473B67Dc06EbBdB038Bc2a8665E5502 (BoredPetsNFT) (type: address),
#      tokenId: 3 (type: uint256),
#      seller: 0x7De63d38Ad58fE9953124E78C1DE7c7178A7bE11 of unknown class (type: address),
#      owner: 0x9B988D3Fb9A2B8E81581Fb1f41E1843538781E43 (Marketplace) (type: address),
#      price: 5000000000000000 (type: uint256)
#    )


#    ---------------------------


#  3 passing (2s)
#  3 failing

#  1) Contract: Marketplace
#       should list nft:
#     AssertionError: Listing fee not transferred: expected <BN: 0> to equal '100000000000000'
#      at Context.<anonymous> (test/marketplace.js:81:12)
#      at processTicksAndRejections (node:internal/process/task_queues:104:5)

#  2) Contract: Marketplace
#       should validate reselling:

#      Wrong kind of exception received
#      + expected - actual

#      -VM Exception while processing transaction: revert ERC721: transfer from incorrect owner -- Reason given: ERC721: transfer from incorrect owner.
#      +Price must be at least 1 wei
      
#      at expectException (node_modules/@openzeppelin/test-helpers/src/expectRevert.js:20:30)
#      at expectRevert (node_modules/@openzeppelin/test-helpers/src/expectRevert.js:75:3)
#      at Context.<anonymous> (test/marketplace.js:123:5)

#  3) Contract: Marketplace
#       should resell nft:
#     AssertionError: NFT contract is not correct: expected undefined to equal '0x2dA783C0C473B67Dc06EbBdB038Bc2a8665E5502'
#      at assertListing (test/marketplace.js:13:10)
#      at Context.<anonymous> (test/marketplace.js:150:5)
#      at processTicksAndRejections (node:internal/process/task_queues:104:5)
```