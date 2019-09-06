package captcha

import "github.com/mojocn/base64Captcha"

const (
	//CaptchaComplexLower complex level lower.
	CaptchaComplexLower = iota
	//CaptchaComplexMedium complex level medium.
	CaptchaComplexMedium
	//CaptchaComplexHigh complex level high.
	CaptchaComplexHigh
)
const (
	//CaptchaModeNumber mode number.
	CaptchaModeNumber = iota
	//CaptchaModeAlphabet mode alphabet.
	CaptchaModeAlphabet
	//CaptchaModeArithmetic mode arithmetic.
	CaptchaModeArithmetic
	//CaptchaModeNumberAlphabet mode mix number and alphabet,this is also default mode.
	CaptchaModeNumberAlphabet
)

//ConfigAudio 音频验证码
type ConfigAudio struct {
	CaptchaLen int
	Language   string
}

//ConfigDigit 数字验证码
type ConfigDigit struct {
	Height     int
	Width      int
	DotCount   int
	CaptchaLen int
	MaxSkew    float64
}

//ConfigCharacter 字符,公式,验证码配置
type ConfigCharacter struct {
	Height             int
	Width              int
	IsUseSimpleFont    bool
	IsShowHollowLine   bool
	IsShowNoiseDot     bool
	IsShowNoiseText    bool
	IsShowSlimeLine    bool
	IsShowSineLine     bool
	CaptchaLen         int
	Mode               int
	ComplexOfNoiseText int
	ComplexOfNoiseDot  int
}

//NewCaptcha 新生产一个音频的验证码
func (audio *ConfigAudio) NewCaptcha(uuid string) (key string, base64String string) {
	config := base64Captcha.ConfigAudio{
		CaptchaLen: audio.CaptchaLen,
		Language:   audio.Language,
	}
	key, audioCap := base64Captcha.GenerateCaptcha(uuid, config)
	base64String = base64Captcha.CaptchaWriteToBase64Encoding(audioCap)
	return key, base64String
}

//NewCaptcha 新生产一个数字验证码
func (digit *ConfigDigit) NewCaptcha(uuid string) (key string, base64String string) {
	config := base64Captcha.ConfigDigit{
		Height:     digit.Height,
		Width:      digit.Width,
		MaxSkew:    digit.MaxSkew,
		DotCount:   digit.DotCount,
		CaptchaLen: digit.CaptchaLen,
	}
	key, audioCap := base64Captcha.GenerateCaptcha(uuid, config)
	base64String = base64Captcha.CaptchaWriteToBase64Encoding(audioCap)
	return key, base64String
}

//NewCaptcha 新生产一个字符,公式证码
func (character *ConfigCharacter) NewCaptcha(uuid string) (key string, base64String string) {
	config := base64Captcha.ConfigCharacter{
		Height:             character.Height,
		Width:              character.Width,
		Mode:               character.Mode,
		ComplexOfNoiseText: character.ComplexOfNoiseText,
		ComplexOfNoiseDot:  character.ComplexOfNoiseDot,
		IsShowHollowLine:   character.IsShowHollowLine,
		IsShowNoiseDot:     character.IsShowNoiseDot,
		IsShowNoiseText:    character.IsShowNoiseText,
		IsShowSlimeLine:    character.IsShowSlimeLine,
		IsShowSineLine:     character.IsShowSineLine,
		CaptchaLen:         character.CaptchaLen,
	}
	key, audioCap := base64Captcha.GenerateCaptcha(uuid, config)
	base64String = base64Captcha.CaptchaWriteToBase64Encoding(audioCap)
	return key, base64String
}

//VerifyCaptcha 验证验证码
func VerifyCaptcha(identifier, verifyValue string) bool {
	return base64Captcha.VerifyCaptcha(identifier, verifyValue)
}
