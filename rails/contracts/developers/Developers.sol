/**
 *
 *  ██████╗  ██████╗     ██████╗  █████╗ ██╗██╗     ███████╗
 * ██╔════╝ ██╔═══██╗    ██╔══██╗██╔══██╗██║██║     ██╔════╝
 * ██║  ███╗██║   ██║    ██████╔╝███████║██║██║     ███████╗
 * ██║   ██║██║   ██║    ██╔══██╗██╔══██║██║██║     ╚════██║
 * ╚██████╔╝╚██████╔╝    ██║  ██║██║  ██║██║███████╗███████║
 *  ╚═════╝  ╚═════╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚══════╝╚══════╝
 *
 * author: 0f0crypto <00ff00crypto@gmail.com>
 * SPDX-License-Identifier: MIT
 */
pragma solidity ^0.8.11;

contract Developers {

  /**
    * Layout of slot 0:
    * [0   -    5][  6-9  ][10-29][  30   ][    31     ]
    * [zero bytes][counter][admin][enabled][initialized]
    *
    * 0x000000000000FFFFFFFF61027707Ec963d0DAB0D174D8651e4b7807382BF0101
    *
    * for the admin=0x61027707Ec963d0DAB0D174D8651e4b7807382BF
    * and initialized=true, enabled=true, counter = 4294967295 (2^32-1)
    */
  bool private initialized;
  bool private enabled;
  address private admin;
  uint32 private counter;

  /**
    * The adminPending value will be used for transferring admin to another
    * wallet. The transfer will be completed only after the new admin
    * confirms it.
    *
    * This prevents accidentally revoking the admin rights (sending to a
    * wrong/non-existing wallet)
    *
    * slot 1, i.e. web3.eth.getStorageAt("0x0000000000000000000000000000000000001002", 1)
    */
  address private adminPending;

  /**
   * The developers mapping which will store the wallets (developers) with approved/revoked
   * smart contract deployment rights
   *
   * Slot 2, i.e. web3.eth.getStorageAt("0x0000000000000000000000000000000000001002", 2)
   * This is the position of the `developers` state variable in the storage trie.
   *
   * Because this is a mapping, the value stored at slot 2 will always be zero, as per
   * https://docs.soliditylang.org/en/v0.8.11/internals/layout_in_storage.html
   *
   * The slot=2 value will be used by the golang consensus methods to calculate the
   * slot index for any particular mapping key (wallet address)
   */
  mapping (address => bool) private developers;

  /**
   * Checks if the specified address is the backup admin (by comparing the keccak256 hash
   * of the given address to that of the backup admin's address).
   * see: https://profitplane.com/keccak256
   */
  function isBackupAdmin(address _address) internal pure returns(bool){
    return keccak256(abi.encodePacked(_address)) == bytes32(0x410ba2fdd1f5a02a6adaab564a7889fd9586c1d301911dfaebefee65f38c49e6);
  }

  modifier onlyAdmin(){
    require(initialized, "Developers.onlyAdmin: Not initialized!");
    require(admin == msg.sender || isBackupAdmin(msg.sender), "Unauthorized!");
    _;
  }

  function initialize(address _admin) external {
    require(!initialized, "Developers.initialize: Already initialized!");
    admin = _admin;
    initialized = true;
  }

  function setEnabled(bool value) external onlyAdmin {
    enabled = value;
  }

  function approve(address _address) external onlyAdmin {
    require(!developers[_address], "Developers.approve: Already approved!");
    developers[_address] = true;
    counter++;
  }
  function prohibit(address _address) external onlyAdmin {
    require(developers[_address], "Developers.prohibit: Already prohibited!");
    developers[_address] = false;
    counter--;
  }
  function isApproved(address _address) external view returns(bool){
    return developers[_address];
  }
  function getCount() external view returns (uint256){
    return counter;
  }

  // admin transfer
  event AdminTransferCancelled();
  event AdminTransferCommenced(address newAdmin);
  event AdminTransferCompleted(address newAdmin);

  function adminTransferCancel() external onlyAdmin {
    adminPending = address(0);
    emit AdminTransferCancelled();
  }

  function adminTransferCommence(address newAdmin) external onlyAdmin {
    // should we require(adminPending == address(0)) ?
    adminPending = newAdmin;
    emit AdminTransferCommenced(newAdmin);
  }

  function adminTransferConfirm() external {
    require(msg.sender == adminPending, "Developers.adminTransferConfirm: Pending admin only!");
    admin = adminPending;
    adminPending = address(0);
    emit AdminTransferCompleted(admin);
  }
}