package client

import (
	"fmt"
	"testing"

	"github.com/LagrangeDev/LagrangeGo/utils"
	binary2 "github.com/LagrangeDev/LagrangeGo/utils/binary"
)

func TestQQClient_FecthQrcode(t *testing.T) {
	//data := utils.GetBytesFromHex("0000031d0002031d0031000000000000000000000000000000000000000000030000003200000000000000000000000000005f5e164f000018e362dbda2f6d7f1dc34f99d73f6a313a87f6a691f707cb9a00030017025489504e470d0a1a0a0000000d4948445200000093000000930100000000a29e16b8000000097048597300000b1300000b1301009a9c1800000206494441544889dd963daeeb201085c7a2a08b3780c436e8d852bc01ff6cc0de121ddb406203a6a34099774872df2b5ec3b4d74a117f91d070e6cc9910ffffd0af6637d1eacdc6f8c225d0246185dbc2f9d6445a1d11af1216cc56d5e9f2ae6989b408d9d3b787531c99e56c617b7b54449394715b9d99186ffdc82261d06f89e6fbf9a7e910c3733b75717a3a14f5b79743ece57a152fd7b60a09f1a380a1f0a326f269253e9d153174fbd955541072f56d96300ef97466452d5dfbcc2216e9e10d399a83bdf547fb61167261dc3edfbe776f96b0c2388c77dd6b215245c25e3e9f64779796d0deb711b03e5594a6de34b83517092b215fa191cf37a9dd9b59c218ed227b3bbe2a24fcccc7282b117d8378aa602ed94c12c6017adb5da7559b557f7d3fc84a84cb6c8976a77cfa2c643826736ce4d293681331cea7e6ddf7d15cc2b79641f622b3756b9ba9a20f9ff9186698c5dae6d81e305a358b84e18c874ba491be86f48f5f061962af724f91ca57fcd1608c71ed79c03d3bcde3abfd2883536e8490c3642bf47f133132e4e9ed7124d94f660f32cf3ba9a3d243b7257e736d9451da02adc041f5ec97303cf029739a19c29b4dc2fa3e770d2ebb021fd1cc12d6f779b585ed153194f610b1be57d3d2bd9677b22c6535cdb52fabada649c8b68ae4b3a5623daa43c4708f9810d81bbfd7b284dd3d032061de7d236d59c246ffc3fd0ef60768294b992e8185210000000049454e44ae426082001c000600000078000200d10068124468747470733a2f2f74787a2e71712e636f6d2f703f6b3d34324c62326939746678334454356e5850326f784f6f6632707048334238756126663d313630303030313631351a2034324c62326939746678334454356e5850326f784f6f6632707048334238756103")
	//decrypted := binary2.NewReader(data)
	//decrypted.ReadBytes(54)
	//retCode := decrypted.ReadU8()
	//qrsig := decrypted.ReadBytesWithLength("u16", false)
	//tlvs := decrypted.ReadTlv()
	//fmt.Println(retCode)
	//for tag, tlv := range tlvs {
	//	fmt.Printf("%d:%x\n", tag, tlv)
	//}
	//fmt.Printf("sig: %x\n", qrsig)
	//
	//if retCode == 0 && tlvs[0x17] != nil {
	//	sig := qrsig
	//	png := tlvs[0x17]
	//	url := string(binary2.NewReader(tlvs[209]).ReadBytesWithLength("u16", true))
	//	os.WriteFile("qrcode.png", png, 666)
	//	fmt.Println("url: ", url)
	//	fmt.Printf("sig: %x\n", sig)
	//}
	urlraw := utils.GetBytesFromHex("124468747470733a2f2f74787a2e71712e636f6d2f703f6b3d34324c62326939746678334454356e5850326f784f6f6632707048334238756126663d313630303030313631351a2034324c62326939746678334454356e5850326f784f6f66327070483342387561")
	reader := binary2.NewReader(urlraw)
	reader.ReadU8()
	url := reader.ReadBytesWithLength("u8", false)
	fmt.Println(string(url))
}
