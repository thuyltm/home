To verify an ECDSA signature, you must possess
- the original signed data
- the public key of the signer
- the signature itself (typically represented as two integer, $r$ and $s$)

Verification involves:
1. hashing the message
2. Computing a point $R^{'}$ using the public key
3. Checking if its $x$-coordinate matches $r$

### Essential Components
- Public Key ($Q$): The signer's public key (a point on the elliptic curve).
- Message Hash ($z$): The hash of the original data using the same algorithm (e.g., SHA-256).
- Signature ($r, s$): The numerical pair produced during signing.
- Curve Parameters: The elliptic curve definition (e.g., secp256k1). 

### Verification Steps (Mathematical Process)
1. Validate Signatures: Ensure $r$ and $s$ are integers in the range $[1, n-1]$, where $n$ is the __order of the curve__.
2. Calculate Message Hash ($z$): Hash the original message to get $z$.
3. Compute Inverse ($w$): Calculate the __modular inverse of the signature__ component $s$: $w=s^{-1}$
4. Calculate $u_1$, $u_2$
    
    1. $u_1 = z \times w$ (mod $n$)
    2. $u_2 = r \times w$ (mod $n$)
5. Calculate Curve Point ($X$): Compute $X=(x_X,y_X)=u_1G+u_2Q$, where $G$ is the generator point.
6. Verify Signature: The signature is valid if $x_X \equiv r \qquad (\text{mod\,} n)$
    , otherwise it is invalid