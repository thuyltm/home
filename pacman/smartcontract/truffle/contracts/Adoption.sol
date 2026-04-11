// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Adoption {
    address[16] public adopters;//arrays have the type "address" and the length 16

    function adopt(uint petId) public returns (uint) {
        require(petId >= 0 && petId <= 15);
        adopters[petId] = msg.sender;//the address of who call this function is denoted by msg.sender
        return petId;
    }

    function getAdopters() public view returns (address[16] memory) {
        return adopters;
    }
}