// Deployed at 0x5E8b7B869976bC7FCF3437D050Fc5ecc8dBd2A6B
// cethacea contract call --contract=0x5E8b7B869976bC7FCF3437D050Fc5ecc8dBd2A6B --chain=http://localhost:8545 --abi=abi.json 'getDataHash' 0
object "DatahashFetcher" {
   code {
      datacopy(0, dataoffset("runtime"), datasize("runtime"))
      return(0, datasize("runtime"))
   }
   object "runtime" {
      code {
         // Match against the keccak of the ABI function signature needed.
         switch shr(0xe0,calldataload(0))
            // bytes4(keccak("getDataHash(uint256)"))
            case 0xa595d8fc {
               // DATAHASH opcode has hex value 0x49
               let kzg_versioned_hash := verbatim_1i_1o(hex"49", calldataload(0x04))
               mstore(0x00,kzg_versioned_hash)
               return(0,0x20)
            }
      }
   }
}