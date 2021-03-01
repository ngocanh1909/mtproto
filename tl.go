package mtproto

import (
	"math/big"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

type TL interface {
	Encode() []byte
}

type TL_msg_container struct {
	items []TL_MT_message
}

type TL_MT_message struct {
	msg_id int64
	seq_no int32
	size   int32
	data   interface{}
}

type TL_req_pq struct {
	nonce []byte
}

type TL_p_q_inner_data struct {
	pq           *big.Int
	p            *big.Int
	q            *big.Int
	nonce        []byte
	server_nonce []byte
	new_nonce    []byte
}
type TL_req_DH_params struct {
	nonce        []byte
	server_nonce []byte
	p            *big.Int
	q            *big.Int
	fp           uint64
	encdata      []byte
}
type TL_client_DH_inner_data struct {
	nonce        []byte
	server_nonce []byte
	retry        int64
	g_b          *big.Int
}
type TL_set_client_DH_params struct {
	nonce        []byte
	server_nonce []byte
	encdata      []byte
}
type TL_resPQ struct {
	nonce        []byte
	server_nonce []byte
	pq           *big.Int
	fingerprints []int64
}

type TL_server_DH_params_ok struct {
	nonce            []byte
	server_nonce     []byte
	encrypted_answer []byte
}

type TL_server_DH_inner_data struct {
	nonce        []byte
	server_nonce []byte
	g            int32
	dh_prime     *big.Int
	g_a          *big.Int
	server_time  int32
}

type TL_new_session_created struct {
	first_msg_id int64
	unique_id    int64
	server_salt  []byte
}
type TL_vector_wallpaper struct {
	Wallpapers []TL // Contact
}
type TL_bad_server_salt struct {
	bad_msg_id      int64
	bad_msg_seqno   int32
	error_code      int32
	new_server_salt []byte
}
type TL_crc32BadMsgNotification struct {
	bad_msg_id    int64
	bad_msg_seqno int32
	error_code    int32
}
type TL_crc_bad_msg_notification struct {
	bad_msg_id    int64
	bad_msg_seqno int32
	error_code    int32
}
func GenerateNonce(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}
func GenerateMessageId() int64 {
	const nano = 1000 * 1000 * 1000
	//FIXME: Windows system clock has time resolution issue. https://github.com/golang/go/issues/17696
	//Remove the sleep when the issue is resolved.
	if strings.Contains(runtime.GOOS, "windows") {
		time.Sleep(2 * time.Millisecond)
	}
	unixnano := time.Now().UnixNano()

	return ((unixnano / nano) << 32) | ((unixnano % nano) & -4)
}
type TL_msgs_ack struct {
	msgIds []int64
}

type TL_rpc_result struct {
	req_msg_id int64
	obj        interface{}
}

type TL_rpc_error struct {
	error_code    int32
	error_message string
}

type TL_dh_gen_ok struct {
	nonce           []byte
	server_nonce    []byte
	new_nonce_hash1 []byte
}

type TL_ping struct {
	ping_id int64
}

type TL_pong struct {
	msg_id  int64
	ping_id int64
}
