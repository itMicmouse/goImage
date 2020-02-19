package goImage

import (
	"container/list"
	"fmt"
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"image"
	"image/draw"
	"os"
)

const channelName = "goImage"

// GoImagePlugin implements flutter.Plugin and handles method.
type GoImagePlugin struct{}

var _ flutter.Plugin = &GoImagePlugin{} // compile-time type check

var listInfo = list.New();

// InitPlugin initializes the plugin.
func (p *GoImagePlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getPlatformVersion", p.handlePlatformVersion)
	return nil
}

func (p *GoImagePlugin) handlePlatformVersion(arguments interface{}) (reply interface{}, err error) {
	return "go-flutter " + flutter.PlatformVersion, nil
}

type ImageInfo struct {
	ID  int64
	funct flutter.ExternalTextureHanlderFunc
}

// InitPluginTexture is used to create and manage backend textures
func (p *GoImagePlugin) InitPluginTexture(registry *flutter.TextureRegistry) error {

	texture := registry.NewTexture()
	texture.ID = 1

	listInfo.PushBack(newImageInfo(1,p.textureHanler1))
	listInfo.PushBack(newImageInfo(2,p.textureHanler2))
	listInfo.PushBack(newImageInfo(3,p.textureHanler3))
	listInfo.PushBack(newImageInfo(4,p.textureHanler4))
	listInfo.PushBack(newImageInfo(5,p.textureHanler5))
	listInfo.PushBack(newImageInfo(6,p.textureHanler6))
	listInfo.PushBack(newImageInfo(7,p.textureHanler7))
	listInfo.PushBack(newImageInfo(8,p.textureHanler8))
	listInfo.PushBack(newImageInfo(9,p.textureHanler9))
	listInfo.PushBack(newImageInfo(10,p.textureHanler10))
	listInfo.PushBack(newImageInfo(11,p.textureHanler11))
	listInfo.PushBack(newImageInfo(12,p.textureHanler12))
	listInfo.PushBack(newImageInfo(13,p.textureHanler13))
	listInfo.PushBack(newImageInfo(14,p.textureHanler14))
	listInfo.PushBack(newImageInfo(15,p.textureHanler15))
	listInfo.PushBack(newImageInfo(16,p.textureHanler16))
	listInfo.PushBack(newImageInfo(17,p.textureHanler17))
	listInfo.PushBack(newImageInfo(18,p.textureHanler18))
	listInfo.PushBack(newImageInfo(19,p.textureHanler19))
	listInfo.PushBack(newImageInfo(20,p.textureHanler20))
	listInfo.PushBack(newImageInfo(21,p.textureHanler21))
	listInfo.PushBack(newImageInfo(22,p.textureHanler22))
	listInfo.PushBack(newImageInfo(23,p.textureHanler23))
	listInfo.PushBack(newImageInfo(24,p.textureHanler24))
	listInfo.PushBack(newImageInfo(25,p.textureHanler25))
	listInfo.PushBack(newImageInfo(26,p.textureHanler26))
	listInfo.PushBack(newImageInfo(27,p.textureHanler27))
	listInfo.PushBack(newImageInfo(28,p.textureHanler28))
	listInfo.PushBack(newImageInfo(29,p.textureHanler29))
	listInfo.PushBack(newImageInfo(30,p.textureHanler30))
	listInfo.PushBack(newImageInfo(31,p.textureHanler31))
	listInfo.PushBack(newImageInfo(32,p.textureHanler32))
	listInfo.PushBack(newImageInfo(33,p.textureHanler33))
	listInfo.PushBack(newImageInfo(34,p.textureHanler34))
	listInfo.PushBack(newImageInfo(35,p.textureHanler35))
	listInfo.PushBack(newImageInfo(36,p.textureHanler36))
	listInfo.PushBack(newImageInfo(37,p.textureHanler37))
	listInfo.PushBack(newImageInfo(38,p.textureHanler38))
	listInfo.PushBack(newImageInfo(39,p.textureHanler39))
	listInfo.PushBack(newImageInfo(40,p.textureHanler40))
	listInfo.PushBack(newImageInfo(41,p.textureHanler41))
	listInfo.PushBack(newImageInfo(42,p.textureHanler42))
	listInfo.PushBack(newImageInfo(43,p.textureHanler43))
	listInfo.PushBack(newImageInfo(44,p.textureHanler44))
	listInfo.PushBack(newImageInfo(45,p.textureHanler45))
	listInfo.PushBack(newImageInfo(46,p.textureHanler46))
	listInfo.PushBack(newImageInfo(47,p.textureHanler47))
	listInfo.PushBack(newImageInfo(48,p.textureHanler48))
	listInfo.PushBack(newImageInfo(49,p.textureHanler49))
	listInfo.PushBack(newImageInfo(50,p.textureHanler50))
	listInfo.PushBack(newImageInfo(51,p.textureHanler51))
	listInfo.PushBack(newImageInfo(52,p.textureHanler52))
	listInfo.PushBack(newImageInfo(53,p.textureHanler53))
	listInfo.PushBack(newImageInfo(54,p.textureHanler54))
	listInfo.PushBack(newImageInfo(55,p.textureHanler55))
	listInfo.PushBack(newImageInfo(56,p.textureHanler56))
	listInfo.PushBack(newImageInfo(57,p.textureHanler57))
	listInfo.PushBack(newImageInfo(58,p.textureHanler58))
	listInfo.PushBack(newImageInfo(59,p.textureHanler59))
	listInfo.PushBack(newImageInfo(60,p.textureHanler60))
	listInfo.PushBack(newImageInfo(61,p.textureHanler61))
	listInfo.PushBack(newImageInfo(62,p.textureHanler62))
	listInfo.PushBack(newImageInfo(63,p.textureHanler63))
	listInfo.PushBack(newImageInfo(64,p.textureHanler64))
	listInfo.PushBack(newImageInfo(65,p.textureHanler65))
	listInfo.PushBack(newImageInfo(66,p.textureHanler66))
	listInfo.PushBack(newImageInfo(67,p.textureHanler67))
	listInfo.PushBack(newImageInfo(68,p.textureHanler68))
	listInfo.PushBack(newImageInfo(69,p.textureHanler69))
	listInfo.PushBack(newImageInfo(70,p.textureHanler70))
	listInfo.PushBack(newImageInfo(71,p.textureHanler71))
	listInfo.PushBack(newImageInfo(72,p.textureHanler72))
	listInfo.PushBack(newImageInfo(73,p.textureHanler73))
	listInfo.PushBack(newImageInfo(74,p.textureHanler74))
	listInfo.PushBack(newImageInfo(75,p.textureHanler75))
	listInfo.PushBack(newImageInfo(76,p.textureHanler76))
	listInfo.PushBack(newImageInfo(77,p.textureHanler77))
	listInfo.PushBack(newImageInfo(78,p.textureHanler78))
	listInfo.PushBack(newImageInfo(79,p.textureHanler79))
	listInfo.PushBack(newImageInfo(80,p.textureHanler80))
	listInfo.PushBack(newImageInfo(81,p.textureHanler81))
	listInfo.PushBack(newImageInfo(82,p.textureHanler82))
	listInfo.PushBack(newImageInfo(83,p.textureHanler83))
	listInfo.PushBack(newImageInfo(84,p.textureHanler84))
	listInfo.PushBack(newImageInfo(85,p.textureHanler85))

	for e := listInfo.Front(); e != nil; e = e.Next() {
		info, ok := (e.Value).(ImageInfo)
		if ok {
			texture3 := registry.NewTexture()
			texture3.ID = info.ID;
			texture3.Register(info.funct)
			fmt.Printf("texture_tutorial:%d", info.ID)
		} else {

		}
	}


	// after 5 seconds, remove the texture from the scene.
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	fmt.Printf("texture_tutorial: UnRegistering texture %v after 5 seconds\n", texture.ID)
	//	err := texture.UnRegister()
	//	if err != nil {
	//		fmt.Printf("texture_tutorial: %v\n", err)
	//	}
	//	texture.FrameAvailable() // redraw
	//}()

	return texture.Register(p.textureHanler)
}

func newImageInfo(id int64, funct flutter.ExternalTextureHanlderFunc) ImageInfo {
	return ImageInfo{id,funct};
}

func (p *GoImagePlugin) textureHanler(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/code.png"
	return getFile(file)
}


func (p *GoImagePlugin) textureHanler1(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/label_add.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler2(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/personnel.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler3(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_screen_blue.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler4(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/warning.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler5(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/key_back.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler6(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/template_line.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler7(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_medical_pay_unselected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler8(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/add_white.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler9(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/loginIconPasswordActive.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler10(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/code.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler11(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/personnel_permission.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler12(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/loginIconUser.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler13(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/add_blue.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler14(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/pop_bg.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler15(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_aili_pay_unselected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler16(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_back.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler17(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_medical_pay_selected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler18(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_more_pay_unselected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler19(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_search.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler20(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/clossButton.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler21(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/search_close.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler22(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/iconInfo.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler23(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/template_add.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler24(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/loginIconPassword.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler25(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_aili_pay_selected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler26(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/popdownDia.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler27(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_shaixuanguanbi.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler28(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/shiliang_normal.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler29(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/set_inspection.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler30(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_wechat_pay_unselected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler31(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_dian.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler32(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/loginIconUserActive.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler33(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/shiliang_selected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler34(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_logo.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler35(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/clander.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler36(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_setting.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler37(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/sexboy.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler38(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_stored_pay_unselected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler39(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_union_pay_unselected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler40(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_heart.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler41(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/diagnosis_right.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler42(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/README.md"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler43(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/set_supplier.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler44(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/set_tag.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler45(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/serverip.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler46(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/template_arrow.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler47(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/zhanghu.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler48(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/loginIconEye.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler49(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_down.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler50(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/logo.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler51(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/loginBgHeader.jpg"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler52(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/bottom_detail.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler53(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/more_template_bg.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler54(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/set_print.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler55(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/personnel1.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler56(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/sexgirl.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler57(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/serveriptrue.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler58(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_home.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler59(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/right_detail.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler60(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_more_pay_selected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler61(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/print-template.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler62(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/set_prescription_module.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler63(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_group.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler64(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/template_cancel.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler65(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_money_pay_unselected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler66(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/search_no_data.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler67(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/right_white.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler68(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_screen.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler69(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_wechat_pay_selected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler70(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_computer.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler71(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_stored_pay_selected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler72(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/loginIconEyeShow.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler73(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/loginIconClose.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler74(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_order.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler75(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/set_member_manager.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler76(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/sealCharge.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler77(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_union_pay_selected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler78(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_delete.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler79(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/label_add@2x.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler80(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/case_no_data.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler81(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/right.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler82(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_chushengnianyu.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler83(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/druglabel_more.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler84(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/icon_money_pay_selected.png"
	return getFile(file)
}
func (p *GoImagePlugin) textureHanler85(width, height int) (bool, *flutter.PixelBuffer) {
	file := "./images/close.png"
	return getFile(file)
}


func getFile(src string) (bool, *flutter.PixelBuffer) {
	imgFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("image %q not found on disk: %v\n", src, err)
		return false, nil
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Printf("error decoding file %s:%v\n", src, err)
		return false, nil
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		fmt.Printf("unsupported stride\n")
		return false, nil
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	return true, &flutter.PixelBuffer{
		Pix:    rgba.Pix,
		Width:  rgba.Bounds().Size().X,
		Height: rgba.Bounds().Size().Y,
	}
}
