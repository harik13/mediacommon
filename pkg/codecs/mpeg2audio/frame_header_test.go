package mpeg2audio

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFrameHeaderUnmarshal(t *testing.T) {
	for _, ca := range []struct {
		name        string
		enc         []byte
		dec         FrameHeader
		sampleCount int
	}{
		{
			"mp2 32k",
			[]byte{
				0xff, 0xfd, 0x48, 0x00, 0x11, 0x11, 0x25, 0x24,
				0x9b, 0x6d, 0xb4, 0x92, 0x24, 0x9a, 0xaa, 0xaa,
				0xaa, 0xaa, 0xaa, 0xaf, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff, 0xf6, 0xb5, 0xad,
				0x6b, 0x5f, 0x3e, 0x7c, 0xf9, 0x6c, 0x5b, 0x16,
				0xc5, 0xb1, 0x6c, 0x5b, 0x1f, 0x3e, 0x7c, 0xf9,
				0xad, 0x6b, 0x5a, 0xd6, 0xb5, 0xad, 0x7c, 0xf9,
				0xf3, 0xe5, 0xb1, 0x6c, 0x5b, 0x16, 0xc5, 0xb1,
				0x6c, 0x7c, 0xf9, 0xf3, 0xe6, 0xb5, 0xad, 0x6b,
				0x5a, 0xd6, 0xb5, 0xf3, 0xe7, 0xcf, 0x96, 0xc5,
				0xb1, 0x6c, 0x5b, 0x16, 0xc5, 0xb1, 0xf3, 0xe7,
				0xcf, 0x9a, 0xd6, 0xb5, 0xad, 0x6b, 0x5a, 0xd7,
				0xcf, 0x9f, 0x3e, 0x5b, 0x16, 0xc5, 0xb1, 0x6c,
				0x5b, 0x16, 0xc7, 0xcf, 0x9f, 0x3e, 0x6b, 0x5a,
				0xd6, 0xb5, 0xad, 0x6b, 0x5f, 0x3e, 0x7c, 0xf9,
				0x6c, 0x5b, 0x16, 0xc5, 0xb1, 0x6c, 0x5b, 0x1f,
				0x3e, 0x7c, 0xf9, 0xad, 0x6b, 0x5a, 0xd6, 0xb5,
				0xad, 0x7c, 0xf9, 0xf3, 0xe5, 0xb1, 0x6c, 0x5b,
				0x16, 0xc5, 0xb1, 0x6c, 0x7c, 0xf9, 0xf3, 0xe6,
				0xb5, 0xad, 0x6b, 0x5a, 0xd6, 0xb5, 0xf3, 0xe7,
				0xcf, 0x96, 0xc5, 0xb1, 0x6c, 0x5b, 0x16, 0xc5,
				0xb1, 0xf3, 0xe7, 0xcf, 0x9a, 0xd6, 0xb5, 0xad,
				0x6b, 0x5a, 0xd7, 0xcf, 0x9f, 0x3e, 0x5b, 0x16,
				0xc5, 0xb1, 0x6c, 0x5b, 0x16, 0xc7, 0xcf, 0x9f,
				0x3e, 0x6b, 0x5a, 0xd6, 0xb5, 0xad, 0x6b, 0x5f,
				0x3e, 0x7c, 0xf9, 0x6c, 0x5b, 0x16, 0xc5, 0xb1,
				0x6c, 0x5b, 0x1f, 0x3e, 0x7c, 0xf9, 0xad, 0x6b,
				0x5a, 0xd6, 0xb5, 0xad, 0x7c, 0xf9, 0xf3, 0xe5,
				0xb1, 0x6c, 0x5b, 0x16, 0xc5, 0xb1, 0x6c, 0x7c,
				0xf9, 0xf3, 0xe6, 0xb5, 0xad, 0x6b, 0x5a, 0xd6,
				0xb5, 0xf3, 0xe7, 0xcf, 0x96, 0xc5, 0xb1, 0x6c,
				0x5b, 0x16, 0xc5, 0xb1, 0xf3, 0xe7, 0xcf, 0x9a,
				0xd6, 0xb5, 0xad, 0x6b, 0x5a, 0xd7, 0xcf, 0x9f,
				0x3e, 0x5b, 0x16, 0xc5, 0xb1, 0x6c, 0x5b, 0x16,
				0xc7, 0xcf, 0x9f, 0x3e, 0x6b, 0x5a, 0xd0, 0x00,
			},
			FrameHeader{
				Layer:       2,
				Bitrate:     64000,
				SampleRate:  32000,
				ChannelMode: ChannelModeStereo,
			},
			1152,
		},
		{
			"mp3 32k",
			[]byte{
				0xff, 0xfb, 0x18, 0x64, 0x00, 0x0f, 0xf0, 0x00,
				0x00, 0x69, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
				0x0d, 0x20, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01,
				0xa4, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x34,
				0x80, 0x00, 0x00, 0x04, 0x4c, 0x41, 0x4d, 0x45,
				0x33, 0x2e, 0x31, 0x30, 0x30, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x4c, 0x41, 0x4d, 0x45, 0x33, 0x2e, 0x31,
				0x30, 0x30, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
			},
			FrameHeader{
				Layer:       3,
				Bitrate:     32000,
				SampleRate:  32000,
				ChannelMode: ChannelModeJointStereo,
			},
			1152,
		},
		{
			"mp3 32k mono",
			[]byte{
				0xff, 0xfb, 0x68, 0xc4, 0xdc, 0x80, 0x16, 0x25,
				0xcd, 0x5f, 0xbd, 0xa5, 0x80, 0x2a, 0x92, 0xbb,
				0x2b, 0xf5, 0xb6, 0x21, 0x3c, 0x7f, 0xdb, 0x79,
				0x6b, 0xab, 0x6f, 0x8d, 0x5f, 0x37, 0xfc, 0xa6,
				0x6c, 0x82, 0xbe, 0xb2, 0x49, 0xb3, 0xe3, 0x55,
				0x53, 0xa1, 0x41, 0x90, 0x38, 0xc6, 0xae, 0xcf,
				0x79, 0x3b, 0xff, 0x33, 0x7c, 0xe7, 0xae, 0x76,
				0xa6, 0x66, 0x36, 0x8a, 0xc6, 0x74, 0xd1, 0xeb,
				0x37, 0x67, 0x4e, 0x86, 0x5c, 0xd9, 0x09, 0x4a,
				0x8a, 0x73, 0x74, 0x6b, 0xa2, 0x33, 0x9b, 0x99,
				0x8e, 0xe8, 0x3b, 0x16, 0xb1, 0x39, 0x24, 0x00,
				0x91, 0x0d, 0x66, 0x93, 0x68, 0x02, 0x9f, 0x33,
				0x40, 0xd0, 0xf6, 0x53, 0x18, 0x19, 0x2f, 0xd3,
				0x36, 0x2b, 0x01, 0xa5, 0x6c, 0xf3, 0x52, 0xab,
				0x2c, 0xf9, 0x5a, 0x2e, 0xd4, 0x64, 0xd0, 0xb8,
				0x03, 0x4a, 0x87, 0xca, 0x6d, 0xc5, 0x43, 0x6a,
				0x1a, 0x57, 0x1f, 0x0f, 0x5b, 0x56, 0xd4, 0x0a,
				0x2b, 0xbf, 0x45, 0xd1, 0xdf, 0x96, 0xf2, 0xef,
				0xda, 0xd9, 0xd6, 0x74, 0x9b, 0x45, 0x38, 0xac,
				0x15, 0x90, 0xce, 0xbf, 0x6a, 0xea, 0x5c, 0x74,
				0x50, 0x8f, 0xb3, 0x23, 0xf9, 0x35, 0x58, 0x90,
				0x1f, 0x4e, 0x7e, 0x59, 0x28, 0x6a, 0xb1, 0x25,
				0x54, 0x8a, 0x1e, 0x6a, 0x12, 0x23, 0x0e, 0x9f,
				0x18, 0x27, 0x56, 0x14, 0x7d, 0x7b, 0xf5, 0x9b,
				0xbe, 0xc5, 0x76, 0x1e, 0x20, 0x90, 0x62, 0xa8,
				0xd3, 0x57, 0xe9, 0x7a, 0xe2, 0xa4, 0x60, 0xfd,
				0xad, 0x94, 0x82, 0x9e, 0xa2, 0x6c, 0x58, 0xe8,
				0x7a, 0x8f, 0xa2, 0x73, 0xab, 0xfb, 0xfe, 0x9b,
				0xf9, 0xbe, 0x2b, 0xfe, 0x7d, 0x62, 0x74, 0xc6,
				0x6a, 0x45, 0xcd, 0x62, 0xe2, 0x34, 0x8f, 0x1e,
				0x91, 0xc1, 0xab, 0x2d, 0x7f, 0xf3, 0xef, 0x30,
				0x70, 0xa0, 0x99, 0x6b, 0xc1, 0xb8, 0x3c, 0x3e,
				0x6f, 0xfa, 0xea, 0x00, 0x14, 0x52, 0xd1, 0x2d,
				0xa1, 0x29, 0x3e, 0x63, 0xa9, 0x01, 0xa3, 0x80,
				0xa1, 0x81, 0xc5, 0x0a, 0x55, 0xb2, 0x26, 0xb9,
				0xa7, 0x5a, 0x5f, 0x0a, 0x6e, 0x49, 0x31, 0xcd,
				0x26, 0x3f, 0x2b, 0x2d, 0x0f, 0xbd, 0x51, 0xe1,
				0x71, 0xd2, 0x61, 0x49, 0x01, 0xab, 0xbe, 0x3f,
				0x42, 0x6e, 0x95, 0x97, 0x73, 0x7f, 0x0e, 0xdd,
				0x38, 0xde, 0x66, 0x27, 0xa0, 0x8d, 0xf8, 0xeb,
				0x75, 0xb0, 0xbf, 0xf8, 0x7d, 0xd7, 0x1a, 0x31,
				0x2d, 0xb6, 0xf9, 0x60, 0xf9, 0x6e, 0xb5, 0xe7,
				0x87, 0x5e, 0x19, 0xa0, 0x3a, 0x38, 0x6d, 0x74,
				0xd9, 0x50, 0x2d, 0x44, 0xd4, 0xa9, 0xec, 0x2e,
				0xa3, 0x1a, 0xc3, 0x4b, 0x3e, 0xc2, 0xc5, 0x7d,
			},
			FrameHeader{
				Layer:       3,
				Bitrate:     80000,
				SampleRate:  32000,
				ChannelMode: ChannelModeMono,
			},
			1152,
		},
		{
			"mp3 44.1k",
			[]byte{
				0xff, 0xfa, 0x52, 0x04, 0xa9, 0xbe, 0xe4, 0x8f,
				0xf0, 0xfd, 0x02, 0xdc, 0x80, 0x00, 0x30, 0x00,
				0x22, 0xc1, 0x5b, 0x90, 0x14, 0x23, 0x24, 0x05,
				0x58, 0x3f, 0x72, 0x02, 0x84, 0xc4, 0xc0, 0xc5,
				0x07, 0xae, 0x40, 0x21, 0xbc, 0x98, 0x90, 0xfa,
				0x3a, 0x2d, 0xda, 0x07, 0xe1, 0x4d, 0xa9, 0x9a,
				0xb8, 0xa2, 0x3b, 0x20, 0xc1, 0xc1, 0xba, 0x08,
				0x94, 0x30, 0x8b, 0xc5, 0x69, 0x51, 0x95, 0xd5,
				0xd7, 0x42, 0x91, 0x65, 0x09, 0xfb, 0x7e, 0x7e,
				0xd9, 0xcf, 0x7f, 0x77, 0x45, 0x03, 0x8d, 0x5c,
				0xcd, 0x52, 0x82, 0x19, 0xbc, 0x94, 0x8c, 0x78,
				0x13, 0xe0, 0x94, 0xc2, 0x96, 0x62, 0x82, 0x20,
				0xb9, 0xf1, 0x3a, 0x05, 0xfa, 0x94, 0x06, 0xbd,
				0xf6, 0x67, 0xa3, 0xca, 0xa5, 0x3a, 0xd5, 0xb5,
				0x34, 0xa9, 0xe8, 0x7e, 0x9f, 0x2f, 0x53, 0xde,
				0x8b, 0xd6, 0x3c, 0x2f, 0x2d, 0xb4, 0x56, 0x0c,
				0xc5, 0x3e, 0x7a, 0xa7, 0x81, 0x5c, 0x35, 0x60,
				0xb3, 0x0c, 0x28, 0x2c, 0x08, 0x06, 0xc0, 0xe0,
				0x3c, 0x0a, 0xfa, 0x1a, 0x6f, 0x43, 0x55, 0xbe,
				0x05, 0x5a, 0x53, 0xae, 0xcb, 0x74, 0xa9, 0xe8,
				0x7e, 0x9f, 0x2f, 0x53, 0xde, 0x8b, 0xd6, 0x20,
				0x36, 0xce, 0xcb, 0xcd, 0x95, 0x15, 0x08, 0xaa,
				0x82, 0x13, 0x51, 0x48, 0xc1, 0x09, 0x28, 0x46,
				0x11, 0x0b, 0x3b, 0x41, 0x34, 0x50, 0x24, 0x18,
				0xa7, 0x72, 0x88, 0x99, 0x49, 0x17, 0x63, 0xac,
				0xa7, 0x98, 0x7e, 0x81, 0x7b, 0x13, 0x9d, 0x7f,
				0xd3,
			},
			FrameHeader{
				Layer:       3,
				Bitrate:     64000,
				SampleRate:  44100,
				Padding:     true,
				ChannelMode: ChannelModeStereo,
			},
			1152,
		},
		{
			"mp3 48k",
			[]byte{
				0xff, 0xfb, 0x14, 0x64, 0x00, 0x0f, 0xf0, 0x00,
				0x00, 0x69, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
				0x0d, 0x20, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01,
				0xa4, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x34,
				0x80, 0x00, 0x00, 0x04, 0x4c, 0x41, 0x4d, 0x45,
				0x33, 0x2e, 0x31, 0x30, 0x30, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0xc0, 0x65, 0xf4, 0xa0, 0x31, 0x8f, 0xce,
				0x8d, 0x46, 0xfc, 0x8c, 0x73, 0xb9, 0x34, 0x3e,
				0xb5, 0x03, 0x39, 0xc0, 0x04, 0x01, 0x98, 0x44,
				0x38, 0xe0, 0x98, 0x10, 0x9b, 0xa8, 0x0f, 0xa8,
			},
			FrameHeader{
				Layer:       3,
				Bitrate:     32000,
				SampleRate:  48000,
				ChannelMode: ChannelModeJointStereo,
			},
			1152,
		},
	} {
		t.Run(ca.name, func(t *testing.T) {
			var h FrameHeader
			err := h.Unmarshal(ca.enc)
			require.NoError(t, err)
			require.Equal(t, ca.dec, h)
			require.Equal(t, len(ca.enc), h.FrameLen())
			require.Equal(t, ca.sampleCount, h.SampleCount())
		})
	}
}

func FuzzFrameHeaderUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var h FrameHeader
		h.Unmarshal(b)
	})
}
