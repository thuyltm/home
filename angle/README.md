### Hardware Security & Trusted Execution
- Apple Secure Enclave: Architecture SEP-based features: key lifecycle, biometric-bound keys, DataProtection classes, and Keychain ACLs across the MDM stack on Apple devices
- Androud StrongBox & TEE: Design StrongBox/TEE-based features, perform key/ID attestation, and assess compatibility across OEM device fleets
- TPM 2.0: Architect desktop MDM agents and TPM-based device identity using PCRs, sealed storage, and EK/AIK certificates
- HSM & Cloud KMS: Architect key management with on-prem HSMs or cloud KMS via PKCS#11/KMIP; plan HA/DR for cryptographic key material

Design integrity verification across the device fleet and handle tampered/rooted/jailbroken endpoints
### Cryptography & Applied Cryptography
- Asymmetric Key Systems (AKS): AES-GCM/GCM-SIV/CBC/CTR, ChaCha20-Poly1305; mastery of nonce management, AEAD design, padding oracle, and timing attacks
- Symmetric Key Systems (SKS): RSA (PKCS#1 v1.5, OAEP, PSS) ≥ 2048 bits with deep understanding of pitfalls; able to advise on migration from RSA to ECC
- Elliptic Curve: ECDSA, EdDSA (Ed25519), ECDH/X25519 across NIST and Curve25519 families; pick curves appropriate to mobile, IoT, or server contexts.
- Hash, MAC, KDF: SHA-2/3, BLAKE2/3, HMAC, and KDFs (HKDF, PBKDF2, Argon2, scrypt) with correct selection per use case
- PQC: Understand ML-KEM (Kyber), ML-DSA (Dilithium), hybrid schemes, and build a PQC roadmap for the MDM product.

Mandate well-vetted libraries (Tink, libsodium, BoringSSL, BouncyCastle, JCA/JCE, CryptoKit) and audit against custom cryptography.

### Data Encryption & Digital Signatures
- Encryption at rest: FDE, FBE (Android), FileVault/APFS, envelope encryption (DEK/KEK), AES-KW; design key storage and rotation strategy
- Encryption in transit: TLS 1.2/1.3, mTLS server↔agent, certificate pinning, OCSP stapling, HSTS; design defenses against downgrade, MITM, BEAST, and CRIME
- PKI and Digital Certificate Management: Design and operate an internal CA for MDM enrollment at fleet scale. X.509, CSR, chain validation, OCSP/CRL, SCEP, EST, ACME
- Digital Signature: Sign commands/policies, firmware/OTA, APK/IPA, with timestamping and JWT/JWS/JWE/COSE for API authentication
- Key Management lifecycle: generation, distribution, rotation, revocation, escrow, destruction; full key lifecycle with separtion of signing and encryption keys, in line with NIST SP 800-57
- Secure channels & protocols: Noise, Signal Protocol, Double Ratchet for secure server↔agent messaging when required.
- Threat modeling & secure SDLC: Embed STRIDE/LINDDUN/abuse-case threat modeling into the SDLC for every feature