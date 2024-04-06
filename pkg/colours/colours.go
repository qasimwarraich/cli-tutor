package colours

type ColourPallete string

const (
	DefaultWelcome ColourPallete = "#FF2FF2"
	DefaultError   ColourPallete = "#FF0000"
	DefaultNote    ColourPallete = "#FFFF00"
	DefaultTip     ColourPallete = "#00FFF0"
	DefaultGuide   ColourPallete = "#00FF00"
	// Bg Shades
	SumiInk0 ColourPallete = "#16161D"
	SumiInk1 ColourPallete = "#181820"
	SumiInk2 ColourPallete = "#1a1a22"
	SumiInk3 ColourPallete = "#1F1F28"
	SumiInk4 ColourPallete = "#2A2A37"
	SumiInk5 ColourPallete = "#363646"
	SumiInk6 ColourPallete = "#54546D" //fg

	// Popup and Floats
	WaveBlue1 ColourPallete = "#223249"
	WaveBlue2 ColourPallete = "#2D4F67"
	// Diff and Git
	WinterGreen  ColourPallete = "#2B3328"
	WinterYellow ColourPallete = "#49443C"
	WinterRed    ColourPallete = "#43242B"
	WinterBlue   ColourPallete = "#252535"
	AutumnGreen  ColourPallete = "#76946A"
	AutumnRed    ColourPallete = "#C34043"
	AutumnYellow ColourPallete = "#DCA561"

	// Diag
	SamuraiRed  ColourPallete = "#E82424"
	RoninYellow ColourPallete = "#FF9E3B"
	WaveAqua1   ColourPallete = "#6A9589"
	DragonBlue  ColourPallete = "#658594"

	// Fg and Comments
	OldWhite  ColourPallete = "#C8C093"
	FujiWhite ColourPallete = "#DCD7BA"
	FujiGray  ColourPallete = "#727169"

	OniViolet     ColourPallete = "#957FB8"
	OniViolet2    ColourPallete = "#b8b4d0"
	CrystalBlue   ColourPallete = "#7E9CD8"
	SpringViolet1 ColourPallete = "#938AA9"
	SpringViolet2 ColourPallete = "#9CABCA"
	SpringBlue    ColourPallete = "#7FB4CA"
	LightBlue     ColourPallete = "#A3D4D5" // unused yet
	WaveAqua2     ColourPallete = "#7AA89F" // improve lightness: desaturated greenish Aqua

	//  WaveAqua2  ColourPallete = "#68AD99"
	//  WaveAqua4  ColourPallete = "#7AA880"
	//  WaveAqua5  ColourPallete = "#6CAF95"
	//  WaveAqua3  ColourPallete = "#68AD99"

	SpringGreen ColourPallete = "#98BB6C"
	BoatYellow1 ColourPallete = "#938056"
	BoatYellow2 ColourPallete = "#C0A36E"
	CarpYellow  ColourPallete = "#E6C384"

	SakuraPink   ColourPallete = "#D27E99"
	WaveRed      ColourPallete = "#E46876"
	PeachRed     ColourPallete = "#FF5D62"
	SurimiOrange ColourPallete = "#FFA066"
	KatanaGray   ColourPallete = "#717C7C"

	DragonBlack0 ColourPallete = "#0d0c0c"
	DragonBlack1 ColourPallete = "#12120f"
	DragonBlack2 ColourPallete = "#1D1C19"
	DragonBlack3 ColourPallete = "#181616"
	DragonBlack4 ColourPallete = "#282727"
	DragonBlack5 ColourPallete = "#393836"
	DragonBlack6 ColourPallete = "#625e5a"

	DragonWhite               ColourPallete = "#c5c9c5"
	DragonGreen               ColourPallete = "#87a987"
	DragonGreen2              ColourPallete = "#8a9a7b"
	DragonPink                ColourPallete = "#a292a3"
	DragonOrange              ColourPallete = "#b6927b"
	DragonOrange2             ColourPallete = "#b98d7b"
	DragonGray                ColourPallete = "#a6a69c"
	DragonGray2               ColourPallete = "#9e9b93"
	DragonGray3               ColourPallete = "#7a8382"
	DragonBlue2               ColourPallete = "#8ba4b0"
	DragonVioletColourPalette ColourPallete = "#8992a7"
	DragonRed                 ColourPallete = "#c4746e"
	DragonAqua                ColourPallete = "#8ea4a2"
	DragonAsh                 ColourPallete = "#737c73"
	DragonTeal                ColourPallete = "#949fb5"
	DragonYellow              ColourPallete = "#c4b28a" //"#a99c8b,"
	// "#8a9aa3"

	LotusInk1    ColourPallete = "#545464"
	LotusInk2    ColourPallete = "#43436c"
	LotusGray    ColourPallete = "#dcd7ba"
	LotusGray2   ColourPallete = "#716e61"
	LotusGray3   ColourPallete = "#8a8980"
	LotusWhite0  ColourPallete = "#d5cea3"
	LotusWhite1  ColourPallete = "#dcd5ac"
	LotusWhite2  ColourPallete = "#e5ddb0"
	LotusWhite3  ColourPallete = "#f2ecbc"
	LotusWhite4  ColourPallete = "#e7dba0"
	LotusWhite5  ColourPallete = "#e4d794"
	LotusViolet1 ColourPallete = "#a09cac"
	LotusViolet2 ColourPallete = "#766b90"
	LotusViolet3 ColourPallete = "#c9cbd1"
	LotusViolet4 ColourPallete = "#624c83"
	LotusBlue1   ColourPallete = "#c7d7e0"
	LotusBlue2   ColourPallete = "#b5cbd2"
	LotusBlue3   ColourPallete = "#9fb5c9"
	LotusBlue4   ColourPallete = "#4d699b"
	LotusBlue5   ColourPallete = "#5d57a3"
	LotusGreen   ColourPallete = "#6f894e"
	LotusGreen2  ColourPallete = "#6e915f"
	LotusGreen3  ColourPallete = "#b7d0ae"
	LotusPink    ColourPallete = "#b35b79"
	LotusOrange  ColourPallete = "#cc6d00"
	LotusOrange2 ColourPallete = "#e98a00"
	LotusYellow  ColourPallete = "#77713f"
	LotusYellow2 ColourPallete = "#836f4a"
	LotusYellow3 ColourPallete = "#de9800"
	LotusYellow4 ColourPallete = "#f9d791"
	LotusRed     ColourPallete = "#c84053"
	LotusRed2    ColourPallete = "#d7474b"
	LotusRed3    ColourPallete = "#e82424"
	LotusRed4    ColourPallete = "#d9a594"
	LotusAqua    ColourPallete = "#597b75"
	LotusAqua2   ColourPallete = "#5e857a"
	LotusTeal1   ColourPallete = "#4e8ca2"
	LotusTeal2   ColourPallete = "#6693bf"
	LotusTeal3   ColourPallete = "#5a7785"
	LotusCyan    ColourPallete = "#d7e3d8 "
)

func ColourAsString(colour ColourPallete) string {
	return string(colour)
}
