// contracts/etherwatt.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EtherWatt {
    string public name = "EtherWatt";
    string public symbol = "EW";
    uint8 public decimals = 1;
    uint256 public totalSupply = 1000000;
    mapping(address => uint256) public balanceOf;

    event Transfer(address indexed from, address indexed to, uint256 value);

    constructor() {
        balanceOf[msg.sender] = totalSupply;
    }

    function reward(address _to, uint256 _value) public {
        require(balanceOf[msg.sender] >= _value && _value > 0);
        balanceOf[msg.sender] -= _value;
        balanceOf[_to] += _value;
        emit Transfer(msg.sender, _to, _value);
    }

    function transfer(address _from, address _to, uint256 _value) public {
        require(balanceOf[_from] >= _value && _value > 0);
        balanceOf[_from] -= _value;
        balanceOf[_to] += _value;
        emit Transfer(_from, _to, _value);
    }

}
