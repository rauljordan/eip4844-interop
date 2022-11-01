// SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.13;

interface DataHashReader {
    function getDataHashes() external view returns (bytes32[] memory);
}

contract ToySequencerInbox {
    DataHashReader dataHashReader;
    event DataHashesParsed(
        uint numHashes,
        bytes32[] hashes
    );
    constructor(address _dataHashReader) {
        dataHashReader = DataHashReader(_dataHashReader);
    }
    function submitBatch() public {
        bytes32[] memory hashes = dataHashReader.getDataHashes();
        emit DataHashesParsed(hashes.length, hashes);
    }
}
