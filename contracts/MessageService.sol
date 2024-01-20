// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;

/**
 * @title MessageService
 * @dev Can write messages & read previous sent messages
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract MessageService {
    
    struct Message {
        address addr;
        string username;
        string text;
        bool read;
    }
    Message[] messages;

    /**
     * @dev write store the message
     * @param username indicates the name of the user
     * @param message indicates the text of the message
     */
    function write(string memory username, string memory message) public {
        if(bytes(message).length > 300) {
            revert("Message limit is 300 characters");
        }

        messages.push(Message(msg.sender, username, message, false));
    }

    /**
     * @dev readAll store the message
     */
    function readAll() public {
        for (uint256 i = 0; i < messages.length; i++) {
            // check unread state
            if (messages[i].read == false) {
                messages[i].read = true; // mark as read
            }
        }
    }

    /**
     * @dev readByPosition store the message
     * @param position indicates the position of the message to read
     */
    function readByPosition(uint256 position) public {
        if (position > messages.length) {
            revert("Indicated position does not exists");
        }

        for (uint256 i = 0; i < messages.length; i++) {
            // check position and unread state
            if (i == position && messages[i].read == false) {
                messages[i].read = true; // mark as read
            }
        }
    }

    /**
     * @dev retrieveAll
     * @return retrive all the messages
     */
    function retrieveAll() public view returns (Message[] memory) {
        Message[] memory result = new Message[](messages.length);
        for (uint256 i = 0; i < messages.length; i++) {
            result[i] = messages[i];
        }
        return result;
    }

    /**
     * @dev retrieveUnread
     * @return retrive unread messages
     */
    function retrieveUnread() public view returns (Message[] memory) {

        // 1. determine the unread count
        uint256 resultCount;

        for (uint i = 0; i < messages.length; i++) {
            if (messages[i].read == false) {
                resultCount++;  
            }
        }

        // 2. create the fixed-length array
        Message[] memory result = new Message[](resultCount);  
        uint256 j;

        // 3. fill the array
        for (uint i = 0; i < messages.length; i++) {
            if (messages[i].read == false) {
                result[j] = messages[i];
                j++;
            }
        }
        
        return result;
    }

}