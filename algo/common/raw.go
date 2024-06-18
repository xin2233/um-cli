package common

import (
	"errors"
	"strings"
)

type RawDecoder struct {
	file     []byte
	audioExt string
}

// NewRawDecoder 
//  @param file 
//  @return DecoderInterface 
func NewRawDecoder(file []byte) DecoderInterface {
	return &RawDecoder{file: file}
}

// Validate 
//  @receiver d 
//  @return error 
func (d *RawDecoder) Validate() error {
	for ext, sniffer := range snifferRegistry {
		if sniffer(d.file) {
			d.audioExt = strings.ToLower(ext)
			return nil
		}
	}
	return errors.New("audio doesn't recognized")
}

// Decode 
//  @receiver d 
//  @return error 
func (d RawDecoder) Decode() error {
	return nil
}

// GetCoverImage 
//  @receiver d 
//  @return []byte 
func (d RawDecoder) GetCoverImage() []byte {
	return nil
}

// GetAudioData 
//  @receiver d 
//  @return []byte 
func (d RawDecoder) GetAudioData() []byte {
	return d.file
}

// GetAudioExt 
//  @receiver d 
//  @return string 
func (d RawDecoder) GetAudioExt() string {
	return d.audioExt
}

// GetMeta 
//  @receiver d 
//  @return MetaInterface 
func (d RawDecoder) GetMeta() MetaInterface {
	return nil
}

// init 
func init() {
	RegisterDecoder("mp3", true, NewRawDecoder)
	RegisterDecoder("flac", true, NewRawDecoder)
	RegisterDecoder("ogg", true, NewRawDecoder)
	RegisterDecoder("m4a", true, NewRawDecoder)
	RegisterDecoder("wav", true, NewRawDecoder)
	RegisterDecoder("wma", true, NewRawDecoder)
	RegisterDecoder("aac", true, NewRawDecoder)
}
