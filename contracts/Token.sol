// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract CashbackRewardsToken is ERC20, AccessControl {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    mapping(address => uint256) private _cashbackReceived;

    RewardPointsToken private rewardPointsToken;

    event CashbackGenerated(address indexed customer, uint256 cashbackAmount);
    event RewardPointsGenerated(address indexed customer, uint256 rewardPoints);
    event ProductPurchased(address indexed buyer, uint256 productValue, uint256 cashbackAmount, uint256 rewardPoints);

    constructor(address rewardPointsAddress) ERC20("CashbackRewardsToken", "CRT") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);

        rewardPointsToken = RewardPointsToken(rewardPointsAddress);
    }

    function mint(address to, uint256 amount) public onlyRole(MINTER_ROLE) {
        _mint(to, amount);
    }

    /**
     * @dev this function is called when a user wants to buy a product, 
            from which he or she will recieve 10% cashback and 10 reward points.
     * @param valueOfProduct The value of the product being purchased (in CRT tokens).
     */
    function buyProduct(uint256 valueOfProduct) public {
        uint256 buyerBalance = balanceOf(_msgSender());
        require(
            buyerBalance >= valueOfProduct,
            "Insufficient balance to purchase product."
        );

        _burn(_msgSender(), valueOfProduct);

        uint256 cashbackAmount = (valueOfProduct * 10) / 100;

        _mint(_msgSender(), cashbackAmount);

        rewardPointsToken.mint(_msgSender(), 10);

        _cashbackReceived[_msgSender()] += cashbackAmount;

        emit ProductPurchased(_msgSender(), valueOfProduct, cashbackAmount, 10);
        emit CashbackGenerated(_msgSender(), cashbackAmount);
        emit RewardPointsGenerated(_msgSender(), 10);
    }

    /**
     * @dev Returns the total cashback received by an address.
     * @param account is the address of the user.
     * @return The total cashback amount received by the user.
     */
    function getTotalCashbackReceived(address account) public view returns (uint256) {
        return _cashbackReceived[account];
    }
}

contract RewardPointsToken is ERC20, AccessControl {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    constructor() ERC20("RewardPointsToken", "RPT") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
    }

    /**
     * @dev this functions is called to reward points to a specified address (customer who bought the product).
     * @param to is the address to mint the 10 reward points to.
     * @param amount is the amount of reward points to mint, 10 for now but can be modified.
     */
    function mint(address to, uint256 amount) public onlyRole(MINTER_ROLE) {
        _mint(to, amount);
    }
}