package mtproto

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"math"
	"math/big"
	"os"
	"runtime/pprof"
)

// DecodeBuf type
type DecodeBuf struct {
	Buf  []byte
	Off  int
	Size int
	Err  error
}

// NewDecodeBuf func
func NewDecodeBuf(b []byte) *DecodeBuf {
	return &DecodeBuf{b, 0, len(b), nil}
}

func (m *DecodeBuf) Long() int64 {
	if m.Err != nil {
		return 0
	}
	if m.Off+8 > m.Size {
		m.Err = errors.New("DecodeLong")
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(m.Buf[m.Off : m.Off+8]))
	m.Off += 8
	if __debug&DebugLevelDecodeDetails != 0 {
		//fmt.Println("Decode::Long::", x)
	}
	return x
}

// Double func
func (m *DecodeBuf) Double() float64 {
	if m.Err != nil {
		return 0
	}
	if m.Off+8 > m.Size {
		m.Err = errors.New("DecodeDouble")
		return 0
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(m.Buf[m.Off : m.Off+8]))
	m.Off += 8
	if __debug&DebugLevelDecodeDetails != 0 {
		//fmt.Println("Decode::Double::", x)
	}
	return x
}

// Int func
func (m *DecodeBuf) Int() int32 {
	if m.Err != nil {
		return 0
	}
	//fmt.Printf("m.Off+4 : %d m.Size : %d\n",m.Off+4,m.Size)
	if m.Off+4 > m.Size {
		m.Err = errors.New("DecodeInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.Buf[m.Off : m.Off+4])
	m.Off += 4
	if __debug&DebugLevelDecodeDetails != 0 {
		//fmt.Println("Decode::Int::", x)
	}
	return int32(x)
}

// UInt func
func (m *DecodeBuf) UInt() uint32 {
	if m.Err != nil {
		return 0
	}
	if m.Off+4 > m.Size {
		m.Err = errors.New("DecodeUInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.Buf[m.Off : m.Off+4])
	m.Off += 4
	if __debug&DebugLevelDecodeDetails != 0 {
		//fmt.Println(fmt.Sprintf("Decode::UInt::%x", x))
	}
	return x
}

// Bytes func
func (m *DecodeBuf) Bytes(size int) []byte {
	if m.Err != nil {
		return nil
	}
	if m.Off+size > m.Size {
		m.Err = errors.New("DecodeBytes")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.Buf[m.Off:m.Off+size])
	m.Off += size
	if __debug&DebugLevelDecodeDetails != 0 {
		if len(x) > 10 {
			fmt.Println("Decode::Bytes::", len(x), x[:10], " ...")
		} else {
			fmt.Println("Decode::Bytes::", len(x), x)
		}

	}
	return x
}

// StringBytes func
func (m *DecodeBuf) StringBytes() []byte {
	if m.Err != nil {
		return nil
	}
	var size, padding int

	if m.Off+1 > m.Size {
		m.Err = errors.New("DecodeStringBytes")
		return nil
	}
	size = int(m.Buf[m.Off])
	m.Off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.Off+3 > m.Size {
			m.Err = errors.New("DecodeStringBytes")
			return nil
		}
		size = int(m.Buf[m.Off]) | int(m.Buf[m.Off+1])<<8 | int(m.Buf[m.Off+2])<<16
		m.Off += 3
		padding = (4 - size%4) & 3
	}

	if m.Off+size > m.Size {
		m.Err = errors.New("DecodeStringBytes: Wrong size")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.Buf[m.Off:m.Off+size])
	m.Off += size

	if m.Off+padding > m.Size {
		m.Err = errors.New("DecodeStringBytes: Wrong padding")
		return nil
	}
	m.Off += padding
	if __debug&DebugLevelDecodeDetails != 0 {
		if len(x) > 10 {
			fmt.Println("Decode::StringBytes::", len(x), x[:10], " ...")
		} else {
			fmt.Println("Decode::StringBytes::", len(x), x)
		}

	}
	return x
}

// String func
func (m *DecodeBuf) String() string {
	b := m.StringBytes()
	if m.Err != nil {
		return ""
	}
	x := string(b)
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::String::", x)
	}
	return x
}

// BigInt func
func (m *DecodeBuf) BigInt() *big.Int {
	b := m.StringBytes()
	if m.Err != nil {
		return nil
	}
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::BigInt::", x)
	}
	return x
}

// VectorInt func
func (m *DecodeBuf) VectorInt() []int32 {
	// constructor := m.UInt()
	constructor := m.Int()
	if m.Err != nil {
		return nil
	}
	if constructor != crc32Vector {
		// m.Err = fmt.Errorf("DecodeVectorInt: Wrong constructor (0x%08x)", constructor)
		m.Err = fmt.Errorf("DecodeVectorInt: Wrong constructor %d", constructor)
		return nil
	}
	size := m.Int()
	if m.Err != nil {
		return nil
	}
	if size < 0 {
		m.Err = errors.New("DecodeVectorInt: Wrong size")
		return nil
	}
	x := make([]int32, size)
	i := int32(0)
	for i < size {
		y := m.Int()
		if m.Err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::VectorInt::", x)
	}
	return x
}

// VectorLong func
func (m *DecodeBuf) VectorLong() []int64 {
	// constructor := m.UInt()
	constructor := m.Int()
	if m.Err != nil {
		return nil
	}
	if constructor != crc32Vector {
		// m.Err = fmt.Errorf("DecodeVectorLong: Wrong constructor (0x%08x)", constructor)
		m.Err = fmt.Errorf("DecodeVectorLong: Wrong constructor %d", constructor)
		return nil
	}
	size := m.Int()
	if m.Err != nil {
		return nil
	}
	if size < 0 {
		m.Err = errors.New("DecodeVectorLong: Wrong size")
		return nil
	}
	x := make([]int64, size)
	i := int32(0)
	for i < size {
		y := m.Long()
		if m.Err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::VectorLong::", x)
	}
	return x
}

// VectorString func
func (m *DecodeBuf) VectorString() []string {
	// constructor := m.UInt()
	constructor := m.Int()
	if m.Err != nil {
		return nil
	}
	if constructor != crc32Vector {
		// m.Err = fmt.Errorf("DecodeVectorString: Wrong constructor (0x%08x)", constructor)
		m.Err = fmt.Errorf("DecodeVectorString: Wrong constructor %d", constructor)
		return nil
	}
	size := m.Int()
	if m.Err != nil {
		return nil
	}
	if size < 0 {
		m.Err = errors.New("DecodeVectorString: Wrong size")
		return nil
	}
	x := make([]string, size)
	i := int32(0)
	for i < size {
		y := m.String()
		if m.Err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::VectorString::", x)
	}
	return x
}

// Bool func
func (m *DecodeBuf) Bool() bool {
	constructor := m.Int()
	if m.Err != nil {
		return false
	}
	switch constructor {
	case crc_boolFalse:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("Decode::Bool::", false)
		}
		return false
	case crc_boolTrue:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("Decode::Bool::", true)
		}
		return true
	}
	return false
}

// Vector func
func (m *DecodeBuf) Vector() []TL {
	// constructor := m.UInt()
	constructor := m.Int()
	glog.V(3).Infof("===== = ==  > >> > >   >>> > > >> > >> > >> >>>>>>>>>>>>>>>>>>>>>>>>>>>>Constructor: %d", constructor)
	if m.Err != nil {
		return nil
	}
	if constructor != crc32Vector {
		// m.Err = fmt.Errorf("DecodeVector: Wrong constructor (0x%08x)", constructor)
		m.Err = fmt.Errorf("DecodeVector: Wrong constructor %d", constructor)
		return nil
	}
	size := m.Int()
	if m.Err != nil {
		return nil
	}
	if size < 0 {
		m.Err = errors.New("DecodeVector: Wrong size")
		return nil
	}
	x := make([]TL, size)
	i := int32(0)
	for i < size {
		y := m.Object()
		if m.Err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::Vector::", x)
	}
	return x
}

// VectorWithConstructor func
func (m *DecodeBuf) VectorWithConstructor(constructor int32) []TL {
	size := m.Int()
	if m.Err != nil {
		return nil
	}
	if size < 0 {
		m.Err = errors.New("DecodeVector: Wrong size")
		return nil
	}
	x := make([]TL, size)
	i := int32(0)
	for i < size {
		y := m.Object()
		if m.Err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::Vector::", x)
	}
	return x
}
// Object func
func (m *DecodeBuf) Object() (r TL) {
	constructor := m.Int()

	//fmt.Printf("Decode object: %x\n", constructor)
	if m.Err != nil {
		return nil
	}

	// if constructor == 1910543603 {
	// 	glog.V(3).Infof("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	// }

	switch constructor {
	case crc32ResPQ:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("reqPQ", constructor)
		}
		r = TL_resPQ{m.Bytes(16), m.Bytes(16), m.BigInt(), m.VectorLong()}

	case crc32ServerDHParamsOk:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("server_DH_params_ok", constructor)
		}
		r = TL_server_DH_params_ok{m.Bytes(16), m.Bytes(16), m.StringBytes()}

	case crc32ServerDHInnerData:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("server_DH_inner_data", constructor)
		}
		r = TL_server_DH_inner_data{
			m.Bytes(16), m.Bytes(16), m.Int(),
			m.BigInt(), m.BigInt(), m.Int(),
		}

	case crc32DHGenOk:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("dh_gen_ok", constructor)
		}
		r = TL_dh_gen_ok{m.Bytes(16), m.Bytes(16), m.Bytes(16)}

	case crc32Ping:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("ping", constructor)
		}
		r = TL_ping{m.Long()}

	case crc32Pong:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("pong", constructor)
		}
		r = TL_pong{m.Long(), m.Long()}

	case crc32MsgContainer:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("msg_container", constructor)
		}
		size := m.Int()
		arr := make([]TL_MT_message, size)
		for i := int32(0); i < size; i++ {
			arr[i] = TL_MT_message{m.Long(), m.Int(), m.Int(), m.Object()}
			//fmt.Println(constructor, arr[i])
			if m.Err != nil {
				fmt.Println(m.Err.Error())
				return nil
			}
		}
		r = TL_msg_container{arr}

	case crc32RPCResult:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("rpc_result")
		}
		r = TL_rpc_result{m.Long(), m.Object()}
		fmt.Printf("Object rpc_result: %v\n",m.Object())

	case crc32RPCError:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("rpc_error", constructor)
		}
		r = TL_rpc_error{m.Int(), m.String()}

	case crc32NewSessionCreated:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("new_session_created", constructor)
		}
		r = TL_new_session_created{m.Long(), m.Long(), m.Bytes(8)}

	case crc32BadServerSalt:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("bad_server_salt", constructor)
		}
		r = TL_bad_server_salt{m.Long(), m.Int(), m.Int(), m.Bytes(8)}

	case crc32BadMsgNotification:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("bad_msg_notification", constructor)
		}
		r = TL_crc32BadMsgNotification{m.Long(), m.Int(), m.Int()}

	case crc32MsgsAck:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("msgs_ack", constructor)
		}
		r = TL_msgs_ack{m.VectorLong()}

	case crc32GzipPacked:
		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println("gzip_packed", constructor)
		}
		obj := make([]byte, 0, 4096)

		var buf bytes.Buffer
		_, _ = buf.Write(m.StringBytes())
		gz, _ := gzip.NewReader(&buf)
		b := make([]byte, 4096)
		for true {
			n, _ := gz.Read(b)
			obj = append(obj, b[0:n]...)
			if n <= 0 {
				break
			}
		}
		d := NewDecodeBuf(obj)
		r = d.Object()

	case crc32Vector:
		r = TL_vector_wallpaper{
			Wallpapers: m.VectorWithConstructor(constructor),
		}
	default:

		if __debug&DebugLevelDecodeDetails != 0 {
			fmt.Println(fmt.Sprintf("default %x", constructor))
		}
		r = m.ObjectGenerated(constructor)

	}

	if m.Err != nil {
		return nil
	}
	return
}

// Flags func
func (m *DecodeBuf) Flags() int32 {
	if m.Err != nil {
		return 0
	}
	if m.Off+4 > m.Size {
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
		m.Err = errors.New("DecodeInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.Buf[m.Off : m.Off+4])
	m.Off += 4
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::Flags::", x)
	}
	return int32(x)
}

// FlaggedLong func
func (m *DecodeBuf) FlaggedLong(flags, f int32) int64 {
	if m.Err != nil {
		return 0
	}
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return 0
	}

	if m.Off+8 > m.Size {
		m.Err = errors.New("DecodeLong")
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(m.Buf[m.Off : m.Off+8]))
	m.Off += 8
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::FlaggedLong::", x)
	}
	return x
}

// FlaggedDouble func
func (m *DecodeBuf) FlaggedDouble(flags, f int32) float64 {
	if m.Err != nil {
		return 0
	}
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return 0
	}

	if m.Off+8 > m.Size {
		m.Err = errors.New("DecodeDouble")
		return 0
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(m.Buf[m.Off : m.Off+8]))
	m.Off += 8
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::FlaggedDouble::", x)
	}
	return x
}

// FlaggedInt func
func (m *DecodeBuf) FlaggedInt(flags, f int32) int32 {
	if m.Err != nil {
		return 0
	}
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return 0
	}

	if m.Off+4 > m.Size {
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
		m.Err = errors.New("DecodeInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.Buf[m.Off : m.Off+4])
	m.Off += 4
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::FlaggedInt::", x)
	}
	return int32(x)
}

// FlaggedString func
func (m *DecodeBuf) FlaggedString(flags, f int32) string {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return ""
	}

	b := m.StringBytes()
	if m.Err != nil {
		return ""
	}
	x := string(b)
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::FlaggedString::", x)
	}
	return x
}

// FlaggedVector func
func (m *DecodeBuf) FlaggedVector(flags, f int32) []TL {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return nil
	}
	// constructor := m.UInt()
	constructor := m.Int()
	if m.Err != nil {
		return nil
	}
	if constructor != crc32Vector {
		// m.Err = fmt.Errorf("DecodeFlaggedVector: Wrong constructor (0x%08x)", constructor)
		m.Err = fmt.Errorf("DecodeFlaggedVector: Wrong constructor %d", constructor)
		return nil
	}
	size := m.Int()
	if m.Err != nil {
		return nil
	}
	if size < 0 {
		m.Err = errors.New("DecodeVector: Wrong size")
		return nil
	}
	x := make([]TL, size)
	i := int32(0)
	for i < size {
		y := m.Object()
		if m.Err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	if __debug&DebugLevelDecodeDetails != 0 {
		fmt.Println("Decode::FlaggedVector::", x)
	}
	return x
}

// FlaggedObject func
func (m *DecodeBuf) FlaggedObject(flags, f int32) (r TL) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return nil
	}
	// constructor := m.UInt()
	constructor := m.Int()
	if m.Err != nil {
		return nil
	}

	switch constructor {

	case crc32ResPQ:
		r = TL_resPQ{m.Bytes(16), m.Bytes(16), m.BigInt(), m.VectorLong()}

	case crc32ServerDHParamsOk:
		r = TL_server_DH_params_ok{m.Bytes(16), m.Bytes(16), m.StringBytes()}

	case crc32ServerDHInnerData:
		r = TL_server_DH_inner_data{
			m.Bytes(16), m.Bytes(16), m.Int(),
			m.BigInt(), m.BigInt(), m.Int(),
		}

	case crc32DHGenOk:
		r = TL_dh_gen_ok{m.Bytes(16), m.Bytes(16), m.Bytes(16)}

	case crc32Ping:
		r = TL_ping{m.Long()}

	case crc32Pong:
		r = TL_pong{m.Long(), m.Long()}

	case crc32MsgContainer:
		size := m.Int()
		arr := make([]TL_MT_message, size)
		for i := int32(0); i < size; i++ {
			arr[i] = TL_MT_message{m.Long(), m.Int(), m.Int(), m.Object()}
			//fmt.Println(constructor, arr[i])
			if m.Err != nil {
				fmt.Println(m.Err.Error())
				return nil
			}
		}
		r = TL_msg_container{arr}

	case crc32RPCResult:
		r = TL_rpc_result{m.Long(), m.Object()}

	case crc32RPCError:
		r = TL_rpc_error{m.Int(), m.String()}

	case crc32NewSessionCreated:
		r = TL_new_session_created{m.Long(), m.Long(), m.Bytes(8)}

	case crc32BadServerSalt:
		r = TL_bad_server_salt{m.Long(), m.Int(), m.Int(), m.Bytes(8)}

	case crc32BadMsgNotification:
		r = TL_crc32BadMsgNotification{m.Long(), m.Int(), m.Int()}

	case crc32MsgsAck:
		r = TL_msgs_ack{m.VectorLong()}

	case crc32GzipPacked:
		obj := make([]byte, 0, 4096)

		var buf bytes.Buffer
		_, _ = buf.Write(m.StringBytes())
		gz, _ := gzip.NewReader(&buf)

		b := make([]byte, 4096)
		for true {
			n, _ := gz.Read(b)
			obj = append(obj, b[0:n]...)
			if n <= 0 {
				break
			}
		}
		d := NewDecodeBuf(obj)
		r = d.Object()

	default:
		r = m.ObjectGenerated(constructor)

	}

	if m.Err != nil {
		return nil
	}
	return
}

// FlaggedStringBytes func
func (m *DecodeBuf) FlaggedStringBytes(flags, f int32) []byte {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return nil
	}
	if m.Err != nil {
		return nil
	}
	var size, padding int

	if m.Off+1 > m.Size {
		m.Err = errors.New("DecodeStringBytes")
		return nil
	}
	size = int(m.Buf[m.Off])
	m.Off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.Off+3 > m.Size {
			m.Err = errors.New("DecodeStringBytes")
			return nil
		}
		size = int(m.Buf[m.Off]) | int(m.Buf[m.Off+1])<<8 | int(m.Buf[m.Off+2])<<16
		m.Off += 3
		padding = (4 - size%4) & 3
	}

	if m.Off+size > m.Size {
		m.Err = errors.New("DecodeStringBytes: Wrong size")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.Buf[m.Off:m.Off+size])
	m.Off += size

	if m.Off+padding > m.Size {
		m.Err = errors.New("DecodeStringBytes: Wrong padding")
		return nil
	}
	m.Off += padding

	if __debug&DebugLevelDecodeDetails != 0 {
		if len(x) > 10 {
			fmt.Println("Decode::FlaggedStringBytes::", len(x), x[:10], " ...")
		} else {
			fmt.Println("Decode::FlaggedStringBytes::", len(x), x)
		}

	}
	return x
}

// dump func
func (m *DecodeBuf) dump() {
	glog.V(3).Info("DecodeBuf dump: ", hex.Dump(m.Buf[m.Off:m.Size]))
}

// toBool func
func ToBool(x TL) bool {
	_, ok := x.(TL_boolTrue)
	return ok
}

// FromBool func
func FromBool(v bool) TL {
	if v {
		return TL_boolTrue{}
	}

	return TL_boolFalse{}
}
