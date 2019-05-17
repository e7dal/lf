/*
 * LF: Global Fully Replicated Key/Value Store
 * Copyright (C) 2018-2019  ZeroTier, Inc.  https://www.zerotier.com/
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 *
 * --
 *
 * You can be released from the requirements of the license by purchasing
 * a commercial license. Buying such a license is mandatory as soon as you
 * develop commercial closed-source software that incorporates or links
 * directly against ZeroTier software without disclosing the source code
 * of your own application.
 */

package lf

// This contains the defaults for the Sol LF network, the global shared LF
// database for Earth and its neighbors in the Sol system. It should be good
// up to Kardashev type II civilization scale.

// SolDefaultNodeURL is the default URL for clients to access Sol (servers operated by ZeroTier, Inc.)
const SolDefaultNodeURL = "https://lf.zerotier.com/"

// SolGenesisRecords are the genesis records for initializing a Sol member node.
var SolGenesisRecords = []byte{0x1, 0x1, 0x87, 0x8, 0xbe, 0x17, 0xb8, 0x7f, 0x40, 0x69, 0x26, 0x21, 0xbd, 0x7, 0x5, 0x47, 0x7f, 0x1, 0xd3, 0x74, 0x1c, 0x0, 0x12, 0x79, 0x29, 0xed, 0x28, 0x53, 0x51, 0x2f, 0x12, 0xe8, 0xc2, 0x7e, 0x74, 0x8a, 0xcd, 0xd0, 0x8c, 0x67, 0x7f, 0x49, 0xf4, 0x5d, 0x78, 0x3f, 0x7f, 0x27, 0xb7, 0x2f, 0x48, 0xf6, 0x8, 0x4c, 0x74, 0xd, 0x9d, 0x38, 0x52, 0x55, 0xf7, 0xbd, 0x64, 0xa6, 0x64, 0x27, 0xd0, 0x96, 0x91, 0x5e, 0xab, 0xe8, 0xda, 0x21, 0xae, 0x58, 0x43, 0x44, 0xde, 0xe1, 0xf9, 0x6f, 0xa9, 0x9a, 0xb8, 0x65, 0x1c, 0xa8, 0x14, 0x70, 0xfc, 0x70, 0x25, 0x5d, 0x5e, 0xdf, 0xab, 0x9, 0x43, 0x8, 0x11, 0x41, 0x4, 0xbe, 0xfd, 0xbd, 0x82, 0x25, 0x3f, 0x51, 0x3, 0xdb, 0x85, 0x42, 0x9, 0x37, 0x8f, 0x5d, 0xf5, 0xae, 0x71, 0x87, 0x47, 0x1d, 0x4, 0x10, 0x5d, 0x3d, 0xf2, 0x8d, 0xdb, 0xd, 0xad, 0x7c, 0x8e, 0x2b, 0x44, 0x6d, 0xda, 0x4a, 0x34, 0xf0, 0xc7, 0x89, 0x26, 0x3b, 0xde, 0xe2, 0xe8, 0x6, 0x6e, 0x8f, 0x1a, 0xf1, 0xf1, 0xc0, 0x45, 0xd9, 0x85, 0x9a, 0x3f, 0x52, 0x99, 0x83, 0xcf, 0x55, 0xda, 0xa0, 0x3d, 0xeb, 0x6b, 0xbe, 0x2f, 0x60, 0xb0, 0x77, 0x2d, 0x88, 0x4b, 0xc5, 0xe6, 0x57, 0x8e, 0xf2, 0xf7, 0x7d, 0xe6, 0xbc, 0x68, 0xba, 0xe6, 0x43, 0xbd, 0xe0, 0x47, 0x1d, 0x4b, 0x8e, 0xa7, 0xd9, 0x66, 0xfe, 0xec, 0x32, 0x97, 0x1, 0xd5, 0x98, 0xb5, 0xfc, 0xe1, 0xf8, 0xb1, 0x4b, 0x86, 0x5, 0xca, 0x51, 0x8c, 0x21, 0x8d, 0x4d, 0x20, 0x2f, 0x7a, 0x5a, 0x88, 0x60, 0x3f, 0xa1, 0x9d, 0xaa, 0x2a, 0x14, 0xf, 0x24, 0x7f, 0x69, 0xb7, 0xa5, 0xf9, 0x5a, 0x41, 0xc9, 0x8b, 0xfc, 0xa9, 0xe5, 0x49, 0x95, 0x34, 0x2, 0x13, 0xe6, 0x65, 0x3f, 0xa1, 0x7f, 0x1f, 0x9c, 0xd6, 0xa0, 0x5c, 0x8d, 0x98, 0x69, 0x7f, 0xc5, 0x41, 0x3d, 0x6a, 0x69, 0x15, 0x66, 0x7d, 0x88, 0x94, 0x94, 0xfe, 0xa4, 0x6b, 0x3f, 0x39, 0x4b, 0xcb, 0xc1, 0x32, 0x51, 0x71, 0xde, 0x5b, 0xcf, 0x17, 0xbb, 0x54, 0x1b, 0x91, 0xef, 0x5, 0x43, 0xb3, 0x90, 0x38, 0x2d, 0xa3, 0x20, 0x6d, 0x1a, 0xa0, 0xb6, 0xda, 0xf1, 0xa3, 0x7a, 0x60, 0xe3, 0x14, 0x90, 0x6, 0x79, 0xed, 0x80, 0xf0, 0x76, 0x58, 0x65, 0x5f, 0x33, 0xe7, 0xe4, 0x8b, 0xa2, 0x26, 0x78, 0x7b, 0x3, 0xd5, 0xd5, 0x84, 0x2, 0x73, 0x2b, 0x9, 0x48, 0xdf, 0xba, 0x52, 0x9, 0x45, 0xc1, 0x90, 0x52, 0xf7, 0xf3, 0xe3, 0x6c, 0x90, 0xcb, 0xb6, 0x7f, 0x75, 0x80, 0xb0, 0x2f, 0x2, 0xdf, 0xb, 0x89, 0x56, 0xec, 0x75, 0xab, 0x3c, 0x35, 0x8, 0x94, 0xf3, 0xb9, 0xfb, 0x91, 0xaf, 0xdb, 0x19, 0x45, 0x25, 0xf, 0x86, 0xf8, 0x44, 0x10, 0x9f, 0x47, 0xc2, 0xfb, 0x60, 0x2b, 0xe0, 0x6b, 0x83, 0xe9, 0x22, 0x76, 0xd5, 0xb7, 0x47, 0xb1, 0x4f, 0xb, 0x42, 0xb3, 0x87, 0xb5, 0x41, 0x4e, 0xd3, 0xa8, 0x1f, 0x44, 0x80, 0x57, 0x3f, 0xa8, 0x87, 0x30, 0x7c, 0x50, 0x45, 0xd1, 0xb8, 0xe9, 0xfc, 0xe4, 0xe2, 0x5, 0x43, 0x11, 0x46, 0x71, 0x4, 0x79, 0x24, 0xe8, 0x4a, 0x72, 0xf8, 0x6d, 0x99, 0x3c, 0x9b, 0x29, 0x62, 0x40, 0xac, 0x67, 0xaa, 0x4e, 0x55, 0x12, 0x8c, 0x63, 0xd, 0x3c, 0xf0, 0xd4, 0x66, 0x89, 0x20, 0x4d, 0x81, 0x8e, 0x30, 0xc5, 0x5c, 0x72, 0x3f, 0x40, 0x73, 0xf5, 0x3b, 0xd6, 0x64, 0x82, 0xf7, 0x96, 0xb1, 0xde, 0x5, 0xa5, 0x25, 0xc, 0xc1, 0x1b, 0xd6, 0x94, 0x8a, 0x50, 0x2d, 0xfd, 0xc, 0xae, 0xde, 0xd4, 0x6f, 0x77, 0x66, 0x10, 0x2e, 0x74, 0xef, 0x58, 0x5a, 0x8e, 0x5c, 0x36, 0x18, 0x82, 0xe, 0x4d, 0xbd, 0xf5, 0x7e, 0x32, 0xdd, 0xdd, 0x9f, 0xcb, 0x8d, 0xbd, 0xec, 0x1e, 0x77, 0x31, 0x72, 0xb6, 0xc8, 0xd, 0x8, 0x74, 0x0, 0xd6, 0x32, 0x84, 0x80, 0xc1, 0xb0, 0x3c, 0x7a, 0x1d, 0x82, 0x37, 0x2c, 0x81, 0x14, 0xcd, 0x94, 0x11, 0x21, 0x2f, 0x94, 0x41, 0xf, 0x6e, 0x1b, 0xe2, 0xf5, 0xb2, 0x9b, 0xe3, 0x53, 0x78, 0x7b, 0x0, 0xd7, 0x41, 0xc8, 0x6c, 0xcb, 0x45, 0x54, 0xec, 0xc5, 0x51, 0x2a, 0x1f, 0x31, 0x63, 0xf4, 0x55, 0x6b, 0xec, 0x71, 0x36, 0xb6, 0xc4, 0x93, 0x2f, 0x54, 0xca, 0xce, 0x18, 0x2, 0xce, 0x29, 0x14, 0x6e, 0xed, 0xd, 0x76, 0x86, 0x76, 0xbb, 0x54, 0x1b, 0x1b, 0x51, 0xa0, 0xdb, 0x87, 0x18, 0x33, 0x7f, 0xa3, 0x5c, 0xfc, 0x65, 0x20, 0x5e, 0x43, 0xd4, 0xdd, 0xc2, 0x45, 0xfb, 0x26, 0xb1, 0x6c, 0xa5, 0xd2, 0x5b, 0x45, 0x60, 0x77, 0xf8, 0x8a, 0x6a, 0x4d, 0xee, 0xe3, 0xfa, 0x96, 0xa4, 0xf3, 0xcf, 0x31, 0xad, 0x26, 0x20, 0xdf, 0x4a, 0x20, 0xd6, 0xe9, 0x34, 0x11, 0xfe, 0x52, 0xcd, 0xfb, 0xd9, 0x4, 0xd2, 0xd8, 0x90, 0x90, 0xf3, 0xd0, 0xef, 0xbd, 0x53, 0x3f, 0x57, 0x2e, 0xb6, 0xc2, 0xf0, 0x12, 0x92, 0x62, 0xf6, 0x5e, 0x34, 0xd0, 0x49, 0x1c, 0x49, 0x60, 0x79, 0xa3, 0x37, 0x9, 0x40, 0x87, 0x87, 0x59, 0x4e, 0x1f, 0xd, 0x90, 0xd3, 0x55, 0xb7, 0x76, 0x4d, 0x90, 0x30, 0xaf, 0x99, 0xc2, 0xde, 0x14, 0x43, 0xe0, 0x53, 0x22, 0xe1, 0xfe, 0x60, 0x2a, 0x81, 0xe7, 0x26, 0xcc, 0x51, 0x99, 0xb, 0x60, 0xfc, 0xaf, 0x71, 0x14, 0x4a, 0x3d, 0x8f, 0x21, 0xed, 0x97, 0x89, 0x23, 0x3e, 0x8c, 0xe0, 0xdf, 0x48, 0x8, 0x57, 0x1a, 0x90, 0xe2, 0xf4, 0x7c, 0x3d, 0xcf, 0x7a, 0x67, 0x11, 0xe2, 0xc4, 0x41, 0x57, 0x18, 0x99, 0x3d, 0xcc, 0x58, 0xa, 0x68, 0x43, 0xb6, 0x5a, 0x73, 0xc, 0x83, 0xe3, 0x6, 0x24, 0xc0, 0x17, 0xe4, 0x22, 0xc2, 0x4d, 0x94, 0xab, 0x59, 0x2e, 0x99, 0xba, 0x85, 0xf8, 0xa0, 0x62, 0x1d, 0x16, 0x9a, 0xcb, 0xed, 0xa, 0x1a, 0xac, 0x94, 0x55, 0x8c, 0x27, 0xed, 0xd7, 0xf0, 0xcd, 0x7f, 0xf4, 0x14, 0x15, 0xac, 0x9d, 0xe7, 0xea, 0x23, 0x7e, 0x95, 0x2d, 0xd3, 0xd2, 0x58, 0x47, 0x22, 0xa0, 0x96, 0xb1, 0x9b, 0xf5, 0xde, 0x49, 0x65, 0x8b, 0xfe, 0x63, 0x93, 0xe2, 0xc6, 0xec, 0x6d, 0x1f, 0x10, 0xbd, 0x78, 0x8, 0x85, 0x5e, 0x1b, 0xdf, 0x54, 0x21, 0x80, 0xbf, 0xb5, 0xa2, 0x4e, 0x6, 0x1f, 0xd2, 0x2a, 0x6d, 0xe1, 0x11, 0x16, 0x3e, 0x13, 0x4, 0x2f, 0x6e, 0x18, 0xd3, 0xb9, 0x3c, 0xea, 0xe5, 0x4e, 0xf1, 0xd8, 0xdd, 0x8b, 0xfd, 0x75, 0x99, 0xbd, 0x3e, 0x98, 0xde, 0xf1, 0xa5, 0xaf, 0x56, 0x77, 0x50, 0xd, 0xdc, 0x61, 0x93, 0x59, 0x22, 0x64, 0x22, 0x82, 0xd6, 0xcf, 0x28, 0x2f, 0x89, 0x54, 0x42, 0x48, 0xae, 0xe6, 0x97, 0xdd, 0xd4, 0x3b, 0xce, 0x2d, 0xcc, 0xf6, 0x36, 0xb2, 0x9b, 0x82, 0x37, 0xc5, 0xa7, 0x6d, 0xe4, 0x84, 0xa7, 0xd3, 0x61, 0x43, 0x2f, 0x3a, 0xf1, 0x50, 0xf0, 0x55, 0x10, 0xea, 0x0, 0xd6, 0x45, 0x9e, 0x55, 0x35, 0x67, 0x92, 0xa5, 0xda, 0xd9, 0xf9, 0x82, 0x8e, 0x67, 0xfe, 0xc0, 0xab, 0xf9, 0xd8, 0x2b, 0x78, 0xc2, 0xff, 0xce, 0xf5, 0x5a, 0xde, 0x1f, 0x83, 0xb3, 0xb5, 0x48, 0x57, 0xe5, 0xef, 0x55, 0xf0, 0x9, 0x4a, 0x2a, 0xb4, 0xe6, 0x14, 0xc8, 0xcc, 0x5f, 0xdd, 0xe9, 0xcc, 0x47, 0xf, 0xa3, 0xa3, 0x51, 0x6, 0x6c, 0x7b, 0xc2, 0xb6, 0xb4, 0x60, 0x95, 0x7e, 0xdb, 0x45, 0x3f, 0x47, 0x8b, 0x6f, 0xc, 0xa9, 0x60, 0xee, 0x71, 0x68, 0x17, 0xbd, 0xdd, 0x4a, 0xa, 0x4, 0x99, 0x18, 0xba, 0xfe, 0x69, 0x62, 0x8a, 0x80, 0x9e, 0xcf, 0xa7, 0x9b, 0xcb, 0x4a, 0x49, 0xa2, 0x8d, 0x17, 0xa4, 0xa1, 0x78, 0x23, 0xd0, 0x24, 0x86, 0x64, 0x0, 0x9a, 0xa0, 0xfc, 0xe6, 0x5, 0x0, 0x1, 0xf, 0x5f, 0x85, 0xb7, 0x2f, 0xf, 0x5f, 0x72, 0xe5, 0xa9, 0x0, 0x7, 0xf4, 0x3a, 0x60, 0x68, 0xcd, 0xc0, 0x26, 0xee, 0x49, 0x23, 0xb, 0xf0, 0xe9, 0x99, 0xbb, 0x95, 0xeb, 0xb6, 0x15, 0x3e, 0x39, 0xd7, 0x2d, 0x8a, 0x13, 0xbf, 0x6d, 0x7f, 0x80, 0x44, 0xba, 0x91, 0xef, 0xa1, 0x64, 0xe1, 0x91, 0x5d, 0xd8, 0xa, 0xf7, 0x4, 0x42, 0xea, 0x27, 0x97, 0x5e, 0x83, 0x86, 0x17, 0x7b, 0xc7, 0x8, 0xc6, 0xb9, 0xc7, 0x6d, 0x65, 0xc7, 0x57, 0x7e, 0xa8, 0x5e, 0xe4, 0x6c, 0x4e, 0x1, 0x80, 0xb0, 0xe2, 0xc3, 0xe, 0xf7, 0x79, 0x3c, 0x79, 0xb5, 0xae, 0xfb, 0xb1, 0xc, 0x64, 0xb8, 0x9, 0x20, 0x2d, 0x89, 0x43, 0x7b, 0x44, 0xe0, 0xdd, 0xba, 0x83, 0xc7, 0xd5, 0x3a, 0x7f, 0xb0, 0x1, 0x1, 0x0, 0x18, 0xba, 0xfe, 0x69, 0x62, 0x8a, 0x80, 0x9e, 0xcf, 0xa7, 0x9b, 0xcb, 0x4a, 0x49, 0xa2, 0x8d, 0x17, 0xa4, 0xa1, 0x78, 0x23, 0xd0, 0x24, 0x86, 0x64, 0x1, 0xd9, 0x81, 0x9c, 0x42, 0x50, 0x74, 0x9a, 0x44, 0x9e, 0xb1, 0xb1, 0x6b, 0x2, 0x1b, 0xfa, 0x6e, 0xd9, 0x48, 0x20, 0x96, 0xc3, 0xe, 0x4b, 0x83, 0xd4, 0x1b, 0x1, 0x13, 0x14, 0xf1, 0x0, 0x1, 0x9b, 0xa0, 0xfc, 0xe6, 0x5, 0x0, 0x1, 0x81, 0x49, 0xe0, 0x97, 0x92, 0x81, 0x49, 0xd6, 0xbe, 0x9e, 0x0, 0x0, 0x1a, 0x4c, 0x60, 0x4c, 0x55, 0x59, 0x96, 0x63, 0xd4, 0x84, 0xc7, 0x3d, 0x37, 0x13, 0xc7, 0xbe, 0x58, 0xab, 0x91, 0xaa, 0x59, 0x6d, 0x47, 0x3, 0x1, 0xce, 0x83, 0xe9, 0x4a, 0x93, 0xc7, 0x1b, 0x53, 0xd8, 0xe9, 0xb6, 0xae, 0xd2, 0xd7, 0x9c, 0x2b, 0xa2, 0x5c, 0x2c, 0xbe, 0x70, 0xa8, 0xe7, 0x8d, 0x83, 0x2a, 0xf4, 0x12, 0x2d, 0x32, 0xc0, 0x78, 0xf7, 0xf4, 0xbb, 0x74, 0x7a, 0xd9, 0x4c, 0x3c, 0x15, 0x7, 0x6d, 0xcf, 0x25, 0xc0, 0xb1, 0x63, 0xeb, 0xa6, 0x13, 0xa5, 0x24, 0x78, 0x41, 0x8f, 0x70, 0x60, 0x27, 0x61, 0xf2, 0x20, 0x44, 0xdd, 0x90, 0x6f, 0xba, 0x2c, 0x9b, 0xcb, 0x4e, 0x82, 0x89, 0x37}
