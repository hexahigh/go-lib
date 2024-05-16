package ansicolor

// Regular
const Red = "\033[0;31m"
const Black = "\033[0;30m"
const Cyan = "\033[36m"
const Green = "\033[0;32m"
const Yellow = "\033[0;33m"
const Blue = "\033[0;34m"
const Purple = "\033[0;35m"
const White = "\033[0;37m"
const None = "\033[0m"
const Reset = None

// Bold
const BlackBold = "\033[1;30m"
const RedBold = "\033[1;31m"
const GreenBold = "\033[1;32m"
const YellowBold = "\033[1;33m"
const BlueBold = "\033[1;34m"
const PurpleBold = "\033[1;35m"
const CyanBold = "\033[1;36m"
const WhiteBold = "\033[1;37m"

// Underline
const BlackUnderline = "\033[4;30m"
const RedUnderline = "\033[4;31m"
const GreenUnderline = "\033[4;32m"
const YellowUnderline = "\033[4;33m"
const BlueUnderline = "\033[4;34m"
const PurpleUnderline = "\033[4;35m"
const CyanUnderline = "\033[4;36m"
const WhiteUnderline = "\033[4;37m"

// Background
const BlackBackground = "\033[40m"
const RedBackground = "\033[41m"
const GreenBackground = "\033[42m"
const YellowBackground = "\033[43m"
const BlueBackground = "\033[44m"
const PurpleBackground = "\033[45m"
const CyanBackground = "\033[46m"
const WhiteBackground = "\033[47m"

// High Intensity
const BlackHighIntensity = "\033[0;90m"
const RedHighIntensity = "\033[0;91m"
const GreenHighIntensity = "\033[0;92m"
const YellowHighIntensity = "\033[0;93m"
const BlueHighIntensity = "\033[0;94m"
const PurpleHighIntensity = "\033[0;95m"
const CyanHighIntensity = "\033[0;96m"
const WhiteHighIntensity = "\033[0;97m"

// Bold High Intensity
const BlackBoldHighIntensity = "\033[1;90m"
const RedBoldHighIntensity = "\033[1;91m"
const GreenBoldHighIntensity = "\033[1;92m"
const YellowBoldHighIntensity = "\033[1;93m"
const BlueBoldHighIntensity = "\033[1;94m"
const PurpleBoldHighIntensity = "\033[1;95m"
const CyanBoldHighIntensity = "\033[1;96m"
const WhiteBoldHighIntensity = "\033[1;97m"

// High Intensity backgrounds
const BlackBackgroundHighIntensity = "\033[0;100m"
const RedBackgroundHighIntensity = "\033[0;101m"
const GreenBackgroundHighIntensity = "\033[0;102m"
const YellowBackgroundHighIntensity = "\033[0;103m"
const BlueBackgroundHighIntensity = "\033[0;104m"
const PurpleBackgroundHighIntensity = "\033[0;105m"
const CyanBackgroundHighIntensity = "\033[0;106m"
const WhiteBackgroundHighIntensity = "\033[0;107m"

//######
// 24bit
//######

// Regular 24bit
const Black24bit = "\033[0;38;2;0;0;0m"
const Red24bit = "\033[0;38;2;255;0;0m"
const Green24bit = "\033[0;38;2;0;255;0m"
const Yellow24bit = "\033[0;38;2;255;255;0m"
const Blue24bit = "\033[0;38;2;0;0;255m"
const Purple24bit = "\033[0;38;2;255;0;255m"
const Cyan24bit = "\033[0;38;2;0;255;255m"
const White24bit = "\033[0;38;2;255;255;255m"
const Orange24bit = "\033[0;38;2;255;165;0m"
const Gray24bit = "\033[0;38;2;128;128;128m"
const Lime24bit = "\033[0;38;2;0;255;0m"

// Bold 24bit
const WhiteBold24bit = "\033[1;38;2;255;255;255m"
const RedBold24bit = "\033[1;38;2;255;0;0m"
const GreenBold24bit = "\033[1;38;2;0;255;0m"
const YellowBold24bit = "\033[1;38;2;255;255;0m"
const BlueBold24bit = "\033[1;38;2;0;0;255m"
const PurpleBold24bit = "\033[1;38;2;255;0;255m"
const CyanBold24bit = "\033[1;38;2;0;255;255m"
const BlackBold24bit = "\033[1;38;2;0;0;0m"
const OrangeBold24bit = "\033[1;38;2;255;165;0m"
const GrayBold24bit = "\033[1;38;2;128;128;128m"
const LimeBold24bit = "\033[1;38;2;0;255;0m"

// High Intensity 24bit
const WhiteHighIntensity24bit = "\033[0;38;2;255;255;255m"
const RedHighIntensity24bit = "\033[0;38;2;255;0;0m"
const GreenHighIntensity24bit = "\033[0;38;2;0;255;0m"
const YellowHighIntensity24bit = "\033[0;38;2;255;255;0m"
const BlueHighIntensity24bit = "\033[0;38;2;0;0;255m"
const PurpleHighIntensity24bit = "\033[0;38;2;255;0;255m"
const CyanHighIntensity24bit = "\033[0;38;2;0;255;255m"
const BlackHighIntensity24bit = "\033[0;38;2;0;0;0m"
const OrangeHighIntensity24bit = "\033[0;38;2;255;165;0m"
const GrayHighIntensity24bit = "\033[0;38;2;128;128;128m"

// Underline 24bit
const WhiteUnderline24bit = "\033[4;38;2;255;255;255m"
const RedUnderline24bit = "\033[4;38;2;255;0;0m"
const GreenUnderline24bit = "\033[4;38;2;0;255;0m"
const YellowUnderline24bit = "\033[4;38;2;255;255;0m"
const BlueUnderline24bit = "\033[4;38;2;0;0;255m"
const PurpleUnderline24bit = "\033[4;38;2;255;0;255m"
const CyanUnderline24bit = "\033[4;38;2;0;255;255m"
const BlackUnderline24bit = "\033[4;38;2;0;0;0m"
const OrangeUnderline24bit = "\033[4;38;2;255;165;0m"
const GrayUnderline24bit = "\033[4;38;2;128;128;128m"

// Background 24bit
const WhiteBackground24bit = "\033[0;48;2;255;255;255m"
const RedBackground24bit = "\033[0;48;2;255;0;0m"
const GreenBackground24bit = "\033[0;48;2;0;255;0m"
const YellowBackground24bit = "\033[0;48;2;255;255;0m"
const BlueBackground24bit = "\033[0;48;2;0;0;255m"
const PurpleBackground24bit = "\033[0;48;2;255;0;255m"
const CyanBackground24bit = "\033[0;48;2;0;255;255m"
const BlackBackground24bit = "\033[0;48;2;0;0;0m"
const OrangeBackground24bit = "\033[0;48;2;255;165;0m"
const GrayBackground24bit = "\033[0;48;2;128;128;128m"

// Bold high intensity 24bit
const WhiteBoldHighIntensity24bit = "\033[1;38;2;255;255;255m"
const RedBoldHighIntensity24bit = "\033[1;38;2;255;0;0m"
const GreenBoldHighIntensity24bit = "\033[1;38;2;0;255;0m"
const YellowBoldHighIntensity24bit = "\033[1;38;2;255;255;0m"
const BlueBoldHighIntensity24bit = "\033[1;38;2;0;0;255m"
const PurpleBoldHighIntensity24bit = "\033[1;38;2;255;0;255m"
const CyanBoldHighIntensity24bit = "\033[1;38;2;0;255;255m"
const BlackBoldHighIntensity24bit = "\033[1;38;2;0;0;0m"
const OrangeBoldHighIntensity24bit = "\033[1;38;2;255;165;0m"
const GrayBoldHighIntensity24bit = "\033[1;38;2;128;128;128m"

// High Intensity backgrounds 24bit
const BlackBackgroundHighIntensity24bit = "\033[0;48;2;0;0;0m"
const RedBackgroundHighIntensity24bit = "\033[0;48;2;255;0;0m"
const GreenBackgroundHighIntensity24bit = "\033[0;48;2;0;255;0m"
const YellowBackgroundHighIntensity24bit = "\033[0;48;2;255;255;0m"
const BlueBackgroundHighIntensity24bit = "\033[0;48;2;0;0;255m"
const PurpleBackgroundHighIntensity24bit = "\033[0;48;2;255;0;255m"
const CyanBackgroundHighIntensity24bit = "\033[0;48;2;0;255;255m"
const WhiteBackgroundHighIntensity24bit = "\033[0;48;2;255;255;255m"
const OrangeBackgroundHighIntensity24bit = "\033[0;48;2;255;165;0m"
const GrayBackgroundHighIntensity24bit = "\033[0;48;2;128;128;128m"
