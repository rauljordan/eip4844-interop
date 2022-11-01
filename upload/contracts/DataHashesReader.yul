// Deployed at 0x6006f425ac71535452a5e23BCF73E26A115f281d
object "DataHashesReader" {
   code {
      datacopy(0, dataoffset("runtime"), datasize("runtime"))
      return(0, datasize("runtime"))
   }
   object "runtime" {
      code {
         // Match against the keccak of the ABI function signature needed.
         switch shr(0xe0,calldataload(0))
            // bytes4(keccak("getDataHashes()"))
            case 0x46115383
            case 0xe83a2d82 {
                // DATAHASH opcode has hex value 0x49
                let i := 0
                for {} true {} {
                    let hash := verbatim_1i_1o(hex"49", i)
                    if iszero(hash) {
                        break
                    }
                    mstore(add(mul(i, 32), 64), hash)
                    i := add(i, 1)
                }
                mstore(0, 32)
                mstore(32, i)
                return(0, add(mul(i, 32), 64))
            }
      }
   }
}