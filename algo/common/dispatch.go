package common

import (
	"path/filepath"
	"strings"
)

type NewDecoderFunc func([]byte) DecoderInterface

type decoderItem struct {
	noop    bool
	decoder NewDecoderFunc
}

var DecoderRegistry = make(map[string][]decoderItem)

// RegisterDecoder 注册decoder
//  @param ext 
//  @param noop 
//  @param dispatchFunc 
func RegisterDecoder(ext string, noop bool, dispatchFunc NewDecoderFunc) {
	DecoderRegistry[ext] = append(DecoderRegistry[ext],
		decoderItem{noop: noop, decoder: dispatchFunc})
}

// GetDecoder 实现了从给定文件名中提取文件扩展名，并根据该扩展名以及是否跳过空操作（noop）的条件，
// 从解码器注册表（DecoderRegistry）中选择相应的解码器函数。函数的返回值是一个 NewDecoderFunc 类型的切片，其中包含了满足条件的解码器函数
// noop : No Operation
//  @param filename 
//  @param skipNoop 可能是一个布尔（boolean）类型的参数，用于指示是否跳过 noop 操作。如果 skipNoop 为 true，则在执行某些操作时，会忽略或跳过那些标记为 noop 的函数或方法
//  @return rs NewDecoderFunc
func GetDecoder(filename string, skipNoop bool) (rs []NewDecoderFunc) {
	ext := strings.ToLower(strings.TrimLeft(filepath.Ext(filename), "."))
	for _, dec := range DecoderRegistry[ext] {
		if skipNoop && dec.noop {
			continue
		}
		rs = append(rs, dec.decoder)
	}
	// 由于Go语言支持命名返回值，rs 作为函数的命名返回值，在 return 被调用时会自动使用 rs 变量的当前值作为返回值
	return
}
