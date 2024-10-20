package belajar

import (
	"fmt"
	"github.com/whatsauth/watoken"
	"testing"
)

// var privatekey = "privatekey"
var publickeyb = "publickey"
var encode = "encode"

func TestGenerateKeyPASETO(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("asoy", privateKey)
	fmt.Println(hasil, err)
}

func TestHashPass(t *testing.T) {
	password := "cihuypass"

	Hashedpass, err := HashPass(password)
	fmt.Println("error : ", err)
	fmt.Println("Hash : ", Hashedpass)
}

func TestHashFunc(t *testing.T) {
	conn := MongoCreateConnection("MONGOSTRING", "tesgis")
	userdata := new(User)
	userdata.Username = "faisal"
	userdata.Password = "ganteng"

	data := GetOneUser(conn, "user", User{
		Username: userdata.Username,
		Password: userdata.Password,
	})
	fmt.Printf("%+v", data)
	fmt.Println(" ")
	hashpass, _ := HashPass(userdata.Password)
	fmt.Println("Hasil hash : ", hashpass)
	compared := CompareHashPass(userdata.Password, data.Password)
	fmt.Println("result : ", compared)
}

func TestTokenEncoder(t *testing.T) {
	conn := MongoCreateConnection("MONGOSTRING", "tesgis")
	privateKey, publicKey := watoken.GenerateKey()
	userdata := new(User)
	userdata.Username = "faisal"
	userdata.Password = "ganteng"

	data := GetOneUser(conn, "user", User{
		Username: userdata.Username,
		Password: userdata.Password,
	})
	fmt.Println("Private Key : ", privateKey)
	fmt.Println("Public Key : ", publicKey)
	fmt.Printf("%+v", data)
	fmt.Println(" ")

	encode := TokenEncoder(data.Username, privateKey)
	fmt.Printf("%+v", encode)
}

func TestInsertUserdata(t *testing.T) {
	conn := MongoCreateConnection("MONGOSTRING", "tesgis")
	password, err := HashPass("passcek")
	fmt.Println("err", err)
	data := InsertUserdata(conn, "faisalash", "role", password)
	fmt.Println(data)
}

func TestDecodeToken(t *testing.T) {
	deco := watoken.DecodeGetId("public",
		"token")
	fmt.Println(deco)
}

func TestCompareUsername(t *testing.T) {
	conn := MongoCreateConnection("MONGOSTRING", "tesgis")
	deco := watoken.DecodeGetId("public",
		"token")
	compare := CompareUsername(conn, "user", deco)
	fmt.Println(compare)
}

func TestEncodeWithRole(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	role := "admin"
	username := "faidiq"
	encoder, err := EncodeWithRole(role, username, privateKey)

	fmt.Println(" error :", err)
	fmt.Println("Private :", privateKey)
	fmt.Println("Public :", publicKey)
	fmt.Println("encode: ", encoder)

}

func TestDecoder2(t *testing.T) {
	pay, err := Decoder(publickeyb, encode)
	user, _ := DecodeGetUser(publickeyb, encode)
	role, _ := DecodeGetRole(publickeyb, encode)
	use, ro := DecodeGetRoleandUser(publickeyb, encode)
	fmt.Println("user :", user)
	fmt.Println("role :", role)
	fmt.Println("user and role :", use, ro)
	fmt.Println("err : ", err)
	fmt.Println("payload : ", pay)
}

// TestGetOTP
func TestGetOTP(t *testing.T) {
	// Ambil secret dari database untuk pengguna tertentu
	username := "faisal" // Ganti dengan username yang sesuai
	secret := GetUserSecret(username) // Pastikan Anda memiliki fungsi untuk mendapatkan secret

	if secret == "" {
		t.Error("Secret untuk pengguna tidak ditemukan")
		return
	}

	// Menghasilkan OTP dari secret
	otp := GenerateOTP(secret) // Pastikan Anda memiliki fungsi untuk menghasilkan OTP

	// Menampilkan OTP
	fmt.Println("OTP untuk pengguna", username, "adalah:", otp)

	// Verifikasi OTP (misalnya, menggunakan OTP yang benar)
	isValid := VerifyOTP(secret, otp) // Pastikan Anda memiliki fungsi untuk memverifikasi OTP
	if !isValid {
		t.Error("OTP yang dihasilkan tidak valid")
	}
}