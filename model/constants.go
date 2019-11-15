package model

type Sizes uint32
type ArrayTable uint32
type TextureShadowCompareFunc uint32
type TextureShadowTextureMode uint32

const (
	INDENT_VALUE uint32 = 2
	BOOL_SIZE    Sizes  = 1
	CHAR_SIZE    Sizes  = 1
	SHORT_SIZE   Sizes  = 2
	INT_SIZE     Sizes  = 4
	LONG_SIZE    Sizes  = 4
	INT64_SIZE   Sizes  = 8
	FLOAT_SIZE   Sizes  = 4
	DOUBLE_SIZE  Sizes  = 8
	GLENUM_SIZE  Sizes  = 4

	ID_BYTE_ARRAY   = 0
	ID_UBYTE_ARRAY  = 1
	ID_SHORT_ARRAY  = 2
	ID_USHORT_ARRAY = 3
	ID_INT_ARRAY    = 4
	ID_UINT_ARRAY   = 5
	ID_FLOAT_ARRAY  = 6
	ID_DOUBLE_ARRAY = 7
	ID_VEC2B_ARRAY  = 8
	ID_VEC3B_ARRAY  = 9
	ID_VEC4B_ARRAY  = 10
	ID_VEC4UB_ARRAY = 11
	ID_VEC2S_ARRAY  = 12
	ID_VEC3S_ARRAY  = 13
	ID_VEC4S_ARRAY  = 14
	ID_VEC2_ARRAY   = 15
	ID_VEC3_ARRAY   = 16
	ID_VEC4_ARRAY   = 17
	ID_VEC2D_ARRAY  = 18
	ID_VEC3D_ARRAY  = 19
	ID_VEC4D_ARRAY  = 20
	ID_VEC2UB_ARRAY = 21
	ID_VEC3UB_ARRAY = 22
	ID_VEC2US_ARRAY = 23
	ID_VEC3US_ARRAY = 24
	ID_VEC4US_ARRAY = 25

	ID_VEC2I_ARRAY  = 26
	ID_VEC3I_ARRAY  = 27
	ID_VEC4I_ARRAY  = 28
	ID_VEC2UI_ARRAY = 29
	ID_VEC3UI_ARRAY = 30
	ID_VEC4UI_ARRAY = 31

	ID_UINT64_ARRAY = 32
	ID_INT64_ARRAY  = 33

	ID_DRAWARRAYS            = 50
	ID_DRAWARRAY_LENGTH      = 51
	ID_DRAWELEMENTS_UBYTE    = 52
	ID_DRAWELEMENTS_USHORT   = 53
	ID_DRAWELEMENTS_UINT     = 54
	GL_ALPHA_TEST            = 0x0BC0
	GL_BLEND                 = 0x0BE2
	GL_COLOR_LOGIC_OP        = 0x0BF2
	GL_COLOR_MATERIAL        = 0x0B57
	GL_CULL_FACE             = 0x0B44
	GL_DEPTH_TEST            = 0x0B71
	GL_FOG                   = 0x0B60
	GL_FRAGMENT_PROGRAM_ARB  = 0x8804
	GL_LINE_STIPPLE          = 0x0B24
	GL_POINT_SMOOTH          = 0x0B10
	GL_POINT_SPRITE_ARB      = 0x8861
	GL_POLYGON_OFFSET_FILL   = 0x8037
	GL_POLYGON_OFFSET_LINE   = 0x2A02
	GL_POLYGON_OFFSET_POINT  = 0x2A01
	GL_POLYGON_STIPPLE       = 0x0B42
	GL_SCISSOR_TEST          = 0x0C11
	GL_STENCIL_TEST          = 0x0B90
	GL_STENCIL_TEST_TWO_SIDE = 0x8910
	GL_VERTEX_PROGRAM_ARB    = 0x8620

	GL_COLOR_SUM      = 0x8458
	GL_LIGHTING       = 0x0B50
	GL_NORMALIZE      = 0x0BA1
	GL_RESCALE_NORMAL = 0x803A

	GL_TEXTURE_1D        = 0x0DE0
	GL_TEXTURE_2D        = 0x0DE1
	GL_TEXTURE_3D        = 0x806F
	GL_TEXTURE_BUFFER    = 0x8C2A
	GL_TEXTURE_CUBE_MAP  = 0x8513
	GL_TEXTURE_RECTANGLE = 0x84F5
	GL_TEXTURE_GEN_Q     = 0x0C63
	GL_TEXTURE_GEN_R     = 0x0C62
	GL_TEXTURE_GEN_S     = 0x0C60
	GL_TEXTURE_GEN_T     = 0x0C61
	GL_TEXTURE_2D_ARRAY  = 0x8C1A

	GL_CLIP_PLANE0 = 0x3000
	GL_CLIP_PLANE1 = 0x3001
	GL_CLIP_PLANE2 = 0x3002
	GL_CLIP_PLANE3 = 0x3003
	GL_CLIP_PLANE4 = 0x3004
	GL_CLIP_PLANE5 = 0x3005

	GL_LIGHT0 = 0x4000
	GL_LIGHT1 = 0x4001
	GL_LIGHT2 = 0x4002
	GL_LIGHT3 = 0x4003
	GL_LIGHT4 = 0x4004
	GL_LIGHT5 = 0x4005
	GL_LIGHT6 = 0x4006
	GL_LIGHT7 = 0x4007

	GL_VERTEX_PROGRAM_POINT_SIZE = 0x8642
	GL_VERTEX_PROGRAM_TWO_SIDE   = 0x8643

	// Functions
	GL_NEVER    = 0x0200
	GL_LESS     = 0x0201
	GL_EQUAL    = 0x0202
	GL_LEQUAL   = 0x0203
	GL_GREATER  = 0x0204
	GL_NOTEQUAL = 0x0205
	GL_GEQUAL   = 0x0206
	GL_ALWAYS   = 0x0207

	// Texture environment states
	GL_REPLACE     = 0x1E01
	GL_MODULATE    = 0x2100
	GL_ADD         = 0x0104
	GL_ADD_SIGNED  = 0x8574
	GL_INTERPOLATE = 0x8575
	GL_SUBTRACT    = 0x84E7
	GL_DOT3_RGB    = 0x86AE
	GL_DOT3_RGBA   = 0x86AF

	GL_CONSTANT      = 0x8576
	GL_PRIMARY_COLOR = 0x8577
	GL_PREVIOUS      = 0x8578
	GL_TEXTURE       = 0x1702
	GL_TEXTURE0      = 0x84C0
	GL_TEXTURE1      = 0x84C1
	GL_TEXTURE2      = 0x84C2
	GL_TEXTURE3      = 0x84C3
	GL_TEXTURE4      = 0x84C4
	GL_TEXTURE5      = 0x84C5
	GL_TEXTURE6      = 0x84C6
	GL_TEXTURE7      = 0x84C7

	GL_COMBINE_ARB        = 0x8570
	GL_COMBINE_RGB_ARB    = 0x8571
	GL_COMBINE_ALPHA_ARB  = 0x8572
	GL_SOURCE0_RGB_ARB    = 0x8580
	GL_SOURCE1_RGB_ARB    = 0x8581
	GL_SOURCE2_RGB_ARB    = 0x8582
	GL_SOURCE0_ALPHA_ARB  = 0x8588
	GL_SOURCE1_ALPHA_ARB  = 0x8589
	GL_SOURCE2_ALPHA_ARB  = 0x858A
	GL_OPERAND0_RGB_ARB   = 0x8590
	GL_OPERAND1_RGB_ARB   = 0x8591
	GL_OPERAND2_RGB_ARB   = 0x8592
	GL_OPERAND0_ALPHA_ARB = 0x8598
	GL_OPERAND1_ALPHA_ARB = 0x8599
	GL_OPERAND2_ALPHA_ARB = 0x859A
	GL_RGB_SCALE_ARB      = 0x8573
	GL_ADD_SIGNED_ARB     = 0x8574
	GL_INTERPOLATE_ARB    = 0x8575
	GL_SUBTRACT_ARB       = 0x84E7
	GL_CONSTANT_ARB       = 0x8576
	GL_PRIMARY_COLOR_ARB  = 0x8577
	GL_PREVIOUS_ARB       = 0x8578

	GL_DOT3_RGB_ARB  = 0x86AE
	GL_DOT3_RGBA_ARB = 0x86AF

	// Texture clamp modes
	GL_CLAMP               = 0x2900
	GL_CLAMP_TO_EDGE       = 0x812F
	GL_CLAMP_TO_BORDER     = 0x812D
	GL_REPEAT              = 0x2901
	GL_MIRROR              = 0x8370
	GL_CLAMP_TO_BORDER_ARB = 0x812D
	GL_MIRRORED_REPEAT_IBM = 0x8370

	// Texture filter modes
	GL_LINEAR                 = 0x2601
	GL_LINEAR_MIPMAP_LINEAR   = 0x2703
	GL_LINEAR_MIPMAP_NEAREST  = 0x2701
	GL_NEAREST                = 0x2600
	GL_NEAREST_MIPMAP_LINEAR  = 0x2702
	GL_NEAREST_MIPMAP_NEAREST = 0x2700

	// Texture formats
	GL_INTENSITY                                 = 0x8049
	GL_LUMINANCE                                 = 0x1909
	GL_ALPHA                                     = 0x1906
	GL_LUMINANCE_ALPHA                           = 0x190A
	GL_RGB                                       = 0x1907
	GL_RGBA                                      = 0x1908
	GL_COMPRESSED_ALPHA_ARB                      = 0x84E9
	GL_COMPRESSED_LUMINANCE_ARB                  = 0x84EA
	GL_COMPRESSED_INTENSITY_ARB                  = 0x84EC
	GL_COMPRESSED_LUMINANCE_ALPHA_ARB            = 0x84EB
	GL_COMPRESSED_RGB_ARB                        = 0x84ED
	GL_COMPRESSED_RGBA_ARB                       = 0x84EE
	GL_COMPRESSED_RGB_S3TC_DXT1_EXT              = 0x83F0
	GL_COMPRESSED_RGBA_S3TC_DXT1_EXT             = 0x83F1
	GL_COMPRESSED_RGBA_S3TC_DXT3_EXT             = 0x83F2
	GL_COMPRESSED_RGBA_S3TC_DXT5_EXT             = 0x83F3
	GL_COMPRESSED_RGB_PVRTC_4BPPV1_IMG           = 0x8C00
	GL_COMPRESSED_RGB_PVRTC_2BPPV1_IMG           = 0x8C01
	GL_COMPRESSED_RGBA_PVRTC_4BPPV1_IMG          = 0x8C02
	GL_COMPRESSED_RGBA_PVRTC_2BPPV1_IMG          = 0x8C03
	GL_ETC1_RGB8_OES                             = 0x8D64
	GL_COMPRESSED_RGB8_ETC2                      = 0x9274
	GL_COMPRESSED_SRGB8_ETC2                     = 0x9275
	GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2  = 0x9276
	GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2 = 0x9277
	GL_COMPRESSED_RGBA8_ETC2_EAC                 = 0x9278
	GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC          = 0x9279
	GL_COMPRESSED_R11_EAC                        = 0x9270
	GL_COMPRESSED_SIGNED_R11_EAC                 = 0x9271
	GL_COMPRESSED_RG11_EAC                       = 0x9272
	GL_COMPRESSED_SIGNED_RG11_EAC                = 0x9273

	// Texture source types
	GL_BYTE           = 0x1400
	GL_SHORT          = 0x1402
	GL_INT            = 0x1404
	GL_FLOAT          = 0x1406
	GL_DOUBLE         = 0x140A
	GL_UNSIGNED_BYTE  = 0x1401
	GL_UNSIGNED_SHORT = 0x1403
	GL_UNSIGNED_INT   = 0x1405

	// Blend values
	GL_DST_ALPHA                = 0x0304
	GL_DST_COLOR                = 0x0306
	GL_ONE                      = 1
	GL_ONE_MINUS_DST_ALPHA      = 0x0305
	GL_ONE_MINUS_DST_COLOR      = 0x0307
	GL_ONE_MINUS_SRC_ALPHA      = 0x0303
	GL_ONE_MINUS_SRC_COLOR      = 0x0301
	GL_SRC_ALPHA                = 0x0302
	GL_SRC_ALPHA_SATURATE       = 0x0308
	GL_SRC_COLOR                = 0x0300
	GL_CONSTANT_COLOR           = 0x8001
	GL_ONE_MINUS_CONSTANT_COLOR = 0x8002
	GL_CONSTANT_ALPHA           = 0x8003
	GL_ONE_MINUS_CONSTANT_ALPHA = 0x8004
	GL_ZERO                     = 0

	// Fog coordinate sources
	GL_COORDINATE = 0x8451
	GL_DEPTH      = 0x8452

	GL_GENERATE_MIPMAP_SGIS      = 0x8191
	GL_GENERATE_MIPMAP_HINT_SGIS = 0x8192

	GL_TEXTURE_COMPRESSION_HINT_ARB = 0x84EF

	// Hint targets
	GL_FOG_HINT                        = 0x0C54
	GL_GENERATE_MIPMAP_HINT            = 0x8192
	GL_LINE_SMOOTH_HINT                = 0x0C52
	GL_PERSPECTIVE_CORRECTION_HINT     = 0x0C50
	GL_POINT_SMOOTH_HINT               = 0x0C51
	GL_POLYGON_SMOOTH_HINT             = 0x0C53
	GL_TEXTURE_COMPRESSION_HINT        = 0x84EF
	GL_FRAGMENT_SHADER_DERIVATIVE_HINT = 0x8B8B

	GL_FOG_COORDINATE = 0x8451
	GL_FRAGMENT_DEPTH = 0x8452

	// Polygon modes
	GL_POINT = 0x1B00
	GL_LINE  = 0x1B01
	GL_FILL  = 0x1B00

	// Misc
	GL_BACK           = 0x0405
	GL_FRONT          = 0x0404
	GL_FRONT_AND_BACK = 0x0408
	GL_FIXED_ONLY     = 0x891D
	GL_FASTEST        = 0x1101
	GL_NICEST         = 0x1101
	GL_DONT_CARE      = 0x1100

	GL_OBJECT_LINEAR = 0x2401
	GL_EYE_LINEAR    = 0x2400
	GL_SPHERE_MAP    = 0x2402

	GL_NORMAL_MAP = 0x8511

	GL_REFLECTION_MAP = 0x8512

	GL_RED   = 0x1903
	GL_GREEN = 0x1904
	GL_BLUE  = 0x1905

	ArrayType ArrayTable = 0

	ByteArrayType  ArrayTable = 1
	ShortArrayType ArrayTable = 2
	IntArrayType   ArrayTable = 3

	UByteArrayType  ArrayTable = 4
	UShortArrayType ArrayTable = 5
	UIntArrayType   ArrayTable = 6

	FloatArrayType  ArrayTable = 7
	DoubleArrayType ArrayTable = 8

	Vec2bArrayType ArrayTable = 9
	Vec3bArrayType ArrayTable = 10
	Vec4bArrayType ArrayTable = 11

	Vec2sArrayType ArrayTable = 12
	Vec3sArrayType ArrayTable = 13
	Vec4sArrayType ArrayTable = 14

	Vec2iArrayType ArrayTable = 15
	Vec3iArrayType ArrayTable = 16
	Vec4iArrayType ArrayTable = 17

	Vec2ubArrayType ArrayTable = 18
	Vec3ubArrayType ArrayTable = 19
	Vec4ubArrayType ArrayTable = 20

	Vec2usArrayType ArrayTable = 21
	Vec3usArrayType ArrayTable = 22
	Vec4usArrayType ArrayTable = 23

	Vec2uiArrayType ArrayTable = 24
	Vec3uiArrayType ArrayTable = 25
	Vec4uiArrayType ArrayTable = 26

	Vec2ArrayType ArrayTable = 27
	Vec3ArrayType ArrayTable = 28
	Vec4ArrayType ArrayTable = 29

	Vec2dArrayType ArrayTable = 30
	Vec3dArrayType ArrayTable = 31
	Vec4dArrayType ArrayTable = 32

	MatrixArrayType  ArrayTable = 33
	MatrixdArrayType ArrayTable = 34

	QuatArrayType ArrayTable = 35

	UInt64ArrayType ArrayTable = 36
	Int64ArrayType  ArrayTable = 37

	LastArrayType ArrayTable = 37

	DRAWARRAYS         = 50
	DRAWARRAYSLENGTH   = 51
	DRAWElEMENTSUBYTE  = 52
	DRAWElEMENTSUSHORT = 53
	DRAWElEMENTSUINT   = 54

	GL_POINTS         = 0x0000
	GL_LINES          = 0x0001
	GL_LINE_STRIP     = 0x0003
	GL_LINE_LOOP      = 0x0002
	GL_TRIANGLES      = 0x0004
	GL_TRIANGLE_STRIP = 0x0005
	GL_TRIANGLE_FAN   = 0x0006
	GL_QUADS          = 0x0007
	GL_QUAD_STRIP     = 0x0008
	GL_POLYGON        = 0x0009

	GL_LINES_ADJACENCY              = 0x000A
	GL_LINES_ADJACENCY_EXT          = 0x000A
	GL_LINE_STRIP_ADJACENCY_EXT     = 0x000B
	GL_LINE_STRIP_ADJACENCY         = 0x000B
	GL_TRIANGLES_ADJACENCY          = 0x000C
	GL_TRIANGLES_ADJACENCY_EXT      = 0x000C
	GL_TRIANGLE_STRIP_ADJACENCY     = 0x000D
	GL_TRIANGLE_STRIP_ADJACENCY_EXT = 0x000D

	GL_PATCHES = 0x000E

	USE_IMAGE_DATA_FORMAT      = 0
	USE_USER_DEFINED_FORMAT    = 1
	USE_ARB_COMPRESSION        = 2
	USE_S3TC_DXT1_COMPRESSION  = 3
	USE_S3TC_DXT3_COMPRESSION  = 4
	USE_S3TC_DXT5_COMPRESSION  = 5
	USE_PVRTC_2BPP_COMPRESSION = 6
	USE_PVRTC_4BPP_COMPRESSION = 7
	USE_ETC_COMPRESSION        = 8
	USE_ETC2_COMPRESSION       = 9
	USE_RGTC1_COMPRESSION      = 10
	USE_RGTC2_COMPRESSION      = 11
	USE_S3TC_DXT1c_COMPRESSION = 12
	USE_S3TC_DXT1a_COMPRESSION = 13

	NEVER    TextureShadowCompareFunc = 0x0200
	LESS     TextureShadowCompareFunc = 0x0201
	EQUAL    TextureShadowCompareFunc = 0x0202
	LEQUAL   TextureShadowCompareFunc = 0x0203
	GREATER  TextureShadowCompareFunc = 0x0204
	NOTEQUAL TextureShadowCompareFunc = 0x0205
	GEQUAL   TextureShadowCompareFunc = 0x0206
	ALWAYS   TextureShadowCompareFunc = 0x0207

	LUMINANCE TextureShadowTextureMode = 0x1909
	INTENSITY TextureShadowTextureMode = 0x8049
	ALPHA     TextureShadowTextureMode = 0x1906
	NONE                               = 0x0000

	GL_RGBA32UI_EXT            = 0x8D70
	GL_RGB32UI_EXT             = 0x8D71
	GL_ALPHA32UI_EXT           = 0x8D72
	GL_INTENSITY32UI_EXT       = 0x8D73
	GL_LUMINANCE32UI_EXT       = 0x8D74
	GL_LUMINANCE_ALPHA32UI_EXT = 0x8D75

	GL_RGBA16UI_EXT            = 0x8D76
	GL_RGB16UI_EXT             = 0x8D77
	GL_ALPHA16UI_EXT           = 0x8D78
	GL_INTENSITY16UI_EXT       = 0x8D79
	GL_LUMINANCE16UI_EXT       = 0x8D7A
	GL_LUMINANCE_ALPHA16UI_EXT = 0x8D7B

	GL_RGBA8UI_EXT            = 0x8D7C
	GL_RGB8UI_EXT             = 0x8D7D
	GL_ALPHA8UI_EXT           = 0x8D7E
	GL_INTENSITY8UI_EXT       = 0x8D7F
	GL_LUMINANCE8UI_EXT       = 0x8D80
	GL_LUMINANCE_ALPHA8UI_EXT = 0x8D81

	GL_RGBA32I_EXT            = 0x8D82
	GL_RGB32I_EXT             = 0x8D83
	GL_ALPHA32I_EXT           = 0x8D84
	GL_INTENSITY32I_EXT       = 0x8D85
	GL_LUMINANCE32I_EXT       = 0x8D86
	GL_LUMINANCE_ALPHA32I_EXT = 0x8D87

	GL_RGBA16I_EXT            = 0x8D88
	GL_RGB16I_EXT             = 0x8D89
	GL_ALPHA16I_EXT           = 0x8D8A
	GL_INTENSITY16I_EXT       = 0x8D8B
	GL_LUMINANCE16I_EXT       = 0x8D8C
	GL_LUMINANCE_ALPHA16I_EXT = 0x8D8D

	GL_RGBA8I_EXT            = 0x8D8E
	GL_RGB8I_EXT             = 0x8D8F
	GL_ALPHA8I_EXT           = 0x8D90
	GL_INTENSITY8I_EXT       = 0x8D91
	GL_LUMINANCE8I_EXT       = 0x8D92
	GL_LUMINANCE_ALPHA8I_EXT = 0x8D93

	GL_RED_INTEGER_EXT             = 0x8D94
	GL_GREEN_INTEGER_EXT           = 0x8D95
	GL_BLUE_INTEGER_EXT            = 0x8D96
	GL_ALPHA_INTEGER_EXT           = 0x8D97
	GL_RGB_INTEGER_EXT             = 0x8D98
	GL_RGBA_INTEGER_EXT            = 0x8D99
	GL_BGR_INTEGER_EXT             = 0x8D9A
	GL_BGRA_INTEGER_EXT            = 0x8D9B
	GL_LUMINANCE_INTEGER_EXT       = 0x8D9C
	GL_LUMINANCE_ALPHA_INTEGER_EXT = 0x8D9D

	GL_RGBA_INTEGER_MODE_EXT = 0x8D9E

	GL_RG         = 0x8227
	GL_RG_INTEGER = 0x8228
	GL_R8         = 0x8229
	GL_R16        = 0x822A
	GL_RG8        = 0x822B
	GL_RG16       = 0x822C
	GL_R16F       = 0x822D
	GL_R32F       = 0x822E
	GL_RG16F      = 0x822F
	GL_RG32F      = 0x8230
	GL_R8I        = 0x8231
	GL_R8UI       = 0x8232
	GL_R16I       = 0x8233
	GL_R16UI      = 0x8234
	GL_R32I       = 0x8235
	GL_R32UI      = 0x8236
	GL_RG8I       = 0x8237
	GL_RG8UI      = 0x8238
	GL_RG16I      = 0x8239
	GL_RG16UI     = 0x823A
	GL_RG32I      = 0x823B
	GL_RG32UI     = 0x823C

	GL_RGBA32F_ARB            = 0x8814
	GL_RGB32F_ARB             = 0x8815
	GL_ALPHA32F_ARB           = 0x8816
	GL_INTENSITY32F_ARB       = 0x8817
	GL_LUMINANCE32F_ARB       = 0x8818
	GL_LUMINANCE_ALPHA32F_ARB = 0x8819
	GL_RGBA16F_ARB            = 0x881A
	GL_RGB16F_ARB             = 0x881B
	GL_ALPHA16F_ARB           = 0x881C
	GL_INTENSITY16F_ARB       = 0x881D
	GL_LUMINANCE16F_ARB       = 0x881E
	GL_LUMINANCE_ALPHA16F_ARB = 0x881F

	GL_COMPRESSED_RED_RGTC1_EXT              = 0x8DBB
	GL_COMPRESSED_SIGNED_RED_RGTC1_EXT       = 0x8DBC
	GL_COMPRESSED_RED_GREEN_RGTC2_EXT        = 0x8DBD
	GL_COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT = 0x8DBE

	GL_BGR                         = 0x80E0
	GL_BGRA                        = 0x80E1
	GL_UNSIGNED_BYTE_3_3_2         = 0x8032
	GL_UNSIGNED_BYTE_2_3_3_REV     = 0x8362
	GL_UNSIGNED_SHORT_5_6_5        = 0x8363
	GL_UNSIGNED_SHORT_5_6_5_REV    = 0x8364
	GL_UNSIGNED_SHORT_4_4_4_4      = 0x8033
	GL_UNSIGNED_SHORT_4_4_4_4_REV  = 0x8365
	GL_UNSIGNED_SHORT_5_5_5_1      = 0x8034
	GL_UNSIGNED_SHORT_1_5_5_5_REV  = 0x8366
	GL_UNSIGNED_INT_8_8_8_8        = 0x8035
	GL_UNSIGNED_INT_8_8_8_8_REV    = 0x8367
	GL_UNSIGNED_INT_10_10_10_2     = 0x8036
	GL_UNSIGNED_INT_2_10_10_10_REV = 0x8368

	GL_ALPHA4              = 0x803B
	GL_ALPHA8              = 0x803C
	GL_ALPHA12             = 0x803D
	GL_ALPHA16             = 0x803E
	GL_BITMAP              = 0x1A00
	GL_COLOR_INDEX         = 0x1900
	GL_INTENSITY12         = 0x804C
	GL_INTENSITY16         = 0x804D
	GL_INTENSITY4          = 0x804A
	GL_INTENSITY8          = 0x804B
	GL_LUMINANCE12         = 0x8041
	GL_LUMINANCE12_ALPHA4  = 0x8046
	GL_LUMINANCE12_ALPHA12 = 0x8047
	GL_LUMINANCE16         = 0x8042
	GL_LUMINANCE16_ALPHA16 = 0x8048
	GL_LUMINANCE4          = 0x803F
	GL_LUMINANCE4_ALPHA4   = 0x8043
	GL_LUMINANCE6_ALPHA2   = 0x8044
	GL_LUMINANCE8          = 0x8040
	GL_LUMINANCE8_ALPHA8   = 0x8045
	GL_RGBA8               = 0x8058
	GL_RGBA16              = 0x805B
	GL_PACK_ROW_LENGTH     = 0x0D02

	GL_DEPTH_COMPONENT16 = 0x81A5
	GL_DEPTH_COMPONENT24 = 0x81A6
	GL_DEPTH_COMPONENT32 = 0x81A7

	GL_KHR_texture_compression_astc_hdr       = 1
	GL_COMPRESSED_RGBA_ASTC_4x4_KHR           = 0x93B0
	GL_COMPRESSED_RGBA_ASTC_5x4_KHR           = 0x93B1
	GL_COMPRESSED_RGBA_ASTC_5x5_KHR           = 0x93B2
	GL_COMPRESSED_RGBA_ASTC_6x5_KHR           = 0x93B3
	GL_COMPRESSED_RGBA_ASTC_6x6_KHR           = 0x93B4
	GL_COMPRESSED_RGBA_ASTC_8x5_KHR           = 0x93B5
	GL_COMPRESSED_RGBA_ASTC_8x6_KHR           = 0x93B6
	GL_COMPRESSED_RGBA_ASTC_8x8_KHR           = 0x93B7
	GL_COMPRESSED_RGBA_ASTC_10x5_KHR          = 0x93B8
	GL_COMPRESSED_RGBA_ASTC_10x6_KHR          = 0x93B9
	GL_COMPRESSED_RGBA_ASTC_10x8_KHR          = 0x93BA
	GL_COMPRESSED_RGBA_ASTC_10x10_KHR         = 0x93BB
	GL_COMPRESSED_RGBA_ASTC_12x10_KHR         = 0x93BC
	GL_COMPRESSED_RGBA_ASTC_12x12_KHR         = 0x93BD
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR   = 0x93D0
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR   = 0x93D1
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR   = 0x93D2
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR   = 0x93D3
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR   = 0x93D4
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR   = 0x93D5
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR   = 0x93D6
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR   = 0x93D7
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR  = 0x93D8
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR  = 0x93D9
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR  = 0x93DA
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR = 0x93DB
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR = 0x93DC
	GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR = 0x93DD

	GL_DEPTH_COMPONENT32F    = 0x8CAC
	GL_DEPTH_COMPONENT32F_NV = 0x8DAB
	GL_DEPTH_COMPONENT       = 0x1902
	GL_STENCIL_INDEX         = 0x1901

	GL_HILO_NV                        = 0x86F4
	GL_DSDT_NV                        = 0x86F5
	GL_DSDT_MAG_NV                    = 0x86F6
	GL_DSDT_MAG_VIB_NV                = 0x86F7
	GL_HILO16_NV                      = 0x86F8
	GL_SIGNED_HILO_NV                 = 0x86F9
	GL_SIGNED_HILO16_NV               = 0x86FA
	GL_SIGNED_RGBA_NV                 = 0x86FB
	GL_SIGNED_RGBA8_NV                = 0x86FC
	GL_SIGNED_RGB_NV                  = 0x86FE
	GL_SIGNED_RGB8_NV                 = 0x86FF
	GL_SIGNED_LUMINANCE_NV            = 0x8701
	GL_SIGNED_LUMINANCE8_NV           = 0x8702
	GL_SIGNED_LUMINANCE_ALPHA_NV      = 0x8703
	GL_SIGNED_LUMINANCE8_ALPHA8_NV    = 0x8704
	GL_SIGNED_ALPHA_NV                = 0x8705
	GL_SIGNED_ALPHA8_NV               = 0x8706
	GL_SIGNED_INTENSITY_NV            = 0x8707
	GL_SIGNED_INTENSITY8_NV           = 0x8708
	GL_DSDT8_NV                       = 0x8709
	GL_DSDT8_MAG8_NV                  = 0x870A
	GL_DSDT8_MAG8_INTENSITY8_NV       = 0x870B
	GL_SIGNED_RGB_UNSIGNED_ALPHA_NV   = 0x870C
	GL_SIGNED_RGB8_UNSIGNED_ALPHA8_NV = 0x870D

	GL_COMPRESSED_ALPHA           = 0x84E9
	GL_COMPRESSED_LUMINANCE       = 0x84EA
	GL_COMPRESSED_LUMINANCE_ALPHA = 0x84EB
	GL_COMPRESSED_INTENSITY       = 0x84EC
	GL_COMPRESSED_RGB             = 0x84ED
	GL_COMPRESSED_RGBA            = 0x84EE
)
