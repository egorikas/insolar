// Code generated by "stringer -type=packetType"; DO NOT EDIT.

package packet

import "strconv"

const _packetType_name = "TypePingTypeStoreTypeFindHostTypeFindValueTypeRPCTypeRelayTypeAuthTypeCheckOriginTypeObtainIPTypeRelayOwnershipTypeKnownOuterHostsTypeCheckNodePrivTypeCascadeSendTypePulse"

var _packetType_index = [...]uint8{0, 8, 17, 29, 42, 49, 58, 66, 81, 93, 111, 130, 147, 162, 171}

func (i packetType) String() string {
	i -= 1
	if i < 0 || i >= packetType(len(_packetType_index)-1) {
		return "packetType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _packetType_name[_packetType_index[i]:_packetType_index[i+1]]
}
