/*
 *
 * Copyright 2014 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpc

import (
	"context"
	"encoding/json"
	"fmt"
)

// Invoke sends the RPC request on the wire and returns after response is
// received.  This is typically called by generated code.
//
// All errors returned by Invoke are compatible with the status package.
func (cc *ClientConn) Invoke(ctx context.Context, method string, args, reply interface{}, opts ...CallOption) error {
	s, _ := json.MarshalIndent(ctx, "", "\t")
	fmt.Printf("Paul - call.go:33 - Invoke - ctx: %v", s)
	s, _ = json.MarshalIndent(method, "", "\t")
	fmt.Printf("Paul - call.go:35 - Invoke - method: %v", s)
	s, _ = json.MarshalIndent(args, "", "\t")
	fmt.Printf("Paul - call.go:37 - Invoke - args: %v", s)
	s, _ = json.MarshalIndent(reply, "", "\t")
	fmt.Printf("Paul - call.go:39 - Invoke - reply: %v", s)
	s, _ = json.MarshalIndent(opts, "", "\t")
	fmt.Printf("Paul - call.go:41 - Invoke - opts: %v", s)
	// allow interceptor to see all applicable call options, which means those
	// configured as defaults from dial option as well as per-call options
	opts = combine(cc.dopts.callOptions, opts)

	if cc.dopts.unaryInt != nil {
		return cc.dopts.unaryInt(ctx, method, args, reply, cc, invoke, opts...)
	}
	return invoke(ctx, method, args, reply, cc, opts...)
}

func combine(o1 []CallOption, o2 []CallOption) []CallOption {
	// we don't use append because o1 could have extra capacity whose
	// elements would be overwritten, which could cause inadvertent
	// sharing (and race conditions) between concurrent calls
	if len(o1) == 0 {
		return o2
	} else if len(o2) == 0 {
		return o1
	}
	ret := make([]CallOption, len(o1)+len(o2))
	copy(ret, o1)
	copy(ret[len(o1):], o2)
	return ret
}

// Invoke sends the RPC request on the wire and returns after response is
// received.  This is typically called by generated code.
//
// DEPRECATED: Use ClientConn.Invoke instead.
func Invoke(ctx context.Context, method string, args, reply interface{}, cc *ClientConn, opts ...CallOption) error {
	return cc.Invoke(ctx, method, args, reply, opts...)
}

var unaryStreamDesc = &StreamDesc{ServerStreams: false, ClientStreams: false}

func invoke(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, opts ...CallOption) error {
	s, _ := json.MarshalIndent(ctx, "", "\t")
	fmt.Printf("Paul - call.go:79 - invoke - ctx: %v", s)
	s, _ = json.MarshalIndent(method, "", "\t")
	fmt.Printf("Paul - call.go:81 - invoke - method: %v", s)
	s, _ = json.MarshalIndent(req, "", "\t")
	fmt.Printf("Paul - call.go:83 - invoke - req: %v", s)
	s, _ = json.MarshalIndent(reply, "", "\t")
	fmt.Printf("Paul - call.go:85 - invoke - reply: %v", s)
	s, _ = json.MarshalIndent(cc, "", "\t")
	fmt.Printf("Paul - call.go:87 - invoke - cc: %v", s)
	s, _ = json.MarshalIndent(opts, "", "\t")
	fmt.Printf("Paul - call.go:89 - invoke - opts: %v", s)
	cs, err := newClientStream(ctx, unaryStreamDesc, cc, method, opts...)
	if err != nil {
		return err
	}
	if err := cs.SendMsg(req); err != nil {
		return err
	}
	return cs.RecvMsg(reply)
}
