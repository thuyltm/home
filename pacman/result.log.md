```sh
% bazel run //pacman:main -- send -from 1NGY8FFLQFJzKxA5YSkfQfGp9EmMXsgf8X -to 19Gi88hPgK9HFp9rEMj9QptTnDyi97RYU9 -amount 6
# 000000977ff451062df0ca2324109a2c8f4ecd0184456df7836bcfa198de2392
% bazel run //pacman:main -- printchain
#============ Block 000000977ff451062df0ca2324109a2c8f4ecd0184456df7836bcfa198de2392 ============
#Prev. block: 000000c11050c80b6dd1c79b54e1ef6ebfd49974fd3dd8377e71dfce3d09215b
#PoW: true

#--- Transaction bd5a56f22749d4d61ee66d737d99ba4e2fc9d88bf205cff81234d717f921873e:
# Input 0:
#    TXID: 3e12718bf367fe3a93a6ced7d6298ee704f7eb1f845180a8f342b7635b84e311
#    Out:       0
#    Signature: aafe414575fc8f210d9db55797e15fc5172d47974e3f59fa4b02b977b642720b971a8aab1c70c97e47718de2fdc6774cd03775e0bff95bfcd35122f4ad8d01a9
#    PubKey:    c1165f57b85f18dcd7c44c7c91c299740e1557b7c1f63bbc857b5fa6cb87721ec800e97bc3063f66880a459b137fd3e06355e996346c9ca2a20337433a2e70d8
# Output 0:
#    Value:  6
#    PubKeyHash: 5ab992cc4e66513bf58b17f7df8b609991219cef
#  Output 1:
#    Value:  4
#    PubKeyHash: e94adcd79cd8eb63c12bed3c1a48fe482fa284ec


#============ Block 000000c11050c80b6dd1c79b54e1ef6ebfd49974fd3dd8377e71dfce3d09215b ============
#Prev. block: 
#PoW: true

#--- Transaction 3e12718bf367fe3a93a6ced7d6298ee704f7eb1f845180a8f342b7635b84e311:
# Input 0:
#    TXID: 
#    Out:       -1
#    Signature: 
#    PubKey:    5468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73
#  Output 0:
#    Value:  10
#    PubKeyHash: e94adcd79cd8eb63c12bed3c1a48fe482fa284ec
% bazel run //pacman:main -- printUTXO
# Key: bd5a56f22749d4d61ee66d737d99ba4e2fc9d88bf205cff81234d717f921873e
# Value: 
#==================
#Out 0
#Value 6
#PubKeyHash 5ab992cc4e66513bf58b17f7df8b609991219cef

#==================
#Out 1
#Value 4
#PubKeyHash e94adcd79cd8eb63c12bed3c1a48fe482fa284ec
% bazel run //pacman:main -- getbalance -address 1NGY8FFLQFJzKxA5YSkfQfGp9EmMXsgf8X
# Balance of '1NGY8FFLQFJzKxA5YSkfQfGp9EmMXsgf8X': 4
```