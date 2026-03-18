```sh
% bazel run //pacman/address:main -- createwallet
# Your new address: 16e98KL3MbKkod5puogH3zUw6jDPgcJW9n
% bazel run //pacman/address:main -- createwallet
# Your new address: 1MMukuFU3uUM7KU92cUEXsNym6yDLhrpUZ
% bazel run //pacman/address:main -- createwallet
# Your new address: 1GeHSmPnbXf76hKUH5Q3qEaYQgSh83M9ab
% bazel run //pacman/address:main -- createblockchain -address 16e98KL3MbKkod5puogH3zUw6jDPgcJW9n
#0000009cd1d2494b13cc03d50a4ff28493a426c6be93c98524884f120eb7433f

#Done!
% bazel run //pacman/address:main -- getbalance -address 16e98KL3MbKkod5puogH3zUw6jDPgcJW9n
#Balance of '16e98KL3MbKkod5puogH3zUw6jDPgcJW9n': 10
% bazel run //pacman/address:main -- send -from 16e98KL3MbKkod5puogH3zUw6jDPgcJW9n -to 1MMukuFU3uUM7KU92cUEXsNym6yDLhrpUZ -amount 5
#00000099dd6d1ca2a1e797cb7dfed2eddab85b50a166b826fbab4d510a826d2e

#Success!
% bazel run //pacman/address:main -- printchain
# ============ Block 00000099dd6d1ca2a1e797cb7dfed2eddab85b50a166b826fbab4d510a826d2e ============
# Prev. block: 0000009cd1d2494b13cc03d50a4ff28493a426c6be93c98524884f120eb7433f
# PoW: true
# 
# --- Transaction 9afe4b6ff25148af5169bd5e0a9b4d596873e60104dc796dfc456422818a56e4:
#   Input 0:
#     TXID: 88e5eb053b4ab251bbb969ba85cdfbe2a5b14929516834afc8cedb4a22bd1785 ---- TRANSACTION ID OF BLOCK 0
#     Out:       0
#     Signature: 7344c4c1b57d177661e01d726b6951aab54c9a81f6b2ddfa1ff516f068d76a118bb6da880d34a249778af8a39141660eef7b64de1d60782831dd39c9cc4bce7b ---Sign with Private Key of Wallet 1
#     PubKey:    d930c603c372fa2d59408e41019195c92a3a36cb6e803e562776804d6aaab1a8d820bbc4514a1cf93f18bd281eb92c4b3a7a4ccdc80ed60f4caffe5da97f11f4 --- Public Key of Wallet 1
#   Output 0:
#     Value:  5
#     PubKeyHash: df56dfc9262cba6e5816250070b21777f0d0dbd0 --- Base58Decode of Address 2 or Hash of publicKey of Wall 1
#   Output 1:
#     Value:  5
#     PubKeyHash: 3ddecd8876590d12ef31dac44799ee1c540d30e2 --- Base58Decode of Address 1 or Hash of publicKey of Wall 1


# ============ Block 0000009cd1d2494b13cc03d50a4ff28493a426c6be93c98524884f120eb7433f ============
# Prev. block: 
# PoW: true

# --- Transaction 88e5eb053b4ab251bbb969ba85cdfbe2a5b14929516834afc8cedb4a22bd1785: -- TRANSACTION ID OF BLOCK 0
#   Input 0:
#     TXID: 
#     Out:       -1
#     Signature: 
#     PubKey:    5468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73
#   Output 0:
#     Value:  10
#     PubKeyHash: 3ddecd8876590d12ef31dac44799ee1c540d30e2 --- Base58Decode of Address 1
% bazel run //pacman/address:main -- send -from 1MMukuFU3uUM7KU92cUEXsNym6yDLhrpUZ -to 1GeHSmPnbXf76hKUH5Q3qEaYQgSh83M9ab -amount 2
# 0000005cb1689c5c46d86ca44c28002517452d815b2bf2bd351cfe05c4574337

# Success!

% bazel run //pacman/address:main -- printchain
# ============ Block 0000005cb1689c5c46d86ca44c28002517452d815b2bf2bd351cfe05c4574337 ============
# Prev. block: 00000099dd6d1ca2a1e797cb7dfed2eddab85b50a166b826fbab4d510a826d2e
# PoW: true
#
#--- Transaction 462fcd0e9a04efca131ff50f996d27afaf2f4ea6950686f6975c01cab153e91f:
#  Input 0:
#    TXID: 9afe4b6ff25148af5169bd5e0a9b4d596873e60104dc796dfc456422818a56e4
#    Out:       0
#    Signature: c2e76f3c66c49cda8b95e5931605e03781e3dfead428d741e9e35d7ad3db5bc1a292fdcaeac6f29f4265560d02feed7e5bfb7e4f8de89aab86ca3d0d8a22623e
#    PubKey:    0f4bdec85b99ae2195eee7385459158fd9be1774cacb5b89b5008629484b1d2fd1da03c9e5403f3821466972cea748152dc4943678f506a32794e28fd18aef5a --- Public Key of Wallet 2
#  Output 0:
#    Value:  2
#    PubKeyHash: ab97011983a7e4c5f9c981267e642b23fb7a0e38 --- Base58Decode of Address 3
#  Output 1:
#    Value:  3
#    PubKeyHash: df56dfc9262cba6e5816250070b21777f0d0dbd0 --- Base58Decode of Address 2


#============ Block 00000099dd6d1ca2a1e797cb7dfed2eddab85b50a166b826fbab4d510a826d2e ============
# Prev. block: 0000009cd1d2494b13cc03d50a4ff28493a426c6be93c98524884f120eb7433f
# PoW: true

# --- Transaction 9afe4b6ff25148af5169bd5e0a9b4d596873e60104dc796dfc456422818a56e4:
# Input 0:
#    TXID: 88e5eb053b4ab251bbb969ba85cdfbe2a5b14929516834afc8cedb4a22bd1785
#    Out:       0
#    Signature: 7344c4c1b57d177661e01d726b6951aab54c9a81f6b2ddfa1ff516f068d76a118bb6da880d34a249778af8a39141660eef7b64de1d60782831dd39c9cc4bce7b
#    PubKey:    d930c603c372fa2d59408e41019195c92a3a36cb6e803e562776804d6aaab1a8d820bbc4514a1cf93f18bd281eb92c4b3a7a4ccdc80ed60f4caffe5da97f11f4 --- Public Key of Wallet 1
#  Output 0:
#    Value:  5
#    PubKeyHash: df56dfc9262cba6e5816250070b21777f0d0dbd0 --- Base58Decode of Address 2
#  Output 1:
#    Value:  5
#    PubKeyHash: 3ddecd8876590d12ef31dac44799ee1c540d30e2 --- Base58Decode of Address 1


#============ Block 0000009cd1d2494b13cc03d50a4ff28493a426c6be93c98524884f120eb7433f ============
# Prev. block: 
# PoW: true

# --- Transaction 88e5eb053b4ab251bbb969ba85cdfbe2a5b14929516834afc8cedb4a22bd1785:
#  Input 0:
#    TXID: 
#    Out:       -1
#    Signature: 
#    PubKey:    5468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73
#  Output 0:
#    Value:  10
#    PubKeyHash: 3ddecd8876590d12ef31dac44799ee1c540d30e2 --- Base58Decode of Address 1
```