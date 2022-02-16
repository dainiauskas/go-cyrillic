package cyrillic

// Converting table
// 185 (¹) - 8470 (№)
// 196 (Ä) - 1044 (Д)
// 197 (Å) - 1045 (Е)
// 201 (É) - 1049 (Й)
// 211 (Ó) - 1059 (У)
// 213 (Õ) - 1061 (Х)
// 214 (Ö) - 1062 (Ц)
// 215 (×) - 1063 (Ч)
// 216 (Ø) - 1025 (Ё)
// 220 (Ü) - 1068 (Ь)
// 223 (ß) - 1071 (Я)
// 228 (ä) - 1076 (д)
// 229 (å) - 1077 (е)
// 233 (é) - 1081 (й)
// 243 (ó) - 1091 (у)
// 245 (õ) - 1093 (х)
// 246 (ö) - 1094 (ц)
// 247 (÷) - 1095 (ч)
// 248 (ø) - 1105 (ё)
// 252 (ü) - 1100 (ь)
// 256 (Ā) - 1042 (В)
// 257 (ā) - 1074 (в)
// 260 (Ą) - 1040 (А)
// 261 (ą) - 1072 (а)
// 262 (Ć) - 1043 (Г)
// 263 (ć) - 1075 (г)
// 268 (Č) - 1048 (И)
// 269 (č) - 1080 (и)
// 274 (Ē) - 1047 (З)
// 275 (ē) - 1079 (з)
// 278 (Ė) - 1051 (Л)
// 279 (ė) - 1083 (л)
// 280 (Ę) - 1046 (Ж)
// 281 (ę) - 1078 (ж)
// 290 (Ģ) - 1052 (М)
// 291 (ģ) - 1084 (м)
// 298 (Ī) - 1054 (О)
// 299 (ī) - 1086 (о)
// 302 (Į) - 1041 (Б)
// 303 (į) - 1073 (б)
// 310 (Ķ) - 1053 (Н)
// 311 (ķ) - 1085 (н)
// 315 (Ļ) - 1055 (П)
// 316 (ļ) - 1087 (п)
// 321 (Ł) - 1065 (Щ)
// 322 (ł) - 1097 (щ)
// 323 (Ń) - 1057 (С)
// 324 (ń) - 1089 (с)
// 325 (Ņ) - 1058 (Т)
// 326 (ņ) - 1090 (т)
// 332 (Ō) - 1060 (Ф)
// 333 (ō) - 1092 (ф)
// 346 (Ś) - 1066 (Ъ)
// 347 (ś) - 1098 (ъ)
// 352 (Š) - 1056 (Р)
// 353 (š) - 1088 (р)
// 362 (Ū) - 1067 (Ы)
// 363 (ū) - 1099 (ы)
// 370 (Ų) - 1064 (Ш)
// 371 (ų) - 1096 (ш)
// 377 (Ź) - 1050 (К)
// 378 (ź) - 1082 (к)
// 379 (Ż) - 1069 (Э)
// 380 (ż) - 1101 (э)
// 381 (Ž) - 1070 (Ю)
// 382 (ž) - 1102 (ю)
// 729 (˙) - 1103 (я)

var (
	rusWin = []rune{
		248, 233, 246, 243, 378, 229, 311, 263, 371, 322, 275, 245, 347,
		333, 363, 257, 261, 316, 353, 299, 279, 228, 281, 380, 729, 247,
		324, 291, 269, 326, 252, 303, 382, 216, 201, 214, 211, 377, 197,
		310, 262, 370, 321, 274, 213, 346, 332, 362, 256, 260, 315, 352,
		298, 278, 196, 280, 379, 223, 215, 323, 290, 268, 325, 220, 302,
		381, 185,
	}

	rusUTF = []rune{
		1105, 1081, 1094, 1091, 1082, 1077, 1085, 1075, 1096, 1097, 1079,
		1093, 1098, 1092, 1099, 1074, 1072, 1087, 1088, 1086, 1083, 1076,
		1078, 1101, 1103, 1095, 1089, 1084, 1080, 1090, 1100, 1073, 1102,
		1025, 1049, 1062, 1059, 1050, 1045, 1053, 1043, 1064, 1065, 1047,
		1061, 1066, 1060, 1067, 1042, 1040, 1055, 1056, 1054, 1051, 1044,
		1046, 1069, 1071, 1063, 1057, 1052, 1048, 1058, 1068, 1041, 1070,
		8470,
	}
)

// ToUTF8 convert to UTF8
func ToUTF8(v string) string {
	r := []rune(v)

	for i, v := range r {
		r[i] = find(v)
	}

	return string(r)
}

func find(v rune) rune {
	for x, e := range rusWin {
		if v == e {
			v = rusUTF[x]
			break
		}
	}

	return v
}
