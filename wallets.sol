pragma solidity ^0.4.23;

contract Ownable {
    address public owner;

    constructor() public {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function transferOwnership(address newOwner) public onlyOwner {
        if (newOwner != address(0)) {
            owner = newOwner;
        }
    }

    function getOwner() public view returns(address) {
        return owner;
    }

}

contract HotWallet {
    address public owner;
    WalletFactory public factory;
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }
    modifier onlyAdmin() {
        require(msg.sender == factory.getOwner());
        _;
    }
    
    constructor(address _admin, address _owner) public {
        require(_admin != address(0) && _owner != address(0));
        factory = WalletFactory(_admin);
        require(factory.isWalletFactory());
        owner = _owner;
    }
    
    function isHotWallet() external pure returns(bool){
        return true;
    }

    // the msg.sender must be factory.owner
    function dyn_call(address _constract, bytes _data) public payable onlyAdmin {
        if (!_constract.call.value(msg.value)(_data)){
            revert();
        }
    }
    
    function withdraw() external onlyOwner{
        require(owner != address(0));
        owner.transfer(address(this).balance);
    }
    
    // other functions
    function() public payable{}

}

contract WalletFactory is Ownable {
    mapping(address => address) public hotwallets;

    function createWallet(address _owner) public onlyOwner {
        require(hotwallets[_owner] == address(0));
        HotWallet w = new HotWallet(address(this), _owner);
        hotwallets[_owner] = address(w);
    }
    
    function isWalletFactory() external pure returns(bool){
        return true;
    }
}
