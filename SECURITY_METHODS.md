### Mutual authentication

Mutual authentication can be accomplised with two types of credentials: usernames and passwords, and public key certificates

Mutual authentication is often employed in the Internet of Things (IoT)

A mutual authentication process may be implemented as follows:
1. Alice sends a message **encrypted with Bob's public key** to show that Alice is a valid user
2. Bob verifies the message:
    1. The message is decrypted with Bob's secret key, giving Alice's ID
    2. Bob checks if the message matches a valid user
3. Bob sends Alice a message back, **encrypted with Alice's public key**, to show that Bob is a valid user
4. Alice verifies the message
    1. The message is decrypted with Alice's secret key, giving Blob;s ID
    2. Alice checks if the message matches a valid user
5. At this point, both parties are verified to be who they claim to be and sage for the other to communicate with. Lastly, Alice and Bob will **create a shared secret key so that the can continue communicate** in a secure way

### Server Name Indication

Server Name Indication (SNI) is an extension to Transport Layer Security (TLS). The extension allows a server to present one of multiple possible certicates on the same IP address and port number.