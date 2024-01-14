func main 


import {
	"fmt"
	"os"
	"img/png"

	"github.com/skip2/go-qrcode"
}

func main() {
	textToEncode : "Hello, World"

	//Generate Code QR
	qrcode,err := qrcodeEncode(textToEncode.qrcode.Medium,256)
	if err != nil {
		fmt.Println("Error Generting QR Code", err)
		return
	}

	//Save QR to file
	filName := "qrcode.png"
	err = SavQRCodeToFile(fileName, qrCode)
	if err != nil {
		fmt.Println("Error Generting QR Code", err)
		return
	}

	fmt.Println("Berhasil di generate", fileName)
}

func SavQRCodeToFile(fileName string, qrCode []byte) error {
	 //Membuat new file
	 file.err := os.Crete(fileName)
	 if err != nil {
		return
	}
	defer file.Close()
	
	//Create new PNG deoder
	img, err := png.Decode(byte,NewReader(qrCode))
	if err != nil {
		return err
	}

	// menyimpan gambar code qr
	err = png.Encode(file,img)
	if err != nil {
		return err
	}

	return nil

}