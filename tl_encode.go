package mtproto

import (
	"encoding/binary"
	"math"
	"math/big"

	"github.com/golang/glog"
)

// EncodeBuf type
type EncodeBuf struct {
	Buf []byte
}

// NewEncodeBuf func
func NewEncodeBuf(cap int) *EncodeBuf {
	return &EncodeBuf{make([]byte, 0, cap)}
}

// Int func
func (e *EncodeBuf) Int(s int32) {
	e.Buf = append(e.Buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.Buf[len(e.Buf)-4:], uint32(s))
}

// UInt func
func (e *EncodeBuf) UInt(s uint32) {
	e.Buf = append(e.Buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.Buf[len(e.Buf)-4:], s)
}

// Long func
func (e *EncodeBuf) Long(s int64) {
	e.Buf = append(e.Buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.Buf[len(e.Buf)-8:], uint64(s))
}

// Double func
func (e *EncodeBuf) Double(s float64) {
	e.Buf = append(e.Buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.Buf[len(e.Buf)-8:], math.Float64bits(s))
}

// String func
func (e *EncodeBuf) String(s string) {
	e.StringBytes([]byte(s))
}

// BigInt func
func (e *EncodeBuf) BigInt(s *big.Int) {
	e.StringBytes(s.Bytes())
}

// StringBytes func
func (e *EncodeBuf) StringBytes(s []byte) {
	var res []byte
	size := len(s)
	if size < 254 {
		nl := 1 + size + (4-(size+1)%4)&3
		res = make([]byte, nl)
		res[0] = byte(size)
		copy(res[1:], s)

	} else {
		nl := 4 + size + (4-size%4)&3
		res = make([]byte, nl)
		binary.LittleEndian.PutUint32(res, uint32(size<<8|254))
		copy(res[4:], s)

	}
	e.Buf = append(e.Buf, res...)
}

// Bytes func
func (e *EncodeBuf) Bytes(s []byte) {
	e.Buf = append(e.Buf, s...)
}

// VectorInt func
func (e *EncodeBuf) VectorInt(v []int32) {
	e.Int(crc32Vector)
	// x := make([]byte, 4+4+len(v)*4)
	x := make([]byte, 4+len(v)*4)
	// binary.LittleEndian.PutUint32(x, crc32Vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	i := 8
	for _, v := range v {
		binary.LittleEndian.PutUint32(x[i:], uint32(v))
		i += 4
	}
	e.Buf = append(e.Buf, x...)
}

// VectorLong func
func (e *EncodeBuf) VectorLong(v []int64) {
	e.Int(crc32Vector)
	// x := make([]byte, 4+len(v)*8)
	x := make([]byte, 4+4+len(v)*8)
	var c = int32(crc32Vector)
	binary.LittleEndian.PutUint32(x, uint32(c))
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	i := 8
	for _, v := range v {
		binary.LittleEndian.PutUint64(x[i:], uint64(v))
		i += 8
	}
	e.Buf = append(e.Buf, x...)
}

// VectorString func
func (e *EncodeBuf) VectorString(v []string) {
	e.Int(crc32Vector)
	// x := make([]byte, 8)
	x := make([]byte, 4)
	binary.LittleEndian.PutUint32(x, uint32(len(v)))
	// binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	e.Buf = append(e.Buf, x...)
	for _, v := range v {
		e.String(v)
	}
}

// Vector func
func (e *EncodeBuf) Vector(v []TL) {
	e.Int(crc32Vector)
	// x := make([]byte, 8)
	x := make([]byte, 4)
	binary.LittleEndian.PutUint32(x, uint32(len(v)))
	// binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	e.Buf = append(e.Buf, x...)
	for _, v := range v {
		e.Buf = append(e.Buf, v.Encode()...)
	}
}

// Encode func
func (e TL_msg_container) Encode() []byte { return nil }

// Encode func
func (e TL_resPQ) Encode() []byte { return nil }

// Encode func
func (e TL_server_DH_params_ok) Encode() []byte { return nil }

// Encode func
func (e TL_server_DH_inner_data) Encode() []byte { return nil }

// Encode func
func (e TL_dh_gen_ok) Encode() []byte { return nil }

// Encode func
func (e TL_rpc_result) Encode() []byte { return nil }

// Encode func
func (e TL_rpc_error) Encode() []byte { return nil }

// Encode func
func (e TL_new_session_created) Encode() []byte { return nil }

// Encode func
func (e TL_bad_server_salt) Encode() []byte { return nil }

// Encode func
func (e TL_crc32BadMsgNotification) Encode() []byte { return nil }

// Encode func

// Encode func
func (e TL_req_pq) Encode() []byte {
	// x := NewEncodeBuf(20)
	x := NewEncodeBuf(512)
	x.Int(crc32ReqPQ)
	// x.UInt(crc32ReqPQ)
	x.Bytes(e.nonce)
	return x.Buf
}

// Encode func
func (e TL_p_q_inner_data) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc32PQInnerData)
	x.BigInt(e.pq)
	x.BigInt(e.p)
	x.BigInt(e.q)
	x.Bytes(e.nonce)
	x.Bytes(e.server_nonce)
	x.Bytes(e.new_nonce)
	return x.Buf
}

// Encode func
func (e TL_req_DH_params) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(crc32ReqDHParams)
	// x.UInt(crc32ReqDHParams)
	x.Bytes(e.nonce)
	x.Bytes(e.server_nonce)
	x.BigInt(e.p)
	x.BigInt(e.q)
	x.Long(int64(e.fp))
	x.StringBytes(e.encdata)
	return x.Buf
}

// Encode func
func (e TL_client_DH_inner_data) Encode() []byte {
	x := NewEncodeBuf(512)
	// x.UInt(crc32ClientDHInnerData)
	x.Int(crc32ClientDHInnerData)
	x.Bytes(e.nonce)
	x.Bytes(e.server_nonce)
	x.Long(e.retry)
	x.BigInt(e.g_b)
	return x.Buf
}

// Encode func
func (e TL_set_client_DH_params) Encode() []byte {
	// x := NewEncodeBuf(256)
	x := NewEncodeBuf(512)
	// x.UInt(crc32SetClientDHParams)
	x.Int(crc32SetClientDHParams)
	x.Bytes(e.nonce)
	x.Bytes(e.server_nonce)
	x.StringBytes(e.encdata)
	return x.Buf
}

// Encode func
func (e TL_ping) Encode() []byte {
	// x := NewEncodeBuf(32)
	x := NewEncodeBuf(512)
	x.Int(crc32Ping)
	// x.UInt(crc32Ping)
	x.Long(e.ping_id)
	return x.Buf
}

// Encode func
func (e TL_pong) Encode() []byte {
	// x := NewEncodeBuf(32)
	x := NewEncodeBuf(512)
	x.Int(crc32Pong)
	// x.UInt(crc32Pong)
	x.Long(e.msg_id)
	x.Long(e.ping_id)
	return x.Buf
}

// Encode func
func (e TL_msgs_ack) Encode() []byte {
	glog.V(3).Infof("TL_msgs_ack encoding!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	// x := NewEncodeBuf(64)
	x := NewEncodeBuf(512)
	x.Int(crc32MsgsAck)
	// x.UInt(crc32MsgsAck)
	x.VectorLong(e.msgIds)
	return x.Buf
}
