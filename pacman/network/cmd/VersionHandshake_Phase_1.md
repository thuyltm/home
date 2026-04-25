Before two nodes can start doing something, they have to exchange their verisons
1. Node A connect to Node B
2. Node A sends information about its version to Node B
3. Node B sends information about its version back to Node A
4. Node B send information ACK message to Node A
5. Node B sets version to the minimum of these 2 versions
6. Node A send ACK message to Node B
7. Node A sets version to the minimum of these 2 versions.