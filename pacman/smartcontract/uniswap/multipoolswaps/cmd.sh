#! /bin/bash
forge init priceoracle
rm -rf .git
git rm -r --cached .
forge install foundry-rs/forge-std